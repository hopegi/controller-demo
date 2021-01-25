package dataworks_public

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

// UpdateUdfFile invokes the dataworks_public.UpdateUdfFile API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateudffile.html
func (client *Client) UpdateUdfFile(request *UpdateUdfFileRequest) (response *UpdateUdfFileResponse, err error) {
	response = CreateUpdateUdfFileResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateUdfFileWithChan invokes the dataworks_public.UpdateUdfFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateudffile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUdfFileWithChan(request *UpdateUdfFileRequest) (<-chan *UpdateUdfFileResponse, <-chan error) {
	responseChan := make(chan *UpdateUdfFileResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateUdfFile(request)
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

// UpdateUdfFileWithCallback invokes the dataworks_public.UpdateUdfFile API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/updateudffile.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateUdfFileWithCallback(request *UpdateUdfFileRequest, callback func(response *UpdateUdfFileResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateUdfFileResponse
		var err error
		defer close(result)
		response, err = client.UpdateUdfFile(request)
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

// UpdateUdfFileRequest is the request struct for api UpdateUdfFile
type UpdateUdfFileRequest struct {
	*requests.RpcRequest
	ReturnValue          string           `position:"Body" name:"ReturnValue"`
	Resources            string           `position:"Body" name:"Resources"`
	FunctionType         string           `position:"Body" name:"FunctionType"`
	CmdDescription       string           `position:"Body" name:"CmdDescription"`
	UdfDescription       string           `position:"Body" name:"UdfDescription"`
	ParameterDescription string           `position:"Body" name:"ParameterDescription"`
	ProjectIdentifier    string           `position:"Body" name:"ProjectIdentifier"`
	Example              string           `position:"Body" name:"Example"`
	ClassName            string           `position:"Body" name:"ClassName"`
	FileFolderPath       string           `position:"Body" name:"FileFolderPath"`
	ProjectId            requests.Integer `position:"Body" name:"ProjectId"`
	FileId               string           `position:"Body" name:"FileId"`
}

// UpdateUdfFileResponse is the response struct for api UpdateUdfFile
type UpdateUdfFileResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
}

// CreateUpdateUdfFileRequest creates a request to invoke UpdateUdfFile API
func CreateUpdateUdfFileRequest() (request *UpdateUdfFileRequest) {
	request = &UpdateUdfFileRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "UpdateUdfFile", "dide", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateUdfFileResponse creates a response to parse from UpdateUdfFile response
func CreateUpdateUdfFileResponse() (response *UpdateUdfFileResponse) {
	response = &UpdateUdfFileResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
