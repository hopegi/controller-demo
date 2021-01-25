package live

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

// RealTimeSnapshotCommand invokes the live.RealTimeSnapshotCommand API synchronously
func (client *Client) RealTimeSnapshotCommand(request *RealTimeSnapshotCommandRequest) (response *RealTimeSnapshotCommandResponse, err error) {
	response = CreateRealTimeSnapshotCommandResponse()
	err = client.DoAction(request, response)
	return
}

// RealTimeSnapshotCommandWithChan invokes the live.RealTimeSnapshotCommand API asynchronously
func (client *Client) RealTimeSnapshotCommandWithChan(request *RealTimeSnapshotCommandRequest) (<-chan *RealTimeSnapshotCommandResponse, <-chan error) {
	responseChan := make(chan *RealTimeSnapshotCommandResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RealTimeSnapshotCommand(request)
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

// RealTimeSnapshotCommandWithCallback invokes the live.RealTimeSnapshotCommand API asynchronously
func (client *Client) RealTimeSnapshotCommandWithCallback(request *RealTimeSnapshotCommandRequest, callback func(response *RealTimeSnapshotCommandResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RealTimeSnapshotCommandResponse
		var err error
		defer close(result)
		response, err = client.RealTimeSnapshotCommand(request)
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

// RealTimeSnapshotCommandRequest is the request struct for api RealTimeSnapshotCommand
type RealTimeSnapshotCommandRequest struct {
	*requests.RpcRequest
	Mode       requests.Integer `position:"Query" name:"Mode"`
	AppName    string           `position:"Query" name:"AppName"`
	StreamName string           `position:"Query" name:"StreamName"`
	DomainName string           `position:"Query" name:"DomainName"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
	Command    string           `position:"Query" name:"Command"`
	Interval   requests.Integer `position:"Query" name:"Interval"`
}

// RealTimeSnapshotCommandResponse is the response struct for api RealTimeSnapshotCommand
type RealTimeSnapshotCommandResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateRealTimeSnapshotCommandRequest creates a request to invoke RealTimeSnapshotCommand API
func CreateRealTimeSnapshotCommandRequest() (request *RealTimeSnapshotCommandRequest) {
	request = &RealTimeSnapshotCommandRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "RealTimeSnapshotCommand", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRealTimeSnapshotCommandResponse creates a response to parse from RealTimeSnapshotCommand response
func CreateRealTimeSnapshotCommandResponse() (response *RealTimeSnapshotCommandResponse) {
	response = &RealTimeSnapshotCommandResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
