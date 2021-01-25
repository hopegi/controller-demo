package cloudwf

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

// GetUserUmengPagePermission invokes the cloudwf.GetUserUmengPagePermission API synchronously
// api document: https://help.aliyun.com/api/cloudwf/getuserumengpagepermission.html
func (client *Client) GetUserUmengPagePermission(request *GetUserUmengPagePermissionRequest) (response *GetUserUmengPagePermissionResponse, err error) {
	response = CreateGetUserUmengPagePermissionResponse()
	err = client.DoAction(request, response)
	return
}

// GetUserUmengPagePermissionWithChan invokes the cloudwf.GetUserUmengPagePermission API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getuserumengpagepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserUmengPagePermissionWithChan(request *GetUserUmengPagePermissionRequest) (<-chan *GetUserUmengPagePermissionResponse, <-chan error) {
	responseChan := make(chan *GetUserUmengPagePermissionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetUserUmengPagePermission(request)
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

// GetUserUmengPagePermissionWithCallback invokes the cloudwf.GetUserUmengPagePermission API asynchronously
// api document: https://help.aliyun.com/api/cloudwf/getuserumengpagepermission.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetUserUmengPagePermissionWithCallback(request *GetUserUmengPagePermissionRequest, callback func(response *GetUserUmengPagePermissionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetUserUmengPagePermissionResponse
		var err error
		defer close(result)
		response, err = client.GetUserUmengPagePermission(request)
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

// GetUserUmengPagePermissionRequest is the request struct for api GetUserUmengPagePermission
type GetUserUmengPagePermissionRequest struct {
	*requests.RpcRequest
	Bid requests.Integer `position:"Query" name:"Bid"`
}

// GetUserUmengPagePermissionResponse is the response struct for api GetUserUmengPagePermission
type GetUserUmengPagePermissionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode int    `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMsg  string `json:"ErrorMsg" xml:"ErrorMsg"`
}

// CreateGetUserUmengPagePermissionRequest creates a request to invoke GetUserUmengPagePermission API
func CreateGetUserUmengPagePermissionRequest() (request *GetUserUmengPagePermissionRequest) {
	request = &GetUserUmengPagePermissionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudwf", "2017-03-28", "GetUserUmengPagePermission", "cloudwf", "openAPI")
	return
}

// CreateGetUserUmengPagePermissionResponse creates a response to parse from GetUserUmengPagePermission response
func CreateGetUserUmengPagePermissionResponse() (response *GetUserUmengPagePermissionResponse) {
	response = &GetUserUmengPagePermissionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
