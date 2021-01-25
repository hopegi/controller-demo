package cloudcallcenter

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

// SilenceTimeoutVn invokes the cloudcallcenter.SilenceTimeoutVn API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/silencetimeoutvn.html
func (client *Client) SilenceTimeoutVn(request *SilenceTimeoutVnRequest) (response *SilenceTimeoutVnResponse, err error) {
	response = CreateSilenceTimeoutVnResponse()
	err = client.DoAction(request, response)
	return
}

// SilenceTimeoutVnWithChan invokes the cloudcallcenter.SilenceTimeoutVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/silencetimeoutvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SilenceTimeoutVnWithChan(request *SilenceTimeoutVnRequest) (<-chan *SilenceTimeoutVnResponse, <-chan error) {
	responseChan := make(chan *SilenceTimeoutVnResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SilenceTimeoutVn(request)
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

// SilenceTimeoutVnWithCallback invokes the cloudcallcenter.SilenceTimeoutVn API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/silencetimeoutvn.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) SilenceTimeoutVnWithCallback(request *SilenceTimeoutVnRequest, callback func(response *SilenceTimeoutVnResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SilenceTimeoutVnResponse
		var err error
		defer close(result)
		response, err = client.SilenceTimeoutVn(request)
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

// SilenceTimeoutVnRequest is the request struct for api SilenceTimeoutVn
type SilenceTimeoutVnRequest struct {
	*requests.RpcRequest
	ConversationId string `position:"Query" name:"ConversationId"`
	InitialContext string `position:"Query" name:"InitialContext"`
	InstanceId     string `position:"Query" name:"InstanceId"`
}

// SilenceTimeoutVnResponse is the response struct for api SilenceTimeoutVn
type SilenceTimeoutVnResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	TextResponse  string `json:"TextResponse" xml:"TextResponse"`
	Interruptible bool   `json:"Interruptible" xml:"Interruptible"`
	Action        string `json:"Action" xml:"Action"`
	ActionParams  string `json:"ActionParams" xml:"ActionParams"`
}

// CreateSilenceTimeoutVnRequest creates a request to invoke SilenceTimeoutVn API
func CreateSilenceTimeoutVnRequest() (request *SilenceTimeoutVnRequest) {
	request = &SilenceTimeoutVnRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "SilenceTimeoutVn", "", "")
	request.Method = requests.GET
	return
}

// CreateSilenceTimeoutVnResponse creates a response to parse from SilenceTimeoutVn response
func CreateSilenceTimeoutVnResponse() (response *SilenceTimeoutVnResponse) {
	response = &SilenceTimeoutVnResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
