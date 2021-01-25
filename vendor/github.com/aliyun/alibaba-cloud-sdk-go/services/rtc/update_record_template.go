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

// UpdateRecordTemplate invokes the rtc.UpdateRecordTemplate API synchronously
func (client *Client) UpdateRecordTemplate(request *UpdateRecordTemplateRequest) (response *UpdateRecordTemplateResponse, err error) {
	response = CreateUpdateRecordTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateRecordTemplateWithChan invokes the rtc.UpdateRecordTemplate API asynchronously
func (client *Client) UpdateRecordTemplateWithChan(request *UpdateRecordTemplateRequest) (<-chan *UpdateRecordTemplateResponse, <-chan error) {
	responseChan := make(chan *UpdateRecordTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateRecordTemplate(request)
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

// UpdateRecordTemplateWithCallback invokes the rtc.UpdateRecordTemplate API asynchronously
func (client *Client) UpdateRecordTemplateWithCallback(request *UpdateRecordTemplateRequest, callback func(response *UpdateRecordTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateRecordTemplateResponse
		var err error
		defer close(result)
		response, err = client.UpdateRecordTemplate(request)
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

// UpdateRecordTemplateRequest is the request struct for api UpdateRecordTemplate
type UpdateRecordTemplateRequest struct {
	*requests.RpcRequest
	Formats           *[]string        `position:"Query" name:"Formats"  type:"Repeated"`
	OssFilePrefix     string           `position:"Query" name:"OssFilePrefix"`
	BackgroundColor   requests.Integer `position:"Query" name:"BackgroundColor"`
	TaskProfile       string           `position:"Query" name:"TaskProfile"`
	LayoutIds         *[]string        `position:"Query" name:"LayoutIds"  type:"Repeated"`
	ShowLog           string           `position:"Query" name:"ShowLog"`
	OssBucket         string           `position:"Query" name:"OssBucket"`
	MnsQueue          string           `position:"Query" name:"MnsQueue"`
	FileSplitInterval requests.Integer `position:"Query" name:"FileSplitInterval"`
	OwnerId           requests.Integer `position:"Query" name:"OwnerId"`
	TemplateId        string           `position:"Query" name:"TemplateId"`
	AppId             string           `position:"Query" name:"AppId"`
	Name              string           `position:"Query" name:"Name"`
	MediaEncode       requests.Integer `position:"Query" name:"MediaEncode"`
}

// UpdateRecordTemplateResponse is the response struct for api UpdateRecordTemplate
type UpdateRecordTemplateResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TemplateId string `json:"TemplateId" xml:"TemplateId"`
}

// CreateUpdateRecordTemplateRequest creates a request to invoke UpdateRecordTemplate API
func CreateUpdateRecordTemplateRequest() (request *UpdateRecordTemplateRequest) {
	request = &UpdateRecordTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("rtc", "2018-01-11", "UpdateRecordTemplate", "rtc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateRecordTemplateResponse creates a response to parse from UpdateRecordTemplate response
func CreateUpdateRecordTemplateResponse() (response *UpdateRecordTemplateResponse) {
	response = &UpdateRecordTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
