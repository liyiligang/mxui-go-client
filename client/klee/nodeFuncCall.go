// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/25 14:36
// Description:

package klee

import (
	"encoding/json"
	"errors"
	"github.com/liyiligang/klee-client-go/klee/typedef"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"reflect"
)

func (client *ManageClient) reqNodeFuncCall(message []byte) error {
	var err error
	req := protoManage.ReqNodeFuncCall{}
	err = req.Unmarshal(message)
	if err != nil {
		return err
	}
	defer func(){
		if err != nil  {
			ans := protoManage.AnsNodeFuncCall{Error: err.Error(), NodeFuncCall: protoManage.NodeFuncCall{
				Base: req.NodeFuncCall.Base, State: protoManage.State_StateError, ManagerID: req.NodeFuncCall.ManagerID,
				FuncID: req.NodeFuncCall.FuncID, ReturnVal: err.Error(), ReturnType: protoManage.NodeFuncReturnType_Error,
			}}
			client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
		}
	}()
	nodeFuncLoad, ok := client.data.nodeFuncMap.Load(req.NodeFuncCall.FuncID)
	if !ok {
		err = errors.New("func callback is non-existent")
		return err
	}

	nodeFunc, ok := nodeFuncLoad.(*NodeFuncRegister)
	if !ok {
		err = errors.New("nodeFuncMap data is error")
		return err
	}

	res, err := client.callFuncByReflect(nodeFunc, req.NodeFuncCall.Parameter)
	if err != nil {
		return err
	}

	returnType := nodeFunc.ReturnType
	baseType := nodeFunc.BaseType
	if returnType == protoManage.NodeFuncReturnType_Unknown {
		rType:=reflect.TypeOf(res)
		returnType, baseType = client.getNodeFuncReturnType(rType)
		if returnType == protoManage.NodeFuncReturnType_Error {
			err, ok = res.(error)
			if !ok {
				return errors.New("data is error")
			}
			return err
		}
	}

	if baseType {
		res = client.packageBaseType(returnType, res)
	}

	var data []byte
	if res != nil {
		data , err = json.Marshal(res)
		if err != nil {
			return err
		}
	}
	ans := protoManage.AnsNodeFuncCall{NodeFuncCall: protoManage.NodeFuncCall{
		Base: req.NodeFuncCall.Base, Parameter: req.NodeFuncCall.Parameter,
		ReturnVal: string(data), ReturnType: returnType, State: protoManage.State_StateNormal,
		ManagerID: req.NodeFuncCall.ManagerID, FuncID: req.NodeFuncCall.FuncID,
	}}
	return client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
}

func (client *ManageClient) callFuncByReflect(nodeFunc *NodeFuncRegister, nodeFuncPara string) (interface{}, error) {
	funcCallRef := reflect.ValueOf(nodeFunc.CallFunc)
	vType:=reflect.TypeOf(nodeFunc.CallFunc)
	var res []reflect.Value
	if vType.NumIn() > 0 {
		in := vType.In(0)
		para := reflect.New(in).Interface()
		err := json.Unmarshal([]byte(nodeFuncPara), para)
		if err != nil {
			return nil, err
		}
		params := make([]reflect.Value, vType.NumIn())
		params[0] = reflect.ValueOf(para).Elem()
		res = funcCallRef.Call(params)
	}else {
		res = funcCallRef.Call(make([]reflect.Value, 0))
	}
	if len(res) > 0 {
		if nodeFunc.ErrorPos >= 0 {
			errVal := res[nodeFunc.ErrorPos].Interface()
			err, ok := errVal.(error)
			if ok {
				if err != nil {
					return nil, err
				}
			}
		}
		return res[0].Interface(), nil
	}
	return nil, nil
}

func (client *ManageClient) packageBaseType(returnType protoManage.NodeFuncReturnType, val interface{}) interface{}{
	if returnType == protoManage.NodeFuncReturnType_Text {
		val = &typedef.NodeFuncReturnText{Data: val}
	}else if returnType == protoManage.NodeFuncReturnType_Json {
		val = &typedef.NodeFuncReturnJson{Data: val}
	}
	return val
}