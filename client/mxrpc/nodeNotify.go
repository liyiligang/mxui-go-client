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
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
)

func (client *Client) SendNodeNotify(msg string, nodeNotifyLevel constant.NodeNotifyLevel, show bool) error {
	node, err := client.getNode()
	if err != nil {
		return err
	}
	nodeNotify := &protoManage.NodeNotify{SenderID: node.Base.ID, SenderType: protoManage.NotifySenderType_NotifySenderTypeNode,
		Message: msg, State: protoManage.State(nodeNotifyLevel), ShowPop: show}
	return client.sendPB(protoManage.Order_NodeNotifyAdd, nodeNotify)
}

func (client *Client) reqNodeNotify(message []byte) error {
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