package dms_enterprise

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

// ListWorkFlowNodes invokes the dms_enterprise.ListWorkFlowNodes API synchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listworkflownodes.html
func (client *Client) ListWorkFlowNodes(request *ListWorkFlowNodesRequest) (response *ListWorkFlowNodesResponse, err error) {
	response = CreateListWorkFlowNodesResponse()
	err = client.DoAction(request, response)
	return
}

// ListWorkFlowNodesWithChan invokes the dms_enterprise.ListWorkFlowNodes API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listworkflownodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListWorkFlowNodesWithChan(request *ListWorkFlowNodesRequest) (<-chan *ListWorkFlowNodesResponse, <-chan error) {
	responseChan := make(chan *ListWorkFlowNodesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListWorkFlowNodes(request)
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

// ListWorkFlowNodesWithCallback invokes the dms_enterprise.ListWorkFlowNodes API asynchronously
// api document: https://help.aliyun.com/api/dms-enterprise/listworkflownodes.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListWorkFlowNodesWithCallback(request *ListWorkFlowNodesRequest, callback func(response *ListWorkFlowNodesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListWorkFlowNodesResponse
		var err error
		defer close(result)
		response, err = client.ListWorkFlowNodes(request)
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

// ListWorkFlowNodesRequest is the request struct for api ListWorkFlowNodes
type ListWorkFlowNodesRequest struct {
	*requests.RpcRequest
	SearchName string           `position:"Query" name:"SearchName"`
	Tid        requests.Integer `position:"Query" name:"Tid"`
}

// ListWorkFlowNodesResponse is the response struct for api ListWorkFlowNodes
type ListWorkFlowNodesResponse struct {
	*responses.BaseResponse
	RequestId     string                           `json:"RequestId" xml:"RequestId"`
	Success       bool                             `json:"Success" xml:"Success"`
	ErrorMessage  string                           `json:"ErrorMessage" xml:"ErrorMessage"`
	ErrorCode     string                           `json:"ErrorCode" xml:"ErrorCode"`
	WorkFlowNodes WorkFlowNodesInListWorkFlowNodes `json:"WorkFlowNodes" xml:"WorkFlowNodes"`
}

// CreateListWorkFlowNodesRequest creates a request to invoke ListWorkFlowNodes API
func CreateListWorkFlowNodesRequest() (request *ListWorkFlowNodesRequest) {
	request = &ListWorkFlowNodesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dms-enterprise", "2018-11-01", "ListWorkFlowNodes", "dmsenterprise", "openAPI")
	return
}

// CreateListWorkFlowNodesResponse creates a response to parse from ListWorkFlowNodes response
func CreateListWorkFlowNodesResponse() (response *ListWorkFlowNodesResponse) {
	response = &ListWorkFlowNodesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
