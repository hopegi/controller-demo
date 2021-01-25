package rtc

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

// StopRecordTask invokes the rtc.StopRecordTask API synchronously
func (client *Client) StopRecordTask(request *StopRecordTaskRequest) (response *StopRecordTaskResponse, err error) {
	response = CreateStopRecordTaskResponse()
	err = client.DoAction(request, response)
	return
}

// StopRecordTaskWithChan invokes the rtc.StopRecordTask API asynchronously
func (client *Client) StopRecordTaskWithChan(request *StopRecordTaskRequest) (<-chan *StopRecordTaskResponse, <-chan error) {
	responseChan := make(chan *StopRecordTaskResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopRecordTask(request)
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

// StopRecordTaskWithCallback invokes the rtc.StopRecordTask API asynchronously
func (client *Client) StopRecordTaskWithCallback(request *StopRecordTaskRequest, callback func(response *StopRecordTaskResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopRecordTaskResponse
		var err error
		defer close(result)
		response, err = client.StopRecordTask(request)
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

// StopRecordTaskRequest is the request struct for api StopRecordTask
type StopRecordTaskRequest struct {
	*requests.RpcRequest
	TaskId  string           `position:"Query" name:"TaskId"`
	ShowLog string           `position:"Query" name:"ShowLog"`
	OwnerId requests.Integer `position:"Query" name:"OwnerId"`
	AppId   string           `position:"Query" name:"AppId"`
}

// StopRecordTaskResponse is the response struct for api StopRecordTask
type StopRecordTaskResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopRecordTaskRequest creates a request to invoke StopRecordTask API
func CreateStopRecordTaskRequest() (request *StopRecordTaskRequest) {
	request = &StopRecordTaskRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "StopRecordTask", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStopRecordTaskResponse creates a response to parse from StopRecordTask response
func CreateStopRecordTaskResponse() (response *StopRecordTaskResponse) {
	response = &StopRecordTaskResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
