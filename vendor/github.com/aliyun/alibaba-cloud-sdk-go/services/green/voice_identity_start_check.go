package green

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

// VoiceIdentityStartCheck invokes the green.VoiceIdentityStartCheck API synchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartcheck.html
func (client *Client) VoiceIdentityStartCheck(request *VoiceIdentityStartCheckRequest) (response *VoiceIdentityStartCheckResponse, err error) {
	response = CreateVoiceIdentityStartCheckResponse()
	err = client.DoAction(request, response)
	return
}

// VoiceIdentityStartCheckWithChan invokes the green.VoiceIdentityStartCheck API asynchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoiceIdentityStartCheckWithChan(request *VoiceIdentityStartCheckRequest) (<-chan *VoiceIdentityStartCheckResponse, <-chan error) {
	responseChan := make(chan *VoiceIdentityStartCheckResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.VoiceIdentityStartCheck(request)
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

// VoiceIdentityStartCheckWithCallback invokes the green.VoiceIdentityStartCheck API asynchronously
// api document: https://help.aliyun.com/api/green/voiceidentitystartcheck.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) VoiceIdentityStartCheckWithCallback(request *VoiceIdentityStartCheckRequest, callback func(response *VoiceIdentityStartCheckResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *VoiceIdentityStartCheckResponse
		var err error
		defer close(result)
		response, err = client.VoiceIdentityStartCheck(request)
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

// VoiceIdentityStartCheckRequest is the request struct for api VoiceIdentityStartCheck
type VoiceIdentityStartCheckRequest struct {
	*requests.RoaRequest
	ClientInfo string `position:"Query" name:"ClientInfo"`
}

// VoiceIdentityStartCheckResponse is the response struct for api VoiceIdentityStartCheck
type VoiceIdentityStartCheckResponse struct {
	*responses.BaseResponse
}

// CreateVoiceIdentityStartCheckRequest creates a request to invoke VoiceIdentityStartCheck API
func CreateVoiceIdentityStartCheckRequest() (request *VoiceIdentityStartCheckRequest) {
	request = &VoiceIdentityStartCheckRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Green", "2018-05-09", "VoiceIdentityStartCheck", "/green/voice/auth/start/check", "green", "openAPI")
	request.Method = requests.POST
	return
}

// CreateVoiceIdentityStartCheckResponse creates a response to parse from VoiceIdentityStartCheck response
func CreateVoiceIdentityStartCheckResponse() (response *VoiceIdentityStartCheckResponse) {
	response = &VoiceIdentityStartCheckResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
