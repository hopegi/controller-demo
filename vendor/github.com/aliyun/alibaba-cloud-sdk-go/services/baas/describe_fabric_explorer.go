package baas

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

// DescribeFabricExplorer invokes the baas.DescribeFabricExplorer API synchronously
// api document: https://help.aliyun.com/api/baas/describefabricexplorer.html
func (client *Client) DescribeFabricExplorer(request *DescribeFabricExplorerRequest) (response *DescribeFabricExplorerResponse, err error) {
	response = CreateDescribeFabricExplorerResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeFabricExplorerWithChan invokes the baas.DescribeFabricExplorer API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricexplorer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricExplorerWithChan(request *DescribeFabricExplorerRequest) (<-chan *DescribeFabricExplorerResponse, <-chan error) {
	responseChan := make(chan *DescribeFabricExplorerResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeFabricExplorer(request)
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

// DescribeFabricExplorerWithCallback invokes the baas.DescribeFabricExplorer API asynchronously
// api document: https://help.aliyun.com/api/baas/describefabricexplorer.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeFabricExplorerWithCallback(request *DescribeFabricExplorerRequest, callback func(response *DescribeFabricExplorerResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeFabricExplorerResponse
		var err error
		defer close(result)
		response, err = client.DescribeFabricExplorer(request)
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

// DescribeFabricExplorerRequest is the request struct for api DescribeFabricExplorer
type DescribeFabricExplorerRequest struct {
	*requests.RpcRequest
	ExUrl          string `position:"Query" name:"ExUrl"`
	ExMethod       string `position:"Query" name:"ExMethod"`
	OrganizationId string `position:"Body" name:"OrganizationId"`
	ExBody         string `position:"Query" name:"ExBody"`
}

// DescribeFabricExplorerResponse is the response struct for api DescribeFabricExplorer
type DescribeFabricExplorerResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      int    `json:"ErrorCode" xml:"ErrorCode"`
	Result         string `json:"Result" xml:"Result"`
	DynamicCode    string `json:"DynamicCode" xml:"DynamicCode"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
}

// CreateDescribeFabricExplorerRequest creates a request to invoke DescribeFabricExplorer API
func CreateDescribeFabricExplorerRequest() (request *DescribeFabricExplorerRequest) {
	request = &DescribeFabricExplorerRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Baas", "2018-12-21", "DescribeFabricExplorer", "baas", "openAPI")
	return
}

// CreateDescribeFabricExplorerResponse creates a response to parse from DescribeFabricExplorer response
func CreateDescribeFabricExplorerResponse() (response *DescribeFabricExplorerResponse) {
	response = &DescribeFabricExplorerResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
