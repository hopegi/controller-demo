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

// AddLiveAudioAuditConfig invokes the live.AddLiveAudioAuditConfig API synchronously
func (client *Client) AddLiveAudioAuditConfig(request *AddLiveAudioAuditConfigRequest) (response *AddLiveAudioAuditConfigResponse, err error) {
	response = CreateAddLiveAudioAuditConfigResponse()
	err = client.DoAction(request, response)
	return
}

// AddLiveAudioAuditConfigWithChan invokes the live.AddLiveAudioAuditConfig API asynchronously
func (client *Client) AddLiveAudioAuditConfigWithChan(request *AddLiveAudioAuditConfigRequest) (<-chan *AddLiveAudioAuditConfigResponse, <-chan error) {
	responseChan := make(chan *AddLiveAudioAuditConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AddLiveAudioAuditConfig(request)
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

// AddLiveAudioAuditConfigWithCallback invokes the live.AddLiveAudioAuditConfig API asynchronously
func (client *Client) AddLiveAudioAuditConfigWithCallback(request *AddLiveAudioAuditConfigRequest, callback func(response *AddLiveAudioAuditConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AddLiveAudioAuditConfigResponse
		var err error
		defer close(result)
		response, err = client.AddLiveAudioAuditConfig(request)
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

// AddLiveAudioAuditConfigRequest is the request struct for api AddLiveAudioAuditConfig
type AddLiveAudioAuditConfigRequest struct {
	*requests.RpcRequest
	OssEndpoint string           `position:"Query" name:"OssEndpoint"`
	OssObject   string           `position:"Query" name:"OssObject"`
	AppName     string           `position:"Query" name:"AppName"`
	StreamName  string           `position:"Query" name:"StreamName"`
	OssBucket   string           `position:"Query" name:"OssBucket"`
	DomainName  string           `position:"Query" name:"DomainName"`
	OwnerId     requests.Integer `position:"Query" name:"OwnerId"`
	BizType     string           `position:"Query" name:"BizType"`
}

// AddLiveAudioAuditConfigResponse is the response struct for api AddLiveAudioAuditConfig
type AddLiveAudioAuditConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAddLiveAudioAuditConfigRequest creates a request to invoke AddLiveAudioAuditConfig API
func CreateAddLiveAudioAuditConfigRequest() (request *AddLiveAudioAuditConfigRequest) {
	request = &AddLiveAudioAuditConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("live", "2016-11-01", "AddLiveAudioAuditConfig", "live", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAddLiveAudioAuditConfigResponse creates a response to parse from AddLiveAudioAuditConfig response
func CreateAddLiveAudioAuditConfigResponse() (response *AddLiveAudioAuditConfigResponse) {
	response = &AddLiveAudioAuditConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
