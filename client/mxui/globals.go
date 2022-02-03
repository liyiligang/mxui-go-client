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
	"github.com/liyiligang/base/component/Jrpc"
	"github.com/liyiligang/mxui-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxui-go-client/schema"
	"google.golang.org/grpc"
	"sync"
	"sync/atomic"
	"time"
)

type ClientConfig struct {
	Addr           string
	PublicKeyPath  string
	CertName       string
	NodeName  	   string
	ConnectTimeOut time.Duration
	RequestTimeOut time.Duration
	KeepaliveTime  time.Duration
	NotifyCall	   func(nodeNotify NodeNotify)
}

type Client struct {
	config    	ClientConfig
	data      	clientData
	conn      	*grpc.ClientConn
	engine    	protoManage.RpcEngineClient
	stream    	atomic.Value
	FuncSchema	*schema.Reflector
	keepalive 	*Jrpc.RpcKeepalive
}

type clientData struct {
	node 			  atomic.Value
	nodeLinkMap	  	  sync.Map
	nodeFuncMap	  	  sync.Map
	nodeReportMap	  sync.Map
}

