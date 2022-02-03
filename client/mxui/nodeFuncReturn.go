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

type NodeFuncReturnText struct {
	Data 			interface{}
}

type NodeFuncReturnJson struct {
	Data 			interface{}
}

type NodeFuncReturnLink struct {
	Name        	string
	Link 			string
	Blank			bool
}

type NodeFuncReturnImage struct {
	URL 			string
	Fit				string
}

type NodeFuncReturnMedia struct {
	URL         	string
	Type 	    	string
	Live			bool
	Loop			bool
}

type NodeFuncReturnFile struct {
	Name        	string
	ID 				int64
}

type NodeFuncReturnCharts struct {
	Data 			map[string]interface{}
}

type NodeFuncReturnTable struct {
	Stripe			bool
	Border			bool
	IndexCol		NodeFuncReturnTableCol
	ShowSummary     bool
	SumText			string
	Col        		[]NodeFuncReturnTableCol
	Row				[]NodeFuncReturnTableRow
}

type NodeFuncReturnTableCol struct {
	Name        	string
	Width			uint32
	Fixed			string
	Align			string
	Resizable		bool
	MergeSameCol	bool
}

type NodeFuncReturnTableRow struct {
	Value			[]interface{}
	MergeSameRow	bool
}

type NodeFuncReturnTableVal struct {
	Data			interface{}
	State			DataState
}

func (table *NodeFuncReturnTable) AddTableCol(col NodeFuncReturnTableCol) {
	table.Col = append(table.Col, col)
}

func (table *NodeFuncReturnTable) AddTableRow(row NodeFuncReturnTableRow) {
	table.Row = append(table.Row, row)
}


