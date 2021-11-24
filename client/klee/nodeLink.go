// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/23 9:45
// Description:

package klee

import (
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
)

type NodeLinkState int32
const (
	NodeLinkStateConnected    NodeLinkState =  NodeLinkState(protoManage.State_StateNormal)
	NodeLinkStateConnecting   NodeLinkState =  NodeLinkState(protoManage.State_StateWarn)
	NodeLinkStateDisconnected NodeLinkState =  NodeLinkState(protoManage.State_StateError)
)


