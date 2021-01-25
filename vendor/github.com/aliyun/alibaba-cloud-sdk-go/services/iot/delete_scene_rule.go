package iot

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

// DeleteSceneRule invokes the iot.DeleteSceneRule API synchronously
// api document: https://help.aliyun.com/api/iot/deletescenerule.html
func (client *Client) DeleteSceneRule(request *DeleteSceneRuleRequest) (response *DeleteSceneRuleResponse, err error) {
	response = CreateDeleteSceneRuleResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteSceneRuleWithChan invokes the iot.DeleteSceneRule API asynchronously
// api document: https://help.aliyun.com/api/iot/deletescenerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSceneRuleWithChan(request *DeleteSceneRuleRequest) (<-chan *DeleteSceneRuleResponse, <-chan error) {
	responseChan := make(chan *DeleteSceneRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteSceneRule(request)
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

// DeleteSceneRuleWithCallback invokes the iot.DeleteSceneRule API asynchronously
// api document: https://help.aliyun.com/api/iot/deletescenerule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteSceneRuleWithCallback(request *DeleteSceneRuleRequest, callback func(response *DeleteSceneRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteSceneRuleResponse
		var err error
		defer close(result)
		response, err = client.DeleteSceneRule(request)
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

// DeleteSceneRuleRequest is the request struct for api DeleteSceneRule
type DeleteSceneRuleRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
	RuleId        string `position:"Query" name:"RuleId"`
}

// DeleteSceneRuleResponse is the response struct for api DeleteSceneRule
type DeleteSceneRuleResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Code         string `json:"Code" xml:"Code"`
}

// CreateDeleteSceneRuleRequest creates a request to invoke DeleteSceneRule API
func CreateDeleteSceneRuleRequest() (request *DeleteSceneRuleRequest) {
	request = &DeleteSceneRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "DeleteSceneRule", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteSceneRuleResponse creates a response to parse from DeleteSceneRule response
func CreateDeleteSceneRuleResponse() (response *DeleteSceneRuleResponse) {
	response = &DeleteSceneRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
