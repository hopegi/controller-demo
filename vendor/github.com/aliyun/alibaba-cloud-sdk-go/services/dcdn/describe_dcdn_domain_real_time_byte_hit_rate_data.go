package dcdn

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

// DescribeDcdnDomainRealTimeByteHitRateData invokes the dcdn.DescribeDcdnDomainRealTimeByteHitRateData API synchronously
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateData(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest) (response *DescribeDcdnDomainRealTimeByteHitRateDataResponse, err error) {
	response = CreateDescribeDcdnDomainRealTimeByteHitRateDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeDcdnDomainRealTimeByteHitRateDataWithChan invokes the dcdn.DescribeDcdnDomainRealTimeByteHitRateData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateDataWithChan(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest) (<-chan *DescribeDcdnDomainRealTimeByteHitRateDataResponse, <-chan error) {
	responseChan := make(chan *DescribeDcdnDomainRealTimeByteHitRateDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeDcdnDomainRealTimeByteHitRateData(request)
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

// DescribeDcdnDomainRealTimeByteHitRateDataWithCallback invokes the dcdn.DescribeDcdnDomainRealTimeByteHitRateData API asynchronously
func (client *Client) DescribeDcdnDomainRealTimeByteHitRateDataWithCallback(request *DescribeDcdnDomainRealTimeByteHitRateDataRequest, callback func(response *DescribeDcdnDomainRealTimeByteHitRateDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeDcdnDomainRealTimeByteHitRateDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeDcdnDomainRealTimeByteHitRateData(request)
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

// DescribeDcdnDomainRealTimeByteHitRateDataRequest is the request struct for api DescribeDcdnDomainRealTimeByteHitRateData
type DescribeDcdnDomainRealTimeByteHitRateDataRequest struct {
	*requests.RpcRequest
	StartTime  string           `position:"Query" name:"StartTime"`
	DomainName string           `position:"Query" name:"DomainName"`
	EndTime    string           `position:"Query" name:"EndTime"`
	OwnerId    requests.Integer `position:"Query" name:"OwnerId"`
}

// DescribeDcdnDomainRealTimeByteHitRateDataResponse is the response struct for api DescribeDcdnDomainRealTimeByteHitRateData
type DescribeDcdnDomainRealTimeByteHitRateDataResponse struct {
	*responses.BaseResponse
	RequestId string                                          `json:"RequestId" xml:"RequestId"`
	Data      DataInDescribeDcdnDomainRealTimeByteHitRateData `json:"Data" xml:"Data"`
}

// CreateDescribeDcdnDomainRealTimeByteHitRateDataRequest creates a request to invoke DescribeDcdnDomainRealTimeByteHitRateData API
func CreateDescribeDcdnDomainRealTimeByteHitRateDataRequest() (request *DescribeDcdnDomainRealTimeByteHitRateDataRequest) {
	request = &DescribeDcdnDomainRealTimeByteHitRateDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dcdn", "2018-01-15", "DescribeDcdnDomainRealTimeByteHitRateData", "", "")
	request.Method = requests.GET
	return
}

// CreateDescribeDcdnDomainRealTimeByteHitRateDataResponse creates a response to parse from DescribeDcdnDomainRealTimeByteHitRateData response
func CreateDescribeDcdnDomainRealTimeByteHitRateDataResponse() (response *DescribeDcdnDomainRealTimeByteHitRateDataResponse) {
	response = &DescribeDcdnDomainRealTimeByteHitRateDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
