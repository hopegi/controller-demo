package adb

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

// AllocateClusterPublicConnection invokes the adb.AllocateClusterPublicConnection API synchronously
// api document: https://help.aliyun.com/api/adb/allocateclusterpublicconnection.html
func (client *Client) AllocateClusterPublicConnection(request *AllocateClusterPublicConnectionRequest) (response *AllocateClusterPublicConnectionResponse, err error) {
	response = CreateAllocateClusterPublicConnectionResponse()
	err = client.DoAction(request, response)
	return
}

// AllocateClusterPublicConnectionWithChan invokes the adb.AllocateClusterPublicConnection API asynchronously
// api document: https://help.aliyun.com/api/adb/allocateclusterpublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateClusterPublicConnectionWithChan(request *AllocateClusterPublicConnectionRequest) (<-chan *AllocateClusterPublicConnectionResponse, <-chan error) {
	responseChan := make(chan *AllocateClusterPublicConnectionResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.AllocateClusterPublicConnection(request)
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

// AllocateClusterPublicConnectionWithCallback invokes the adb.AllocateClusterPublicConnection API asynchronously
// api document: https://help.aliyun.com/api/adb/allocateclusterpublicconnection.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) AllocateClusterPublicConnectionWithCallback(request *AllocateClusterPublicConnectionRequest, callback func(response *AllocateClusterPublicConnectionResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *AllocateClusterPublicConnectionResponse
		var err error
		defer close(result)
		response, err = client.AllocateClusterPublicConnection(request)
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

// AllocateClusterPublicConnectionRequest is the request struct for api AllocateClusterPublicConnection
type AllocateClusterPublicConnectionRequest struct {
	*requests.RpcRequest
	ResourceOwnerId        requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ConnectionStringPrefix string           `position:"Query" name:"ConnectionStringPrefix"`
	ResourceOwnerAccount   string           `position:"Query" name:"ResourceOwnerAccount"`
	DBClusterId            string           `position:"Query" name:"DBClusterId"`
	OwnerAccount           string           `position:"Query" name:"OwnerAccount"`
	OwnerId                requests.Integer `position:"Query" name:"OwnerId"`
}

// AllocateClusterPublicConnectionResponse is the response struct for api AllocateClusterPublicConnection
type AllocateClusterPublicConnectionResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateAllocateClusterPublicConnectionRequest creates a request to invoke AllocateClusterPublicConnection API
func CreateAllocateClusterPublicConnectionRequest() (request *AllocateClusterPublicConnectionRequest) {
	request = &AllocateClusterPublicConnectionRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("adb", "2019-03-15", "AllocateClusterPublicConnection", "ads", "openAPI")
	request.Method = requests.POST
	return
}

// CreateAllocateClusterPublicConnectionResponse creates a response to parse from AllocateClusterPublicConnection response
func CreateAllocateClusterPublicConnectionResponse() (response *AllocateClusterPublicConnectionResponse) {
	response = &AllocateClusterPublicConnectionResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
