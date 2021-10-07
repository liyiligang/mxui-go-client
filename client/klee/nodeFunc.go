// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/23 15:52
// Description:

package klee

import (
	"context"
	"errors"
	"github.com/alecthomas/jsonschema"
	"github.com/liyiligang/klee-client-go/klee/typedef"
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

type NodeFuncRegister struct {
	Name 			string
	CallFunc 		interface{}
	Level 			NodeFuncLevel
	ReturnType      protoManage.NodeFuncReturnType
	BaseType		bool
	ErrorPos		int
}

func (client *ManageClient) RegisterNodeFunc(nodeFunc NodeFuncRegister) error {
	err := client.nodeFuncRegisterCheck(nodeFunc.CallFunc)
	if err != nil {
		return err
	}

	var schema string
	rType:=reflect.TypeOf(nodeFunc.CallFunc)
	if rType.NumIn() > 0 {
		schema, err = client.getNodeFuncJsonSchema(rType.In(0))
		if err != nil {
			return err
		}
	}

	if rType.NumOut() > 0 {
		nodeFunc.ReturnType, nodeFunc.BaseType = client.getNodeFuncReturnType(rType.Out(0))
	}
	nodeFunc.ErrorPos  = client.getNodeFuncReturnErrorPos(rType)

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
	client.data.nodeFuncMap.Store(resNodeFunc.Base.ID, &nodeFunc)
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
	//if vType.NumIn() != 1 {
	//	return errors.New("CallFunc 参数数量必须为1")
	//}
	if vType.NumIn() > 0 {
		if vType.In(0).Kind() == reflect.Ptr {
			if vType.In(0).Elem().Kind() != reflect.Struct {
				return errors.New("函数参数必须是 *reflect.Struct 类型")
			}
		}else if vType.In(0).Kind() != reflect.Struct {
			return errors.New("函数参数必须是 reflect.Struct 类型")
		}
	}

	//if vType.NumOut() != 1 {
	//	return errors.New("CallFunc 返回值数量必须为1")
	//}
	return nil
}

func (client *ManageClient) getNodeFuncJsonSchema(rType reflect.Type) (string, error) {
	schema := jsonschema.ReflectFromType(rType)
	byte, err := schema.MarshalJSON()
	return string(byte), err
}

func (client *ManageClient) getNodeFuncReturnType(rType reflect.Type) (protoManage.NodeFuncReturnType, bool){
	if rType == nil {
		return protoManage.NodeFuncReturnType_NotReturn, false
	}
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
	}
	if rType.String() == "error" {
		return protoManage.NodeFuncReturnType_NotReturn, false
	}
	switch rType {
	case reflect.TypeOf(typedef.NodeFuncReturnText{}):
		return protoManage.NodeFuncReturnType_Text, false
	case reflect.TypeOf(typedef.NodeFuncReturnJson{}):
		return protoManage.NodeFuncReturnType_Json, false
	case reflect.TypeOf(typedef.NodeFuncReturnLink{}):
		return protoManage.NodeFuncReturnType_Link, false
	case reflect.TypeOf(typedef.NodeFuncReturnImage{}):
		return protoManage.NodeFuncReturnType_Image, false
	case reflect.TypeOf(typedef.NodeFuncReturnMedia{}):
		return protoManage.NodeFuncReturnType_Media, false
	case reflect.TypeOf(typedef.NodeFuncReturnFile{}):
		return protoManage.NodeFuncReturnType_File, false
	case reflect.TypeOf(typedef.NodeFuncReturnCharts{}):
		return protoManage.NodeFuncReturnType_Charts, false
	case reflect.TypeOf(typedef.NodeFuncReturnTable{}):
		return protoManage.NodeFuncReturnType_Table, false
	case reflect.TypeOf(errors.New("")).Elem():
		return protoManage.NodeFuncReturnType_Error, false
	}
	switch rType.Kind() {
	case reflect.Bool:
	case reflect.Int:
	case reflect.Int8:
	case reflect.Int16:
	case reflect.Int32:
	case reflect.Int64:
	case reflect.Uint:
	case reflect.Uint8:
	case reflect.Uint16:
	case reflect.Uint32:
	case reflect.Uint64:
	case reflect.Float32:
	case reflect.Float64:
	case reflect.String:
		return protoManage.NodeFuncReturnType_Text, true
	case reflect.Array:
	case reflect.Map:
	case reflect.Slice:
	case reflect.Struct:
		return protoManage.NodeFuncReturnType_Json, true
	}
	return protoManage.NodeFuncReturnType_Unsure, false
}

func (client *ManageClient) getNodeFuncReturnErrorPos(rType reflect.Type) int {
	for i:=0; i<rType.NumOut(); i++ {
		res := rType.Out(i)
		if res != nil {
			if res.Kind() == reflect.Ptr {
				res = res.Elem()
			}
			if res.String() == "error" {
				return i
			}
		}
	}
	return -1
}