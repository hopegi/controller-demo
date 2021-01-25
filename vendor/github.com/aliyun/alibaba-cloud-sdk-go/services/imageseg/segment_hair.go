package imageseg

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

// SegmentHair invokes the imageseg.SegmentHair API synchronously
func (client *Client) SegmentHair(request *SegmentHairRequest) (response *SegmentHairResponse, err error) {
	response = CreateSegmentHairResponse()
	err = client.DoAction(request, response)
	return
}

// SegmentHairWithChan invokes the imageseg.SegmentHair API asynchronously
func (client *Client) SegmentHairWithChan(request *SegmentHairRequest) (<-chan *SegmentHairResponse, <-chan error) {
	responseChan := make(chan *SegmentHairResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SegmentHair(request)
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

// SegmentHairWithCallback invokes the imageseg.SegmentHair API asynchronously
func (client *Client) SegmentHairWithCallback(request *SegmentHairRequest, callback func(response *SegmentHairResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SegmentHairResponse
		var err error
		defer close(result)
		response, err = client.SegmentHair(request)
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

// SegmentHairRequest is the request struct for api SegmentHair
type SegmentHairRequest struct {
	*requests.RpcRequest
	ImageURL string `position:"Query" name:"ImageURL"`
}

// SegmentHairResponse is the response struct for api SegmentHair
type SegmentHairResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateSegmentHairRequest creates a request to invoke SegmentHair API
func CreateSegmentHairRequest() (request *SegmentHairRequest) {
	request = &SegmentHairRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imageseg", "2019-12-30", "SegmentHair", "imageseg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSegmentHairResponse creates a response to parse from SegmentHair response
func CreateSegmentHairResponse() (response *SegmentHairResponse) {
	response = &SegmentHairResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
