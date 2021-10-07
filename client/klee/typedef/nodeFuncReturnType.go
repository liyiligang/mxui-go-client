// Copyright 2021 The Authors. All rights reserved.
// Author: liyiligang
// Date: 2021/10/06 16:04
// Description:

package typedef

import "github.com/liyiligang/klee-client-go/protoFiles/protoManage"

type NodeFuncReturnText struct {
	Data 			interface{}
}

type NodeFuncReturnJson struct {
	Data 			interface{}
}

type NodeFuncReturnLink struct {
	Name        	string
	Link 			string
	AutoOpen		bool
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
	AutoPlay		bool
}

type NodeFuncReturnFile struct {
	Name        	string
	Data 			[]byte
	AutoSave		bool
}

type NodeFuncReturnCharts struct {
	Data 			map[string]interface{}
}

type NodeFuncReturnTableCol struct {
	Name        	string
	Width			uint32
	Type            string
	Fixed			string
	Align			string
	Resizable		bool
	MergeSameCol	bool
}

type NodeFuncReturnTableRow struct {
	Data			[]interface{}
	State			protoManage.State
	MergeSameRow	bool
}

type NodeFuncReturnTable struct {
	Stripe			bool
	Border			bool
	ShowSummary     bool
	ShowIndex		bool
	SumText			string
	Col        		[]NodeFuncReturnTableCol
	Row				[]NodeFuncReturnTableRow
}
