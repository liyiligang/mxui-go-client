// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/21 10:32
// Description: rpc client call

package klee

import (
	"errors"
	"github.com/liyiligang/base/commonConst"
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
)

func (client *ManageClient) RpcServeConnected(rpcKeepalive *Jrpc.RpcKeepalive, isReConnect bool) {
	if isReConnect {
		err := client.initManageClientStream()
		if err != nil {
			client.RpcStreamError("rpc stream init error", err)
		}
	}
}

func (client *ManageClient) RpcServeDisconnected(rpcKeepalive *Jrpc.RpcKeepalive, isCloseByUser bool) {
	defer func(){
		if isCloseByUser {
			client.closeConn()
		}
	}()
}

func (client *ManageClient) RpcStreamConnect(stream *Jrpc.RpcStream) (interface{}, error) {
	pbByte, err := client.getNodeStreamByte()
	if err != nil {
		return 0, err
	}
	stream.WriteRpcStreamClientHeader(pbByte)
	return commonConst.ManageNodeID, nil
}

func (client *ManageClient) RpcStreamConnected(stream *Jrpc.RpcStream) error {
	client.setRpcStream(stream)
	err := client.nodeOnline(stream.GetRpcContext().RpcStreamServerHeader)
	if err != nil{
		return err
	}
	return nil
}

func (client *ManageClient) RpcStreamClosed(stream *Jrpc.RpcStream) {
	client.setRpcStream(nil)
}

func (client *ManageClient) RpcStreamReceiver(stream *Jrpc.RpcStream, recv interface{}) {
	var err error
	defer func(){
		if err != nil && client.config.NotifyCall != nil {
			client.config.NotifyCall(protoManage.NodeNotify{
				SenderType:           protoManage.NotifySenderType_NotifySenderTypeNode,
				Message:              err.Error(),
				State:                protoManage.State_StateError,
			})
		}
	}()
	res, ok := recv.(*protoManage.Message)
	if !ok {
		err = errors.New("rpc stream消息断言错误")
		return
	}
	switch res.Order {
	case protoManage.Order_NodeFuncCallReq:
		err = client.reqNodeFuncCall(res.Message)
		break
	case protoManage.Order_NodeNotifyError:
		err = client.reqNodeNotify(res.Message)
		break
	default:
		err = errors.New("rpc stream指令错误：" +  Jtool.Int64ToString(int64(res.Order)))
	}
}

func (client *ManageClient) RpcStreamError(text string, err error) {
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



