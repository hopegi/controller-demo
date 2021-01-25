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

// DeleteBusiness invokes the dataworks_public.DeleteBusiness API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletebusiness.html
func (client *Client) DeleteBusiness(request *DeleteBusinessRequest) (response *DeleteBusinessResponse, err error) {
	response = CreateDeleteBusinessResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteBusinessWithChan invokes the dataworks_public.DeleteBusiness API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletebusiness.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBusinessWithChan(request *DeleteBusinessRequest) (<-chan *DeleteBusinessResponse, <-chan error) {
	responseChan := make(chan *DeleteBusinessResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteBusiness(request)
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

// DeleteBusinessWithCallback invokes the dataworks_public.DeleteBusiness API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/deletebusiness.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteBusinessWithCallback(request *DeleteBusinessRequest, callback func(response *DeleteBusinessResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteBusinessResponse
		var err error
		defer close(result)
		response, err = client.DeleteBusiness(request)
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

// DeleteBusinessRequest is the request struct for api DeleteBusiness
type DeleteBusinessRequest struct {
	*requests.RpcRequest
	BusinessId        requests.Integer `position:"Body" name:"BusinessId"`
	ProjectId         requests.Integer `position:"Body" name:"ProjectId"`
	ProjectIdentifier string           `position:"Body" name:"ProjectIdentifier"`
}

// DeleteBusinessResponse is the response struct for api DeleteBusiness
type DeleteBusinessResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateDeleteBusinessRequest creates a request to invoke DeleteBusiness API
func CreateDeleteBusinessRequest() (request *DeleteBusinessRequest) {
	request = &DeleteBusinessRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "DeleteBusiness", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteBusinessResponse creates a response to parse from DeleteBusiness response
func CreateDeleteBusinessResponse() (response *DeleteBusinessResponse) {
	response = &DeleteBusinessResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
