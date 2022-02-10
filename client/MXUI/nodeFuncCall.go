/*
 * Copyright 2021 liyiligang.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package MXUI

import (
	"encoding/json"
	"errors"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"reflect"
)

func (client *Client) reqNodeFuncCall(message []byte) error {
	var err error
	req := protoManage.ReqNodeFuncCall{}
	err = req.Unmarshal(message)
	if err != nil {
		return err
	}
	err = client.nodeFuncCall(&req)
	if err != nil  {
		ans := protoManage.AnsNodeFuncCall{Error: err.Error(), NodeFuncCall: protoManage.NodeFuncCall{
			Base: req.NodeFuncCall.Base, State: protoManage.State_StateError, RequesterID: req.NodeFuncCall.RequesterID,
			RequesterName: req.NodeFuncCall.RequesterName, FuncID: req.NodeFuncCall.FuncID,
			ReturnVal: err.Error(), ReturnType: protoManage.NodeFuncReturnType_Error,
		}}
		return client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
	}
	return nil
}

func (client *Client) nodeFuncCall(req *protoManage.ReqNodeFuncCall) error {
	nodeFuncLoad, ok := client.data.nodeFuncMap.Load(req.NodeFuncCall.FuncID)
	if !ok {
		return errors.New("function info callback is not found")
	}
	nodeFunc, ok := nodeFuncLoad.(*NodeFuncRegister)
	if !ok {
		return errors.New("function info assert fail with *NodeFuncRegister")
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
		RequesterID: req.NodeFuncCall.RequesterID, RequesterName: req.NodeFuncCall.RequesterName,
		FuncID: req.NodeFuncCall.FuncID,
	}}
	return client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
}

func (client *Client) callFuncByReflect(nodeFunc *NodeFuncRegister, nodeFuncPara string) (interface{}, error) {
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

func (client *Client) packageBaseType(returnType protoManage.NodeFuncReturnType, val interface{}) interface{}{
	if returnType == protoManage.NodeFuncReturnType_Text {
		val = &NodeFuncReturnText{Data: val}
	}else if returnType == protoManage.NodeFuncReturnType_Json {
		val = &NodeFuncReturnJson{Data: val}
	}
	return val
}