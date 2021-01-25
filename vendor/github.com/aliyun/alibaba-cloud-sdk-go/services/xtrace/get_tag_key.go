package xtrace

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

// GetTagKey invokes the xtrace.GetTagKey API synchronously
// api document: https://help.aliyun.com/api/xtrace/gettagkey.html
func (client *Client) GetTagKey(request *GetTagKeyRequest) (response *GetTagKeyResponse, err error) {
	response = CreateGetTagKeyResponse()
	err = client.DoAction(request, response)
	return
}

// GetTagKeyWithChan invokes the xtrace.GetTagKey API asynchronously
// api document: https://help.aliyun.com/api/xtrace/gettagkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagKeyWithChan(request *GetTagKeyRequest) (<-chan *GetTagKeyResponse, <-chan error) {
	responseChan := make(chan *GetTagKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetTagKey(request)
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

// GetTagKeyWithCallback invokes the xtrace.GetTagKey API asynchronously
// api document: https://help.aliyun.com/api/xtrace/gettagkey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetTagKeyWithCallback(request *GetTagKeyRequest, callback func(response *GetTagKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetTagKeyResponse
		var err error
		defer close(result)
		response, err = client.GetTagKey(request)
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

// GetTagKeyRequest is the request struct for api GetTagKey
type GetTagKeyRequest struct {
	*requests.RpcRequest
	EndTime     requests.Integer `position:"Query" name:"EndTime"`
	ServiceName string           `position:"Query" name:"ServiceName"`
	StartTime   requests.Integer `position:"Query" name:"StartTime"`
	SpanName    string           `position:"Query" name:"SpanName"`
}

// GetTagKeyResponse is the response struct for api GetTagKey
type GetTagKeyResponse struct {
	*responses.BaseResponse
	RequestId string  `json:"RequestId" xml:"RequestId"`
	TagKeys   TagKeys `json:"TagKeys" xml:"TagKeys"`
}

// CreateGetTagKeyRequest creates a request to invoke GetTagKey API
func CreateGetTagKeyRequest() (request *GetTagKeyRequest) {
	request = &GetTagKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("xtrace", "2019-08-08", "GetTagKey", "", "")
	request.Method = requests.POST
	return
}

// CreateGetTagKeyResponse creates a response to parse from GetTagKey response
func CreateGetTagKeyResponse() (response *GetTagKeyResponse) {
	response = &GetTagKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
