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

// BindDriverToEdgeInstance invokes the iot.BindDriverToEdgeInstance API synchronously
// api document: https://help.aliyun.com/api/iot/binddrivertoedgeinstance.html
func (client *Client) BindDriverToEdgeInstance(request *BindDriverToEdgeInstanceRequest) (response *BindDriverToEdgeInstanceResponse, err error) {
	response = CreateBindDriverToEdgeInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// BindDriverToEdgeInstanceWithChan invokes the iot.BindDriverToEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/binddrivertoedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindDriverToEdgeInstanceWithChan(request *BindDriverToEdgeInstanceRequest) (<-chan *BindDriverToEdgeInstanceResponse, <-chan error) {
	responseChan := make(chan *BindDriverToEdgeInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.BindDriverToEdgeInstance(request)
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

// BindDriverToEdgeInstanceWithCallback invokes the iot.BindDriverToEdgeInstance API asynchronously
// api document: https://help.aliyun.com/api/iot/binddrivertoedgeinstance.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) BindDriverToEdgeInstanceWithCallback(request *BindDriverToEdgeInstanceRequest, callback func(response *BindDriverToEdgeInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *BindDriverToEdgeInstanceResponse
		var err error
		defer close(result)
		response, err = client.BindDriverToEdgeInstance(request)
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

// BindDriverToEdgeInstanceRequest is the request struct for api BindDriverToEdgeInstance
type BindDriverToEdgeInstanceRequest struct {
	*requests.RpcRequest
	DriverId      string `position:"Query" name:"DriverId"`
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	DriverVersion string `position:"Query" name:"DriverVersion"`
	OrderId       string `position:"Query" name:"OrderId"`
	InstanceId    string `position:"Query" name:"InstanceId"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// BindDriverToEdgeInstanceResponse is the response struct for api BindDriverToEdgeInstance
type BindDriverToEdgeInstanceResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateBindDriverToEdgeInstanceRequest creates a request to invoke BindDriverToEdgeInstance API
func CreateBindDriverToEdgeInstanceRequest() (request *BindDriverToEdgeInstanceRequest) {
	request = &BindDriverToEdgeInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "BindDriverToEdgeInstance", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateBindDriverToEdgeInstanceResponse creates a response to parse from BindDriverToEdgeInstance response
func CreateBindDriverToEdgeInstanceResponse() (response *BindDriverToEdgeInstanceResponse) {
	response = &BindDriverToEdgeInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
