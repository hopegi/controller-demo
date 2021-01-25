package vod

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

// GetMediaAuditResult invokes the vod.GetMediaAuditResult API synchronously
func (client *Client) GetMediaAuditResult(request *GetMediaAuditResultRequest) (response *GetMediaAuditResultResponse, err error) {
	response = CreateGetMediaAuditResultResponse()
	err = client.DoAction(request, response)
	return
}

// GetMediaAuditResultWithChan invokes the vod.GetMediaAuditResult API asynchronously
func (client *Client) GetMediaAuditResultWithChan(request *GetMediaAuditResultRequest) (<-chan *GetMediaAuditResultResponse, <-chan error) {
	responseChan := make(chan *GetMediaAuditResultResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMediaAuditResult(request)
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

// GetMediaAuditResultWithCallback invokes the vod.GetMediaAuditResult API asynchronously
func (client *Client) GetMediaAuditResultWithCallback(request *GetMediaAuditResultRequest, callback func(response *GetMediaAuditResultResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMediaAuditResultResponse
		var err error
		defer close(result)
		response, err = client.GetMediaAuditResult(request)
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

// GetMediaAuditResultRequest is the request struct for api GetMediaAuditResult
type GetMediaAuditResultRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	ResourceRealOwnerId  requests.Integer `position:"Query" name:"ResourceRealOwnerId"`
	OwnerId              string           `position:"Query" name:"OwnerId"`
	MediaId              string           `position:"Query" name:"MediaId"`
}

// GetMediaAuditResultResponse is the response struct for api GetMediaAuditResult
type GetMediaAuditResultResponse struct {
	*responses.BaseResponse
	RequestId        string           `json:"RequestId" xml:"RequestId"`
	MediaAuditResult MediaAuditResult `json:"MediaAuditResult" xml:"MediaAuditResult"`
}

// CreateGetMediaAuditResultRequest creates a request to invoke GetMediaAuditResult API
func CreateGetMediaAuditResultRequest() (request *GetMediaAuditResultRequest) {
	request = &GetMediaAuditResultRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetMediaAuditResult", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetMediaAuditResultResponse creates a response to parse from GetMediaAuditResult response
func CreateGetMediaAuditResultResponse() (response *GetMediaAuditResultResponse) {
	response = &GetMediaAuditResultResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
