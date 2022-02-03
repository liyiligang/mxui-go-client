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

import "github.com/liyiligang/mxui-go-client/protoFiles/protoManage"

const ConstRpcClientSendBroadcast = -1
const ConstRpcClientMaxMsgSize = 10*1024*1024
const ConstManageNodeID = 0

type NodeState int32
const (
	NodeStateNormal   NodeState =  NodeState(protoManage.State_StateNormal)
	NodeStateAbnormal NodeState =  NodeState(protoManage.State_StateWarn)
	NodeStateClose    NodeState =  NodeState(protoManage.State_StateError)
)

type UserLevel int32
const (
	UserLevelLevelVisitor      	UserLevel =   1
	UserLevelLevelMember       	UserLevel =   2
	UserLevelLevelManager      	UserLevel =   3
	UserLevelLevelSuperManager 	UserLevel =   4
)

type DataState int32
const (
	DataStateInfo     	 DataState =  DataState(protoManage.State_StateUnknow)
	DataStateSuccess	 DataState =  DataState(protoManage.State_StateNormal)
	DataStateWarn   	 DataState =  DataState(protoManage.State_StateWarn)
	DataStateError  	 DataState =  DataState(protoManage.State_StateError)
)

type NodeReportType int32
const (
	NodeReportTypeTable    NodeReportType =  NodeReportType(protoManage.NodeReportType_NodeReportTypeTable)
	NodeReportTypeLine 	   NodeReportType =  NodeReportType(protoManage.NodeReportType_NodeReportTypeLine)
)

type NodeNotifyLevel int32
const (
	NodeNotifyLevelInfo    NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateUnknow)
	NodeNotifyLevelSuccess NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateNormal)
	NodeNotifyLevelWarn    NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateWarn)
	NodeNotifyLevelError   NodeNotifyLevel =  NodeNotifyLevel(protoManage.State_StateError)
)