package ccc

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

// GenerateAgentStatistic is a nested struct in ccc response
type GenerateAgentStatistic struct {
	TotalLoggedInTime int64    `json:"TotalLoggedInTime" xml:"TotalLoggedInTime"`
	RecordDate        string   `json:"RecordDate" xml:"RecordDate"`
	SkillGroupIds     string   `json:"SkillGroupIds" xml:"SkillGroupIds"`
	AverageReadyTime  int64    `json:"AverageReadyTime" xml:"AverageReadyTime"`
	SkillGroupNames   string   `json:"SkillGroupNames" xml:"SkillGroupNames"`
	MaxReadyTime      int64    `json:"MaxReadyTime" xml:"MaxReadyTime"`
	AgentName         string   `json:"AgentName" xml:"AgentName"`
	AgentId           string   `json:"AgentId" xml:"AgentId"`
	LoginName         string   `json:"LoginName" xml:"LoginName"`
	InstanceId        string   `json:"InstanceId" xml:"InstanceId"`
	TotalReadyTime    int64    `json:"TotalReadyTime" xml:"TotalReadyTime"`
	TotalBreakTime    int64    `json:"TotalBreakTime" xml:"TotalBreakTime"`
	OccupancyRate     float64  `json:"OccupancyRate" xml:"OccupancyRate"`
	Inbound           Inbound  `json:"Inbound" xml:"Inbound"`
	Overall           Overall  `json:"Overall" xml:"Overall"`
	Outbound          Outbound `json:"Outbound" xml:"Outbound"`
}
