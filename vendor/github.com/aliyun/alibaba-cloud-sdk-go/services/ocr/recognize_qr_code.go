package ocr

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

// RecognizeQrCode invokes the ocr.RecognizeQrCode API synchronously
// api document: https://help.aliyun.com/api/ocr/recognizeqrcode.html
func (client *Client) RecognizeQrCode(request *RecognizeQrCodeRequest) (response *RecognizeQrCodeResponse, err error) {
	response = CreateRecognizeQrCodeResponse()
	err = client.DoAction(request, response)
	return
}

// RecognizeQrCodeWithChan invokes the ocr.RecognizeQrCode API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizeqrcode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeQrCodeWithChan(request *RecognizeQrCodeRequest) (<-chan *RecognizeQrCodeResponse, <-chan error) {
	responseChan := make(chan *RecognizeQrCodeResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.RecognizeQrCode(request)
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

// RecognizeQrCodeWithCallback invokes the ocr.RecognizeQrCode API asynchronously
// api document: https://help.aliyun.com/api/ocr/recognizeqrcode.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) RecognizeQrCodeWithCallback(request *RecognizeQrCodeRequest, callback func(response *RecognizeQrCodeResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *RecognizeQrCodeResponse
		var err error
		defer close(result)
		response, err = client.RecognizeQrCode(request)
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

// RecognizeQrCodeRequest is the request struct for api RecognizeQrCode
type RecognizeQrCodeRequest struct {
	*requests.RpcRequest
	Tasks *[]RecognizeQrCodeTasks `position:"Body" name:"Tasks"  type:"Repeated"`
}

// RecognizeQrCodeTasks is a repeated param struct in RecognizeQrCodeRequest
type RecognizeQrCodeTasks struct {
	ImageURL string `name:"ImageURL"`
}

// RecognizeQrCodeResponse is the response struct for api RecognizeQrCode
type RecognizeQrCodeResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateRecognizeQrCodeRequest creates a request to invoke RecognizeQrCode API
func CreateRecognizeQrCodeRequest() (request *RecognizeQrCodeRequest) {
	request = &RecognizeQrCodeRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ocr", "2019-12-30", "RecognizeQrCode", "ocr", "openAPI")
	request.Method = requests.POST
	return
}

// CreateRecognizeQrCodeResponse creates a response to parse from RecognizeQrCode response
func CreateRecognizeQrCodeResponse() (response *RecognizeQrCodeResponse) {
	response = &RecognizeQrCodeResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
