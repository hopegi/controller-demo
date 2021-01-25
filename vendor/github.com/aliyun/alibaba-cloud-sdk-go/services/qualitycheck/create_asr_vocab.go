package qualitycheck

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

// CreateAsrVocab invokes the qualitycheck.CreateAsrVocab API synchronously
// api document: https://help.aliyun.com/api/qualitycheck/createasrvocab.html
func (client *Client) CreateAsrVocab(request *CreateAsrVocabRequest) (response *CreateAsrVocabResponse, err error) {
	response = CreateCreateAsrVocabResponse()
	err = client.DoAction(request, response)
	return
}

// CreateAsrVocabWithChan invokes the qualitycheck.CreateAsrVocab API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createasrvocab.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAsrVocabWithChan(request *CreateAsrVocabRequest) (<-chan *CreateAsrVocabResponse, <-chan error) {
	responseChan := make(chan *CreateAsrVocabResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateAsrVocab(request)
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

// CreateAsrVocabWithCallback invokes the qualitycheck.CreateAsrVocab API asynchronously
// api document: https://help.aliyun.com/api/qualitycheck/createasrvocab.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateAsrVocabWithCallback(request *CreateAsrVocabRequest, callback func(response *CreateAsrVocabResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateAsrVocabResponse
		var err error
		defer close(result)
		response, err = client.CreateAsrVocab(request)
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

// CreateAsrVocabRequest is the request struct for api CreateAsrVocab
type CreateAsrVocabRequest struct {
	*requests.RpcRequest
	ResourceOwnerId requests.Integer `position:"Query" name:"ResourceOwnerId"`
	JsonStr         string           `position:"Query" name:"JsonStr"`
}

// CreateAsrVocabResponse is the response struct for api CreateAsrVocab
type CreateAsrVocabResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	Data      string `json:"Data" xml:"Data"`
}

// CreateCreateAsrVocabRequest creates a request to invoke CreateAsrVocab API
func CreateCreateAsrVocabRequest() (request *CreateAsrVocabRequest) {
	request = &CreateAsrVocabRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Qualitycheck", "2019-01-15", "CreateAsrVocab", "", "")
	return
}

// CreateCreateAsrVocabResponse creates a response to parse from CreateAsrVocab response
func CreateCreateAsrVocabResponse() (response *CreateAsrVocabResponse) {
	response = &CreateAsrVocabResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
