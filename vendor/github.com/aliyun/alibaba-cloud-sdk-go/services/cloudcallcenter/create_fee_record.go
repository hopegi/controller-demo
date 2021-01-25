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

// CreateFeeRecord invokes the cloudcallcenter.CreateFeeRecord API synchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfeerecord.html
func (client *Client) CreateFeeRecord(request *CreateFeeRecordRequest) (response *CreateFeeRecordResponse, err error) {
	response = CreateCreateFeeRecordResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFeeRecordWithChan invokes the cloudcallcenter.CreateFeeRecord API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfeerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFeeRecordWithChan(request *CreateFeeRecordRequest) (<-chan *CreateFeeRecordResponse, <-chan error) {
	responseChan := make(chan *CreateFeeRecordResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFeeRecord(request)
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

// CreateFeeRecordWithCallback invokes the cloudcallcenter.CreateFeeRecord API asynchronously
// api document: https://help.aliyun.com/api/cloudcallcenter/createfeerecord.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateFeeRecordWithCallback(request *CreateFeeRecordRequest, callback func(response *CreateFeeRecordResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFeeRecordResponse
		var err error
		defer close(result)
		response, err = client.CreateFeeRecord(request)
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

// CreateFeeRecordRequest is the request struct for api CreateFeeRecord
type CreateFeeRecordRequest struct {
	*requests.RpcRequest
	TaobaoMainUid requests.Integer `position:"Query" name:"taobaoMainUid"`
	Callee        string           `position:"Query" name:"callee"`
	EndTime       requests.Integer `position:"Query" name:"endTime"`
	StartTime     requests.Integer `position:"Query" name:"startTime"`
	TaobaoUid     requests.Integer `position:"Query" name:"taobaoUid"`
	MainRamId     requests.Integer `position:"Query" name:"mainRamId"`
	RamId         requests.Integer `position:"Query" name:"ramId"`
	StatusCode    string           `position:"Query" name:"statusCode"`
	Duration      requests.Integer `position:"Query" name:"duration"`
	Caller        string           `position:"Query" name:"caller"`
	BizId         string           `position:"Query" name:"bizId"`
	OrigDuration  requests.Integer `position:"Query" name:"origDuration"`
}

// CreateFeeRecordResponse is the response struct for api CreateFeeRecord
type CreateFeeRecordResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	Message        string `json:"Message" xml:"Message"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateCreateFeeRecordRequest creates a request to invoke CreateFeeRecord API
func CreateCreateFeeRecordRequest() (request *CreateFeeRecordRequest) {
	request = &CreateFeeRecordRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("CloudCallCenter", "2017-07-05", "CreateFeeRecord", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateFeeRecordResponse creates a response to parse from CreateFeeRecord response
func CreateCreateFeeRecordResponse() (response *CreateFeeRecordResponse) {
	response = &CreateFeeRecordResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
