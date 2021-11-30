// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/24 14:30
// Description:

package klee

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/liyiligang/klee-client-go/klee/typedef"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"time"
)

type NodeReportRegister struct {
	Name 					string
	Type					protoManage.NodeReportType
	CallFunc 				NodeReportCallFunc
	CallInterval 			time.Duration
	Schema					NodeReportSchema
	Level 					NodeReportLevel
}

type NodeReportSchema struct {
	CategoryList			[]NodeReportCategory
}

type NodeReportCategory struct {
	Name			string
	Width			uint32
}

type NodeReportCallFunc func() (*typedef.NodeReportData, error)

type nodeReportMapVal struct {
	nodeReportID			int64
	cancel 					context.CancelFunc
}

type NodeReportLevel int32
const (
	NodeReportLevelVisitor      NodeReportLevel =   1
	NodeReportLevelMember       NodeReportLevel =   2
	NodeReportLevelManager      NodeReportLevel =   3
	NodeReportLevelSuperManager NodeReportLevel =   4
)

func (client *ManageClient) RegisterNodeReport(nodeReport NodeReportRegister) error {
	node, err := client.GetNode()
	if err != nil {
		return err
	}
	if nodeReport.CallFunc == nil {
		return errors.New("nodeReport callFunc is nil")
	}
	schema, err := client.getNodeReportSchemaJson(&nodeReport.Schema)
	if err != nil {
		return err
	}
	callName := client.getFuncName(nodeReport.CallFunc)
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

func (client *ManageClient) getNodeReportSchemaJson(schema *NodeReportSchema) (string, error){
	if schema == nil {
		return "", errors.New("schema is nil")
	}
	data, err := json.Marshal(schema)
	return string(data), err
}

func (client *ManageClient) startTicker(interval time.Duration, nodeReport *protoManage.NodeReport, callFunc NodeReportCallFunc) context.CancelFunc {
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

func (client *ManageClient) stopTicker(nodeReportName string){
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


