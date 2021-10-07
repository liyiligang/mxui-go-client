// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/06/18 10:37
// Description: manage client

package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/liyiligang/klee-client-go/klee"
	"github.com/liyiligang/klee-client-go/klee/typedef"
	"github.com/liyiligang/klee-client-go/protoFiles/protoManage"
	"math/rand"
	"time"
)

var manageClient *klee.ManageClient

type TestNode struct {
	Node 		int 		`jsonschema:"title=节点ID,default=1"`
	NodeName 	string		`jsonschema:"title=节点名,default=" jsonschema_extras:"ui:options='type':'textarea'.'rows':6"`
}

type TestUser struct {
	ID       int       		`jsonschema:"title=ID,default=1"`
	Name     string    		`json:"name,omitempty" jsonschema:"title=姓名,default="`
	Sex		 bool	   		`jsonschema:"title=性别,default=false" jsonschema_extras:"ui:options='activeText':'男'.'inactiveText':'女'"`
	Node	 TestNode
	Age		 []int     		`jsonschema:"title=年龄"`
	Date	 time.Time		`jsonschema:"title=日期,default=2020-01-16T02:11:11.000Z"`
	Color    string         `json:"fav_color,omitempty" jsonschema:"title=颜色,enum=red,enum=green,enum=blue,default=green" jsonschema_extras:"enumNames=红色,enumNames=绿色,enumNames=蓝色"`
}

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
	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "文本测试",
		CallFunc: testRectFunc1,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "Json测试",
		CallFunc: testRectFunc2,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "链接测试",
		CallFunc: testRectFunc3,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "媒体测试",
		CallFunc: testRectFunc4,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "文件测试",
		CallFunc: testRectFunc5,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "表格测试",
		CallFunc: testRectFunc6,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "图表测试",
		CallFunc: testRectFunc7,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "多返回值测试",
		CallFunc: testRectFunc8,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "错误测试",
		CallFunc: testRectFunc9,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "无值",
		CallFunc: testRectFunc10,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(klee.NodeFuncRegister{
		Name:     "图片测试",
		CallFunc: testRectFunc11,
		Level:    klee.NodeFuncLevelSuperManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	////node report
	//err = manageClient.RegisterNodeReport("报告测试1", testReport, 3*time.Second, klee.NodeReportLevelVisitor)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////node report manual update
	//err = manageClient.UpdateReportVal("testReport", 1, klee.NodeReportValLevelNormal)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	////node notify
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

func testRectFunc1(str *TestUser) string {
	return "aaaaaa"
}

func testRectFunc2(str *TestUser) (*TestUser, error) {
	str.Name = "<>"
	return str, errors.New("哈哈哈")
}

func testRectFunc3(str *typedef.NodeFuncReturnLink) *typedef.NodeFuncReturnLink {
	str.Name = "百度"
	str.Link="https://tool.lu/"
	return str
}

func testRectFunc4(str *typedef.NodeFuncReturnMedia) *typedef.NodeFuncReturnMedia {
	//https://10.0.2.54:9080/cloudAppFile/temp/2.flv
	//http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4
	//http://220.161.87.62:8800/hls/1/index.m3u8
	return str
}

func testRectFunc5(str *TestUser) typedef.NodeFuncReturnFile {

	//f, _ := os.OpenFile("C:\\Users\\49341\\Desktop\\kk.html", os.O_RDONLY,0600)
	//defer f.Close()
	//data, _ := ioutil.ReadAll(f)

	aa := typedef.NodeFuncReturnFile{Name: "文件测试", Data: []byte("测试文本"), AutoSave: str.Sex}
	return aa
}

func testRectFunc6(str *TestUser) (*typedef.NodeFuncReturnTable, error) {
	aa := typedef.NodeFuncReturnTable{
		Stripe: str.Sex,
		Border: str.Sex,
		ShowIndex: str.Sex,
		ShowSummary:false,
		SumText:"",
		Col: []typedef.NodeFuncReturnTableCol{
			typedef.NodeFuncReturnTableCol{
				Name: "编号",
				Width: 100,
				Type: "index",
			},
			typedef.NodeFuncReturnTableCol{
				Name: "姓名",
				Width: 200,
				Resizable:str.Sex,
				MergeSameCol:str.Sex,
			},
			typedef.NodeFuncReturnTableCol{
				Name: "年龄",
				Width: 100,
				Align: str.Name,
			},
			typedef.NodeFuncReturnTableCol{
				Name: "性别",
				Width: 100,
			},
		},
		Row: []typedef.NodeFuncReturnTableRow{
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"可莉", 10, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"七七", 1000, "女"},
				State: protoManage.State_StateUnknow,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 201, "女"},
				State: protoManage.State_StateNormal,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 202, "女"},
				State: protoManage.State_StateNormal,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 203, "女"},
				State: protoManage.State_StateNormal,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 203, "女"},
				State: protoManage.State_StateNormal,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 205, "女"},
				State: protoManage.State_StateNormal,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"原石", "原石", "原石"},
				State: protoManage.State_StateNormal,
				MergeSameRow:true,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"香菱", 15, "女"},
				State: protoManage.State_StateWarn,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"行秋", 300, "男"},
				State: protoManage.State_StateError,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"行秋", 300, "男"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"行秋", 300, "男"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"可莉", 10, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"七七", 1000, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"莫娜", 200, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"香菱", 15, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"香菱", 15, "女"},
				State: protoManage.State_StateError,
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"香菱", 15, "女"},
			},
			typedef.NodeFuncReturnTableRow{
				Data: []interface{}{"行秋", 300, "男"},
			},
		},
	}
	return &aa, nil
}

func testRectFunc7(str *TestUser) typedef.NodeFuncReturnCharts {
	return getEcharts()
}

type cc struct {
	Name string	`jsonschema:"title=名称,default="`
}

type dd struct {
	Name string
}

func testRectFunc8(str *TestUser) interface{} {
	if str.Sex {
		return nil
	}else {
		return errors.New("啦啦啦")
	}
}

func testRectFunc9(str *TestUser) error {
	if str.Sex {
		return nil
	}else {
		return errors.New("哇哦哇哦哇哦")
	}
}

func testRectFunc10() {
	return
}

func testRectFunc11(str *TestUser) typedef.NodeFuncReturnImage {
	return typedef.NodeFuncReturnImage{URL: "https://webstatic.mihoyo.com/ys/event/e20210901-fab/images/poster1.bccfc913.jpg",
	Fit: str.Name}
}

var testVal = 0.0
func testReport() (float64, klee.NodeReportValLevel) {
	testVal += 1
	return testVal, klee.NodeReportValLevelNormal
}

func generateBarItems() []opts.LineData {
	items := make([]opts.LineData, 0)
	for i := 0; i < 7; i++ {
		items = append(items, opts.LineData{Value: rand.Intn(300)})
	}
	return items
}

func getEcharts() typedef.NodeFuncReturnCharts {
	// create a new bar instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
		Title:    "测试",
		Subtitle: "111",
	}))

	// Put data into instance
	line.SetXAxis([]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}).
		AddSeries("Category A", generateBarItems()).
		AddSeries("Category B", generateBarItems())
	// Where the magic happens
	buf:= &bytes.Buffer{}
	line.Render(buf)
	dd := typedef.NodeFuncReturnCharts{Data: line.JSON()}
	return dd
}


