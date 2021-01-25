package iot

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

// GenerateOTAUploadURL invokes the iot.GenerateOTAUploadURL API synchronously
// api document: https://help.aliyun.com/api/iot/generateotauploadurl.html
func (client *Client) GenerateOTAUploadURL(request *GenerateOTAUploadURLRequest) (response *GenerateOTAUploadURLResponse, err error) {
	response = CreateGenerateOTAUploadURLResponse()
	err = client.DoAction(request, response)
	return
}

// GenerateOTAUploadURLWithChan invokes the iot.GenerateOTAUploadURL API asynchronously
// api document: https://help.aliyun.com/api/iot/generateotauploadurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateOTAUploadURLWithChan(request *GenerateOTAUploadURLRequest) (<-chan *GenerateOTAUploadURLResponse, <-chan error) {
	responseChan := make(chan *GenerateOTAUploadURLResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GenerateOTAUploadURL(request)
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

// GenerateOTAUploadURLWithCallback invokes the iot.GenerateOTAUploadURL API asynchronously
// api document: https://help.aliyun.com/api/iot/generateotauploadurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GenerateOTAUploadURLWithCallback(request *GenerateOTAUploadURLRequest, callback func(response *GenerateOTAUploadURLResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GenerateOTAUploadURLResponse
		var err error
		defer close(result)
		response, err = client.GenerateOTAUploadURL(request)
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

// GenerateOTAUploadURLRequest is the request struct for api GenerateOTAUploadURL
type GenerateOTAUploadURLRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// GenerateOTAUploadURLResponse is the response struct for api GenerateOTAUploadURL
type GenerateOTAUploadURLResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Data         Data   `json:"Data" xml:"Data"`
}

// CreateGenerateOTAUploadURLRequest creates a request to invoke GenerateOTAUploadURL API
func CreateGenerateOTAUploadURLRequest() (request *GenerateOTAUploadURLRequest) {
	request = &GenerateOTAUploadURLRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GenerateOTAUploadURL", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGenerateOTAUploadURLResponse creates a response to parse from GenerateOTAUploadURL response
func CreateGenerateOTAUploadURLResponse() (response *GenerateOTAUploadURLResponse) {
	response = &GenerateOTAUploadURLResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
