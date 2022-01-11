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
	"context"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
	"google.golang.org/grpc"
)

func InitManageClient(config ClientConfig) (*Client, error) {
	var err error
	client := &Client{config: config}
	client.conn, err = Jrpc.GrpcClientInit(Jrpc.RpcClientConfig{
		RpcBaseConfig:  Jrpc.RpcBaseConfig{
			Addr: config.Addr,
			PublicKeyPath: config.PublicKeyPath,
		},
		CertName:       config.CertName,
		ConnectTimeOut: config.ConnectTimeOut,
		ClientOption:  []grpc.DialOption{grpc.WithBlock(), grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(constant.ConstRpcClientMaxMsgSize), grpc.MaxCallRecvMsgSize(constant.ConstRpcClientMaxMsgSize)),
			grpc.WithUnaryInterceptor(client.rpcUnaryInterceptor), grpc.WithStreamInterceptor(client.rpcStreamInterceptor),
		},
	})
	if err != nil {
		return nil, err
	}
	client.engine = protoManage.NewRpcEngineClient(client.conn)
	protoNode, err := client.engine.RegisterNode(context.Background(), &protoManage.Node{Name:client.config.NodeName})
	if err != nil {
		return nil, err
	}
	client.setNode(*protoNode)

	err = client.initManageClientStream()
	if err != nil {
		return nil, err
	}
	client.keepalive = &Jrpc.RpcKeepalive{
		Conn: client.conn,
		KeepaliveTime: config.KeepaliveTime,
	}
	err = Jrpc.RegisterRpcKeepalive(client.keepalive, client)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (client *Client) Close() {
	client.keepalive.Close()
}

func (client *Client) rpcUnaryInterceptor (ctx context.Context, method string, req, reply interface{},
cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error{
	ctxTimeOut, _ := context.WithTimeout(ctx, client.config.RequestTimeOut)
	return invoker(ctxTimeOut, method, req, reply, cc, opts...)
}

func (client *Client) rpcStreamInterceptor (ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn,
	method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
	return streamer(ctx, desc, cc, method, opts...)
}

func (client *Client) closeConn() {
	err := client.conn.Close()
	if err != nil {
		client.RpcStreamError("rpc close error", err)
	}
}
