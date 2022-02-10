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
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/liyiligang/base/component/Jtool"
	"github.com/liyiligang/mxui-go-client/mxui"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func LoadExampleMethod(client *mxui.Client){
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
	returnEmpty(client)
	returnTextExample(client)
	returnObjectExample(client)
	returnLinkExample(client)
	returnImageExample(client)
	returnMediaExample(client)
	returnFileExample(client)
	returnTableExample(client)
	returnChartExample(client)
	returnErrorExample(client)
	returnMultiTypeExample(client)
}

func baseExample(client *mxui.Client){
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

	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "基础表单(Base form)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func defaultValExample(client *mxui.Client){
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

	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "默认值(Form with default)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func uiExample(client *mxui.Client){

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

	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "定制UI(Custom UI)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func dateTimeExample(client *mxui.Client){
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

	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "日期时间(Date time)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}


func checkExample(client *mxui.Client){
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

	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "表单校验(Form with check)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func fileExample(client *mxui.Client){
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
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "文件上传(File upload)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func containerExample(client *mxui.Client){
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
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "容器(Container)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func assistExample(client *mxui.Client){

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
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "辅助选项(Assist options)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnEmpty(client *mxui.Client){
	type returnForm struct {
		Name      	string    		`schema:"title=姓名(name)"`
		Age       	int       		`schema:"title=年龄(age)"`
		Birthday	string			`schema:"title=生日(birthday),format=date"`
	}
	callFunc := func (form *returnForm) {
		fmt.Println(form)
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "无返回值(No return value)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnTextExample(client *mxui.Client){
	type returnForm struct {
		Name      	string    		`schema:"title=姓名(name)"`
		Age       	int64       	`schema:"title=年龄(age)"`
		Birthday	string			`schema:"title=生日(birthday),format=date"`
	}
	callFunc := func (form *returnForm) interface{} {
		return "My name: " + form.Name + "\nMy age: "+ strconv.FormatInt(form.Age, 10) +
			"\nMy birthday: " + form.Birthday
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "文本预览(Text preview)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnObjectExample(client *mxui.Client){
	type returnForm struct {
		Name      	string    		`schema:"title=姓名(name)"`
		Age       	int64       	`schema:"title=年龄(age)"`
		Birthday	string			`schema:"title=生日(birthday),format=date"`
	}
	callFunc := func (form *returnForm) *returnForm {
		return form
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "对象预览(Object preview)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnLinkExample(client *mxui.Client){
	type returnForm struct {
		LinkName      	string    		`schema:"title=链接名(link name),default=百度(baidu)"`
		Link      		string    		`schema:"title=链接Url(link Url),default=https://www.baidu.com"`
	}
	callFunc := func (form *returnForm) mxui.NodeFuncReturnLink {
		return mxui.NodeFuncReturnLink{Name: form.LinkName, Link: form.Link, Blank: true}
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "链接(Link)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnImageExample(client *mxui.Client){
	type returnForm struct {
		Url      	string    	`schema:"title=图片Url(Image Url),default=https://webstatic.mihoyo.com/ys/event/e20210901-fab/images/poster1.bccfc913.jpg"`
		ViewMode    string      `schema:"title=预览模式(View Mode),default=fill,enum=fill;contain;cover;none;scale-down,enumNames=填充(fill);包含(contain);覆盖(cover);无(none);缩放(scale-down)"`
	}
	callFunc := func (form *returnForm) *mxui.NodeFuncReturnImage {
		return &mxui.NodeFuncReturnImage{URL: form.Url, Fit: form.ViewMode}
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "图片预览(Image view)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnMediaExample(client *mxui.Client){
	type returnForm struct {
		Url      	string    		`schema:"title=媒体Url(media Url),default=http://clips.vorwaerts-gmbh.de/big_buck_bunny.mp4"`
		IsLive		bool	   		`schema:"title=直播(is live)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		Loop		bool	   		`schema:"title=循环播放(loop playback)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
	}
	callFunc := func (form *returnForm) mxui.NodeFuncReturnMedia {
		return mxui.NodeFuncReturnMedia{URL: form.Url, Live:form.IsLive, Loop: form.Loop}
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "媒体播放(Media play)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnFileExample(client *mxui.Client){
	type returnForm struct {
		File     	 string    		`schema:"title=文件上传(file upload)" ui:"{'ui:widget':'UploadFile','ui:btnText':'上传文件(upload file)'}"`
	}
	callFunc := func (form *returnForm) *mxui.NodeFuncReturnFile {
		name, data, err := client.DownloadNodeResourceWithBytes(form.File)
		if err != nil{
			return nil
		}
		cData, err := Jtool.CompressByteWithZip(name, data)
		if err != nil{
			return nil
		}
		r, err := client.UploadNodeResourceWithBytes(name+".zip", cData)
		if err != nil{
			return nil
		}
		return &mxui.NodeFuncReturnFile{ID: r.Base.ID, Name:r.Name}
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "文件下载(File download)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnTableExample(client *mxui.Client){
	type Table struct {
		Name      	string    		`schema:"title=姓名(name),default=Jin yun"`
		Age       	int       		`schema:"title=年龄(age),default=20"`
		Birthday	string			`schema:"title=生日(birthday),format=date,default=1996-10-13"`
	}

	type returnForm struct {
		ShowStripe		bool	   		`schema:"title=显示斑马纹(show stripe)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		ShowBorder		bool	   		`schema:"title=显示边框(show border)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		ShowIndex		bool	   		`schema:"title=显示编号(show ID)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		ShowColor		bool	   		`schema:"title=显示颜色(show color)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		MergeCol		bool			`schema:"title=合并相同的行(merge same row)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		MergeRow		bool			`schema:"title=合并相同的列(merge same col)" ui:"{'ui:options':{'activeText':'是(yes)','inactiveText':'否(no)'}}"`
		Align			string			`schema:"title=对齐方式(align),default=left,enum=left;center;right,enumNames=靠左(left);居中(center);靠右(right)"`
		Table			[]Table			`schema:"title=添加用户信息(add user info)"`
	}
	callFunc := func (form *returnForm) *mxui.NodeFuncReturnTable {
		table := mxui.NodeFuncReturnTable{
			Stripe: form.ShowStripe,
			Border: form.ShowBorder,
		}
		if form.ShowIndex{
			table.IndexCol = mxui.NodeFuncReturnTableCol{Name: "ID", Width: 60, Align: form.Align}
		}
		table.AddTableCol(mxui.NodeFuncReturnTableCol{
			Name: "姓名(name)",
			Width: 120,
			Align: form.Align,
			MergeSameCol: form.MergeCol,
		})
		table.AddTableCol(mxui.NodeFuncReturnTableCol{
			Name: "年龄(age)",
			Width: 100,
			Align: form.Align,
			MergeSameCol: form.MergeCol,
		})
		table.AddTableCol(mxui.NodeFuncReturnTableCol{
			Name: "生日(birthday)",
			Align: form.Align,
			MergeSameCol: form.MergeCol,
		})
		for _, v := range form.Table {
			if form.ShowColor{
				table.AddTableRow(mxui.NodeFuncReturnTableRow{
					Value: []interface{}{mxui.NodeFuncReturnTableVal{
						Data:  v.Name,
						State: mxui.DataState(rand.Intn(5)),
					}, v.Age, v.Birthday},
					MergeSameRow: form.MergeRow,
				})
			}else{
				table.AddTableRow(mxui.NodeFuncReturnTableRow{
					Value: []interface{}{v.Name, v.Age, v.Birthday},
					MergeSameRow: form.MergeRow,
				})
			}
		}
		return &table
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "表格预览(Table view)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnChartExample(client *mxui.Client){
	type Chart struct {
		Name      	string    		`schema:"title=姓名(name),required"`
		Age       	int       		`schema:"title=年龄(age),default=20"`
	}
	type returnForm struct {
		ChartType   string      	`schema:"title=图表类型(Chart type),default=line,enum=line;bar;pie,enumNames=折线图(line chart);柱状图(bar chart);饼图(pie chart);"`
		Chart		[]Chart			`schema:"title=添加用户信息(add user info)"`
	}
	returnVal := &mxui.NodeFuncReturnCharts{}
	callFunc := func (form *returnForm) *mxui.NodeFuncReturnCharts {
		switch form.ChartType {
		case "line":
			var xAxis []string
			var yAxis []opts.LineData
			for _, v := range form.Chart {
				xAxis = append(xAxis, v.Name)
				yAxis = append(yAxis, opts.LineData{Value: v.Age})
			}
			line := charts.NewLine()
			line.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
				Title:    "用户年龄分布(User age distribution)",
			}),	charts.WithTooltipOpts(opts.Tooltip{Show: true}))
			line.SetXAxis(xAxis).AddSeries("User", yAxis).SetSeriesOptions(
				charts.WithLabelOpts(opts.Label{
					Show: true,
				}),
			)
			buf:= &bytes.Buffer{}
			_ = line.Render(buf)
			returnVal.Data = line.JSON()
			break
		case "bar":
			var xAxis []string
			var yAxis []opts.BarData
			for _, v := range form.Chart {
				xAxis = append(xAxis, v.Name)
				yAxis = append(yAxis, opts.BarData{Value: v.Age})
			}
			bar := charts.NewBar()
			bar.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
				Title:    "用户年龄分布(User age distribution)",
			}), charts.WithTooltipOpts(opts.Tooltip{Show: true}))
			bar.SetXAxis(xAxis).AddSeries("User", yAxis)
			buf:= &bytes.Buffer{}
			_ = bar.Render(buf)
			returnVal.Data = bar.JSON()
			break
		case "pie":
			var pieData []opts.PieData
			for _, v := range form.Chart {
				pieData = append(pieData, opts.PieData{Name: v.Name, Value: v.Age})
			}
			pie := charts.NewPie()
			pie.SetGlobalOptions(charts.WithTitleOpts(opts.Title{
				Title:    "用户年龄分布(User age distribution)",
			}),	charts.WithTooltipOpts(opts.Tooltip{Show: true}))
			pie.AddSeries("User", pieData)
			buf:= &bytes.Buffer{}
			_ = pie.Render(buf)
			returnVal.Data = pie.JSON()
			break

		}
		return returnVal
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "图表预览(Chart view)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnMultiTypeExample(client *mxui.Client){
	type returnForm struct {
		ReturnType   string      	`schema:"title=返回类型(Return type),default=text,enum=text;object;file;null;error,enumNames=文本(text);对象(object);文件(file);空(null);错误(error)"`
		Name      	string    		`schema:"title=姓名(name),default=Jin yun"`
		Age       	int64       	`schema:"title=年龄(age),default=20"`
		Birthday	string			`schema:"title=生日(birthday),format=date,default=1996-10-13"`
	}
	callFunc := func (form *returnForm) interface{} {
		switch form.ReturnType {
		case "text":
			return "My name: " + form.Name + "\nMy age: "+ strconv.FormatInt(form.Age, 10) +
				"\nMy birthday: " + form.Birthday
		case "object":
			return form
		case "file":
			data, _ := json.Marshal(form)
			r, err := client.UploadNodeResourceWithBytes("用户信息(userInfo).txt", data)
			if err != nil{
				return err
			}
			return &mxui.NodeFuncReturnFile{ID: r.Base.ID, Name:r.Name}
		case "null":
			return nil
		case "error":
			return errors.New("用户信息错误(User information error)")
		}
		return nil
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "动态类型(Dynamic type)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func returnErrorExample(client *mxui.Client){
	type returnForm struct {
		Name      	string    		`schema:"title=姓名(name),default=Jin yun"`
		Age       	int64       	`schema:"title=年龄(age),default=20"`
		Birthday	string			`schema:"title=生日(birthday),format=date,default=1996-10-13"`
	}
	callFunc := func (form *returnForm) (string, error) {
		if form.Age < 500 {
			return "", errors.New("年龄必须大于500(age must be greater than 500)")
		}
		if form.Age > 1000 {
			return "", errors.New("年龄必须小于1000(Age must be less than 1000)")
		}
		return "My name: " + form.Name + "\nMy age: "+ strconv.FormatInt(form.Age, 10) +
			"\nMy birthday: " + form.Birthday, nil
	}
	err := client.RegisterNodeFunc(mxui.NodeFuncRegister{
		Name:     "错误类型(error type)",
		CallFunc: callFunc,
		Level:    mxui.UserLevelLevelVisitor,
	})
	if err != nil {
		fmt.Println(err)
	}
}
