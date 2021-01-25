package dbs

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

// DescribeBackupPlanList invokes the dbs.DescribeBackupPlanList API synchronously
func (client *Client) DescribeBackupPlanList(request *DescribeBackupPlanListRequest) (response *DescribeBackupPlanListResponse, err error) {
	response = CreateDescribeBackupPlanListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupPlanListWithChan invokes the dbs.DescribeBackupPlanList API asynchronously
func (client *Client) DescribeBackupPlanListWithChan(request *DescribeBackupPlanListRequest) (<-chan *DescribeBackupPlanListResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupPlanListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupPlanList(request)
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

// DescribeBackupPlanListWithCallback invokes the dbs.DescribeBackupPlanList API asynchronously
func (client *Client) DescribeBackupPlanListWithCallback(request *DescribeBackupPlanListRequest, callback func(response *DescribeBackupPlanListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupPlanListResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupPlanList(request)
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

// DescribeBackupPlanListRequest is the request struct for api DescribeBackupPlanList
type DescribeBackupPlanListRequest struct {
	*requests.RpcRequest
	ClientToken      string           `position:"Query" name:"ClientToken"`
	BackupPlanId     string           `position:"Query" name:"BackupPlanId"`
	PageNum          requests.Integer `position:"Query" name:"PageNum"`
	OwnerId          string           `position:"Query" name:"OwnerId"`
	BackupPlanStatus string           `position:"Query" name:"BackupPlanStatus"`
	BackupPlanName   string           `position:"Query" name:"BackupPlanName"`
	PageSize         requests.Integer `position:"Query" name:"PageSize"`
	Region           string           `position:"Query" name:"Region"`
}

// DescribeBackupPlanListResponse is the response struct for api DescribeBackupPlanList
type DescribeBackupPlanListResponse struct {
	*responses.BaseResponse
	Success        bool                          `json:"Success" xml:"Success"`
	ErrCode        string                        `json:"ErrCode" xml:"ErrCode"`
	ErrMessage     string                        `json:"ErrMessage" xml:"ErrMessage"`
	HttpStatusCode int                           `json:"HttpStatusCode" xml:"HttpStatusCode"`
	RequestId      string                        `json:"RequestId" xml:"RequestId"`
	TotalPages     int                           `json:"TotalPages" xml:"TotalPages"`
	PageSize       int                           `json:"PageSize" xml:"PageSize"`
	PageNum        int                           `json:"PageNum" xml:"PageNum"`
	TotalElements  int                           `json:"TotalElements" xml:"TotalElements"`
	Items          ItemsInDescribeBackupPlanList `json:"Items" xml:"Items"`
}

// CreateDescribeBackupPlanListRequest creates a request to invoke DescribeBackupPlanList API
func CreateDescribeBackupPlanListRequest() (request *DescribeBackupPlanListRequest) {
	request = &DescribeBackupPlanListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeBackupPlanList", "cbs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupPlanListResponse creates a response to parse from DescribeBackupPlanList response
func CreateDescribeBackupPlanListResponse() (response *DescribeBackupPlanListResponse) {
	response = &DescribeBackupPlanListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
