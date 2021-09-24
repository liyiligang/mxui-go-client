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

type NodeFuncCallState int32
const (
	NodeFuncCallStateTimeout      NodeFuncCallState =   1
	NodeFuncCallStateSuccess 	  NodeFuncCallState =   2
	NodeFuncCallStateWarn    	  NodeFuncCallState =   3
	NodeFuncCallStateError   	  NodeFuncCallState =   4
)

type NodeFuncReturnLink struct {
	Name        string
	Link 		string
	AutoOpen	bool
}

type NodeFuncReturnMedia struct {
	URL         string
	Type 	    string
	Live		bool
	Loop		bool
	AutoPlay	bool
}

type NodeFuncReturnFile struct {
	Name        	string
	Content 		[]byte
	AutoDownload	bool
}

type NodeFuncReturnTableCol struct {
	Name        	string
	Width			uint32
	Type            string
	Fixed			string
	Align			string
	Resizable		bool
	MergeSameCol	bool
}

type NodeFuncReturnTableRow struct {
	Data			[]interface{}
	State			protoManage.State
	MergeSameRow	bool
}

type NodeFuncReturnTable struct {
	Stripe			bool
	Border			bool
	ShowSummary     bool
	ShowIndex		bool
	SumText			string
	Col        		[]NodeFuncReturnTableCol
	Row				[]NodeFuncReturnTableRow
}

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
		ReturnVal: string(byte), ReturnType: res.Type, State: protoManage.State(res.State),
		ManagerID: req.NodeFuncCall.ManagerID, FuncID: req.NodeFuncCall.FuncID,
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