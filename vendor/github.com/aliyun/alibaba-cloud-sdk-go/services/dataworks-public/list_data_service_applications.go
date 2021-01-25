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

// ListDataServiceApplications invokes the dataworks_public.ListDataServiceApplications API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/listdataserviceapplications.html
func (client *Client) ListDataServiceApplications(request *ListDataServiceApplicationsRequest) (response *ListDataServiceApplicationsResponse, err error) {
	response = CreateListDataServiceApplicationsResponse()
	err = client.DoAction(request, response)
	return
}

// ListDataServiceApplicationsWithChan invokes the dataworks_public.ListDataServiceApplications API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listdataserviceapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDataServiceApplicationsWithChan(request *ListDataServiceApplicationsRequest) (<-chan *ListDataServiceApplicationsResponse, <-chan error) {
	responseChan := make(chan *ListDataServiceApplicationsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListDataServiceApplications(request)
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

// ListDataServiceApplicationsWithCallback invokes the dataworks_public.ListDataServiceApplications API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/listdataserviceapplications.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListDataServiceApplicationsWithCallback(request *ListDataServiceApplicationsRequest, callback func(response *ListDataServiceApplicationsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListDataServiceApplicationsResponse
		var err error
		defer close(result)
		response, err = client.ListDataServiceApplications(request)
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

// ListDataServiceApplicationsRequest is the request struct for api ListDataServiceApplications
type ListDataServiceApplicationsRequest struct {
	*requests.RpcRequest
	PageSize      requests.Integer `position:"Body" name:"PageSize"`
	TenantId      requests.Integer `position:"Body" name:"TenantId"`
	ProjectIdList string           `position:"Body" name:"ProjectIdList"`
	PageNumber    requests.Integer `position:"Body" name:"PageNumber"`
}

// ListDataServiceApplicationsResponse is the response struct for api ListDataServiceApplications
type ListDataServiceApplicationsResponse struct {
	*responses.BaseResponse
	ErrorCode      string                            `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string                            `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int                               `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                            `json:"RequestId" xml:"RequestId"`
	Success        bool                              `json:"Success" xml:"Success"`
	Data           DataInListDataServiceApplications `json:"Data" xml:"Data"`
}

// CreateListDataServiceApplicationsRequest creates a request to invoke ListDataServiceApplications API
func CreateListDataServiceApplicationsRequest() (request *ListDataServiceApplicationsRequest) {
	request = &ListDataServiceApplicationsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "ListDataServiceApplications", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListDataServiceApplicationsResponse creates a response to parse from ListDataServiceApplications response
func CreateListDataServiceApplicationsResponse() (response *ListDataServiceApplicationsResponse) {
	response = &ListDataServiceApplicationsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
