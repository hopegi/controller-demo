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

// MarkAuditContentItem invokes the green.MarkAuditContentItem API synchronously
// api document: https://help.aliyun.com/api/green/markauditcontentitem.html
func (client *Client) MarkAuditContentItem(request *MarkAuditContentItemRequest) (response *MarkAuditContentItemResponse, err error) {
	response = CreateMarkAuditContentItemResponse()
	err = client.DoAction(request, response)
	return
}

// MarkAuditContentItemWithChan invokes the green.MarkAuditContentItem API asynchronously
// api document: https://help.aliyun.com/api/green/markauditcontentitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkAuditContentItemWithChan(request *MarkAuditContentItemRequest) (<-chan *MarkAuditContentItemResponse, <-chan error) {
	responseChan := make(chan *MarkAuditContentItemResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.MarkAuditContentItem(request)
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

// MarkAuditContentItemWithCallback invokes the green.MarkAuditContentItem API asynchronously
// api document: https://help.aliyun.com/api/green/markauditcontentitem.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) MarkAuditContentItemWithCallback(request *MarkAuditContentItemRequest, callback func(response *MarkAuditContentItemResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *MarkAuditContentItemResponse
		var err error
		defer close(result)
		response, err = client.MarkAuditContentItem(request)
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

// MarkAuditContentItemRequest is the request struct for api MarkAuditContentItem
type MarkAuditContentItemRequest struct {
	*requests.RpcRequest
	AuditIllegalReasons string `position:"Query" name:"AuditIllegalReasons"`
	SourceIp            string `position:"Query" name:"SourceIp"`
	AuditResult         string `position:"Query" name:"AuditResult"`
	Ids                 string `position:"Query" name:"Ids"`
	Lang                string `position:"Query" name:"Lang"`
}

// MarkAuditContentItemResponse is the response struct for api MarkAuditContentItem
type MarkAuditContentItemResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateMarkAuditContentItemRequest creates a request to invoke MarkAuditContentItem API
func CreateMarkAuditContentItemRequest() (request *MarkAuditContentItemRequest) {
	request = &MarkAuditContentItemRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Green", "2017-08-23", "MarkAuditContentItem", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateMarkAuditContentItemResponse creates a response to parse from MarkAuditContentItem response
func CreateMarkAuditContentItemResponse() (response *MarkAuditContentItemResponse) {
	response = &MarkAuditContentItemResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
