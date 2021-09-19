// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/25 14:36
// Description:

package klee

import (
	"encoding/json"
	"errors"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"reflect"
)

type NodeFuncCallLevel int32
const (
	NodeFuncCallLevelTimeout      NodeFuncCallLevel =   1
	NodeFuncCallLevelLevelSuccess NodeFuncCallLevel =   2
	NodeFuncCallLevelLevelWarn    NodeFuncCallLevel =   3
	NodeFuncCallLevelLevelError   NodeFuncCallLevel =   4
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
				Base: req.NodeFuncCall.Base, State: protoManage.State_StateError,
				ManagerID: req.NodeFuncCall.ManagerID, FuncID: req.NodeFuncCall.FuncID,
			}}
			client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
		}
	}()
	funcCall, ok := client.data.nodeFuncMap.Load(req.NodeFuncCall.FuncID)
	if !ok {
		err = errors.New("func callback is non-existent")
		return err
	}
	res, err := client.callFuncByReflect(funcCall, req.NodeFuncCall.Parameter)
	if err != nil {
		return err
	}
	byte , err := json.Marshal(res.Value)
	if err != nil {
		return err
	}
	ans := protoManage.AnsNodeFuncCall{NodeFuncCall: protoManage.NodeFuncCall{
		Base: req.NodeFuncCall.Base, Parameter: req.NodeFuncCall.Parameter,
		ReturnVal: string(byte), State: protoManage.State(res.Level), ManagerID: req.NodeFuncCall.ManagerID,
		FuncID: req.NodeFuncCall.FuncID,
	}}
	return client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
}

func (client *ManageClient) callFuncByReflect(funcCall interface{}, nodeFuncPara string) (*NodeFuncResponse, error) {
	vType:=reflect.TypeOf(funcCall)
	in := vType.In(0)
	para := reflect.New(in).Interface()
	err := json.Unmarshal([]byte(nodeFuncPara), para)
	if err != nil {
		return nil, err
	}
	funcCallRef := reflect.ValueOf(funcCall)
	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(para).Elem()
	res := funcCallRef.Call(params)
	if len(res) != 1 {
		return nil, errors.New("返回值数量错误")
	}
	resVal, ok := res[0].Interface().(*NodeFuncResponse)
	if !ok {
		return nil, errors.New("返回值断言失败")
	}
	return resVal, nil
}