// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/24 17:41
// Description:

package klee

import "github.com/liyiligang/klee-client-go/protoFiles/protoManage"

type NodeNotifyLevel int32
const (
	NodeNotifyLevelInfo    NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateUnknow)
	NodeNotifyLevelSuccess NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateNormal)
	NodeNotifyLevelWarn    NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateWarn)
	NodeNotifyLevelError   NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateError)
)

func (client *ManageClient) SendNodeNotify(msg string, nodeNotifyLevel NodeNotifyLevel, show bool) error {
	node, err := client.GetNode()
	if err != nil {
		return err
	}
	nodeNotify := &protoManage.NodeNotify{SenderID: node.Base.ID, SenderType: protoManage.NotifySenderType_NotifySenderTypeNode,
		Message: msg, State: protoManage.State(nodeNotifyLevel), ShowPop: show}
	return client.sendPB(protoManage.Order_NodeNotifyAdd, nodeNotify)
}

func (client *ManageClient) reqNodeNotify(message []byte) error {
	req := protoManage.NodeNotify{}
	err := req.Unmarshal(message)
	if err != nil {
		return err
	}
	if client.config.NotifyCall != nil {
		client.config.NotifyCall(req)
	}
	return nil
}