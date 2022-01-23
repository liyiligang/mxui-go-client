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

package example

import (
	"fmt"
	"github.com/liyiligang/mxrpc-go-client/mxrpc"
	"github.com/liyiligang/mxrpc-go-client/typedef/constant"
	"strings"
	"time"
)

func LoadExampleMethod(client *mxrpc.Client){
	//parameter
	baseExample(client)
	defaultValExample(client)
	uiExample(client)
	dateTimeExample(client)
	fileExample(client)
	containerExample(client)
	checkExample(client)
	assistExample(client)

	//return val
}

func baseExample(client *mxrpc.Client){
	type baseForm struct {
		Name      	string    		`schema:"title=姓名(name)"`
		Age       	int       		`schema:"title=年龄(age)"`
		Weight      float64       	`schema:"title=体重(weight)[kg]"`
		Sex		  	bool	   		`schema:"title=是否已婚(married)"`
		Birthday	string			`schema:"title=生日(birthday),format=date"`
		Occupation  string      	`schema:"title=职业(occupation),enum=teacher;sales;doctor,enumNames=教师(teacher);销售(sales);医生(doctor)"`
		Interest  	[]string      	`schema:"title=爱好(interest),uniqueItems=true,enum=basketball;piano;sing,enumNames=篮球(basketball);钢琴(piano);唱歌(sing)"`
	}

	callFunc := func (form *baseForm) *baseForm {
		return form
	}

	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "基础表单(base form)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func defaultValExample(client *mxrpc.Client){
	type defaultValForm struct {
		Name      	string    		`schema:"title=姓名(name),default=Jin yun"`
		Age       	int       		`schema:"title=年龄(age),default=20"`
		Weight      float64       	`schema:"title=体重(weight)[kg],default=62.5"`
		Sex		  	bool	   		`schema:"title=是否已婚(married),default=false"`
		Birthday	string			`schema:"title=生日(birthday),format=date,default=2022-01-16"`
		Occupation  string      	`schema:"title=职业(occupation),default=sales,enum=teacher;sales;doctor,enumNames=教师(teacher);销售(sales);医生(doctor)"`
		Interest  	[]string      	`schema:"title=爱好(interest),default=piano;sing,uniqueItems=true,enum=basketball;piano;sing,enumNames=篮球(basketball);钢琴(piano);唱歌(sing)"`
	}

	callFunc := func (form *defaultValForm) *defaultValForm {
		return form
	}

	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "默认值(form with default)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func uiExample(client *mxrpc.Client){

	type boolCheck struct {
		Golang		  	bool	   		`ui:"{'ui:widget':'el-checkbox','ui:options':{'label':'golang'}}"`
		JavaScript		bool	   		`ui:"{'ui:widget':'el-checkbox','ui:options':{'label':'javaScript'}}"`
		C		  		bool	   		`ui:"{'ui:widget':'el-checkbox','ui:options':{'label':'c++'}}"`
	}

	type uiForm struct {
		Tooltips 		string    		`schema:"title=输入提示(tooltips)" ui:"{'ui:options':{'placeholder': '请输入内容(Please enter)'}}"`
		CustomWidth 	string    		`schema:"title=自定义宽度(custom width)" ui:"{'ui:options':{'width':'70%'}}"`
		CustomStyles    string    		`schema:"title=自定义样式(custom styles)" ui:"{'ui:options':{'style':{'boxShadow':'0 0 6px 2px #2b9939'}}}"`
		Password     	string    		`schema:"title=密码类型(password type)" ui:"{'ui:options':{'type':'password'}}"`
		Textarea     	string    		`schema:"title=多行文本(textarea)" ui:"{'ui:options':{'type':'textarea','rows':3}}"`
		Precision      	float64       	`schema:"title=数值精度(number precision)" ui:"{'ui:options':{'precision':2}}"`
		Step      		int       		`schema:"title=数值步长(number step)" ui:"{'ui:options':{'step':10}}"`
		Slider       	int       		`schema:"title=数值滑块(number slider),minimum=20,maximum=100" ui:"{'ui:widget':'el-slider'}"`
		BoolSwitch		bool	   		`schema:"title=Bool开关(bool switch)" ui:"{'ui:options':{'activeText':'开(open)','inactiveText':'关(close)'}}"`
		BoolRadio		bool	   		`schema:"title=Bool单选(bool radio),enumNames=是(yes);否(no)" ui:"{'ui:widget':'RadioWidget'}"`
		BoolSelect		bool	   		`schema:"title=Bool选择器(bool select),enumNames=是(yes);否(no)" ui:"{'ui:widget':'SelectWidget'}"`
		BoolCheck		boolCheck		`schema:"title=Bool多选(bool multiple)"`
		Radio  			string      	`schema:"title=单选(radio),enum=teacher;sales;doctor,enumNames=教师(teacher);销售(sales);医生(doctor)" ui:"{'ui:widget':'RadioWidget'}"`
		MultiSelect  	[]string      	`schema:"title=多选选择器(multi select),uniqueItems=true,enum=teacher;sales;doctor,enumNames=教师(teacher);销售(sales);医生(doctor)" ui:"{'ui:widget':'SelectWidget'}"`
	}

	callFunc := func (form *uiForm) *uiForm {
		return form
	}

	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name:     "定制UI(custom UI)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func dateTimeExample(client *mxrpc.Client){
	type dateTimeForm struct {
		Time	  			string			`schema:"title=时间选择器(time picker),format=time"`
		Date	  			string			`schema:"title=日期选择器(date picker),format=date"`
		DateTime	  		string			`schema:"title=时间日期选择器(date time picker),format=date-time"`
		DateStamp	  		int64			`schema:"title=日期选择器(date picker)[timeStamp],format=date"`
		DateTimeStamp	  	int64			`schema:"title=时间日期选择器(date time picker)[timeStamp],format=date-time"`
		DateRange	  		[]string		`schema:"title=日期范围选择器(date range picker),format=date"`
		DateTimeRange		[]string		`schema:"title=时间日期范围选择器(date time range picker),format=date-time"`
		DateStampRange	  	[]int64			`schema:"title=日期范围选择器(date range picker)[timeStamp],format=date"`
		DateTimeStampRange	[]int64			`schema:"title=时间日期范围选择器(date time range picker)[timeStamp],format=date-time"`
	}

	callFunc := func (form *dateTimeForm) *dateTimeForm {
		dateTime, _ := time.Parse("2006-01-02T15:04:05.000Z", form.DateTime)
		form.DateTime = dateTime.Local().String()
		if len(form.DateTimeRange) == 2{
			dateTimeStart, _ := time.Parse("2006-01-02T15:04:05.000Z", form.DateTimeRange[0])
			dateTimeEnd, _ := time.Parse("2006-01-02T15:04:05.000Z", form.DateTimeRange[1])
			form.DateTimeRange[0] = dateTimeStart.Local().String()
			form.DateTimeRange[1] = dateTimeEnd.Local().String()
		}
		//timeStart := time.Unix(form.DateTimeRange[0]/1000, 0)
		//timeEnd := time.Unix(form.DateTimeRange[0]/1000, 0)
		return form
	}

	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name: "日期时间(date time)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}


func checkExample(client *mxrpc.Client){
	type checkForm struct {
		Multiple       	int       	`schema:"title=指定数值倍数(specify value multiple),multipleOf=2"`
		MaxMinVal       int       	`schema:"title=限制数值范围(limit value range)[10-100],minimum=10,maximum=100"`
		Required		string		`schema:"title=必填选项(required options),required"`
		MaxMinLength    string    	`schema:"title=限制字符串长度(limit string length)[10-20],minLength=10,maxLength=20"`
		Pattern    		string    	`schema:"title=正则表达式(regular expression)[手机号码(phone number)],pattern=^1[3|4|5|7|8][0-9]{9}$"`
		Email    		string    	`schema:"title=限制为Email格式(email format),format=email"`
		HostName    	string    	`schema:"title=限制为Hostname格式(hostname format),format=hostname"`
		Ipv4    		string    	`schema:"title=限制为Ipv4格式(ipv4 format),format=ipv4"`
		Ipv6    		string    	`schema:"title=限制为Ipv6格式(ipv6 format),format=ipv6"`
		Uri    			string    	`schema:"title=限制为Uri格式(uri format),format=uri"`
		ArrayRange		[]bool		`schema:"title=限制数组大小(limit array size),minItems=2,maxItems=5"`
	}

	callFunc := func (form *checkForm) *checkForm {
		return form
	}

	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name: "校验(form with check)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func fileExample(client *mxrpc.Client){
	type fileForm struct {
		File     	 string    		`schema:"title=文件上传(file upload)" ui:"{'ui:widget':'UploadFile','ui:btnText':'上传文件(upload file)'}"`
		FileList     []string    	`schema:"title=文件批量上传(file batch upload)" ui:"{'ui:widget':'UploadFile','ui:btnText':'上传文件(upload file)'}"`
	}
	callFunc := func (form *fileForm) *fileForm {
		strList := strings.Split(form.File, "_")
		if len(strList) >= 2 {
			form.File = strList[1]
		}
		for i, f := range form.FileList {
			strList = strings.Split(f, "_")
			if len(strList) >= 2 {
				form.FileList[i] = strList[1]
			}
		}
		//download file
		//client.DownloadNodeResourceWithFile(form.File, "../file")
		return form
	}
	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name: "文件上传(file upload)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func containerExample(client *mxrpc.Client){
	type combineA struct {
		String      string    		`schema:"title=字符串(string),default=123"`
		Int      	int    			`schema:"title=整型(integer),default=123"`
	}
	type combineB struct {
		Float      	float64    		`schema:"title=浮点(float),default=1.23"`
		Bool      	bool    		`schema:"title=布尔(bool),default=true"`
	}
	type Date struct {
		Date			string			`schema:"format=date"`
	}
	type containerForm struct {
		CombineA  		combineA		`schema:"title=组合A(combine A)"`
		CombineB  		combineB		`schema:"title=组合B(combine B)"`
		ArrayString		[]string		`schema:"title=字符数组(string array)"`
		ArrayInt		[]int			`schema:"title=整型数组(integer array),default=123"`
		ArrayBool		[]bool			`schema:"title=布尔数组(boolean array),default=true;false"`
		ArraySelect		[]string		`schema:"title=选择器数组(select array),enum=teacher;sales;doctor,enumNames=教师(teacher);销售(sales);医生(doctor)"`
		Date			[]Date			`schema:"title=日期选择器数组(date select array)"`
		StructInt		[]combineA		`schema:"title=结构体数组(struct array)"`
	}
	callFunc := func (form *containerForm) *containerForm {
		return form
	}
	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name: "容器(container)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func assistExample(client *mxrpc.Client){

	type StructPtr struct {
		FloatPtr      *float64    	`schema:"title=浮点指针(float pointer)"`
		BoolPtr       *bool    		`schema:"title=布尔指针(bool pointer)"`
	}
	type assistForm struct {
		Name      	string    		`schema:"title=提示(tips),description=请输入你的姓名(please enter your name)"`
		Age      	int       		`schema:"title=提示(tips),description=请输入你的年龄(please enter your age)："`
		Ignore      string    		`schema:"-"`		//hide UI
		StringPtr   *string    		`schema:"title=字符指针(string pointer)"`
		IntPtr      *int    		`schema:"title=整型指针(integer pointer)"`
		StructPtr   *StructPtr		`schema:"title=结构体指针(struct pointer)"`
	}
	callFunc := func (form *assistForm) *assistForm {
		return form
	}
	err := client.RegisterNodeFunc(mxrpc.NodeFuncRegister{
		Name: "辅助设置(assist options)",
		CallFunc: callFunc,
		Level:    constant.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}