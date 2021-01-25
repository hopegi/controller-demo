package edas

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

// UnbindSlb invokes the edas.UnbindSlb API synchronously
func (client *Client) UnbindSlb(request *UnbindSlbRequest) (response *UnbindSlbResponse, err error) {
	response = CreateUnbindSlbResponse()
	err = client.DoAction(request, response)
	return
}

// UnbindSlbWithChan invokes the edas.UnbindSlb API asynchronously
func (client *Client) UnbindSlbWithChan(request *UnbindSlbRequest) (<-chan *UnbindSlbResponse, <-chan error) {
	responseChan := make(chan *UnbindSlbResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UnbindSlb(request)
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

// UnbindSlbWithCallback invokes the edas.UnbindSlb API asynchronously
func (client *Client) UnbindSlbWithCallback(request *UnbindSlbRequest, callback func(response *UnbindSlbResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UnbindSlbResponse
		var err error
		defer close(result)
		response, err = client.UnbindSlb(request)
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

// UnbindSlbRequest is the request struct for api UnbindSlb
type UnbindSlbRequest struct {
	*requests.RoaRequest
	SlbId string `position:"Query" name:"SlbId"`
	AppId string `position:"Query" name:"AppId"`
	Type  string `position:"Query" name:"Type"`
}

// UnbindSlbResponse is the response struct for api UnbindSlb
type UnbindSlbResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Data      string `json:"Data" xml:"Data"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateUnbindSlbRequest creates a request to invoke UnbindSlb API
func CreateUnbindSlbRequest() (request *UnbindSlbRequest) {
	request = &UnbindSlbRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "UnbindSlb", "/pop/app/unbind_slb_json", "edas", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUnbindSlbResponse creates a response to parse from UnbindSlb response
func CreateUnbindSlbResponse() (response *UnbindSlbResponse) {
	response = &UnbindSlbResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
