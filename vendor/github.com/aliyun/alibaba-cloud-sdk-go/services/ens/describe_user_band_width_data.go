package ens

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

// DescribeUserBandWidthData invokes the ens.DescribeUserBandWidthData API synchronously
func (client *Client) DescribeUserBandWidthData(request *DescribeUserBandWidthDataRequest) (response *DescribeUserBandWidthDataResponse, err error) {
	response = CreateDescribeUserBandWidthDataResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeUserBandWidthDataWithChan invokes the ens.DescribeUserBandWidthData API asynchronously
func (client *Client) DescribeUserBandWidthDataWithChan(request *DescribeUserBandWidthDataRequest) (<-chan *DescribeUserBandWidthDataResponse, <-chan error) {
	responseChan := make(chan *DescribeUserBandWidthDataResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeUserBandWidthData(request)
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

// DescribeUserBandWidthDataWithCallback invokes the ens.DescribeUserBandWidthData API asynchronously
func (client *Client) DescribeUserBandWidthDataWithCallback(request *DescribeUserBandWidthDataRequest, callback func(response *DescribeUserBandWidthDataResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeUserBandWidthDataResponse
		var err error
		defer close(result)
		response, err = client.DescribeUserBandWidthData(request)
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

// DescribeUserBandWidthDataRequest is the request struct for api DescribeUserBandWidthData
type DescribeUserBandWidthDataRequest struct {
	*requests.RpcRequest
	Isp         string `position:"Query" name:"Isp"`
	StartTime   string `position:"Query" name:"StartTime"`
	EnsRegionId string `position:"Query" name:"EnsRegionId"`
	Period      string `position:"Query" name:"Period"`
	EndTime     string `position:"Query" name:"EndTime"`
	Version     string `position:"Query" name:"Version"`
	InstanceId  string `position:"Query" name:"InstanceId"`
}

// DescribeUserBandWidthDataResponse is the response struct for api DescribeUserBandWidthData
type DescribeUserBandWidthDataResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Code        int         `json:"Code" xml:"Code"`
	MonitorData MonitorData `json:"MonitorData" xml:"MonitorData"`
}

// CreateDescribeUserBandWidthDataRequest creates a request to invoke DescribeUserBandWidthData API
func CreateDescribeUserBandWidthDataRequest() (request *DescribeUserBandWidthDataRequest) {
	request = &DescribeUserBandWidthDataRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Ens", "2017-11-10", "DescribeUserBandWidthData", "ens", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeUserBandWidthDataResponse creates a response to parse from DescribeUserBandWidthData response
func CreateDescribeUserBandWidthDataResponse() (response *DescribeUserBandWidthDataResponse) {
	response = &DescribeUserBandWidthDataResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
