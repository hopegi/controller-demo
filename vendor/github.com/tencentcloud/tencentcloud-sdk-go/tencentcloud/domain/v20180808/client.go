// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20180808

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-08-08"

type Client struct {
    common.Client
}

// Deprecated
func NewClientWithSecretId(secretId, secretKey, region string) (client *Client, err error) {
    cpf := profile.NewClientProfile()
    client = &Client{}
    client.Init(region).WithSecretId(secretId, secretKey).WithProfile(cpf)
    return
}

func NewClient(credential *common.Credential, region string, clientProfile *profile.ClientProfile) (client *Client, err error) {
    client = &Client{}
    client.Init(region).
        WithCredential(credential).
        WithProfile(clientProfile)
    return
}


func NewCheckBatchStatusRequest() (request *CheckBatchStatusRequest) {
    request = &CheckBatchStatusRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CheckBatchStatus")
    return
}

func NewCheckBatchStatusResponse() (response *CheckBatchStatusResponse) {
    response = &CheckBatchStatusResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( CheckBatchStatus ) 用于检查批量任务状态 。
// 
// 默认接口请求频率限制：20次/秒。
func (c *Client) CheckBatchStatus(request *CheckBatchStatusRequest) (response *CheckBatchStatusResponse, err error) {
    if request == nil {
        request = NewCheckBatchStatusRequest()
    }
    response = NewCheckBatchStatusResponse()
    err = c.Send(request, response)
    return
}

func NewCheckDomainRequest() (request *CheckDomainRequest) {
    request = &CheckDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CheckDomain")
    return
}

func NewCheckDomainResponse() (response *CheckDomainResponse) {
    response = &CheckDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查域名是否可以注册。
func (c *Client) CheckDomain(request *CheckDomainRequest) (response *CheckDomainResponse, err error) {
    if request == nil {
        request = NewCheckDomainRequest()
    }
    response = NewCheckDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateDomainBatchRequest() (request *CreateDomainBatchRequest) {
    request = &CreateDomainBatchRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "CreateDomainBatch")
    return
}

func NewCreateDomainBatchResponse() (response *CreateDomainBatchResponse) {
    response = &CreateDomainBatchResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 ( CreateDomainBatch ) 用于批量注册域名 。
// 
// 默认接口请求频率限制：20次/秒。
func (c *Client) CreateDomainBatch(request *CreateDomainBatchRequest) (response *CreateDomainBatchResponse, err error) {
    if request == nil {
        request = NewCreateDomainBatchRequest()
    }
    response = NewCreateDomainBatchResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainBaseInfoRequest() (request *DescribeDomainBaseInfoRequest) {
    request = &DescribeDomainBaseInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainBaseInfo")
    return
}

func NewDescribeDomainBaseInfoResponse() (response *DescribeDomainBaseInfoResponse) {
    response = &DescribeDomainBaseInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (  DescribeDomainBaseInfo) 获取域名基础信息。
// 
// 默认接口请求频率限制：20次/秒。
func (c *Client) DescribeDomainBaseInfo(request *DescribeDomainBaseInfoRequest) (response *DescribeDomainBaseInfoResponse, err error) {
    if request == nil {
        request = NewDescribeDomainBaseInfoRequest()
    }
    response = NewDescribeDomainBaseInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainNameListRequest() (request *DescribeDomainNameListRequest) {
    request = &DescribeDomainNameListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainNameList")
    return
}

func NewDescribeDomainNameListResponse() (response *DescribeDomainNameListResponse) {
    response = &DescribeDomainNameListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (  DescribeDomainNameList ) 获取域名列表。
// 
// 默认接口请求频率限制：20次/秒。
func (c *Client) DescribeDomainNameList(request *DescribeDomainNameListRequest) (response *DescribeDomainNameListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainNameListRequest()
    }
    response = NewDescribeDomainNameListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDomainPriceListRequest() (request *DescribeDomainPriceListRequest) {
    request = &DescribeDomainPriceListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeDomainPriceList")
    return
}

func NewDescribeDomainPriceListResponse() (response *DescribeDomainPriceListResponse) {
    response = &DescribeDomainPriceListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 按照域名后缀获取对应的价格列表
func (c *Client) DescribeDomainPriceList(request *DescribeDomainPriceListRequest) (response *DescribeDomainPriceListResponse, err error) {
    if request == nil {
        request = NewDescribeDomainPriceListRequest()
    }
    response = NewDescribeDomainPriceListResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeTemplateListRequest() (request *DescribeTemplateListRequest) {
    request = &DescribeTemplateListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("domain", APIVersion, "DescribeTemplateList")
    return
}

func NewDescribeTemplateListResponse() (response *DescribeTemplateListResponse) {
    response = &DescribeTemplateListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 本接口 (DescribeTemplateList) 用于获取模板列表。
// 
// 默认接口请求频率限制：20次/秒。
func (c *Client) DescribeTemplateList(request *DescribeTemplateListRequest) (response *DescribeTemplateListResponse, err error) {
    if request == nil {
        request = NewDescribeTemplateListRequest()
    }
    response = NewDescribeTemplateListResponse()
    err = c.Send(request, response)
    return
}
