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
	"encoding/json"
	"errors"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"reflect"
)

type NodeReportData struct {
	ValueList 		[]NodeReportVal
}

type NodeReportVal struct {
	Value			interface{}
	State   		protoManage.State
}

func (client *Client) execCallReport(report *nodeReportTicker) error{
	nodeReportData, err := client.callReportByReflect(report)
	if err != nil {
		return err
	}
	val, err := client.getNodeReportDataJson(nodeReportData)
	if err != nil {
		return err
	}
	nodeReportVal := &protoManage.NodeReportVal{ReportID: report.nodeReport.Base.ID, Value: val}
	return client.sendPB(protoManage.Order_NodeReportUpdateVal, nodeReportVal)
}

func (client *Client) UpdateReportVal(name string, nodeReportData interface{}) error{
	v, ok := client.data.nodeReportMap.Load(name)
	if !ok {
		return errors.New("node report name is not found")
	}
	nodeReport, ok := v.(nodeReportMapVal)
	if !ok {
		return errors.New("value assert fail with nodeReportMapVal")
	}
	val, err := client.getNodeReportDataJson(nodeReportData)
	if err != nil {
		return err
	}
	nodeReportVal := &protoManage.NodeReportVal{ReportID: nodeReport.nodeReportID, Value: val}
	return client.sendPB(protoManage.Order_NodeReportUpdateVal, nodeReportVal)
}

func (client *Client) callReportByReflect(report *nodeReportTicker) (interface{}, error) {
	var res []reflect.Value
	res = report.callFuncValue.Call(report.callFuncPara)
	var err error
	ok := false
	if res[1].Interface() != nil {
		err, ok = res[1].Interface().(error)
		if !ok {
			return nil, errors.New("value assert fail with error")
		}
	}
	return res[0].Interface(), err
}

func (client *Client) getNodeReportDataJson(nodeReportData interface{}) (string, error){
	if nodeReportData == nil {
		return "", errors.New("node report data must not be nil")
	}
	data, err := json.Marshal(nodeReportData)
	return string(data), err
}