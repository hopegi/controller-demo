package cloudesl

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

// DeletePlanogramShelf invokes the cloudesl.DeletePlanogramShelf API synchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteplanogramshelf.html
func (client *Client) DeletePlanogramShelf(request *DeletePlanogramShelfRequest) (response *DeletePlanogramShelfResponse, err error) {
	response = CreateDeletePlanogramShelfResponse()
	err = client.DoAction(request, response)
	return
}

// DeletePlanogramShelfWithChan invokes the cloudesl.DeletePlanogramShelf API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteplanogramshelf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePlanogramShelfWithChan(request *DeletePlanogramShelfRequest) (<-chan *DeletePlanogramShelfResponse, <-chan error) {
	responseChan := make(chan *DeletePlanogramShelfResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeletePlanogramShelf(request)
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

// DeletePlanogramShelfWithCallback invokes the cloudesl.DeletePlanogramShelf API asynchronously
// api document: https://help.aliyun.com/api/cloudesl/deleteplanogramshelf.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeletePlanogramShelfWithCallback(request *DeletePlanogramShelfRequest, callback func(response *DeletePlanogramShelfResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeletePlanogramShelfResponse
		var err error
		defer close(result)
		response, err = client.DeletePlanogramShelf(request)
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

// DeletePlanogramShelfRequest is the request struct for api DeletePlanogramShelf
type DeletePlanogramShelfRequest struct {
	*requests.RpcRequest
	BeAutoRefresh requests.Boolean `position:"Body" name:"BeAutoRefresh"`
	StoreId       string           `position:"Body" name:"StoreId"`
	Shelf         string           `position:"Body" name:"Shelf"`
}

// DeletePlanogramShelfResponse is the response struct for api DeletePlanogramShelf
type DeletePlanogramShelfResponse struct {
	*responses.BaseResponse
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	Message        string `json:"Message" xml:"Message"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	Code           string `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
}

// CreateDeletePlanogramShelfRequest creates a request to invoke DeletePlanogramShelf API
func CreateDeletePlanogramShelfRequest() (request *DeletePlanogramShelfRequest) {
	request = &DeletePlanogramShelfRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("cloudesl", "2020-02-01", "DeletePlanogramShelf", "cloudesl", "openAPI")
	return
}

// CreateDeletePlanogramShelfResponse creates a response to parse from DeletePlanogramShelf response
func CreateDeletePlanogramShelfResponse() (response *DeletePlanogramShelfResponse) {
	response = &DeletePlanogramShelfResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
