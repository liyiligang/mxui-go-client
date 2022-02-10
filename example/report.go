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
	"github.com/liyiligang/mxui-go-client/mxui"
	"math/rand"
	"time"
)

func LoadExampleReport(client *mxui.Client){
	tableExample(client)
	lineChartExample(client)
}

func tableExample(client *mxui.Client){
	type tableReport struct {
		Changsha mxui.NodeReportValue `schema:"title=长沙(Chang Sha)"`
		NewYork  int                  `schema:"title=纽约(New York)"`
		Paris    int                  `schema:"title=巴黎(Paris)"`
	}
	callFunc := func() (*tableReport, error) {
		cs :=  rand.Intn(15) - 5
		state := mxui.DataStateSuccess
		if cs < 0 {
			state = mxui.DataStateInfo
		}
		if cs > 5 {
			state = mxui.DataStateWarn
		}
		return &tableReport{Changsha: mxui.NodeReportValue{Data: cs, State: state},
			NewYork: rand.Intn(10)+10, Paris: rand.Intn(10)+20}, nil
	}
	err := client.RegisterNodeReport(mxui.NodeReportRegister{
		Name:         "气温-表格(temperature-table)",
		Type:         mxui.NodeReportTypeTable,
		CallFunc:     callFunc,
		CallInterval: time.Second*2,
		Level:        mxui.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func lineChartExample(client *mxui.Client){
	type lineReport struct {
		Changsha mxui.NodeReportValue `schema:"title=长沙(Chang Sha)"`
		NewYork  int                  `schema:"title=纽约(New York)"`
		Paris    int                  `schema:"title=巴黎(Paris)"`
	}

	callFunc := func() (*lineReport, error) {
		cs :=  rand.Intn(15) - 5
		state := mxui.DataStateSuccess
		if cs < 0 {
			state = mxui.DataStateInfo
		}
		if cs > 5 {
			state = mxui.DataStateWarn
		}
		return &lineReport{Changsha: mxui.NodeReportValue{Data: cs, State: state},
			NewYork: rand.Intn(10)+10, Paris: rand.Intn(10)+20}, nil
	}

	err := client.RegisterNodeReport(mxui.NodeReportRegister{
		Name:         "气温-折线(temperature-line)",
		Type:         mxui.NodeReportTypeLine,
		CallFunc:     callFunc,
		CallInterval: time.Second*2,
		Level:        mxui.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}
}

