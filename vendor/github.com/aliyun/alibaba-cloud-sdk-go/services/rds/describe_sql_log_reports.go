package rds

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

// DescribeSQLLogReports invokes the rds.DescribeSQLLogReports API synchronously
// api document: https://help.aliyun.com/api/rds/describesqllogreports.html
func (client *Client) DescribeSQLLogReports(request *DescribeSQLLogReportsRequest) (response *DescribeSQLLogReportsResponse, err error) {
	response = CreateDescribeSQLLogReportsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeSQLLogReportsWithChan invokes the rds.DescribeSQLLogReports API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogReportsWithChan(request *DescribeSQLLogReportsRequest) (<-chan *DescribeSQLLogReportsResponse, <-chan error) {
	responseChan := make(chan *DescribeSQLLogReportsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeSQLLogReports(request)
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

// DescribeSQLLogReportsWithCallback invokes the rds.DescribeSQLLogReports API asynchronously
// api document: https://help.aliyun.com/api/rds/describesqllogreports.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeSQLLogReportsWithCallback(request *DescribeSQLLogReportsRequest, callback func(response *DescribeSQLLogReportsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeSQLLogReportsResponse
		var err error
		defer close(result)
		response, err = client.DescribeSQLLogReports(request)
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

// DescribeSQLLogReportsRequest is the request struct for api DescribeSQLLogReports
type DescribeSQLLogReportsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	StartTime            string           `position:"Query" name:"StartTime"`
	PageNumber           requests.Integer `position:"Query" name:"PageNumber"`
	PageSize             requests.Integer `position:"Query" name:"PageSize"`
	DBInstanceId         string           `position:"Query" name:"DBInstanceId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	EndTime              string           `position:"Query" name:"EndTime"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeSQLLogReportsResponse is the response struct for api DescribeSQLLogReports
type DescribeSQLLogReportsResponse struct {
	*responses.BaseResponse
	RequestId        string                       `json:"RequestId" xml:"RequestId"`
	TotalRecordCount int                          `json:"TotalRecordCount" xml:"TotalRecordCount"`
	PageNumber       int                          `json:"PageNumber" xml:"PageNumber"`
	PageRecordCount  int                          `json:"PageRecordCount" xml:"PageRecordCount"`
	Items            ItemsInDescribeSQLLogReports `json:"Items" xml:"Items"`
}

// CreateDescribeSQLLogReportsRequest creates a request to invoke DescribeSQLLogReports API
func CreateDescribeSQLLogReportsRequest() (request *DescribeSQLLogReportsRequest) {
	request = &DescribeSQLLogReportsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DescribeSQLLogReports", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeSQLLogReportsResponse creates a response to parse from DescribeSQLLogReports response
func CreateDescribeSQLLogReportsResponse() (response *DescribeSQLLogReportsResponse) {
	response = &DescribeSQLLogReportsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
