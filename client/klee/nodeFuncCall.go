// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/25 14:36
// Description:

package klee

import (
	"errors"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
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
	v, ok := client.data.nodeFuncMap.Load(req.NodeFuncCall.FuncID)
	if !ok {
		err = errors.New("func callback is non-existent")
		return err
	}
	callFunc, ok := v.(CallFuncDef)
	if !ok {
		err = errors.New("callFunc data format is error, its type should be CallFuncDef")
		return err
	}
	res, state := callFunc(req.NodeFuncCall.Parameter)
	ans := protoManage.AnsNodeFuncCall{NodeFuncCall: protoManage.NodeFuncCall{
		Base: req.NodeFuncCall.Base, Parameter: req.NodeFuncCall.Parameter,
		ReturnVal: res, State: protoManage.State(state), ManagerID: req.NodeFuncCall.ManagerID,
		FuncID: req.NodeFuncCall.FuncID,
	}}
	return client.sendPB(protoManage.Order_NodeFuncCallAns, &ans)
}
