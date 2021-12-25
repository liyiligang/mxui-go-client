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
	"bytes"
	"errors"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxrpc-go-client/mxrpc"
	"github.com/liyiligang/mxrpc-go-client/protoFiles/protoManage"
	"github.com/liyiligang/mxrpc-go-client/typedef"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
	"math/rand"
	"time"
)

var manageClient *mxrpc.Client

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
	ColorM   []string       `json:"fav_colorM,omitempty" jsonschema:"title=颜色M" jsonschema_extras:"ui:widget=CheckboxesWidget,uniqueItems=1"`
}

type TestJsonExtras struct {
	Color []string		`schema:"title=颜色,uniqueItems=true,enum=red;green;blue,enumNames=红色;绿色;蓝色,default=green;red" ui:"{\"ui:widget\":\"CheckboxesWidget\"}"`
	Color1 	string		`schema:"title=颜色,enum=red;green;blue,enumNames=红色;绿色;蓝色,default=green"`
	Text 	string		`schema:"title=节点名,default=" ui:"{\"ui:options\":{\"type\":\"textarea\",\"rows\":6}}"`
}

type FileUpload struct {
	Name     	 string    		`json:"name,omitempty" schema:"title=姓名,default="`
	File     	 string    		`json:"file,omitempty" schema:"title=文件上传" ui:"{\"ui:widget\":\"UploadFile\",\"ui:btnText\":\"上传文件\"}"`
	FileList     []string    	`schema:"title=文件批量上传" ui:"{\"ui:widget\":\"UploadFile\",\"ui:btnText\":\"批量上传文件\"}"`
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

	//node func
	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "文本测试",
		CallFunc: testRectFunc1,
		Level:    constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "Json测试",
		CallFunc: testRectFunc2,
		Level:    constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "链接测试",
		CallFunc: testRectFunc3,
		Level:    constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "媒体测试",
		CallFunc: testRectFunc4,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "文件测试",
		CallFunc: testRectFunc5,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "表格测试",
		CallFunc: testRectFunc6,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "图表测试",
		CallFunc: testRectFunc7,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "动态返回值测试",
		CallFunc: testRectFunc8,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "错误测试",
		CallFunc: testRectFunc9,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "无值",
		CallFunc: testRectFunc10,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "图片测试",
		CallFunc: testRectFunc11,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "多参数多返回值测试",
		CallFunc: testRectFunc12,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "文件上传",
		CallFunc: testRectFunc13,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "通用json测试",
		CallFunc: testRectFunc14,
		Level:   constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	//node report
	err = manageClient.RegisterNodeReport(mxrpc.NodeReportRegister{
		Name: "表格报告",
		Type: protoManage.NodeReportType_NodeReportTypeTable,
		CallFunc: testReport1,
		CallInterval:time.Second*2,
		Schema: mxrpc.NodeReportSchema{
			CategoryList:[]mxrpc.NodeReportCategory{
				mxrpc.NodeReportCategory{
					Name: "阳光城",
					Width: 100,
				},
				mxrpc.NodeReportCategory{
					Name: "麓谷",
					Width: 100,
				},
				mxrpc.NodeReportCategory{
					Name: "梅溪湖",
					Width: 100,
				},
			},
		},
		Level: constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	err = manageClient.RegisterNodeReport(mxrpc.NodeReportRegister{
		Name: "折线报告",
		Type: protoManage.NodeReportType_NodeReportTypeLine,
		CallFunc: testReport2,
		CallInterval:time.Second*2,
		Schema: mxrpc.NodeReportSchema{
			CategoryList:[]mxrpc.NodeReportCategory{
				mxrpc.NodeReportCategory{
					Name: "阳光城",
					Width: 100,
				},
				mxrpc.NodeReportCategory{
					Name: "麓谷",
					Width: 100,
				},
				mxrpc.NodeReportCategory{
					Name: "梅溪湖",
					Width: 100,
				},
			},
		},
		Level: constant.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}

	////node report manual update
	//err = manageClient.UpdateReportVal("testReport", 1, klee.NodeReportValLevelNormal)
	//if err != nil {
	//	fmt.Println(err)
	//}
	//
	//node notify
	testNotify()

	select {}
}


func initClient() (*mxrpc.Client, error) {
	c, err := mxrpc.InitManageClient(mxrpc.ClientConfig{
		Addr:":888",
		PublicKeyPath:"../store/cert/grpc/ca_cert.pem",
		CertName: "x.test.example.com",
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
	return time.Now().String()
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

	aa := typedef.NodeFuncReturnFile{Name: "文件测试", URL: "测试文本", AutoSave: str.Sex}
	return aa
}

func testRectFunc6(str *TestUser) (*typedef.NodeFuncReturnTable, error) {
	aa := typedef.NodeFuncReturnTable{
		//Stripe: str.Sex,
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
				//MergeSameCol:str.Sex,
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
				//MergeSameRow:true,
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

func testRectFunc12(str *TestUser) (int, int, string, bool, error) {
	if str.Sex {
		return 1 , 1, "a", false, nil
	}else {
		return 2 , 2, "a", false, errors.New("欸嘿")
	}
}

func testRectFunc13(file *FileUpload) string {
	return file.File
}

func testRectFunc14(str *TestJsonExtras) string {
	return "str.Color"
}

var testVal1 = 0.0
func testReport1() (*mxrpc.NodeReportData, error) {
	testVal1 += 1
	return &mxrpc.NodeReportData{ValueList:[]mxrpc.NodeReportVal{
		mxrpc.NodeReportVal{
			Value: testVal1+0.12,
		},
		mxrpc.NodeReportVal{
			Value: testVal1+1,
			State: protoManage.State_StateError,
		},
		mxrpc.NodeReportVal{
			Value: testVal1*2,
			State: protoManage.State_StateNormal,
		},
	}}, nil
}

var testVal2 = 5000
var testVal22 = 0
var testVal33 = 100
func testReport2() (*mxrpc.NodeReportData, error) {
	r1, _ := Jtool.GetRandInt(1, 10000)

	if testVal22 > 10000 {
		testVal33 = -100
	}

	if testVal22 < 100 {
		testVal33 = 100
	}
	testVal22 += testVal33

	state1 := protoManage.State_StateNormal
	if r1 > 8000 {
		state1 = protoManage.State_StateError
	}else if r1 > 5000 {
		state1 = protoManage.State_StateWarn
	}else if r1 < 2000 {
		state1 = protoManage.State_StateUnknow
	}

	state2 := protoManage.State_StateNormal
	if testVal22 > 8000 {
		state2 = protoManage.State_StateError
	}else if testVal22 > 5000 {
		state2 = protoManage.State_StateWarn
	}else if testVal22 < 2000 {
		state2 = protoManage.State_StateUnknow
	}

	return &mxrpc.NodeReportData{ValueList:[]mxrpc.NodeReportVal{
		mxrpc.NodeReportVal{
			Value: testVal2,
		},

		mxrpc.NodeReportVal{
			Value: r1,
			State: state1,
		},
		mxrpc.NodeReportVal{
			Value: testVal22,
			State: state2,
		},
	}}, nil
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


func testNotify()  {
	go func() {
		for  {
			r1, _ := Jtool.GetRandInt(1, 4)
			r2 := Jtool.GetRandChinese(5, 20)
			err := manageClient.SendNodeNotify(r2, constant.NodeNotifyLevel(r1), false)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(1*time.Second)
		}
	}()
}