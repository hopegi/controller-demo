package cloudauth

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

// InitFaceVerify invokes the cloudauth.InitFaceVerify API synchronously
func (client *Client) InitFaceVerify(request *InitFaceVerifyRequest) (response *InitFaceVerifyResponse, err error) {
	response = CreateInitFaceVerifyResponse()
	err = client.DoAction(request, response)
	return
}

// InitFaceVerifyWithChan invokes the cloudauth.InitFaceVerify API asynchronously
func (client *Client) InitFaceVerifyWithChan(request *InitFaceVerifyRequest) (<-chan *InitFaceVerifyResponse, <-chan error) {
	responseChan := make(chan *InitFaceVerifyResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InitFaceVerify(request)
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

// InitFaceVerifyWithCallback invokes the cloudauth.InitFaceVerify API asynchronously
func (client *Client) InitFaceVerifyWithCallback(request *InitFaceVerifyRequest, callback func(response *InitFaceVerifyResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InitFaceVerifyResponse
		var err error
		defer close(result)
		response, err = client.InitFaceVerify(request)
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

// InitFaceVerifyRequest is the request struct for api InitFaceVerify
type InitFaceVerifyRequest struct {
	*requests.RpcRequest
	ProductCode            string           `position:"Query" name:"ProductCode"`
	FaceContrastPicture    string           `position:"Body" name:"FaceContrastPicture"`
	UserId                 string           `position:"Query" name:"UserId"`
	CertifyId              string           `position:"Query" name:"CertifyId"`
	CertNo                 string           `position:"Query" name:"CertNo"`
	OuterOrderNo           string           `position:"Query" name:"OuterOrderNo"`
	CertType               string           `position:"Query" name:"CertType"`
	FaceContrastPictureUrl string           `position:"Query" name:"FaceContrastPictureUrl"`
	Model                  string           `position:"Body" name:"Model"`
	MetaInfo               string           `position:"Query" name:"MetaInfo"`
	OssObjectName          string           `position:"Query" name:"OssObjectName"`
	CertName               string           `position:"Query" name:"CertName"`
	Ip                     string           `position:"Query" name:"Ip"`
	Mobile                 string           `position:"Query" name:"Mobile"`
	SceneId                requests.Integer `position:"Query" name:"SceneId"`
	OssBucketName          string           `position:"Query" name:"OssBucketName"`
	ReturnUrl              string           `position:"Query" name:"ReturnUrl"`
}

// InitFaceVerifyResponse is the response struct for api InitFaceVerify
type InitFaceVerifyResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	Message      string       `json:"Message" xml:"Message"`
	Code         string       `json:"Code" xml:"Code"`
	ResultObject ResultObject `json:"ResultObject" xml:"ResultObject"`
}

// CreateInitFaceVerifyRequest creates a request to invoke InitFaceVerify API
func CreateInitFaceVerifyRequest() (request *InitFaceVerifyRequest) {
	request = &InitFaceVerifyRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cloudauth", "2019-03-07", "InitFaceVerify", "cloudauth", "openAPI")
	request.Method = requests.POST
	return
}

// CreateInitFaceVerifyResponse creates a response to parse from InitFaceVerify response
func CreateInitFaceVerifyResponse() (response *InitFaceVerifyResponse) {
	response = &InitFaceVerifyResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
