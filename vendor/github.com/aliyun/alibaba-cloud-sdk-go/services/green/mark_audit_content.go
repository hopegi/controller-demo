package green

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

// MarkAuditContent invokes the green.MarkAuditContent API synchronously
// api document: https://help.aliyun.com/api/green/markauditcontent.html
func (client *Client) MarkAuditContent(request *MarkAuditContentRequest) (response *MarkAuditContentResponse, err error) {
	response = CreateMarkAuditContentResponse()
	err = client.DoAction(request, response)
	return
}

// MarkAuditContentWithChan invokes the green.MarkAuditContent API asynchronously
// api document: https://help.aliyun.com/api/green/markauditcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkAuditContentWithChan(request *MarkAuditContentRequest) (<-chan *MarkAuditContentResponse, <-chan error) {
	responseChan := make(chan *MarkAuditContentResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MarkAuditContent(request)
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

// MarkAuditContentWithCallback invokes the green.MarkAuditContent API asynchronously
// api document: https://help.aliyun.com/api/green/markauditcontent.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkAuditContentWithCallback(request *MarkAuditContentRequest, callback func(response *MarkAuditContentResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MarkAuditContentResponse
		var err error
		defer close(result)
		response, err = client.MarkAuditContent(request)
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

// MarkAuditContentRequest is the request struct for api MarkAuditContent
type MarkAuditContentRequest struct {
	*requests.RpcRequest
	AuditIllegalReasons string `position:"Query" name:"AuditIllegalReasons"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	AuditResult         string `position:"Query" name:"AuditResult"`
	Ids                 string `position:"Query" name:"Ids"`
}

// MarkAuditContentResponse is the response struct for api MarkAuditContent
type MarkAuditContentResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMarkAuditContentRequest creates a request to invoke MarkAuditContent API
func CreateMarkAuditContentRequest() (request *MarkAuditContentRequest) {
	request = &MarkAuditContentRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "MarkAuditContent", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMarkAuditContentResponse creates a response to parse from MarkAuditContent response
func CreateMarkAuditContentResponse() (response *MarkAuditContentResponse) {
	response = &MarkAuditContentResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
