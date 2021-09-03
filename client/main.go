// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/18 10:37
// Description: manage client

package main

import (
	"fmt"
	"github.com/liyiligang/klee-client-go/klee"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"time"
)

var manageClient *klee.ManageClient

func main() {
	//example
	//link
	var err error
	manageClient, err = initClient()
	if err !=nil {
		fmt.Println("link error: ", err)
		return
	}

	//node link
	err = manageClient.UpdateNodeLink(15, klee.NodeLinkStateConnected)
	if err != nil {
		fmt.Println(err)
	}

	//node func
	//err = manageClient.RegisterNodeFunc("testFunc", testFunc, klee.NodeFuncLevelVisitor)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	err = manageClient.RegisterNodeFunc("方法测试5", testFunc, klee.NodeFuncLevelSuperManager)
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc("方法测试6", testFunc, klee.NodeFuncLevelVisitor)
	if err != nil {
		fmt.Println(err)
	}

	//node report
	err = manageClient.RegisterNodeReport("报告测试1", testReport, 3*time.Second, klee.NodeReportLevelVisitor)
	if err != nil {
		fmt.Println(err)
	}

	//node report manual update
	err = manageClient.UpdateReportVal("testReport", 1, klee.NodeReportValLevelNormal)
	if err != nil {
		fmt.Println(err)
	}

	//node notify
	//err = manageClient.SendNodeNotify("testNotify", klee.NodeNotifyLevelWarn)
	//if err != nil {
	//	fmt.Println(err)
	//}

	select {}
}


func initClient() (*klee.ManageClient, error) {
	c, err := klee.InitManageClient(klee.ManageClientConfig{
		Addr:":888",
		PublicKeyPath:"../store/cert/grpc/ca_cert.pem",
		CertName: "x.test.example.com",
		NodeGroupName: "TestGroup",
		NodeTypeName: "TestType",
		NodeName: "TestNode",
		ConnectTimeOut: time.Second * 5,
		RequestTimeOut: time.Second * 5,
		KeepaliveTime: time.Second * 1,
		NotifyCall: notifyCall,
	})
	if err != nil {
		return nil, err
	}
	return c, nil
}

func notifyCall(nodeNotify protoManage.NodeNotify) {
	fmt.Println("receive node notify: ", nodeNotify.Message)
}

var testFuncVal = 2
func testFunc(str string) (string, klee.NodeFuncCallLevel) {
	testFuncVal++
	if testFuncVal >= 5 {
		testFuncVal = 2
	}
	manageClient.UpdateNode(klee.NodeState(testFuncVal))
	manageClient.UpdateNodeLink(15, klee.NodeLinkState(testFuncVal))
	return "567890", klee.NodeFuncCallLevelLevelSuccess
}


var testVal = 0.0
func testReport() (float64, klee.NodeReportValLevel) {
	testVal += 1
	return testVal, klee.NodeReportValLevelNormal
}