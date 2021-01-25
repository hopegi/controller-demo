package dataworks_public

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

// GetQualityRule invokes the dataworks_public.GetQualityRule API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/getqualityrule.html
func (client *Client) GetQualityRule(request *GetQualityRuleRequest) (response *GetQualityRuleResponse, err error) {
	response = CreateGetQualityRuleResponse()
	err = client.DoAction(request, response)
	return
}

// GetQualityRuleWithChan invokes the dataworks_public.GetQualityRule API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getqualityrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQualityRuleWithChan(request *GetQualityRuleRequest) (<-chan *GetQualityRuleResponse, <-chan error) {
	responseChan := make(chan *GetQualityRuleResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetQualityRule(request)
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

// GetQualityRuleWithCallback invokes the dataworks_public.GetQualityRule API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getqualityrule.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetQualityRuleWithCallback(request *GetQualityRuleRequest, callback func(response *GetQualityRuleResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetQualityRuleResponse
		var err error
		defer close(result)
		response, err = client.GetQualityRule(request)
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

// GetQualityRuleRequest is the request struct for api GetQualityRule
type GetQualityRuleRequest struct {
	*requests.RpcRequest
	ProjectName string           `position:"Body" name:"ProjectName"`
	RuleId      requests.Integer `position:"Body" name:"RuleId"`
}

// GetQualityRuleResponse is the response struct for api GetQualityRule
type GetQualityRuleResponse struct {
	*responses.BaseResponse
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetQualityRuleRequest creates a request to invoke GetQualityRule API
func CreateGetQualityRuleRequest() (request *GetQualityRuleRequest) {
	request = &GetQualityRuleRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetQualityRule", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetQualityRuleResponse creates a response to parse from GetQualityRule response
func CreateGetQualityRuleResponse() (response *GetQualityRuleResponse) {
	response = &GetQualityRuleResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
