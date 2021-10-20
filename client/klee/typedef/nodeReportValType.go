// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/10/09 9:30
// Description:

package typedef

import "github.com/liyiligang/klee-client-go/protoFiles/protoManage"

type NodeReportValLevel int32
const (
	NodeReportValUnknown     NodeReportValLevel =   1
	NodeReportValLevelNormal NodeReportValLevel =   2
	NodeReportValLevelWarn   NodeReportValLevel =   3
	NodeReportValLevelError  NodeReportValLevel =   4
)

type NodeReportData struct {
	ValueList 		[]NodeReportVal
}

type NodeReportVal struct {
	Value			interface{}
	State   		protoManage.State
}