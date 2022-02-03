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

package main

import (
	"fmt"
	"github.com/liyiligang/mxui-go-client/example"
	"github.com/liyiligang/mxui-go-client/mxui"
	"time"
)

func main() {
	client, err := initClient()
	if err != nil {
		fmt.Println("server connect error: ", err)
		return
	}
	LoadExample(client)
	select {}
}

func initClient() (*mxui.Client, error) {
	c, err := mxui.InitManageClient(mxui.ClientConfig{
		Addr:":888",
		PublicKeyPath:"../store/cert/grpc/ca_cert.pem",
		CertName: "x.test.example.com",
		NodeName: "例子(example)",
		ConnectTimeOut: time.Second * 5,
		RequestTimeOut: time.Second * 5,
		KeepaliveTime: time.Second * 1,
		NotifyCall: func (nodeNotify mxui.NodeNotify){
			fmt.Println("receive node notify: ", nodeNotify.Message)
		},
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func LoadExample(client *mxui.Client) {
	example.LoadExampleMethod(client)
	example.LoadExampleReport(client)
	example.LoadExampleNotify(client)
}