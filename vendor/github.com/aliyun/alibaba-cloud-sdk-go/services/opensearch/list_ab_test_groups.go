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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListABTestGroups invokes the opensearch.ListABTestGroups API synchronously
// api document: https://help.aliyun.com/api/opensearch/listabtestgroups.html
func (client *Client) ListABTestGroups(request *ListABTestGroupsRequest) (response *ListABTestGroupsResponse, err error) {
	response = CreateListABTestGroupsResponse()
	err = client.DoAction(request, response)
	return
}

// ListABTestGroupsWithChan invokes the opensearch.ListABTestGroups API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listabtestgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListABTestGroupsWithChan(request *ListABTestGroupsRequest) (<-chan *ListABTestGroupsResponse, <-chan error) {
	responseChan := make(chan *ListABTestGroupsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListABTestGroups(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListABTestGroupsWithCallback invokes the opensearch.ListABTestGroups API asynchronously
// api document: https://help.aliyun.com/api/opensearch/listabtestgroups.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListABTestGroupsWithCallback(request *ListABTestGroupsRequest, callback func(response *ListABTestGroupsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListABTestGroupsResponse
		var err error
		defer close(result)
		response, err = client.ListABTestGroups(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListABTestGroupsRequest is the request struct for api ListABTestGroups
type ListABTestGroupsRequest struct {
	*requests.RoaRequest
	SceneId          requests.Integer `position:"Path" name:"sceneId"`
	AppGroupIdentity string           `position:"Path" name:"appGroupIdentity"`
}

// ListABTestGroupsResponse is the response struct for api ListABTestGroups
type ListABTestGroupsResponse struct {
	*responses.BaseResponse
	RequestId string       `json:"requestId" xml:"requestId"`
	Result    []ResultItem `json:"result" xml:"result"`
}

// CreateListABTestGroupsRequest creates a request to invoke ListABTestGroups API
func CreateListABTestGroupsRequest() (request *ListABTestGroupsRequest) {
	request = &ListABTestGroupsRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("OpenSearch", "2017-12-25", "ListABTestGroups", "/v4/openapi/app-groups/[appGroupIdentity]/scenes/[sceneId]/groups", "opensearch", "openAPI")
	request.Method = requests.GET
	return
}

// CreateListABTestGroupsResponse creates a response to parse from ListABTestGroups response
func CreateListABTestGroupsResponse() (response *ListABTestGroupsResponse) {
	response = &ListABTestGroupsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
