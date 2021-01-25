package vcs

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

// UnbindCorpGroup invokes the vcs.UnbindCorpGroup API synchronously
func (client *Client) UnbindCorpGroup(request *UnbindCorpGroupRequest) (response *UnbindCorpGroupResponse, err error) {
	response = CreateUnbindCorpGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindCorpGroupWithChan invokes the vcs.UnbindCorpGroup API asynchronously
func (client *Client) UnbindCorpGroupWithChan(request *UnbindCorpGroupRequest) (<-chan *UnbindCorpGroupResponse, <-chan error) {
	responseChan := make(chan *UnbindCorpGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindCorpGroup(request)
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

// UnbindCorpGroupWithCallback invokes the vcs.UnbindCorpGroup API asynchronously
func (client *Client) UnbindCorpGroupWithCallback(request *UnbindCorpGroupRequest, callback func(response *UnbindCorpGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindCorpGroupResponse
		var err error
		defer close(result)
		response, err = client.UnbindCorpGroup(request)
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

// UnbindCorpGroupRequest is the request struct for api UnbindCorpGroup
type UnbindCorpGroupRequest struct {
	*requests.RpcRequest
	CorpId      string `position:"Body" name:"CorpId"`
	CorpGroupId string `position:"Body" name:"CorpGroupId"`
}

// UnbindCorpGroupResponse is the response struct for api UnbindCorpGroup
type UnbindCorpGroupResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateUnbindCorpGroupRequest creates a request to invoke UnbindCorpGroup API
func CreateUnbindCorpGroupRequest() (request *UnbindCorpGroupRequest) {
	request = &UnbindCorpGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "UnbindCorpGroup", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindCorpGroupResponse creates a response to parse from UnbindCorpGroup response
func CreateUnbindCorpGroupResponse() (response *UnbindCorpGroupResponse) {
	response = &UnbindCorpGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
