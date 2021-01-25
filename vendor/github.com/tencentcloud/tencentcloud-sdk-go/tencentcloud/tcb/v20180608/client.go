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

package v20180608

import (
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-06-08"

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


func NewCheckTcbServiceRequest() (request *CheckTcbServiceRequest) {
    request = &CheckTcbServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CheckTcbService")
    return
}

func NewCheckTcbServiceResponse() (response *CheckTcbServiceResponse) {
    response = &CheckTcbServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 检查是否开通Tcb服务
func (c *Client) CheckTcbService(request *CheckTcbServiceRequest) (response *CheckTcbServiceResponse, err error) {
    if request == nil {
        request = NewCheckTcbServiceRequest()
    }
    response = NewCheckTcbServiceResponse()
    err = c.Send(request, response)
    return
}

func NewCommonServiceAPIRequest() (request *CommonServiceAPIRequest) {
    request = &CommonServiceAPIRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CommonServiceAPI")
    return
}

func NewCommonServiceAPIResponse() (response *CommonServiceAPIResponse) {
    response = &CommonServiceAPIResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// TCB云API统一入口
func (c *Client) CommonServiceAPI(request *CommonServiceAPIRequest) (response *CommonServiceAPIResponse, err error) {
    if request == nil {
        request = NewCommonServiceAPIRequest()
    }
    response = NewCommonServiceAPIResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAuthDomainRequest() (request *CreateAuthDomainRequest) {
    request = &CreateAuthDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateAuthDomain")
    return
}

func NewCreateAuthDomainResponse() (response *CreateAuthDomainResponse) {
    response = &CreateAuthDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 增加安全域名
func (c *Client) CreateAuthDomain(request *CreateAuthDomainRequest) (response *CreateAuthDomainResponse, err error) {
    if request == nil {
        request = NewCreateAuthDomainRequest()
    }
    response = NewCreateAuthDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreateHostingDomainRequest() (request *CreateHostingDomainRequest) {
    request = &CreateHostingDomainRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateHostingDomain")
    return
}

func NewCreateHostingDomainResponse() (response *CreateHostingDomainResponse) {
    response = &CreateHostingDomainResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建托管域名
func (c *Client) CreateHostingDomain(request *CreateHostingDomainRequest) (response *CreateHostingDomainResponse, err error) {
    if request == nil {
        request = NewCreateHostingDomainRequest()
    }
    response = NewCreateHostingDomainResponse()
    err = c.Send(request, response)
    return
}

func NewCreatePostpayPackageRequest() (request *CreatePostpayPackageRequest) {
    request = &CreatePostpayPackageRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreatePostpayPackage")
    return
}

func NewCreatePostpayPackageResponse() (response *CreatePostpayPackageResponse) {
    response = &CreatePostpayPackageResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 开通后付费资源
func (c *Client) CreatePostpayPackage(request *CreatePostpayPackageRequest) (response *CreatePostpayPackageResponse, err error) {
    if request == nil {
        request = NewCreatePostpayPackageRequest()
    }
    response = NewCreatePostpayPackageResponse()
    err = c.Send(request, response)
    return
}

func NewCreateStaticStoreRequest() (request *CreateStaticStoreRequest) {
    request = &CreateStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "CreateStaticStore")
    return
}

func NewCreateStaticStoreResponse() (response *CreateStaticStoreResponse) {
    response = &CreateStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建静态托管资源，包括COS和CDN，异步任务创建，查看创建结果需要根据DescribeStaticStore接口来查看
func (c *Client) CreateStaticStore(request *CreateStaticStoreRequest) (response *CreateStaticStoreResponse, err error) {
    if request == nil {
        request = NewCreateStaticStoreRequest()
    }
    response = NewCreateStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteEndUserRequest() (request *DeleteEndUserRequest) {
    request = &DeleteEndUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DeleteEndUser")
    return
}

func NewDeleteEndUserResponse() (response *DeleteEndUserResponse) {
    response = &DeleteEndUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除终端用户
func (c *Client) DeleteEndUser(request *DeleteEndUserRequest) (response *DeleteEndUserResponse, err error) {
    if request == nil {
        request = NewDeleteEndUserRequest()
    }
    response = NewDeleteEndUserResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAuthDomainsRequest() (request *DescribeAuthDomainsRequest) {
    request = &DescribeAuthDomainsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeAuthDomains")
    return
}

func NewDescribeAuthDomainsResponse() (response *DescribeAuthDomainsResponse) {
    response = &DescribeAuthDomainsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取安全域名列表
func (c *Client) DescribeAuthDomains(request *DescribeAuthDomainsRequest) (response *DescribeAuthDomainsResponse, err error) {
    if request == nil {
        request = NewDescribeAuthDomainsRequest()
    }
    response = NewDescribeAuthDomainsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseBuildServiceRequest() (request *DescribeCloudBaseBuildServiceRequest) {
    request = &DescribeCloudBaseBuildServiceRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseBuildService")
    return
}

func NewDescribeCloudBaseBuildServiceResponse() (response *DescribeCloudBaseBuildServiceResponse) {
    response = &DescribeCloudBaseBuildServiceResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取云托管代码上传url
func (c *Client) DescribeCloudBaseBuildService(request *DescribeCloudBaseBuildServiceRequest) (response *DescribeCloudBaseBuildServiceResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseBuildServiceRequest()
    }
    response = NewDescribeCloudBaseBuildServiceResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeCloudBaseRunVersionSnapshotRequest() (request *DescribeCloudBaseRunVersionSnapshotRequest) {
    request = &DescribeCloudBaseRunVersionSnapshotRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeCloudBaseRunVersionSnapshot")
    return
}

func NewDescribeCloudBaseRunVersionSnapshotResponse() (response *DescribeCloudBaseRunVersionSnapshotResponse) {
    response = &DescribeCloudBaseRunVersionSnapshotResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询版本历史
func (c *Client) DescribeCloudBaseRunVersionSnapshot(request *DescribeCloudBaseRunVersionSnapshotRequest) (response *DescribeCloudBaseRunVersionSnapshotResponse, err error) {
    if request == nil {
        request = NewDescribeCloudBaseRunVersionSnapshotRequest()
    }
    response = NewDescribeCloudBaseRunVersionSnapshotResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeDatabaseACLRequest() (request *DescribeDatabaseACLRequest) {
    request = &DescribeDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeDatabaseACL")
    return
}

func NewDescribeDatabaseACLResponse() (response *DescribeDatabaseACLResponse) {
    response = &DescribeDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取数据库权限
func (c *Client) DescribeDatabaseACL(request *DescribeDatabaseACLRequest) (response *DescribeDatabaseACLResponse, err error) {
    if request == nil {
        request = NewDescribeDatabaseACLRequest()
    }
    response = NewDescribeDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUserLoginStatisticRequest() (request *DescribeEndUserLoginStatisticRequest) {
    request = &DescribeEndUserLoginStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUserLoginStatistic")
    return
}

func NewDescribeEndUserLoginStatisticResponse() (response *DescribeEndUserLoginStatisticResponse) {
    response = &DescribeEndUserLoginStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取环境终端用户新增与登录信息
func (c *Client) DescribeEndUserLoginStatistic(request *DescribeEndUserLoginStatisticRequest) (response *DescribeEndUserLoginStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserLoginStatisticRequest()
    }
    response = NewDescribeEndUserLoginStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUserStatisticRequest() (request *DescribeEndUserStatisticRequest) {
    request = &DescribeEndUserStatisticRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUserStatistic")
    return
}

func NewDescribeEndUserStatisticResponse() (response *DescribeEndUserStatisticResponse) {
    response = &DescribeEndUserStatisticResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取终端用户总量与平台分布情况
func (c *Client) DescribeEndUserStatistic(request *DescribeEndUserStatisticRequest) (response *DescribeEndUserStatisticResponse, err error) {
    if request == nil {
        request = NewDescribeEndUserStatisticRequest()
    }
    response = NewDescribeEndUserStatisticResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEndUsersRequest() (request *DescribeEndUsersRequest) {
    request = &DescribeEndUsersRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEndUsers")
    return
}

func NewDescribeEndUsersResponse() (response *DescribeEndUsersResponse) {
    response = &DescribeEndUsersResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取终端用户列表
func (c *Client) DescribeEndUsers(request *DescribeEndUsersRequest) (response *DescribeEndUsersResponse, err error) {
    if request == nil {
        request = NewDescribeEndUsersRequest()
    }
    response = NewDescribeEndUsersResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvFreeQuotaRequest() (request *DescribeEnvFreeQuotaRequest) {
    request = &DescribeEnvFreeQuotaRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvFreeQuota")
    return
}

func NewDescribeEnvFreeQuotaResponse() (response *DescribeEnvFreeQuotaResponse) {
    response = &DescribeEnvFreeQuotaResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询后付费免费配额信息
func (c *Client) DescribeEnvFreeQuota(request *DescribeEnvFreeQuotaRequest) (response *DescribeEnvFreeQuotaResponse, err error) {
    if request == nil {
        request = NewDescribeEnvFreeQuotaRequest()
    }
    response = NewDescribeEnvFreeQuotaResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvLimitRequest() (request *DescribeEnvLimitRequest) {
    request = &DescribeEnvLimitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvLimit")
    return
}

func NewDescribeEnvLimitResponse() (response *DescribeEnvLimitResponse) {
    response = &DescribeEnvLimitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询环境个数上限
func (c *Client) DescribeEnvLimit(request *DescribeEnvLimitRequest) (response *DescribeEnvLimitResponse, err error) {
    if request == nil {
        request = NewDescribeEnvLimitRequest()
    }
    response = NewDescribeEnvLimitResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeEnvsRequest() (request *DescribeEnvsRequest) {
    request = &DescribeEnvsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeEnvs")
    return
}

func NewDescribeEnvsResponse() (response *DescribeEnvsResponse) {
    response = &DescribeEnvsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取环境列表，含环境下的各个资源信息。尤其是各资源的唯一标识，是请求各资源的关键参数
func (c *Client) DescribeEnvs(request *DescribeEnvsRequest) (response *DescribeEnvsResponse, err error) {
    if request == nil {
        request = NewDescribeEnvsRequest()
    }
    response = NewDescribeEnvsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeExtraPkgBillingInfoRequest() (request *DescribeExtraPkgBillingInfoRequest) {
    request = &DescribeExtraPkgBillingInfoRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeExtraPkgBillingInfo")
    return
}

func NewDescribeExtraPkgBillingInfoResponse() (response *DescribeExtraPkgBillingInfoResponse) {
    response = &DescribeExtraPkgBillingInfoResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取增值包计费相关信息
func (c *Client) DescribeExtraPkgBillingInfo(request *DescribeExtraPkgBillingInfoRequest) (response *DescribeExtraPkgBillingInfoResponse, err error) {
    if request == nil {
        request = NewDescribeExtraPkgBillingInfoRequest()
    }
    response = NewDescribeExtraPkgBillingInfoResponse()
    err = c.Send(request, response)
    return
}

func NewDescribePostpayPackageFreeQuotasRequest() (request *DescribePostpayPackageFreeQuotasRequest) {
    request = &DescribePostpayPackageFreeQuotasRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribePostpayPackageFreeQuotas")
    return
}

func NewDescribePostpayPackageFreeQuotasResponse() (response *DescribePostpayPackageFreeQuotasResponse) {
    response = &DescribePostpayPackageFreeQuotasResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 获取后付费免费额度
func (c *Client) DescribePostpayPackageFreeQuotas(request *DescribePostpayPackageFreeQuotasRequest) (response *DescribePostpayPackageFreeQuotasResponse, err error) {
    if request == nil {
        request = NewDescribePostpayPackageFreeQuotasRequest()
    }
    response = NewDescribePostpayPackageFreeQuotasResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeQuotaDataRequest() (request *DescribeQuotaDataRequest) {
    request = &DescribeQuotaDataRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DescribeQuotaData")
    return
}

func NewDescribeQuotaDataResponse() (response *DescribeQuotaDataResponse) {
    response = &DescribeQuotaDataResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询指定指标的配额使用量
func (c *Client) DescribeQuotaData(request *DescribeQuotaDataRequest) (response *DescribeQuotaDataResponse, err error) {
    if request == nil {
        request = NewDescribeQuotaDataRequest()
    }
    response = NewDescribeQuotaDataResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyEnvRequest() (request *DestroyEnvRequest) {
    request = &DestroyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyEnv")
    return
}

func NewDestroyEnvResponse() (response *DestroyEnvResponse) {
    response = &DestroyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁环境
func (c *Client) DestroyEnv(request *DestroyEnvRequest) (response *DestroyEnvResponse, err error) {
    if request == nil {
        request = NewDestroyEnvRequest()
    }
    response = NewDestroyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewDestroyStaticStoreRequest() (request *DestroyStaticStoreRequest) {
    request = &DestroyStaticStoreRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "DestroyStaticStore")
    return
}

func NewDestroyStaticStoreResponse() (response *DestroyStaticStoreResponse) {
    response = &DestroyStaticStoreResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 销毁静态托管资源，该接口创建异步销毁任务，资源最终状态可从DestroyStaticStore接口查看
func (c *Client) DestroyStaticStore(request *DestroyStaticStoreRequest) (response *DestroyStaticStoreResponse, err error) {
    if request == nil {
        request = NewDestroyStaticStoreRequest()
    }
    response = NewDestroyStaticStoreResponse()
    err = c.Send(request, response)
    return
}

func NewModifyDatabaseACLRequest() (request *ModifyDatabaseACLRequest) {
    request = &ModifyDatabaseACLRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyDatabaseACL")
    return
}

func NewModifyDatabaseACLResponse() (response *ModifyDatabaseACLResponse) {
    response = &ModifyDatabaseACLResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改数据库权限
func (c *Client) ModifyDatabaseACL(request *ModifyDatabaseACLRequest) (response *ModifyDatabaseACLResponse, err error) {
    if request == nil {
        request = NewModifyDatabaseACLRequest()
    }
    response = NewModifyDatabaseACLResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEndUserRequest() (request *ModifyEndUserRequest) {
    request = &ModifyEndUserRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEndUser")
    return
}

func NewModifyEndUserResponse() (response *ModifyEndUserResponse) {
    response = &ModifyEndUserResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 管理终端用户
func (c *Client) ModifyEndUser(request *ModifyEndUserRequest) (response *ModifyEndUserResponse, err error) {
    if request == nil {
        request = NewModifyEndUserRequest()
    }
    response = NewModifyEndUserResponse()
    err = c.Send(request, response)
    return
}

func NewModifyEnvRequest() (request *ModifyEnvRequest) {
    request = &ModifyEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ModifyEnv")
    return
}

func NewModifyEnvResponse() (response *ModifyEnvResponse) {
    response = &ModifyEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 更新环境信息
func (c *Client) ModifyEnv(request *ModifyEnvRequest) (response *ModifyEnvResponse, err error) {
    if request == nil {
        request = NewModifyEnvRequest()
    }
    response = NewModifyEnvResponse()
    err = c.Send(request, response)
    return
}

func NewReinstateEnvRequest() (request *ReinstateEnvRequest) {
    request = &ReinstateEnvRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("tcb", APIVersion, "ReinstateEnv")
    return
}

func NewReinstateEnvResponse() (response *ReinstateEnvResponse) {
    response = &ReinstateEnvResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 针对已隔离的免费环境，可以通过本接口将其恢复访问。
func (c *Client) ReinstateEnv(request *ReinstateEnvRequest) (response *ReinstateEnvResponse, err error) {
    if request == nil {
        request = NewReinstateEnvRequest()
    }
    response = NewReinstateEnvResponse()
    err = c.Send(request, response)
    return
}
