package retailcloud

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

// ListNodeLabels invokes the retailcloud.ListNodeLabels API synchronously
// api document: https://help.aliyun.com/api/retailcloud/listnodelabels.html
func (client *Client) ListNodeLabels(request *ListNodeLabelsRequest) (response *ListNodeLabelsResponse, err error) {
	response = CreateListNodeLabelsResponse()
	err = client.DoAction(request, response)
	return
}

// ListNodeLabelsWithChan invokes the retailcloud.ListNodeLabels API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listnodelabels.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodeLabelsWithChan(request *ListNodeLabelsRequest) (<-chan *ListNodeLabelsResponse, <-chan error) {
	responseChan := make(chan *ListNodeLabelsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListNodeLabels(request)
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

// ListNodeLabelsWithCallback invokes the retailcloud.ListNodeLabels API asynchronously
// api document: https://help.aliyun.com/api/retailcloud/listnodelabels.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListNodeLabelsWithCallback(request *ListNodeLabelsRequest, callback func(response *ListNodeLabelsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListNodeLabelsResponse
		var err error
		defer close(result)
		response, err = client.ListNodeLabels(request)
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

// ListNodeLabelsRequest is the request struct for api ListNodeLabels
type ListNodeLabelsRequest struct {
	*requests.RpcRequest
	LabelKey   string           `position:"Query" name:"LabelKey"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
}

// ListNodeLabelsResponse is the response struct for api ListNodeLabels
type ListNodeLabelsResponse struct {
	*responses.BaseResponse
	Code       int                     `json:"Code" xml:"Code"`
	ErrorMsg   string                  `json:"ErrorMsg" xml:"ErrorMsg"`
	PageNumber int                     `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                     `json:"PageSize" xml:"PageSize"`
	RequestId  string                  `json:"RequestId" xml:"RequestId"`
	TotalCount int64                   `json:"TotalCount" xml:"TotalCount"`
	Data       []ListNodeLabelResponse `json:"Data" xml:"Data"`
}

// CreateListNodeLabelsRequest creates a request to invoke ListNodeLabels API
func CreateListNodeLabelsRequest() (request *ListNodeLabelsRequest) {
	request = &ListNodeLabelsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("retailcloud", "2018-03-13", "ListNodeLabels", "", "")
	request.Method = requests.POST
	return
}

// CreateListNodeLabelsResponse creates a response to parse from ListNodeLabels response
func CreateListNodeLabelsResponse() (response *ListNodeLabelsResponse) {
	response = &ListNodeLabelsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
