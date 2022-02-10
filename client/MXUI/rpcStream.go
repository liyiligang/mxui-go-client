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

package MXUI

import (
	"errors"
	"github.com/golang/protobuf/proto"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"google.golang.org/grpc"
)

func (client *Client) initManageClientStream() error {
	rpcStream, err := Jrpc.GrpcStreamClientInit(new(protoManage.Message), Jrpc.RpcStreamCall{
		RpcStreamConnect:client.RpcStreamConnect,
		RpcStreamConnected:client.RpcStreamConnected,
		RpcStreamClosed:client.RpcStreamClosed,
		RpcStreamReceiver:client.RpcStreamReceiver,
		RpcStreamError:client.RpcStreamError,
	})
	if err != nil {
		return err
	}
	node, err := client.getNode()
	if err != nil {
		return err
	}
	bytes, err := node.Marshal()
	if err != nil {
		return err
	}
	channel, err := client.engine.RpcChannel(Jrpc.SetRpcStreamClientHeader(bytes), grpc.WaitForReady(true))
	if err != nil {
		return err
	}
	return rpcStream.GrpcStreamClientRun(channel)
}

func (client *Client) closeStream() {
	rpcStream, err := client.getRpcStream()
	if err != nil {
		client.RpcStreamError("rpc stream close error", err)
		return
	}
	rpcStream.Close(false)
}

func (client *Client) sendPB(order protoManage.Order, pb proto.Message) error {
	pbByte, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return client.send(order, pbByte)
}

func (client *Client) send(order protoManage.Order, data []byte) error {
	rpcStream, err := client.getRpcStream()
	if err != nil {
		return err
	}
	rpcStream.SendData(&protoManage.Message{Order: order, Message: data})
	return nil
}

func (client *Client) setRpcStream(stream *Jrpc.RpcStream){
	client.stream.Store(stream)
}

func (client *Client) getRpcStream() (*Jrpc.RpcStream, error){
	val := client.stream.Load()
	if val == nil {
		return nil, errors.New("rpc stream is not found")
	}
	stream, ok := val.(*Jrpc.RpcStream)
	if !ok {
		return nil, errors.New("rpc stream assert fail with *Jrpc.RpcStream")
	}
	if stream == nil {
		return nil, errors.New("rpc stream must not be nil")
	}
	return stream, nil
}
