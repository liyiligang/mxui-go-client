// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/25 17:18
// Description:

package klee

import (
	"encoding/json"
	"errors"
	"github.com/liyiligang/klee-client-go/klee/typedef"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
)

func (client *ManageClient) execCallReport(nodeReport *protoManage.NodeReport, callFunc NodeReportCallFunc) error{
	nodeReportData, err := callFunc()
	if err != nil {
		return err
	}
	val, err := client.getNodeReportDataJson(nodeReportData)
	if err != nil {
		return err
	}
	nodeReportVal := &protoManage.NodeReportVal{ReportID: nodeReport.Base.ID, Value: val}
	return client.sendPB(protoManage.Order_NodeReportUpdateVal, nodeReportVal)
}

func (client *ManageClient) UpdateReportVal(name string, nodeReportData *typedef.NodeReportData) error{
	v, ok := client.data.nodeReportMap.Load(name)
	if !ok {
		return errors.New("nodeReport name is non-existent")
	}
	nodeReport, ok := v.(nodeReportMapVal)
	if !ok {
		return errors.New("val data format is error, its type should be nodeReportMapVal")
	}
	val, err := client.getNodeReportDataJson(nodeReportData)
	if err != nil {
		return err
	}
	nodeReportVal := &protoManage.NodeReportVal{ReportID: nodeReport.nodeReportID, Value: val}
	return client.sendPB(protoManage.Order_NodeReportUpdateVal, nodeReportVal)
}

func (client *ManageClient) getNodeReportDataJson(nodeReportData *typedef.NodeReportData) (string, error){
	if nodeReportData == nil {
		return "", errors.New("nodeReportData is nil")
	}
	data, err := json.Marshal(nodeReportData)
	return string(data), err
}