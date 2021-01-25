package alimt

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

// GetDetectLanguage invokes the alimt.GetDetectLanguage API synchronously
func (client *Client) GetDetectLanguage(request *GetDetectLanguageRequest) (response *GetDetectLanguageResponse, err error) {
	response = CreateGetDetectLanguageResponse()
	err = client.DoAction(request, response)
	return
}

// GetDetectLanguageWithChan invokes the alimt.GetDetectLanguage API asynchronously
func (client *Client) GetDetectLanguageWithChan(request *GetDetectLanguageRequest) (<-chan *GetDetectLanguageResponse, <-chan error) {
	responseChan := make(chan *GetDetectLanguageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDetectLanguage(request)
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

// GetDetectLanguageWithCallback invokes the alimt.GetDetectLanguage API asynchronously
func (client *Client) GetDetectLanguageWithCallback(request *GetDetectLanguageRequest, callback func(response *GetDetectLanguageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDetectLanguageResponse
		var err error
		defer close(result)
		response, err = client.GetDetectLanguage(request)
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

// GetDetectLanguageRequest is the request struct for api GetDetectLanguage
type GetDetectLanguageRequest struct {
	*requests.RpcRequest
	SourceText string `position:"Body" name:"SourceText"`
}

// GetDetectLanguageResponse is the response struct for api GetDetectLanguage
type GetDetectLanguageResponse struct {
	*responses.BaseResponse
	RequestId        string `json:"RequestId" xml:"RequestId"`
	DetectedLanguage string `json:"DetectedLanguage" xml:"DetectedLanguage"`
}

// CreateGetDetectLanguageRequest creates a request to invoke GetDetectLanguage API
func CreateGetDetectLanguageRequest() (request *GetDetectLanguageRequest) {
	request = &GetDetectLanguageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alimt", "2018-10-12", "GetDetectLanguage", "alimt", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDetectLanguageResponse creates a response to parse from GetDetectLanguage response
func CreateGetDetectLanguageResponse() (response *GetDetectLanguageResponse) {
	response = &GetDetectLanguageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
