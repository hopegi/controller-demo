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

// ListOverviewRealTimeReport invokes the cloudcallcenter.ListOverviewRealTimeReport API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewrealtimereport.html
func (client *Client) ListOverviewRealTimeReport(request *ListOverviewRealTimeReportRequest) (response *ListOverviewRealTimeReportResponse, err error) {
	response = CreateListOverviewRealTimeReportResponse()
	err = client.DoAction(request, response)
	return
}

// ListOverviewRealTimeReportWithChan invokes the cloudcallcenter.ListOverviewRealTimeReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewrealtimereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOverviewRealTimeReportWithChan(request *ListOverviewRealTimeReportRequest) (<-chan *ListOverviewRealTimeReportResponse, <-chan error) {
	responseChan := make(chan *ListOverviewRealTimeReportResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListOverviewRealTimeReport(request)
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

// ListOverviewRealTimeReportWithCallback invokes the cloudcallcenter.ListOverviewRealTimeReport API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/listoverviewrealtimereport.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListOverviewRealTimeReportWithCallback(request *ListOverviewRealTimeReportRequest, callback func(response *ListOverviewRealTimeReportResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListOverviewRealTimeReportResponse
		var err error
		defer close(result)
		response, err = client.ListOverviewRealTimeReport(request)
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

// ListOverviewRealTimeReportRequest is the request struct for api ListOverviewRealTimeReport
type ListOverviewRealTimeReportRequest struct {
	*requests.RpcRequest
	InstanceId string `position:"Query" name:"InstanceId"`
	Dimension  string `position:"Query" name:"Dimension"`
}

// ListOverviewRealTimeReportResponse is the response struct for api ListOverviewRealTimeReport
type ListOverviewRealTimeReportResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Data           string `json:"Data" xml:"Data"`
}

// CreateListOverviewRealTimeReportRequest creates a request to invoke ListOverviewRealTimeReport API
func CreateListOverviewRealTimeReportRequest() (request *ListOverviewRealTimeReportRequest) {
	request = &ListOverviewRealTimeReportRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "ListOverviewRealTimeReport", "", "")
	request.Method = requests.POST
	return
}

// CreateListOverviewRealTimeReportResponse creates a response to parse from ListOverviewRealTimeReport response
func CreateListOverviewRealTimeReportResponse() (response *ListOverviewRealTimeReportResponse) {
	response = &ListOverviewRealTimeReportResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
