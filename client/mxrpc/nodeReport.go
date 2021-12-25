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

package mxrpc

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
	"time"
)

type NodeReportCallFunc func() (*NodeReportData, error)

type NodeReportRegister struct {
	Name 					string
	Type					protoManage.NodeReportType
	CallFunc 				NodeReportCallFunc
	CallInterval 			time.Duration
	Schema					NodeReportSchema
	Level 					constant.UserLevel
}

type NodeReportSchema struct {
	CategoryList			[]NodeReportCategory
}

type NodeReportCategory struct {
	Name			string
	Width			uint32
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
	if nodeReport.CallFunc == nil {
		return errors.New("node report register function must not be nil")
	}
	schema, err := client.getNodeReportSchemaJson(&nodeReport.Schema)
	if err != nil {
		return err
	}
	callName := Jtool.GetFuncName(nodeReport.CallFunc)
	protoNodeReport := protoManage.NodeReport{NodeID: node.Base.ID, Name: nodeReport.Name,
		Func: callName, Schema: schema, Type: nodeReport.Type, Level:protoManage.Level(nodeReport.Level),
		Interval:nodeReport.CallInterval.Milliseconds()}
	ctx, _ := context.WithTimeout(context.Background(), client.config.RequestTimeOut)
	resNodeReport, err := client.engine.RegisterNodeReport(ctx, &protoNodeReport)
	if err != nil {
		return err
	}
	client.stopTicker(resNodeReport.Name)
	var cancel context.CancelFunc
	if nodeReport.CallInterval > 0{
		cancel = client.startTicker(nodeReport.CallInterval, resNodeReport, nodeReport.CallFunc)
	}
	client.data.nodeReportMap.Store(resNodeReport.Name, nodeReportMapVal{
		nodeReportID: resNodeReport.Base.ID, cancel: cancel})
	return nil
}

func (client *Client) getNodeReportSchemaJson(schema *NodeReportSchema) (string, error){
	if schema == nil {
		return "", errors.New("node report schema must not be nil")
	}
	data, err := json.Marshal(schema)
	return string(data), err
}

func (client *Client) startTicker(interval time.Duration, nodeReport *protoManage.NodeReport, callFunc NodeReportCallFunc) context.CancelFunc {
	ticker := time.NewTicker(interval)
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				err := client.execCallReport(nodeReport, callFunc)
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


