                                                  // Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/23 15:52
// Description:

package klee

import (
	"context"
	"errors"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
)

type NodeFuncLevel int32
const (
	NodeFuncLevelVisitor      NodeFuncLevel =   1
	NodeFuncLevelMember       NodeFuncLevel =   2
	NodeFuncLevelManager      NodeFuncLevel =   3
	NodeFuncLevelSuperManager NodeFuncLevel =   4
)

type CallFuncDef func(string) (string, NodeFuncCallLevel)

func (client *ManageClient) RegisterNodeFunc(name string, callFunc CallFuncDef, nodeFuncLevel NodeFuncLevel) error {
	if callFunc == nil {
		return errors.New("callFunc is nil")
	}
	node, err := client.GetNode()
	if err != nil {
		return err
	}
	callName := client.getFuncName(callFunc)
	nodeFunc := protoManage.NodeFunc{NodeID: node.Base.ID, Name: name, Func: callName, State: protoManage.State(nodeFuncLevel)}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	resNodeFunc, err := client.engine.RegisterNodeFunc(ctx, &nodeFunc)
	if err != nil {
		return err
	}
	client.data.nodeFuncMap.Store(resNodeFunc.Base.ID, callFunc)
	return nil
}

