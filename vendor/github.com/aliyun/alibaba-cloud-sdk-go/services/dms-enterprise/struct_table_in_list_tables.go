package dms_enterprise

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// TableInListTables is a nested struct in dms_enterprise response
type TableInListTables struct {
	TableId         string                    `json:"TableId" xml:"TableId"`
	DatabaseId      string                    `json:"DatabaseId" xml:"DatabaseId"`
	TableName       string                    `json:"TableName" xml:"TableName"`
	TableSchemaName string                    `json:"TableSchemaName" xml:"TableSchemaName"`
	Engine          string                    `json:"Engine" xml:"Engine"`
	Encoding        string                    `json:"Encoding" xml:"Encoding"`
	TableType       string                    `json:"TableType" xml:"TableType"`
	NumRows         int64                     `json:"NumRows" xml:"NumRows"`
	StoreCapacity   int64                     `json:"StoreCapacity" xml:"StoreCapacity"`
	TableGuid       string                    `json:"TableGuid" xml:"TableGuid"`
	Description     string                    `json:"Description" xml:"Description"`
	OwnerIdList     OwnerIdListInListTables   `json:"OwnerIdList" xml:"OwnerIdList"`
	OwnerNameList   OwnerNameListInListTables `json:"OwnerNameList" xml:"OwnerNameList"`
}
