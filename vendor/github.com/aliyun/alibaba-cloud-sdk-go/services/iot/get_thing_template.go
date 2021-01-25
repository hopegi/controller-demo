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

// GetThingTemplate invokes the iot.GetThingTemplate API synchronously
// api document: https://help.aliyun.com/api/iot/getthingtemplate.html
func (client *Client) GetThingTemplate(request *GetThingTemplateRequest) (response *GetThingTemplateResponse, err error) {
	response = CreateGetThingTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// GetThingTemplateWithChan invokes the iot.GetThingTemplate API asynchronously
// api document: https://help.aliyun.com/api/iot/getthingtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThingTemplateWithChan(request *GetThingTemplateRequest) (<-chan *GetThingTemplateResponse, <-chan error) {
	responseChan := make(chan *GetThingTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetThingTemplate(request)
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

// GetThingTemplateWithCallback invokes the iot.GetThingTemplate API asynchronously
// api document: https://help.aliyun.com/api/iot/getthingtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetThingTemplateWithCallback(request *GetThingTemplateRequest, callback func(response *GetThingTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetThingTemplateResponse
		var err error
		defer close(result)
		response, err = client.GetThingTemplate(request)
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

// GetThingTemplateRequest is the request struct for api GetThingTemplate
type GetThingTemplateRequest struct {
	*requests.RpcRequest
	CategoryKey     string `position:"Query" name:"CategoryKey"`
	ResourceGroupId string `position:"Query" name:"ResourceGroupId"`
	IotInstanceId   string `position:"Query" name:"IotInstanceId"`
	ApiProduct      string `position:"Body" name:"ApiProduct"`
	ApiRevision     string `position:"Body" name:"ApiRevision"`
}

// GetThingTemplateResponse is the response struct for api GetThingTemplate
type GetThingTemplateResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Code           string `json:"Code" xml:"Code"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	ThingModelJSON string `json:"ThingModelJSON" xml:"ThingModelJSON"`
}

// CreateGetThingTemplateRequest creates a request to invoke GetThingTemplate API
func CreateGetThingTemplateRequest() (request *GetThingTemplateRequest) {
	request = &GetThingTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "GetThingTemplate", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetThingTemplateResponse creates a response to parse from GetThingTemplate response
func CreateGetThingTemplateResponse() (response *GetThingTemplateResponse) {
	response = &GetThingTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
