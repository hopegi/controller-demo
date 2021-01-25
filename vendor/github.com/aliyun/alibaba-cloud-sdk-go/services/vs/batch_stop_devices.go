package vs

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

// BatchStopDevices invokes the vs.BatchStopDevices API synchronously
func (client *Client) BatchStopDevices(request *BatchStopDevicesRequest) (response *BatchStopDevicesResponse, err error) {
	response = CreateBatchStopDevicesResponse()
	err = client.DoAction(request, response)
	return
}

// BatchStopDevicesWithChan invokes the vs.BatchStopDevices API asynchronously
func (client *Client) BatchStopDevicesWithChan(request *BatchStopDevicesRequest) (<-chan *BatchStopDevicesResponse, <-chan error) {
	responseChan := make(chan *BatchStopDevicesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BatchStopDevices(request)
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

// BatchStopDevicesWithCallback invokes the vs.BatchStopDevices API asynchronously
func (client *Client) BatchStopDevicesWithCallback(request *BatchStopDevicesRequest, callback func(response *BatchStopDevicesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BatchStopDevicesResponse
		var err error
		defer close(result)
		response, err = client.BatchStopDevices(request)
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

// BatchStopDevicesRequest is the request struct for api BatchStopDevices
type BatchStopDevicesRequest struct {
	*requests.RpcRequest
	StartTime string           `position:"Query" name:"StartTime"`
	Id        string           `position:"Query" name:"Id"`
	ShowLog   string           `position:"Query" name:"ShowLog"`
	OwnerId   requests.Integer `position:"Query" name:"OwnerId"`
}

// BatchStopDevicesResponse is the response struct for api BatchStopDevices
type BatchStopDevicesResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Results   []Result `json:"Results" xml:"Results"`
}

// CreateBatchStopDevicesRequest creates a request to invoke BatchStopDevices API
func CreateBatchStopDevicesRequest() (request *BatchStopDevicesRequest) {
	request = &BatchStopDevicesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vs", "2018-12-12", "BatchStopDevices", "vs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBatchStopDevicesResponse creates a response to parse from BatchStopDevices response
func CreateBatchStopDevicesResponse() (response *BatchStopDevicesResponse) {
	response = &BatchStopDevicesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
