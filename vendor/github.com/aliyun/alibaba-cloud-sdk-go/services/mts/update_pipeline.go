package mts

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

// UpdatePipeline invokes the mts.UpdatePipeline API synchronously
// api document: https://help.aliyun.com/api/mts/updatepipeline.html
func (client *Client) UpdatePipeline(request *UpdatePipelineRequest) (response *UpdatePipelineResponse, err error) {
	response = CreateUpdatePipelineResponse()
	err = client.DoAction(request, response)
	return
}

// UpdatePipelineWithChan invokes the mts.UpdatePipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/updatepipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePipelineWithChan(request *UpdatePipelineRequest) (<-chan *UpdatePipelineResponse, <-chan error) {
	responseChan := make(chan *UpdatePipelineResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdatePipeline(request)
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

// UpdatePipelineWithCallback invokes the mts.UpdatePipeline API asynchronously
// api document: https://help.aliyun.com/api/mts/updatepipeline.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdatePipelineWithCallback(request *UpdatePipelineRequest, callback func(response *UpdatePipelineResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdatePipelineResponse
		var err error
		defer close(result)
		response, err = client.UpdatePipeline(request)
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

// UpdatePipelineRequest is the request struct for api UpdatePipeline
type UpdatePipelineRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	Role                 string           `position:"Query" name:"Role"`
	State                string           `position:"Query" name:"State"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	NotifyConfig         string           `position:"Query" name:"NotifyConfig"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	PipelineId           string           `position:"Query" name:"PipelineId"`
	Name                 string           `position:"Query" name:"Name"`
}

// UpdatePipelineResponse is the response struct for api UpdatePipeline
type UpdatePipelineResponse struct {
	*responses.BaseResponse
	RequestId string   `json:"RequestId" xml:"RequestId"`
	Pipeline  Pipeline `json:"Pipeline" xml:"Pipeline"`
}

// CreateUpdatePipelineRequest creates a request to invoke UpdatePipeline API
func CreateUpdatePipelineRequest() (request *UpdatePipelineRequest) {
	request = &UpdatePipelineRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Mts", "2014-06-18", "UpdatePipeline", "", "")
	return
}

// CreateUpdatePipelineResponse creates a response to parse from UpdatePipeline response
func CreateUpdatePipelineResponse() (response *UpdatePipelineResponse) {
	response = &UpdatePipelineResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
