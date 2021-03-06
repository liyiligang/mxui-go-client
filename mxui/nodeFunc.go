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

package mxui

import (
	"context"
	"errors"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"reflect"
)

type NodeFuncRegister struct {
	Name 			string
	CallFunc 		interface{}
	Level 			UserLevel
	returnType      protoManage.NodeFuncReturnType
	baseType		bool
	errorPos		int
}

func (client *Client) RegisterNodeFunc(nodeFunc NodeFuncRegister) error {
	err := client.nodeFuncCallRegisterCheck(nodeFunc.CallFunc)
	if err != nil {
		return err
	}
	if nodeFunc.Level < UserLevelLevelVisitor || nodeFunc.Level > UserLevelLevelSuperManager {
		nodeFunc.Level = UserLevelLevelVisitor
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
		nodeFunc.returnType, nodeFunc.baseType = client.getNodeFuncReturnType(rType.Out(0))
	}
	nodeFunc.errorPos  = client.getNodeFuncReturnErrorPos(rType)

	node, err := client.getNode()
	if err != nil {
		return err
	}
	callName := Jtool.GetFuncName(nodeFunc.CallFunc)
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


func (client *Client) nodeFuncCallRegisterCheck(callFunc interface{}) error {
	if callFunc == nil {
		return errors.New("register function must not be nil")
	}
	vType:=reflect.TypeOf(callFunc)
	if vType.Kind() != reflect.Func {
		return errors.New("register function kind must be reflect.Func")
	}
	if vType.NumIn() > 1 {
		return errors.New("register function parameters cannot be greater than 1")
	}
	if vType.NumIn() > 0 {
		if vType.In(0).Kind() == reflect.Ptr {
			if vType.In(0).Elem().Kind() != reflect.Struct {
				return errors.New("register function parameters kind must be reflect.Struct or *reflect.Struct")
			}
		}else if vType.In(0).Kind() != reflect.Struct {
			return errors.New("register function parameters kind must be reflect.Struct or *reflect.Struct")
		}
	}
	return nil
}

func (client *Client) getNodeFuncJsonSchema(rType reflect.Type) (string, error) {
	funcSchema := client.FuncSchema.ReflectFromType(rType)
	bytes, err := funcSchema.MarshalJSON()
	return string(bytes), err
}

func (client *Client) getNodeFuncReturnType(rType reflect.Type) (protoManage.NodeFuncReturnType, bool){
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
	case reflect.TypeOf(NodeFuncReturnText{}):
		return protoManage.NodeFuncReturnType_Text, false
	case reflect.TypeOf(NodeFuncReturnJson{}):
		return protoManage.NodeFuncReturnType_Json, false
	case reflect.TypeOf(NodeFuncReturnLink{}):
		return protoManage.NodeFuncReturnType_Link, false
	case reflect.TypeOf(NodeFuncReturnImage{}):
		return protoManage.NodeFuncReturnType_Image, false
	case reflect.TypeOf(NodeFuncReturnMedia{}):
		return protoManage.NodeFuncReturnType_Media, false
	case reflect.TypeOf(NodeFuncReturnFile{}):
		return protoManage.NodeFuncReturnType_File, false
	case reflect.TypeOf(NodeFuncReturnCharts{}):
		return protoManage.NodeFuncReturnType_Charts, false
	case reflect.TypeOf(NodeFuncReturnTable{}):
		return protoManage.NodeFuncReturnType_Table, false
	case reflect.TypeOf(errors.New("")).Elem():
		return protoManage.NodeFuncReturnType_Error, false
	}
	switch rType.Kind() {
	case reflect.Bool,
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Float32,
	reflect.Float64,
	reflect.String:
		return protoManage.NodeFuncReturnType_Text, true
	case reflect.Array,
	reflect.Map,
	reflect.Slice,
	reflect.Struct:
		return protoManage.NodeFuncReturnType_Json, true
	}
	return protoManage.NodeFuncReturnType_Unknown, false
}

func (client *Client) getNodeFuncReturnErrorPos(rType reflect.Type) int {
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