package vcs

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

// ListEventAlgorithmDetails invokes the vcs.ListEventAlgorithmDetails API synchronously
func (client *Client) ListEventAlgorithmDetails(request *ListEventAlgorithmDetailsRequest) (response *ListEventAlgorithmDetailsResponse, err error) {
	response = CreateListEventAlgorithmDetailsResponse()
	err = client.DoAction(request, response)
	return
}

// ListEventAlgorithmDetailsWithChan invokes the vcs.ListEventAlgorithmDetails API asynchronously
func (client *Client) ListEventAlgorithmDetailsWithChan(request *ListEventAlgorithmDetailsRequest) (<-chan *ListEventAlgorithmDetailsResponse, <-chan error) {
	responseChan := make(chan *ListEventAlgorithmDetailsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListEventAlgorithmDetails(request)
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

// ListEventAlgorithmDetailsWithCallback invokes the vcs.ListEventAlgorithmDetails API asynchronously
func (client *Client) ListEventAlgorithmDetailsWithCallback(request *ListEventAlgorithmDetailsRequest, callback func(response *ListEventAlgorithmDetailsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListEventAlgorithmDetailsResponse
		var err error
		defer close(result)
		response, err = client.ListEventAlgorithmDetails(request)
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

// ListEventAlgorithmDetailsRequest is the request struct for api ListEventAlgorithmDetails
type ListEventAlgorithmDetailsRequest struct {
	*requests.RpcRequest
	SourceId     string           `position:"Body" name:"SourceId"`
	CorpId       string           `position:"Body" name:"CorpId"`
	ExtendValue  string           `position:"Body" name:"ExtendValue"`
	EndTime      string           `position:"Body" name:"EndTime"`
	StartTime    string           `position:"Body" name:"StartTime"`
	PageNumber   requests.Integer `position:"Body" name:"PageNumber"`
	RecordId     string           `position:"Body" name:"RecordId"`
	EventValue   string           `position:"Body" name:"EventValue"`
	DataSourceId string           `position:"Body" name:"DataSourceId"`
	PageSize     requests.Integer `position:"Body" name:"PageSize"`
	EventType    string           `position:"Body" name:"EventType"`
}

// ListEventAlgorithmDetailsResponse is the response struct for api ListEventAlgorithmDetails
type ListEventAlgorithmDetailsResponse struct {
	*responses.BaseResponse
	Code       string  `json:"Code" xml:"Code"`
	Message    string  `json:"Message" xml:"Message"`
	PageNumber int     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int     `json:"PageSize" xml:"PageSize"`
	RequestId  string  `json:"RequestId" xml:"RequestId"`
	Success    string  `json:"Success" xml:"Success"`
	TotalCount int     `json:"TotalCount" xml:"TotalCount"`
	Data       []Datas `json:"Data" xml:"Data"`
}

// CreateListEventAlgorithmDetailsRequest creates a request to invoke ListEventAlgorithmDetails API
func CreateListEventAlgorithmDetailsRequest() (request *ListEventAlgorithmDetailsRequest) {
	request = &ListEventAlgorithmDetailsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListEventAlgorithmDetails", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListEventAlgorithmDetailsResponse creates a response to parse from ListEventAlgorithmDetails response
func CreateListEventAlgorithmDetailsResponse() (response *ListEventAlgorithmDetailsResponse) {
	response = &ListEventAlgorithmDetailsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
