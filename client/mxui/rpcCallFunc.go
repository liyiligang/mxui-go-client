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
	"errors"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxui-go-client/typedef/constant"
)

func (client *Client) RpcServeConnected(rpcKeepalive *Jrpc.RpcKeepalive, isReConnect bool) {
	if isReConnect {
		err := client.initManageClientStream()
		if err != nil {
			client.RpcStreamError("rpc stream reconnect error", err)
		}
	}
	err := client.updateNodeState(constant.NodeStateNormal)
	if err != nil {
		client.RpcStreamError("update node state error", err)
	}
}

func (client *Client) RpcServeDisconnected(rpcKeepalive *Jrpc.RpcKeepalive, isCloseByUser bool) {
	if isCloseByUser {
		client.closeConn()
	}
	err := client.updateNodeState(constant.NodeStateClose)
	if err != nil {
		client.RpcStreamError("update node state error", err)
	}
}

func (client *Client) RpcStreamConnect(stream *Jrpc.RpcStream) (interface{}, error) {
	return constant.ConstManageNodeID, nil
}

func (client *Client) RpcStreamConnected(stream *Jrpc.RpcStream) error {
	client.setRpcStream(stream)
	return nil
}

func (client *Client) RpcStreamClosed(stream *Jrpc.RpcStream) error {
	client.setRpcStream(nil)
	return nil
}

func (client *Client) RpcStreamReceiver(stream *Jrpc.RpcStream, recv interface{}) error {
	res, ok := recv.(*protoManage.Message)
	if !ok {
		return errors.New("recv assert fail with *protoManage.Message")
	}
	var err error
	switch res.Order {
	case protoManage.Order_NodeFuncCallReq:
		err = client.reqNodeFuncCall(res.Message)
		break
	case protoManage.Order_NodeNotifyError:
		err = client.reqNodeNotify(res.Message)
		break
	default:
		err = errors.New("rpc order is invalid with number " +  Jtool.Int64ToString(int64(res.Order)))
	}
	return err
}

func (client *Client) RpcStreamError(text string, err error) {
	msg := text + ": "
	if err != nil {
		msg += err.Error()
	}
	if client.config.NotifyCall != nil {
		client.config.NotifyCall(protoManage.NodeNotify{
			SenderType:           protoManage.NotifySenderType_NotifySenderTypeNode,
			Message:              msg,
			State:                protoManage.State_StateError,
		})
	}
}



