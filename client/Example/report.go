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

package Example

import (
	"fmt"
	"github.com/liyiligang/mxui-go-client/mxui"
	"math/rand"
	"time"
)

func LoadExampleReport(client *MXUI.Client){
	tableExample(client)
	lineChartExample(client)
}

func tableExample(client *MXUI.Client){
	type tableReport struct {
		Changsha MXUI.NodeReportValue `schema:"title=长沙(Chang Sha)"`
		NewYork  int                  `schema:"title=纽约(New York)"`
		Paris    int                  `schema:"title=巴黎(Paris)"`
	}
	callFunc := func() (*tableReport, error) {
		cs :=  rand.Intn(15) - 5
		state := MXUI.DataStateSuccess
		if cs < 0 {
			state = MXUI.DataStateInfo
		}
		if cs > 5 {
			state = MXUI.DataStateWarn
		}
		return &tableReport{Changsha: MXUI.NodeReportValue{Data: cs, State: state},
			NewYork: rand.Intn(10)+10, Paris: rand.Intn(10)+20}, nil
	}
	err := client.RegisterNodeReport(MXUI.NodeReportRegister{
		Name:         "气温-表格(temperature-table)",
		Type:         MXUI.NodeReportTypeTable,
		CallFunc:     callFunc,
		CallInterval: time.Second*2,
		Level:        MXUI.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}
}

func lineChartExample(client *MXUI.Client){
	type lineReport struct {
		Changsha MXUI.NodeReportValue `schema:"title=长沙(Chang Sha)"`
		NewYork  int                  `schema:"title=纽约(New York)"`
		Paris    int                  `schema:"title=巴黎(Paris)"`
	}

	callFunc := func() (*lineReport, error) {
		cs :=  rand.Intn(15) - 5
		state := MXUI.DataStateSuccess
		if cs < 0 {
			state = MXUI.DataStateInfo
		}
		if cs > 5 {
			state = MXUI.DataStateWarn
		}
		return &lineReport{Changsha: MXUI.NodeReportValue{Data: cs, State: state},
			NewYork: rand.Intn(10)+10, Paris: rand.Intn(10)+20}, nil
	}

	err := client.RegisterNodeReport(MXUI.NodeReportRegister{
		Name:         "气温-折线(temperature-line)",
		Type:         MXUI.NodeReportTypeLine,
		CallFunc:     callFunc,
		CallInterval: time.Second*2,
		Level:        MXUI.UserLevelLevelManager,
	})
	if err != nil {
		fmt.Println(err)
	}
}

