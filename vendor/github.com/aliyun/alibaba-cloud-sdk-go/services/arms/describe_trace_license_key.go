package arms

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

// DescribeTraceLicenseKey invokes the arms.DescribeTraceLicenseKey API synchronously
// api document: https://help.aliyun.com/api/arms/describetracelicensekey.html
func (client *Client) DescribeTraceLicenseKey(request *DescribeTraceLicenseKeyRequest) (response *DescribeTraceLicenseKeyResponse, err error) {
	response = CreateDescribeTraceLicenseKeyResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeTraceLicenseKeyWithChan invokes the arms.DescribeTraceLicenseKey API asynchronously
// api document: https://help.aliyun.com/api/arms/describetracelicensekey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceLicenseKeyWithChan(request *DescribeTraceLicenseKeyRequest) (<-chan *DescribeTraceLicenseKeyResponse, <-chan error) {
	responseChan := make(chan *DescribeTraceLicenseKeyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeTraceLicenseKey(request)
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

// DescribeTraceLicenseKeyWithCallback invokes the arms.DescribeTraceLicenseKey API asynchronously
// api document: https://help.aliyun.com/api/arms/describetracelicensekey.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeTraceLicenseKeyWithCallback(request *DescribeTraceLicenseKeyRequest, callback func(response *DescribeTraceLicenseKeyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeTraceLicenseKeyResponse
		var err error
		defer close(result)
		response, err = client.DescribeTraceLicenseKey(request)
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

// DescribeTraceLicenseKeyRequest is the request struct for api DescribeTraceLicenseKey
type DescribeTraceLicenseKeyRequest struct {
	*requests.RpcRequest
}

// DescribeTraceLicenseKeyResponse is the response struct for api DescribeTraceLicenseKey
type DescribeTraceLicenseKeyResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	LicenseKey string `json:"LicenseKey" xml:"LicenseKey"`
}

// CreateDescribeTraceLicenseKeyRequest creates a request to invoke DescribeTraceLicenseKey API
func CreateDescribeTraceLicenseKeyRequest() (request *DescribeTraceLicenseKeyRequest) {
	request = &DescribeTraceLicenseKeyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "DescribeTraceLicenseKey", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeTraceLicenseKeyResponse creates a response to parse from DescribeTraceLicenseKey response
func CreateDescribeTraceLicenseKeyResponse() (response *DescribeTraceLicenseKeyResponse) {
	response = &DescribeTraceLicenseKeyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
