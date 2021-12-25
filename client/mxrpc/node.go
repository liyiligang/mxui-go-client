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
	"errors"
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
)

func (client *Client) getNode() (*protoManage.Node, error) {
	val := client.data.node.Load()
	if val == nil {
		return nil, errors.New("node data is not found")
	}
	node, ok := val.(protoManage.Node)
	if !ok {
		return nil, errors.New("node assert fail with protoManage.Node")
	}
	return &node, nil
}

func (client *Client) setNode(node protoManage.Node){
	client.data.node.Store(node)
}

func (client *Client) updateNodeState(nodeState constant.NodeState) error {
	node, err := client.getNode()
	if err != nil {
		return err
	}
	node.State = protoManage.State(nodeState)
	client.setNode(*node)
	return nil
}