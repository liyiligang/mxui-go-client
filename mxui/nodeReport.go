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
	"encoding/json"
	"errors"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxui-go-client/schema"
	"reflect"
	"time"
)

type NodeReportRegister struct {
	Name 					string
	Type					NodeReportType
	CallFunc 				interface{}
	CallInterval 			time.Duration
	Level 					UserLevel
}

type NodeReportValue struct {
	Data			interface{}
	State			DataState
}

type nodeReportSchema struct {
	CategoryList			[]schema.TableSchema
}

type nodeReportTicker struct {
	interval 			time.Duration
	nodeReport 			*protoManage.NodeReport
	callFunc 			interface{}
	callFuncValue		reflect.Value
	callFuncPara 		[]reflect.Value
}

type nodeReportMapVal struct {
	nodeReportID			int64
	cancel 					context.CancelFunc
}

func (client *Client) RegisterNodeReport(nodeReport NodeReportRegister) error {
	node, err := client.getNode()
	if err != nil {
		return err
	}
	err = client.nodeReportCallRegisterCheck(nodeReport.CallFunc)
	if err != nil {
		return err
	}
	if nodeReport.Level < UserLevelLevelVisitor || nodeReport.Level > UserLevelLevelSuperManager {
		nodeReport.Level = UserLevelLevelVisitor
	}
	var nodeReportSchema string
	rType:=reflect.TypeOf(nodeReport.CallFunc)
	if rType.NumOut() > 0 {
		nodeReportSchema, err = client.getNodeReportSchemaJson(rType.Out(0))
		if err != nil {
			return err
		}
	}

	callName := Jtool.GetFuncName(nodeReport.CallFunc)
	protoNodeReport := protoManage.NodeReport{NodeID: node.Base.ID, Name: nodeReport.Name,
		Func: callName, Schema: nodeReportSchema, Type: protoManage.NodeReportType(nodeReport.Type),
		Level:protoManage.Level(nodeReport.Level), Interval:nodeReport.CallInterval.Milliseconds()}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	resNodeReport, err := client.engine.RegisterNodeReport(ctx, &protoNodeReport)
	if err != nil {
		return err
	}
	client.stopTicker(resNodeReport.Name)
	var cancel context.CancelFunc
	if nodeReport.CallInterval > 0{
		cancel = client.startTicker(&nodeReportTicker{
			interval: nodeReport.CallInterval,
			nodeReport: resNodeReport,
			callFunc: nodeReport.CallFunc,
			callFuncValue: reflect.ValueOf(nodeReport.CallFunc),
			callFuncPara: make([]reflect.Value, 0)})
	}
	client.data.nodeReportMap.Store(resNodeReport.Name, nodeReportMapVal{
		nodeReportID: resNodeReport.Base.ID, cancel: cancel})
	return nil
}

func (client *Client) getNodeReportSchemaJson(rType reflect.Type) (string, error){
	reportSchema := nodeReportSchema{}
	if rType.Kind() == reflect.Ptr {
		rType = rType.Elem()
	}
	for i :=0; i< rType.NumField(); i++ {
		tableSchema := schema.ParseTagWithTableSchema(rType.Field(i))
		reportSchema.CategoryList = append(reportSchema.CategoryList, tableSchema)
	}
	data, err := json.Marshal(reportSchema)
	return string(data), err
}

func (client *Client) nodeReportCallRegisterCheck(callFunc interface{}) error {
	if callFunc == nil {
		return errors.New("register function must not be nil")
	}
	vType:=reflect.TypeOf(callFunc)
	if vType.Kind() != reflect.Func {
		return errors.New("register function kind must be reflect.Func")
	}
	if vType.NumIn() != 0 {
		return errors.New("register function number of parameters must be 0")
	}
	if vType.NumOut() != 2 {
		return errors.New("register function number of return value must be 2")
	}
	if vType.Out(0).Kind() == reflect.Ptr {
		if vType.Out(0).Elem().Kind() != reflect.Struct {
			return errors.New("register function first return value kind must be reflect.Struct or *reflect.Struct")
		}
	}else if vType.Out(0).Kind() != reflect.Struct {
		return errors.New("register function first return value kind must be reflect.Struct or *reflect.Struct")
	}
	if vType.Out(1).String() != "error" {
		return errors.New("register function second return value kind must be error")
	}
	return nil
}

func (client *Client) startTicker(report *nodeReportTicker) context.CancelFunc {
	ticker := time.NewTicker(report.interval)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := client.execCallReport(report)
				if err != nil {
					client.RpcStreamError("node report call error", err)
				}
			case <-ctx.Done():
				return
			}
		}
	}()
	return cancel
}

func (client *Client) stopTicker(nodeReportName string){
	v, ok := client.data.nodeReportMap.Load(nodeReportName)
	if ok {
		val, ok := v.(nodeReportMapVal)
		if ok {
			if val.cancel != nil {
				val.cancel()
			}
		}
	}
}


