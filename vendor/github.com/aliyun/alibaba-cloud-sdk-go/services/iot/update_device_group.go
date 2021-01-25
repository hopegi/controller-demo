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

// UpdateDeviceGroup invokes the iot.UpdateDeviceGroup API synchronously
// api document: https://help.aliyun.com/api/iot/updatedevicegroup.html
func (client *Client) UpdateDeviceGroup(request *UpdateDeviceGroupRequest) (response *UpdateDeviceGroupResponse, err error) {
	response = CreateUpdateDeviceGroupResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDeviceGroupWithChan invokes the iot.UpdateDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/updatedevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDeviceGroupWithChan(request *UpdateDeviceGroupRequest) (<-chan *UpdateDeviceGroupResponse, <-chan error) {
	responseChan := make(chan *UpdateDeviceGroupResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDeviceGroup(request)
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

// UpdateDeviceGroupWithCallback invokes the iot.UpdateDeviceGroup API asynchronously
// api document: https://help.aliyun.com/api/iot/updatedevicegroup.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDeviceGroupWithCallback(request *UpdateDeviceGroupRequest, callback func(response *UpdateDeviceGroupResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDeviceGroupResponse
		var err error
		defer close(result)
		response, err = client.UpdateDeviceGroup(request)
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

// UpdateDeviceGroupRequest is the request struct for api UpdateDeviceGroup
type UpdateDeviceGroupRequest struct {
	*requests.RpcRequest
	IotInstanceId string `position:"Query" name:"IotInstanceId"`
	GroupId       string `position:"Query" name:"GroupId"`
	GroupDesc     string `position:"Query" name:"GroupDesc"`
	ApiProduct    string `position:"Body" name:"ApiProduct"`
	ApiRevision   string `position:"Body" name:"ApiRevision"`
}

// UpdateDeviceGroupResponse is the response struct for api UpdateDeviceGroup
type UpdateDeviceGroupResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	Success      bool   `json:"Success" xml:"Success"`
	Code         string `json:"Code" xml:"Code"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
}

// CreateUpdateDeviceGroupRequest creates a request to invoke UpdateDeviceGroup API
func CreateUpdateDeviceGroupRequest() (request *UpdateDeviceGroupRequest) {
	request = &UpdateDeviceGroupRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "UpdateDeviceGroup", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateDeviceGroupResponse creates a response to parse from UpdateDeviceGroup response
func CreateUpdateDeviceGroupResponse() (response *UpdateDeviceGroupResponse) {
	response = &UpdateDeviceGroupResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
