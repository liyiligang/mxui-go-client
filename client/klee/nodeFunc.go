// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/23 15:52
// Description:

package klee

import (
	"context"
	"errors"
	"github.com/alecthomas/jsonschema"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"reflect"
)

type NodeFuncLevel int32
const (
	NodeFuncLevelVisitor      NodeFuncLevel =   1
	NodeFuncLevelMember       NodeFuncLevel =   2
	NodeFuncLevelManager      NodeFuncLevel =   3
	NodeFuncLevelSuperManager NodeFuncLevel =   4
)

type CallFuncDef func(string) (string, NodeFuncCallLevel)

type NodeFuncRegister struct {
	Name 			string
	CallFunc 		interface{}
	Level 			NodeFuncLevel
}

type NodeFuncResponse struct {
	Value 		interface{}
	Level 		NodeFuncCallLevel
}

func (client *ManageClient) RegisterNodeFunc(nodeFunc NodeFuncRegister) error {
	err := client.nodeFuncRegisterCheck(nodeFunc.CallFunc)
	if err != nil {
		return err
	}
	schema, err := client.getNodeFuncJsonSchema(nodeFunc.CallFunc)
	if err != nil {
		return err
	}
	node, err := client.GetNode()
	if err != nil {
		return err
	}
	callName := client.getFuncName(nodeFunc.CallFunc)
	protoNodeFunc := protoManage.NodeFunc{NodeID: node.Base.ID, Name: nodeFunc.Name,
		Func: callName, Schema: schema, Level: protoManage.Level(nodeFunc.Level)}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	resNodeFunc, err := client.engine.RegisterNodeFunc(ctx, &protoNodeFunc)
	if err != nil {
		return err
	}
	client.data.nodeFuncMap.Store(resNodeFunc.Base.ID, nodeFunc.CallFunc)
	return nil
}


func (client *ManageClient) nodeFuncRegisterCheck(callFunc interface{}) error {
	if callFunc == nil {
		return errors.New("CallFunc 不能为nil值")
	}
	vType:=reflect.TypeOf(callFunc)
	if vType.Kind() != reflect.Func {
		return errors.New("CallFunc 必须是 reflect.Func 类型")
	}
	if vType.NumIn() != 1 {
		return errors.New("CallFunc 参数数量必须为1")
	}
	if vType.In(0).Kind() == reflect.Ptr {
		if vType.In(0).Elem().Kind() != reflect.Struct {
			return errors.New("函数参数必须是 *reflect.Struct 类型")
		}
	}else if vType.In(0).Kind() != reflect.Struct {
		return errors.New("函数参数必须是 reflect.Struct 类型")
	}

	if vType.NumOut() != 1 {
		return errors.New("CallFunc 返回值数量必须为1")
	}
	if vType.Out(0) != reflect.TypeOf(&NodeFuncResponse{}) {
		return errors.New("CallFunc 返回值必须是 &NodeFuncResponse{}")
	}
	return nil
}

func (client *ManageClient) getNodeFuncJsonSchema(callFunc interface{}) (string, error) {
	vType:=reflect.TypeOf(callFunc)
	schema := jsonschema.ReflectFromType(vType.In(0))
	byte, err := schema.MarshalJSON()
	return string(byte), err
}