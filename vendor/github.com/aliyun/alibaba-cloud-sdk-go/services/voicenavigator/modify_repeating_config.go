package voicenavigator

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

// ModifyRepeatingConfig invokes the voicenavigator.ModifyRepeatingConfig API synchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyrepeatingconfig.html
func (client *Client) ModifyRepeatingConfig(request *ModifyRepeatingConfigRequest) (response *ModifyRepeatingConfigResponse, err error) {
	response = CreateModifyRepeatingConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ModifyRepeatingConfigWithChan invokes the voicenavigator.ModifyRepeatingConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyrepeatingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRepeatingConfigWithChan(request *ModifyRepeatingConfigRequest) (<-chan *ModifyRepeatingConfigResponse, <-chan error) {
	responseChan := make(chan *ModifyRepeatingConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ModifyRepeatingConfig(request)
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

// ModifyRepeatingConfigWithCallback invokes the voicenavigator.ModifyRepeatingConfig API asynchronously
// api document: https://help.aliyun.com/api/voicenavigator/modifyrepeatingconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ModifyRepeatingConfigWithCallback(request *ModifyRepeatingConfigRequest, callback func(response *ModifyRepeatingConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ModifyRepeatingConfigResponse
		var err error
		defer close(result)
		response, err = client.ModifyRepeatingConfig(request)
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

// ModifyRepeatingConfigRequest is the request struct for api ModifyRepeatingConfig
type ModifyRepeatingConfigRequest struct {
	*requests.RpcRequest
	Utterances *[]string `position:"Query" name:"Utterances"  type:"Repeated"`
	InstanceId string    `position:"Query" name:"InstanceId"`
}

// ModifyRepeatingConfigResponse is the response struct for api ModifyRepeatingConfig
type ModifyRepeatingConfigResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateModifyRepeatingConfigRequest creates a request to invoke ModifyRepeatingConfig API
func CreateModifyRepeatingConfigRequest() (request *ModifyRepeatingConfigRequest) {
	request = &ModifyRepeatingConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("VoiceNavigator", "2018-06-12", "ModifyRepeatingConfig", "voicebot", "openAPI")
	return
}

// CreateModifyRepeatingConfigResponse creates a response to parse from ModifyRepeatingConfig response
func CreateModifyRepeatingConfigResponse() (response *ModifyRepeatingConfigResponse) {
	response = &ModifyRepeatingConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
