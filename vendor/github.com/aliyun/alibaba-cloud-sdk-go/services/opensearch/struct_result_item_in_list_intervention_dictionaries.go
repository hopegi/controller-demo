package opensearch

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

// ResultItemInListInterventionDictionaries is a nested struct in opensearch response
type ResultItemInListInterventionDictionaries struct {
	Id       int    `json:"id" xml:"id"`
	Name     string `json:"name" xml:"name"`
	Type     string `json:"type" xml:"type"`
	Analyzer string `json:"analyzer" xml:"analyzer"`
	Created  int    `json:"created" xml:"created"`
	Updated  int    `json:"updated" xml:"updated"`
}
