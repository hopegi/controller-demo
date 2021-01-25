package scdn

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

// StopScdnDomain invokes the scdn.StopScdnDomain API synchronously
func (client *Client) StopScdnDomain(request *StopScdnDomainRequest) (response *StopScdnDomainResponse, err error) {
	response = CreateStopScdnDomainResponse()
	err = client.DoAction(request, response)
	return
}

// StopScdnDomainWithChan invokes the scdn.StopScdnDomain API asynchronously
func (client *Client) StopScdnDomainWithChan(request *StopScdnDomainRequest) (<-chan *StopScdnDomainResponse, <-chan error) {
	responseChan := make(chan *StopScdnDomainResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StopScdnDomain(request)
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

// StopScdnDomainWithCallback invokes the scdn.StopScdnDomain API asynchronously
func (client *Client) StopScdnDomainWithCallback(request *StopScdnDomainRequest, callback func(response *StopScdnDomainResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StopScdnDomainResponse
		var err error
		defer close(result)
		response, err = client.StopScdnDomain(request)
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

// StopScdnDomainRequest is the request struct for api StopScdnDomain
type StopScdnDomainRequest struct {
	*requests.RpcRequest
	DomainName    string           `position:"Query" name:"DomainName"`
	OwnerId       requests.Integer `position:"Query" name:"OwnerId"`
	SecurityToken string           `position:"Query" name:"SecurityToken"`
}

// StopScdnDomainResponse is the response struct for api StopScdnDomain
type StopScdnDomainResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateStopScdnDomainRequest creates a request to invoke StopScdnDomain API
func CreateStopScdnDomainRequest() (request *StopScdnDomainRequest) {
	request = &StopScdnDomainRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "StopScdnDomain", "", "")
	request.Method = requests.POST
	return
}

// CreateStopScdnDomainResponse creates a response to parse from StopScdnDomain response
func CreateStopScdnDomainResponse() (response *StopScdnDomainResponse) {
	response = &StopScdnDomainResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
