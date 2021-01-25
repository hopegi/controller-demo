package devops_rdc

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

// BatchInsertMembers invokes the devops_rdc.BatchInsertMembers API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/batchinsertmembers.html
func (client *Client) BatchInsertMembers(request *BatchInsertMembersRequest) (response *BatchInsertMembersResponse, err error) {
	response = CreateBatchInsertMembersResponse()
	err = client.DoAction(request, response)
	return
}

// BatchInsertMembersWithChan invokes the devops_rdc.BatchInsertMembers API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/batchinsertmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchInsertMembersWithChan(request *BatchInsertMembersRequest) (<-chan *BatchInsertMembersResponse, <-chan error) {
	responseChan := make(chan *BatchInsertMembersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchInsertMembers(request)
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

// BatchInsertMembersWithCallback invokes the devops_rdc.BatchInsertMembers API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/batchinsertmembers.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BatchInsertMembersWithCallback(request *BatchInsertMembersRequest, callback func(response *BatchInsertMembersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchInsertMembersResponse
		var err error
		defer close(result)
		response, err = client.BatchInsertMembers(request)
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

// BatchInsertMembersRequest is the request struct for api BatchInsertMembers
type BatchInsertMembersRequest struct {
	*requests.RpcRequest
	Members string `position:"Body" name:"Members"`
	RealPk  string `position:"Body" name:"RealPk"`
	OrgId   string `position:"Body" name:"OrgId"`
}

// BatchInsertMembersResponse is the response struct for api BatchInsertMembers
type BatchInsertMembersResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       bool   `json:"Object" xml:"Object"`
}

// CreateBatchInsertMembersRequest creates a request to invoke BatchInsertMembers API
func CreateBatchInsertMembersRequest() (request *BatchInsertMembersRequest) {
	request = &BatchInsertMembersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "BatchInsertMembers", "", "")
	request.Method = requests.POST
	return
}

// CreateBatchInsertMembersResponse creates a response to parse from BatchInsertMembers response
func CreateBatchInsertMembersResponse() (response *BatchInsertMembersResponse) {
	response = &BatchInsertMembersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
