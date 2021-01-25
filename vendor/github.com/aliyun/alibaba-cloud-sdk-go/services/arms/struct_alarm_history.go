package arms

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

// AlarmHistory is a nested struct in arms response
type AlarmHistory struct {
	Id                int64  `json:"Id" xml:"Id"`
	StrategyId        string `json:"StrategyId" xml:"StrategyId"`
	UserId            string `json:"UserId" xml:"UserId"`
	Target            string `json:"Target" xml:"Target"`
	Phones            string `json:"Phones" xml:"Phones"`
	Emails            string `json:"Emails" xml:"Emails"`
	AlarmTime         int64  `json:"AlarmTime" xml:"AlarmTime"`
	AlarmType         int    `json:"AlarmType" xml:"AlarmType"`
	AlarmResponseCode int    `json:"AlarmResponseCode" xml:"AlarmResponseCode"`
	AlarmContent      string `json:"AlarmContent" xml:"AlarmContent"`
	AlarmSources      string `json:"AlarmSources" xml:"AlarmSources"`
}
