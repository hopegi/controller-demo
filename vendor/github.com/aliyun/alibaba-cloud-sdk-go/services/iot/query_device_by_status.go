package iot

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

// QueryDeviceByStatus invokes the iot.QueryDeviceByStatus API synchronously
// api document: https://help.aliyun.com/api/iot/querydevicebystatus.html
func (client *Client) QueryDeviceByStatus(request *QueryDeviceByStatusRequest) (response *QueryDeviceByStatusResponse, err error) {
	response = CreateQueryDeviceByStatusResponse()
	err = client.DoAction(request, response)
	return
}

// QueryDeviceByStatusWithChan invokes the iot.QueryDeviceByStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicebystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceByStatusWithChan(request *QueryDeviceByStatusRequest) (<-chan *QueryDeviceByStatusResponse, <-chan error) {
	responseChan := make(chan *QueryDeviceByStatusResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryDeviceByStatus(request)
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

// QueryDeviceByStatusWithCallback invokes the iot.QueryDeviceByStatus API asynchronously
// api document: https://help.aliyun.com/api/iot/querydevicebystatus.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) QueryDeviceByStatusWithCallback(request *QueryDeviceByStatusRequest, callback func(response *QueryDeviceByStatusResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryDeviceByStatusResponse
		var err error
		defer close(result)
		response, err = client.QueryDeviceByStatus(request)
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

// QueryDeviceByStatusRequest is the request struct for api QueryDeviceByStatus
type QueryDeviceByStatusRequest struct {
	*requests.RpcRequest
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	IotInstanceId   string           `position:"Query" name:"IotInstanceId"`
	PageSize        requests.Integer `position:"Query" name:"PageSize"`
	CurrentPage     requests.Integer `position:"Query" name:"CurrentPage"`
	ProductKey      string           `position:"Query" name:"ProductKey"`
	ApiProduct      string           `position:"Body" name:"ApiProduct"`
	ApiRevision     string           `position:"Body" name:"ApiRevision"`
	BizTenantId     string           `position:"Query" name:"BizTenantId"`
	Status          requests.Integer `position:"Query" name:"Status"`
}

// QueryDeviceByStatusResponse is the response struct for api QueryDeviceByStatus
type QueryDeviceByStatusResponse struct {
	*responses.BaseResponse
	RequestId    string                    `json:"RequestId" xml:"RequestId"`
	Success      bool                      `json:"Success" xml:"Success"`
	Code         string                    `json:"Code" xml:"Code"`
	ErrorMessage string                    `json:"ErrorMessage" xml:"ErrorMessage"`
	Page         int                       `json:"Page" xml:"Page"`
	PageSize     int                       `json:"PageSize" xml:"PageSize"`
	PageCount    int                       `json:"PageCount" xml:"PageCount"`
	Total        int                       `json:"Total" xml:"Total"`
	Data         DataInQueryDeviceByStatus `json:"Data" xml:"Data"`
}

// CreateQueryDeviceByStatusRequest creates a request to invoke QueryDeviceByStatus API
func CreateQueryDeviceByStatusRequest() (request *QueryDeviceByStatusRequest) {
	request = &QueryDeviceByStatusRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Iot", "2018-01-20", "QueryDeviceByStatus", "iot", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryDeviceByStatusResponse creates a response to parse from QueryDeviceByStatus response
func CreateQueryDeviceByStatusResponse() (response *QueryDeviceByStatusResponse) {
	response = &QueryDeviceByStatusResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
