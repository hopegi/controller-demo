package devops_rdc

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

// CreateDevopsProject invokes the devops_rdc.CreateDevopsProject API synchronously
// api document: https://help.aliyun.com/api/devops-rdc/createdevopsproject.html
func (client *Client) CreateDevopsProject(request *CreateDevopsProjectRequest) (response *CreateDevopsProjectResponse, err error) {
	response = CreateCreateDevopsProjectResponse()
	err = client.DoAction(request, response)
	return
}

// CreateDevopsProjectWithChan invokes the devops_rdc.CreateDevopsProject API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/createdevopsproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDevopsProjectWithChan(request *CreateDevopsProjectRequest) (<-chan *CreateDevopsProjectResponse, <-chan error) {
	responseChan := make(chan *CreateDevopsProjectResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateDevopsProject(request)
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

// CreateDevopsProjectWithCallback invokes the devops_rdc.CreateDevopsProject API asynchronously
// api document: https://help.aliyun.com/api/devops-rdc/createdevopsproject.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateDevopsProjectWithCallback(request *CreateDevopsProjectRequest, callback func(response *CreateDevopsProjectResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateDevopsProjectResponse
		var err error
		defer close(result)
		response, err = client.CreateDevopsProject(request)
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

// CreateDevopsProjectRequest is the request struct for api CreateDevopsProject
type CreateDevopsProjectRequest struct {
	*requests.RpcRequest
	Name        string `position:"Body" name:"Name"`
	Description string `position:"Body" name:"Description"`
	OrgId       string `position:"Body" name:"OrgId"`
}

// CreateDevopsProjectResponse is the response struct for api CreateDevopsProject
type CreateDevopsProjectResponse struct {
	*responses.BaseResponse
	RequestId    string `json:"RequestId" xml:"RequestId"`
	ErrorCode    string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage string `json:"ErrorMessage" xml:"ErrorMessage"`
	Success      bool   `json:"Success" xml:"Success"`
	Object       string `json:"Object" xml:"Object"`
}

// CreateCreateDevopsProjectRequest creates a request to invoke CreateDevopsProject API
func CreateCreateDevopsProjectRequest() (request *CreateDevopsProjectRequest) {
	request = &CreateDevopsProjectRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("devops-rdc", "2020-03-03", "CreateDevopsProject", "", "")
	request.Method = requests.POST
	return
}

// CreateCreateDevopsProjectResponse creates a response to parse from CreateDevopsProject response
func CreateCreateDevopsProjectResponse() (response *CreateDevopsProjectResponse) {
	response = &CreateDevopsProjectResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
