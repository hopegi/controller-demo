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

// DescribeScdnDomainUvData invokes the scdn.DescribeScdnDomainUvData API synchronously
func (client *Client) DescribeScdnDomainUvData(request *DescribeScdnDomainUvDataRequest) (response *DescribeScdnDomainUvDataResponse, err error) {
	response = CreateDescribeScdnDomainUvDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeScdnDomainUvDataWithChan invokes the scdn.DescribeScdnDomainUvData API asynchronously
func (client *Client) DescribeScdnDomainUvDataWithChan(request *DescribeScdnDomainUvDataRequest) (<-chan *DescribeScdnDomainUvDataResponse, <-chan error) {
	responseChan := make(chan *DescribeScdnDomainUvDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeScdnDomainUvData(request)
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

// DescribeScdnDomainUvDataWithCallback invokes the scdn.DescribeScdnDomainUvData API asynchronously
func (client *Client) DescribeScdnDomainUvDataWithCallback(request *DescribeScdnDomainUvDataRequest, callback func(response *DescribeScdnDomainUvDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeScdnDomainUvDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeScdnDomainUvData(request)
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

// DescribeScdnDomainUvDataRequest is the request struct for api DescribeScdnDomainUvData
type DescribeScdnDomainUvDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeScdnDomainUvDataResponse is the response struct for api DescribeScdnDomainUvData
type DescribeScdnDomainUvDataResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	DomainName     string         `json:"DomainName" xml:"DomainName"`
	StartTime      string         `json:"StartTime" xml:"StartTime"`
	UvDataInterval UvDataInterval `json:"UvDataInterval" xml:"UvDataInterval"`
}

// CreateDescribeScdnDomainUvDataRequest creates a request to invoke DescribeScdnDomainUvData API
func CreateDescribeScdnDomainUvDataRequest() (request *DescribeScdnDomainUvDataRequest) {
	request = &DescribeScdnDomainUvDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("scdn", "2017-11-15", "DescribeScdnDomainUvData", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeScdnDomainUvDataResponse creates a response to parse from DescribeScdnDomainUvData response
func CreateDescribeScdnDomainUvDataResponse() (response *DescribeScdnDomainUvDataResponse) {
	response = &DescribeScdnDomainUvDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
