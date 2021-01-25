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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ModifySkillGroupOutboundNumbers invokes the ccc.ModifySkillGroupOutboundNumbers API synchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupoutboundnumbers.html
func (client *Client) ModifySkillGroupOutboundNumbers(request *ModifySkillGroupOutboundNumbersRequest) (response *ModifySkillGroupOutboundNumbersResponse, err error) {
	response = CreateModifySkillGroupOutboundNumbersResponse()
	err = client.DoAction(request, response)
	return
}

// ModifySkillGroupOutboundNumbersWithChan invokes the ccc.ModifySkillGroupOutboundNumbers API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupoutboundnumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupOutboundNumbersWithChan(request *ModifySkillGroupOutboundNumbersRequest) (<-chan *ModifySkillGroupOutboundNumbersResponse, <-chan error) {
	responseChan := make(chan *ModifySkillGroupOutboundNumbersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifySkillGroupOutboundNumbers(request)
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

// ModifySkillGroupOutboundNumbersWithCallback invokes the ccc.ModifySkillGroupOutboundNumbers API asynchronously
// api document: https://help.aliyun.com/api/ccc/modifyskillgroupoutboundnumbers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifySkillGroupOutboundNumbersWithCallback(request *ModifySkillGroupOutboundNumbersRequest, callback func(response *ModifySkillGroupOutboundNumbersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifySkillGroupOutboundNumbersResponse
		var err error
		defer close(result)
		response, err = client.ModifySkillGroupOutboundNumbers(request)
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

// ModifySkillGroupOutboundNumbersRequest is the request struct for api ModifySkillGroupOutboundNumbers
type ModifySkillGroupOutboundNumbersRequest struct {
	*requests.RpcRequest
	OperationType         requests.Integer `position:"Query" name:"OperationType"`
	InstanceId            string           `position:"Query" name:"InstanceId"`
	OutboundPhoneNumberId *[]string        `position:"Query" name:"OutboundPhoneNumberId"  type:"Repeated"`
	SkillGroupId          string           `position:"Query" name:"SkillGroupId"`
}

// ModifySkillGroupOutboundNumbersResponse is the response struct for api ModifySkillGroupOutboundNumbers
type ModifySkillGroupOutboundNumbersResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateModifySkillGroupOutboundNumbersRequest creates a request to invoke ModifySkillGroupOutboundNumbers API
func CreateModifySkillGroupOutboundNumbersRequest() (request *ModifySkillGroupOutboundNumbersRequest) {
	request = &ModifySkillGroupOutboundNumbersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CCC", "2017-07-05", "ModifySkillGroupOutboundNumbers", "", "")
	return
}

// CreateModifySkillGroupOutboundNumbersResponse creates a response to parse from ModifySkillGroupOutboundNumbers response
func CreateModifySkillGroupOutboundNumbersResponse() (response *ModifySkillGroupOutboundNumbersResponse) {
	response = &ModifySkillGroupOutboundNumbersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
