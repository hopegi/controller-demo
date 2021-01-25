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

package v20191112

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Alias struct {

	// 别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 别名的全局唯一资源标识符
	AliasArn *string `json:"AliasArn,omitempty" name:"AliasArn"`

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`

	// 创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 上次修改此数据对象的时间
	LastUpdatedTime *string `json:"LastUpdatedTime,omitempty" name:"LastUpdatedTime"`
}

type Asset struct {

	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包可运行的操作系统，暂时只支持CentOS7.16
	OperateSystem *string `json:"OperateSystem,omitempty" name:"OperateSystem"`

	// 生成包状态，0代表上传中，1代表上传失败，2代表上传成功
	Stauts *int64 `json:"Stauts,omitempty" name:"Stauts"`

	// 生成包大小
	Size *string `json:"Size,omitempty" name:"Size"`

	// 生成包创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 生成包绑定的Fleet个数，最小值为0
	BindFleetNum *int64 `json:"BindFleetNum,omitempty" name:"BindFleetNum"`
}

type AssetCredentials struct {

	// 临时证书密钥ID
	TmpSecretId *string `json:"TmpSecretId,omitempty" name:"TmpSecretId"`

	// 临时证书密钥Key
	TmpSecretKey *string `json:"TmpSecretKey,omitempty" name:"TmpSecretKey"`

	// 临时证书Token
	Token *string `json:"Token,omitempty" name:"Token"`
}

type AttachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 云联网账号 Uin
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 Id
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`
}

func (r *AttachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AttachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *AttachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *AttachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CcnInstanceSets struct {

	// 云联网账号 Uin
	// 注意：此字段可能返回 null，表示取不到有效值。
	AccountId *string `json:"AccountId,omitempty" name:"AccountId"`

	// 云联网 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	CcnId *string `json:"CcnId,omitempty" name:"CcnId"`

	// 云联网关联时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 云联网实例名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 云联网状态：申请中、已连接、已过期、已拒绝、已删除、失败的、关联中、解关联中、解关联失败
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`
}

type CreateAliasRequest struct {
	*tchttp.BaseRequest

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *CreateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 别名对象
		Alias *Alias `json:"Alias,omitempty" name:"Alias"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssetRequest struct {
	*tchttp.BaseRequest

	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包可运行的操作系统，暂时只有CentOS7.16
	OperateSystem *string `json:"OperateSystem,omitempty" name:"OperateSystem"`
}

func (r *CreateAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成包ID
		AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFleetRequest struct {
	*tchttp.BaseRequest

	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 描述，最小长度0，最大长度100
	Description *string `json:"Description,omitempty" name:"Description"`

	// 网络配置
	InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions" list`

	// 服务器类型，例如S5.LARGE8
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队类型，目前只支持ON_DEMAND类型
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器舰队名称，最小长度1，最大长度50
	Name *string `json:"Name,omitempty" name:"Name"`

	// 保护策略：不保护NoProtection、完全保护FullProtection、时限保护TimeLimitProtection
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// VPC 网络 Id，弃用，对等链接已不再使用
	PeerVpcId *string `json:"PeerVpcId,omitempty" name:"PeerVpcId"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 进程配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

	// VPC 子网，弃用，对等链接已不再使用
	SubNetId *string `json:"SubNetId,omitempty" name:"SubNetId"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`
}

func (r *CreateFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFleetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateFleetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队属性
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetAttributes *FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateFleetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话队列名称，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务部署（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations" list`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies" list`

	// 超时时间（单位秒，默认值为600秒）
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

func (r *CreateGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话队列
		GameServerSessionQueue *GameServerSessionQueue `json:"GameServerSessionQueue,omitempty" name:"GameServerSessionQueue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 别名ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 创建者ID，最大长度不超过1024个ASCII字符
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 舰队ID。每个请求需要指定别名ID 或者舰队 ID，如果两个同时指定时，优先选择舰队 ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性，最大长度不超过16组
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话属性详情，最大长度不超过4096个ASCII字符
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话自定义ID，最大长度不超过4096个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 幂等token，最大长度不超过48个ASCII字符
	IdempotencyToken *string `json:"IdempotencyToken,omitempty" name:"IdempotencyToken"`

	// 游戏服务器会话名称，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *CreateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Credentials struct {

	// ssh私钥
	Secret *string `json:"Secret,omitempty" name:"Secret"`

	// 用户名
	UserName *string `json:"UserName,omitempty" name:"UserName"`
}

type DeleteAliasRequest struct {
	*tchttp.BaseRequest

	// 要删除的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *DeleteAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAssetRequest struct {
	*tchttp.BaseRequest

	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DeleteAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFleetRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DeleteFleetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFleetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteFleetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteFleetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteFleetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGameServerSessionQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteGameServerSessionQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`
}

func (r *DeleteScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAliasRequest struct {
	*tchttp.BaseRequest

	// 要检索的队列别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *DescribeAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 别名对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Alias *Alias `json:"Alias,omitempty" name:"Alias"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetRequest struct {
	*tchttp.BaseRequest

	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`
}

func (r *DescribeAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成包信息
		Asset *Asset `json:"Asset,omitempty" name:"Asset"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsRequest struct {
	*tchttp.BaseRequest

	// 生成包可部署地域
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 偏移，代表页数，与asset实际数量相关
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 前端界面每页显示的最大条数，不超过100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`

	// 搜索条件，支持包ID或包名字过滤
	Filter *string `json:"Filter,omitempty" name:"Filter"`
}

func (r *DescribeAssetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAssetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生成包总数
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 生成包列表
		Assets []*Asset `json:"Assets,omitempty" name:"Assets" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAssetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAssetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 云联网实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		CcnInstanceSets []*CcnInstanceSets `json:"CcnInstanceSets,omitempty" name:"CcnInstanceSets" list`

		// 云联网实例个数，最小值为0
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetAttributesRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds" list`

	// 结果返回最大数量，最小值0，最大值1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队属性
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetAttributes []*FleetAttributes `json:"FleetAttributes,omitempty" name:"FleetAttributes" list`

		// 服务器舰队总数，最小值0
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetCapacityRequest struct {
	*tchttp.BaseRequest

	// 服务部署 Id列表
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds" list`

	// 结果返回最大数量
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署容量配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetCapacity []*FleetCapacity `json:"FleetCapacity,omitempty" name:"FleetCapacity" list`

		// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetEventsRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 分页时返回服务部署事件的数量，默认为20，最大值为100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 分页时的数据偏移量，默认为0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 返回的事件列表
		Events []*Event `json:"Events,omitempty" name:"Events" list`

		// 返回的事件总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetPortSettingsRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeFleetPortSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetPortSettingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetPortSettingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 安全组信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		InboundPermissions []*InboundPermission `json:"InboundPermissions,omitempty" name:"InboundPermissions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetPortSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetPortSettingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticDetailsRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetStatisticDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署统计详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		DetailList []*FleetStatisticDetail `json:"DetailList,omitempty" name:"DetailList" list`

		// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 统计时间类型
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetStatisticDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticFlowsRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeFleetStatisticFlowsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticFlowsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticFlowsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 流量统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UsedFlowList []*FleetStatisticFlows `json:"UsedFlowList,omitempty" name:"UsedFlowList" list`

		// 时长统计列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		UsedTimeList []*FleetStatisticTimes `json:"UsedTimeList,omitempty" name:"UsedTimeList" list`

		// 记录总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 统计时间类型，取值：小时和天
	// 注意：此字段可能返回 null，表示取不到有效值。
		TimeType *string `json:"TimeType,omitempty" name:"TimeType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetStatisticFlowsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticFlowsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticSummaryRequest struct {
	*tchttp.BaseRequest

	// 服务部署 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 查询开始时间
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 查询结束时间
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

func (r *DescribeFleetStatisticSummaryRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticSummaryRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetStatisticSummaryResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`

		// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetStatisticSummaryResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetStatisticSummaryResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetUtilizationRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Ids
	FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds" list`
}

func (r *DescribeFleetUtilizationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetUtilizationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeFleetUtilizationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队利用率
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetUtilization []*FleetUtilization `json:"FleetUtilization,omitempty" name:"FleetUtilization" list`

		// 总数，最小值0
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeFleetUtilizationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeFleetUtilizationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话详情列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessionDetails []*GameServerSessionDetail `json:"GameServerSessionDetails,omitempty" name:"GameServerSessionDetails" list`

		// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *DescribeGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionQueuesRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话队列名称数组，单个名字长度1~128
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *DescribeGameServerSessionQueuesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionQueuesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionQueuesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话队列数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessionQueues []*GameServerSessionQueue `json:"GameServerSessionQueues,omitempty" name:"GameServerSessionQueues" list`

		// 游戏服务器会话队列总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionQueuesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionQueuesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 游戏服务器会话状态(ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR)
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`
}

func (r *DescribeGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLimitRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceLimitRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLimitRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceLimitResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 限额
		Limit *int64 `json:"Limit,omitempty" name:"Limit"`

		// 详细信息
		ExtraInfos []*ExtraInfos `json:"ExtraInfos,omitempty" name:"ExtraInfos" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceLimitResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceLimitResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypesRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeInstanceTypesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstanceTypesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器实例类型列表
		InstanceTypeList []*InstanceTypeInfo `json:"InstanceTypeList,omitempty" name:"InstanceTypeList" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstanceTypesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstanceTypesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesExtendRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 返回结果偏移，最小值0
	Offset *int64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *int64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesExtendRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesExtendRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesExtendResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Instances []*InstanceExtend `json:"Instances,omitempty" name:"Instances" list`

		// 梳理信息总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesExtendResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesExtendResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// CVM实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 结果返回最大数量，最小值0，最大值100
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回结果偏移，最小值0
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		Instances []*Instance `json:"Instances,omitempty" name:"Instances" list`

		// 结果返回最大数量
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 玩家ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 玩家会话状态（RESERVED,ACTIVE,COMPLETED,TIMEDOUT）
	PlayerSessionStatusFilter *string `json:"PlayerSessionStatusFilter,omitempty" name:"PlayerSessionStatusFilter"`
}

func (r *DescribePlayerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePlayerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 玩家会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions" list`

		// 页偏移，用于查询下一页，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePlayerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePlayerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuntimeConfigurationRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DescribeRuntimeConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuntimeConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeRuntimeConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署运行配置信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeRuntimeConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeRuntimeConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 状态过滤条件，取值：ACTIVE表示活跃
	StatusFilter *string `json:"StatusFilter,omitempty" name:"StatusFilter"`

	// 返回结果偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 结果返回最大数量，最小值0，最大值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeScalingPoliciesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeScalingPoliciesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 动态扩缩容配置策略数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScalingPolicies []*ScalingPolicy `json:"ScalingPolicies,omitempty" name:"ScalingPolicies" list`

		// 动态扩缩容配置策略总数
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeScalingPoliciesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeScalingPoliciesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaRequest struct {
	*tchttp.BaseRequest

	// 资源类型
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`
}

func (r *DescribeUserQuotaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额资源信息
	// 注意：此字段可能返回 null，表示取不到有效值。
		QuotaResource *QuotaResource `json:"QuotaResource,omitempty" name:"QuotaResource"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserQuotaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotasRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeUserQuotasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeUserQuotasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 配额信息列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		QuotaResource []*QuotaResource `json:"QuotaResource,omitempty" name:"QuotaResource" list`

		// 配额信息列表总数，最小值0
		Total *uint64 `json:"Total,omitempty" name:"Total"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeUserQuotasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeUserQuotasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DesiredPlayerSession struct {

	// 与玩家会话关联的唯一玩家标识
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 开发人员定义的玩家数据
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

type DetachCcnInstancesRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`
}

func (r *DetachCcnInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachCcnInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DetachCcnInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DetachCcnInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DetachCcnInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Event struct {

	// 事件代码，支持以下的事件代码
	// 
	// - FLEET_CREATED 
	// - FLEET_STATE_DOWNLOADING 
	// - FLEET_BINARY_DOWNLOAD_FAILED 
	// - FLEET_CREATION_EXTRACTING_BUILD
	// - FLEET_CREATION_VALIDATING_RUNTIME_CONFIG
	// - FLEET_STATE_VALIDATING
	// - FLEET_STATE_BUILDING 
	// - FLEET_STATE_ACTIVATING
	// - FLEET_STATE_ACTIVE
	// - FLEET_SCALING_EVENT
	// - FLEET_STATE_ERROR
	// - FLEET_VALIDATION_LAUNCH_PATH_NOT_FOUND
	// - FLEET_ACTIVATION_FAILED_NO_INSTANCES
	// - FLEET_VPC_PEERING_SUCCEEDED
	// - FLEET_VPC_PEERING_FAILED
	// - FLEET_VPC_PEERING_DELETE
	// - FLEET_INITIALIZATION_FAILED
	// - FLEET_DELETED
	// - FLEET_STATE_DELETING
	// - FLEET_ACTIVATION_FAILED
	// - GAME_SESSION_ACTIVATION_TIMEOUT
	EventCode *string `json:"EventCode,omitempty" name:"EventCode"`

	// 事件的唯一标识 ID
	EventId *string `json:"EventId,omitempty" name:"EventId"`

	// 事件的发生时间，UTC 时间格式
	EventTime *string `json:"EventTime,omitempty" name:"EventTime"`

	// 事件的消息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 事件相关的日志存储路径
	// 注意：此字段可能返回 null，表示取不到有效值。
	PreSignedLogUrl *string `json:"PreSignedLogUrl,omitempty" name:"PreSignedLogUrl"`

	// 事件对应的资源对象唯一标识 ID，例如服务器舰队 ID
	ResourceId *string `json:"ResourceId,omitempty" name:"ResourceId"`
}

type ExtraInfos struct {

	// 实例类型，例如S5.LARGE8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 实例限额数
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalInstances *uint64 `json:"TotalInstances,omitempty" name:"TotalInstances"`
}

type FleetAttributes struct {

	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 创建服务器舰队时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 服务器舰队资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetArn *string `json:"FleetArn,omitempty" name:"FleetArn"`

	// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队类型，目前只支持ON_DEMAND
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetType *string `json:"FleetType,omitempty" name:"FleetType"`

	// 服务器类型，例如S5.LARGE8
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器舰队名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 游戏会话保护策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	NewGameServerSessionProtectionPolicy *string `json:"NewGameServerSessionProtectionPolicy,omitempty" name:"NewGameServerSessionProtectionPolicy"`

	// 操作系统类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 资源创建限制策略
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 状态：新建、下载中、验证中、生成中、激活中、活跃、异常、删除中、结束
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 服务器舰队停止状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	StoppedActions []*string `json:"StoppedActions,omitempty" name:"StoppedActions" list`

	// 服务器舰队终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// 时限保护超时时间，默认60分钟，最小值5，最大值1440
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionProtectionTimeLimit *uint64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`

	// 计费状态：未开通、已开通、异常、欠费隔离、销毁、解冻
	// 注意：此字段可能返回 null，表示取不到有效值。
	BillingStatus *string `json:"BillingStatus,omitempty" name:"BillingStatus"`
}

type FleetCapacity struct {

	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 服务器实例统计数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceCounts *InstanceCounts `json:"InstanceCounts,omitempty" name:"InstanceCounts"`

	// 服务器伸缩容间隔
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingInterval *uint64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`
}

type FleetStatisticDetail struct {

	// 舰队ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例IP
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceIP *string `json:"InstanceIP,omitempty" name:"InstanceIP"`

	// 开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`

	// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`
}

type FleetStatisticFlows struct {

	// 总流量，单位MB
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedFlowMegaBytes *float64 `json:"TotalUsedFlowMegaBytes,omitempty" name:"TotalUsedFlowMegaBytes"`

	// 统计开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`
}

type FleetStatisticTimes struct {

	// 统计开始时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginTime *string `json:"BeginTime,omitempty" name:"BeginTime"`

	// 统计总时长，单位秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalUsedTimeSeconds *string `json:"TotalUsedTimeSeconds,omitempty" name:"TotalUsedTimeSeconds"`
}

type FleetUtilization struct {

	// 游戏会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveGameServerSessionCount *uint64 `json:"ActiveGameServerSessionCount,omitempty" name:"ActiveGameServerSessionCount"`

	// 活跃进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveServerProcessCount *uint64 `json:"ActiveServerProcessCount,omitempty" name:"ActiveServerProcessCount"`

	// 当前游戏玩家数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 最大玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`
}

type GameProperty struct {

	// 属性名称，最大长度不超过32个ASCII字符
	Key *string `json:"Key,omitempty" name:"Key"`

	// 属性值，最大长度不超过96个ASCII字符
	Value *string `json:"Value,omitempty" name:"Value"`
}

type GameServerSession struct {

	// 游戏服务器会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 创建者ID，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreatorId *string `json:"CreatorId,omitempty" name:"CreatorId"`

	// 当前玩家数量，最小值不小于0
	CurrentPlayerSessionCount *uint64 `json:"CurrentPlayerSessionCount,omitempty" name:"CurrentPlayerSessionCount"`

	// CVM的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏属性，最大长度不超过16组
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话属性详情，最大长度不超过4096个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// CVM IP地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 对战进程详情，最大长度不超过400000个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略（ACCEPT_ALL,DENY_ALL）
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 端口号，最小值不小于1，最大值不超过60000
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏服务器会话状态（ACTIVE,ACTIVATING,TERMINATED,TERMINATING,ERROR）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 游戏服务器会话状态附加信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	StatusReason *string `json:"StatusReason,omitempty" name:"StatusReason"`

	// 终止的时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`

	// 实例类型，最大长度不超过128个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// 当前自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	CurrentCustomCount *int64 `json:"CurrentCustomCount,omitempty" name:"CurrentCustomCount"`

	// 最大自定义数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxCustomCount *int64 `json:"MaxCustomCount,omitempty" name:"MaxCustomCount"`

	// 权重
	// 注意：此字段可能返回 null，表示取不到有效值。
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`

	// 会话可用性状态，是否被屏蔽（Enable,Disable）
	// 注意：此字段可能返回 null，表示取不到有效值。
	AvailabilityStatus *string `json:"AvailabilityStatus,omitempty" name:"AvailabilityStatus"`
}

type GameServerSessionDetail struct {

	// 游戏服务器会话
	GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

	// 保护策略，可选（NoProtection,FullProtection）
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

type GameServerSessionPlacement struct {

	// 部署Id
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 服务部署组名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 玩家延迟
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`

	// 服务部署状态
	Status *string `json:"Status,omitempty" name:"Status"`

	// 分配给正在运行游戏会话的实例的DNS标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 游戏会话Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏会话名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 服务部署区域
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionRegion *string `json:"GameServerSessionRegion,omitempty" name:"GameServerSessionRegion"`

	// 游戏属性
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏会话数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 运行游戏会话的实例的IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 运行游戏会话的实例的端口号
	// 注意：此字段可能返回 null，表示取不到有效值。
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 游戏匹配数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	MatchmakerData *string `json:"MatchmakerData,omitempty" name:"MatchmakerData"`

	// 部署的玩家游戏数据
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlacedPlayerSessions []*PlacedPlayerSession `json:"PlacedPlayerSessions,omitempty" name:"PlacedPlayerSessions" list`

	// 开始时间
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type GameServerSessionQueue struct {

	// 服务部署组名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 服务部署组资源
	GameServerSessionQueueArn *string `json:"GameServerSessionQueueArn,omitempty" name:"GameServerSessionQueueArn"`

	// 目的fleet（可为别名）列表
	// 注意：此字段可能返回 null，表示取不到有效值。
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations" list`

	// 延迟策略集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies" list`

	// 超时时间
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

type GameServerSessionQueueDestination struct {

	// 服务部署组目的的资源描述
	// 注意：此字段可能返回 null，表示取不到有效值。
	DestinationArn *string `json:"DestinationArn,omitempty" name:"DestinationArn"`

	// 服务部署组目的的状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetStatus *string `json:"FleetStatus,omitempty" name:"FleetStatus"`
}

type GetGameServerSessionLogUrlRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID，最小长度不小于1个ASCII字符，最大长度不超过48个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`
}

func (r *GetGameServerSessionLogUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetGameServerSessionLogUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 日志下载URL，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
		PreSignedUrl *string `json:"PreSignedUrl,omitempty" name:"PreSignedUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetGameServerSessionLogUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetGameServerSessionLogUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessRequest struct {
	*tchttp.BaseRequest

	// 服务部署Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例Id
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`
}

func (r *GetInstanceAccessRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetInstanceAccessResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例登录所需要的凭据
		InstanceAccess *InstanceAccess `json:"InstanceAccess,omitempty" name:"InstanceAccess"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetInstanceAccessResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetInstanceAccessResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadCredentialsRequest struct {
	*tchttp.BaseRequest

	// 生成包所在地域，详见产品支持的 [地域列表](https://cloud.tencent.com/document/api/1165/42053#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)
	AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

	// 生成包的ZIP包名，例如：server.zip
	BucketKey *string `json:"BucketKey,omitempty" name:"BucketKey"`
}

func (r *GetUploadCredentialsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadCredentialsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadCredentialsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 上传文件授权信息Auth
		BucketAuth *string `json:"BucketAuth,omitempty" name:"BucketAuth"`

		// Bucket名字
		BucketName *string `json:"BucketName,omitempty" name:"BucketName"`

		// 生成包所在地域
		AssetRegion *string `json:"AssetRegion,omitempty" name:"AssetRegion"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUploadCredentialsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadCredentialsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadFederationTokenRequest struct {
	*tchttp.BaseRequest
}

func (r *GetUploadFederationTokenRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadFederationTokenRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type GetUploadFederationTokenResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 临时证书的过期时间，Unix 时间戳，精确到秒
		ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`

		// 临时证书
		AssetCredentials *AssetCredentials `json:"AssetCredentials,omitempty" name:"AssetCredentials"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *GetUploadFederationTokenResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *GetUploadFederationTokenResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InboundPermission struct {

	// 起始端口号，最小值1025
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 段范围，CIDR 方式划分
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型：TCP或者UDP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号，最大值60000
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type InboundPermissionAuthorization struct {

	// 起始端口号
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 端范围，CIDR方式划分
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type InboundPermissionRevocations struct {

	// 起始端口号
	FromPort *uint64 `json:"FromPort,omitempty" name:"FromPort"`

	// IP 端范围，CIDR 方式换分
	IpRange *string `json:"IpRange,omitempty" name:"IpRange"`

	// 协议类型：UDP或者TCP
	Protocol *string `json:"Protocol,omitempty" name:"Protocol"`

	// 终止端口号
	ToPort *uint64 `json:"ToPort,omitempty" name:"ToPort"`
}

type Instance struct {

	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// IP地址
	// 注意：此字段可能返回 null，表示取不到有效值。
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// dns
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 操作系统
	// 注意：此字段可能返回 null，表示取不到有效值。
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 创建时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`
}

type InstanceAccess struct {

	// 访问实例所需要的凭据
	Credentials *Credentials `json:"Credentials,omitempty" name:"Credentials"`

	// 服务部署Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例公网IP
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 操作系统
	OperatingSystem *string `json:"OperatingSystem,omitempty" name:"OperatingSystem"`
}

type InstanceCounts struct {

	// 活跃的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Active *uint64 `json:"Active,omitempty" name:"Active"`

	// 期望的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Desired *uint64 `json:"Desired,omitempty" name:"Desired"`

	// 空闲的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Idle *uint64 `json:"Idle,omitempty" name:"Idle"`

	// 服务器实例数最大限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxiNum *uint64 `json:"MaxiNum,omitempty" name:"MaxiNum"`

	// 服务器实例数最小限制
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniNum *uint64 `json:"MiniNum,omitempty" name:"MiniNum"`

	// 已开始创建，但未激活的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Pending *uint64 `json:"Pending,omitempty" name:"Pending"`

	// 结束中的服务器实例数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Terminating *uint64 `json:"Terminating,omitempty" name:"Terminating"`
}

type InstanceExtend struct {

	// 实例信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Instance *Instance `json:"Instance,omitempty" name:"Instance"`

	// 实例状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	State *string `json:"State,omitempty" name:"State"`

	// 健康进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	HealthyProcessCnt *int64 `json:"HealthyProcessCnt,omitempty" name:"HealthyProcessCnt"`

	// 活跃进程数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ActiveProcessCnt *int64 `json:"ActiveProcessCnt,omitempty" name:"ActiveProcessCnt"`

	// 当前游戏会话总数
	// 注意：此字段可能返回 null，表示取不到有效值。
	GameSessionCnt *int64 `json:"GameSessionCnt,omitempty" name:"GameSessionCnt"`

	// 最大游戏会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxGameSessionCnt *int64 `json:"MaxGameSessionCnt,omitempty" name:"MaxGameSessionCnt"`

	// 当前玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerSessionCnt *int64 `json:"PlayerSessionCnt,omitempty" name:"PlayerSessionCnt"`

	// 最大玩家会话数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaxPlayerSessionCnt *int64 `json:"MaxPlayerSessionCnt,omitempty" name:"MaxPlayerSessionCnt"`
}

type InstanceTypeInfo struct {

	// 类型名，例如“标准型SA1”
	TypeName *string `json:"TypeName,omitempty" name:"TypeName"`

	// 类型，例如"SA1.SMALL1"
	InstanceType *string `json:"InstanceType,omitempty" name:"InstanceType"`

	// CPU，例如1核就是1
	Cpu *uint64 `json:"Cpu,omitempty" name:"Cpu"`

	// 内存，例如2G就是2
	Memory *uint64 `json:"Memory,omitempty" name:"Memory"`

	// 网络收到包,例如25万PPS就是25
	NetworkCard *uint64 `json:"NetworkCard,omitempty" name:"NetworkCard"`
}

type JoinGameServerSessionBatchRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID列表，最小1组，最大25组
	PlayerIds []*string `json:"PlayerIds,omitempty" name:"PlayerIds" list`

	// 玩家自定义数据
	PlayerDataMap *PlayerDataMap `json:"PlayerDataMap,omitempty" name:"PlayerDataMap"`
}

func (r *JoinGameServerSessionBatchRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionBatchRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionBatchResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 玩家会话列表，最大25组
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlayerSessions []*PlayerSession `json:"PlayerSessions,omitempty" name:"PlayerSessions" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionBatchResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionBatchResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 玩家ID，最大长度1024个ASCII字符
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家自定义数据，最大长度2048个ASCII字符
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`
}

func (r *JoinGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type JoinGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 玩家会话
	// 注意：此字段可能返回 null，表示取不到有效值。
		PlayerSession *PlayerSession `json:"PlayerSession,omitempty" name:"PlayerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *JoinGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *JoinGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAliasesRequest struct {
	*tchttp.BaseRequest

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 路由策略类型，有效值SIMPLE|TERMINAL
	RoutingStrategyType *string `json:"RoutingStrategyType,omitempty" name:"RoutingStrategyType"`

	// 要返回的最大结果数，最小值1
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移，默认0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 排序字段，例如CreationTime
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 排序方式，有效值asc|desc
	OrderWay *string `json:"OrderWay,omitempty" name:"OrderWay"`
}

func (r *ListAliasesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAliasesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListAliasesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 别名对象数组
	// 注意：此字段可能返回 null，表示取不到有效值。
		Aliases []*Alias `json:"Aliases,omitempty" name:"Aliases" list`

		// 总数
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListAliasesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListAliasesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFleetsRequest struct {
	*tchttp.BaseRequest

	// 生成包 Id
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 结果返回最大值，最小值0，最大值1000
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 结果返回偏移，最小值0
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`
}

func (r *ListFleetsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFleetsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ListFleetsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署 Id 列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetIds []*string `json:"FleetIds,omitempty" name:"FleetIds" list`

		// 服务部署 Id 总数，最小值0
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ListFleetsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ListFleetsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PlacedPlayerSession struct {

	// 玩家Id
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话Id
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`
}

type PlayerDataMap struct {

	// 玩家自定义数据键，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	Key *string `json:"Key,omitempty" name:"Key"`

	// 玩家自定义数据值，最小长度不小于1个ASCII字符，最大长度不超过2048个ASCII字符
	Value *string `json:"Value,omitempty" name:"Value"`
}

type PlayerLatency struct {

	// 玩家Id
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 延迟对应的区域名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	RegionIdentifier *string `json:"RegionIdentifier,omitempty" name:"RegionIdentifier"`

	// 毫秒级延迟
	LatencyInMilliseconds *uint64 `json:"LatencyInMilliseconds,omitempty" name:"LatencyInMilliseconds"`
}

type PlayerLatencyPolicy struct {

	// 任意player允许的最大延迟，单位：毫秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	MaximumIndividualPlayerLatencyMilliseconds *uint64 `json:"MaximumIndividualPlayerLatencyMilliseconds,omitempty" name:"MaximumIndividualPlayerLatencyMilliseconds"`

	// 放置新GameServerSession时强制实施策略的时间长度，单位：秒
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyDurationSeconds *uint64 `json:"PolicyDurationSeconds,omitempty" name:"PolicyDurationSeconds"`
}

type PlayerSession struct {

	// 玩家会话创建时间
	CreationTime *string `json:"CreationTime,omitempty" name:"CreationTime"`

	// 游戏服务器会话运行的DNS标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	DnsName *string `json:"DnsName,omitempty" name:"DnsName"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 游戏服务器会话运行的CVM地址
	IpAddress *string `json:"IpAddress,omitempty" name:"IpAddress"`

	// 玩家自定义数据，最大长度2048个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerData *string `json:"PlayerData,omitempty" name:"PlayerData"`

	// 玩家ID，最大长度1024个ASCII字符
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayerId *string `json:"PlayerId,omitempty" name:"PlayerId"`

	// 玩家会话ID
	PlayerSessionId *string `json:"PlayerSessionId,omitempty" name:"PlayerSessionId"`

	// 端口号，最小值不小于1，最大值不超过60000
	Port *uint64 `json:"Port,omitempty" name:"Port"`

	// 玩家会话的状态（RESERVED = 1,ACTIVE = 2,COMPLETED = 3,TIMEDOUT = 4）
	Status *string `json:"Status,omitempty" name:"Status"`

	// 玩家会话终止时间
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerminationTime *string `json:"TerminationTime,omitempty" name:"TerminationTime"`
}

type PutScalingPolicyRequest struct {
	*tchttp.BaseRequest

	// 基于规则的扩缩容配置服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 策略名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 扩缩容调整值
	ScalingAdjustment *int64 `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 扩缩容调整类型
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 扩缩容指标阈值
	Threshold *float64 `json:"Threshold,omitempty" name:"Threshold"`

	// 策略比较符，取值：>,>=,<,<=
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 持续时间长度（分钟）
	EvaluationPeriods *int64 `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 指标名称
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型，取值：TargetBased表示基于目标的策略；RuleBased表示基于规则的策略
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 扩缩容配置类型
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

func (r *PutScalingPolicyRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutScalingPolicyRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PutScalingPolicyResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 规则名称
	// 注意：此字段可能返回 null，表示取不到有效值。
		Name *string `json:"Name,omitempty" name:"Name"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PutScalingPolicyResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PutScalingPolicyResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type QuotaResource struct {

	// 资源类型，1生成包、2服务部署、3别名、4游戏服务器队列、5实例
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResourceType *uint64 `json:"ResourceType,omitempty" name:"ResourceType"`

	// 总额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	HardLimit *uint64 `json:"HardLimit,omitempty" name:"HardLimit"`

	// 剩余额度
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remaining *uint64 `json:"Remaining,omitempty" name:"Remaining"`

	// 额外信息，可能为空
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExtraInfo *string `json:"ExtraInfo,omitempty" name:"ExtraInfo"`
}

type ResolveAliasRequest struct {
	*tchttp.BaseRequest

	// 要获取fleetId的别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`
}

func (r *ResolveAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResolveAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResolveAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 别名指向的fleet的唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResolveAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResolveAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResourceCreationLimitPolicy struct {

	// 创建数量，最小值1，默认2
	NewGameServerSessionsPerCreator *uint64 `json:"NewGameServerSessionsPerCreator,omitempty" name:"NewGameServerSessionsPerCreator"`

	// 单位时间，最小值1，默认3
	PolicyPeriodInMinutes *uint64 `json:"PolicyPeriodInMinutes,omitempty" name:"PolicyPeriodInMinutes"`
}

type RoutingStrategy struct {

	// 别名的路由策略的类型SIMPLE/TERMINAL
	Type *string `json:"Type,omitempty" name:"Type"`

	// 别名指向的队列的唯一标识符
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 与终端路由策略一起使用的消息文本
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type RuntimeConfiguration struct {

	// 游戏会话进程超时
	GameServerSessionActivationTimeoutSeconds *uint64 `json:"GameServerSessionActivationTimeoutSeconds,omitempty" name:"GameServerSessionActivationTimeoutSeconds"`

	// 最大游戏会话数
	MaxConcurrentGameServerSessionActivations *uint64 `json:"MaxConcurrentGameServerSessionActivations,omitempty" name:"MaxConcurrentGameServerSessionActivations"`

	// 服务进程配置
	ServerProcesses []*ServerProcesse `json:"ServerProcesses,omitempty" name:"ServerProcesses" list`
}

type ScalingPolicy struct {

	// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 名称
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 状态
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustment *string `json:"ScalingAdjustment,omitempty" name:"ScalingAdjustment"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ScalingAdjustmentType *string `json:"ScalingAdjustmentType,omitempty" name:"ScalingAdjustmentType"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComparisonOperator *string `json:"ComparisonOperator,omitempty" name:"ComparisonOperator"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	Threshold *string `json:"Threshold,omitempty" name:"Threshold"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	EvaluationPeriods *string `json:"EvaluationPeriods,omitempty" name:"EvaluationPeriods"`

	// 保留参数
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetricName *string `json:"MetricName,omitempty" name:"MetricName"`

	// 策略类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	PolicyType *string `json:"PolicyType,omitempty" name:"PolicyType"`

	// 基于规则的配置
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetConfiguration *TargetConfiguration `json:"TargetConfiguration,omitempty" name:"TargetConfiguration"`
}

type SearchGameServerSessionsRequest struct {
	*tchttp.BaseRequest

	// 别名ID
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 单次查询记录数上限
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 页偏移，用于查询下一页
	NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

	// 搜索条件表达式，支持如下变量
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	// hasAvailablePlayerSessions 是否有可用玩家数 String 取值true或false
	// gameServerSessionProperties 游戏会话属性 String
	// 
	// 表达式String类型 等于=，不等于<>判断
	// 表示Number类型支持 =,<>,>,>=,<,<=
	FilterExpression *string `json:"FilterExpression,omitempty" name:"FilterExpression"`

	// 排序条件关键字
	// 支持排序字段
	// gameServerSessionName 游戏会话名称 String
	// gameServerSessionId 游戏会话ID String
	// maximumSessions 最大的玩家会话数 Number
	// creationTimeMillis 创建时间，单位：毫秒 Number
	// playerSessionCount 当前玩家会话数 Number
	SortExpression *string `json:"SortExpression,omitempty" name:"SortExpression"`
}

func (r *SearchGameServerSessionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchGameServerSessionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话列表
	// 注意：此字段可能返回 null，表示取不到有效值。
		GameServerSessions []*GameServerSession `json:"GameServerSessions,omitempty" name:"GameServerSessions" list`

		// 页偏移，用于查询下一页
	// 注意：此字段可能返回 null，表示取不到有效值。
		NextToken *string `json:"NextToken,omitempty" name:"NextToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchGameServerSessionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchGameServerSessionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ServerProcesse struct {

	// 并发执行数量
	ConcurrentExecutions *uint64 `json:"ConcurrentExecutions,omitempty" name:"ConcurrentExecutions"`

	// 启动路径：/local/game/ 或 C:\game\
	LaunchPath *string `json:"LaunchPath,omitempty" name:"LaunchPath"`

	// 启动参数
	Parameters *string `json:"Parameters,omitempty" name:"Parameters"`
}

type SetServerWeightRequest struct {
	*tchttp.BaseRequest

	// 服务舰队ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 权重
	Weight *int64 `json:"Weight,omitempty" name:"Weight"`
}

func (r *SetServerWeightRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetServerWeightRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SetServerWeightResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SetServerWeightResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SetServerWeightResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartFleetActionsRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 扩展策略，为空或者AUTO_SCALING
	Actions []*string `json:"Actions,omitempty" name:"Actions" list`
}

func (r *StartFleetActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartFleetActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartFleetActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartFleetActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartFleetActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 开始部署游戏服务器会话的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`

	// 游戏服务器会话队列名称
	GameServerSessionQueueName *string `json:"GameServerSessionQueueName,omitempty" name:"GameServerSessionQueueName"`

	// 游戏服务器允许同时连接到游戏会话的最大玩家数量
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 玩家游戏会话信息
	DesiredPlayerSessions []*DesiredPlayerSession `json:"DesiredPlayerSessions,omitempty" name:"DesiredPlayerSessions" list`

	// 玩家游戏会话属性
	GameProperties []*GameProperty `json:"GameProperties,omitempty" name:"GameProperties" list`

	// 游戏服务器会话数据
	GameServerSessionData *string `json:"GameServerSessionData,omitempty" name:"GameServerSessionData"`

	// 游戏服务器会话名称
	GameServerSessionName *string `json:"GameServerSessionName,omitempty" name:"GameServerSessionName"`

	// 玩家延迟
	PlayerLatencies []*PlayerLatency `json:"PlayerLatencies,omitempty" name:"PlayerLatencies" list`
}

func (r *StartGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StartGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StartGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StartGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopFleetActionsRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队扩展策略，为空或者AUTO_SCALING
	Actions []*string `json:"Actions,omitempty" name:"Actions" list`
}

func (r *StopFleetActionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopFleetActionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopFleetActionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopFleetActionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopFleetActionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话放置的唯一标识符
	PlacementId *string `json:"PlacementId,omitempty" name:"PlacementId"`
}

func (r *StopGameServerSessionPlacementRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopGameServerSessionPlacementResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 游戏服务器会话放置
		GameServerSessionPlacement *GameServerSessionPlacement `json:"GameServerSessionPlacement,omitempty" name:"GameServerSessionPlacement"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopGameServerSessionPlacementResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopGameServerSessionPlacementResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type TargetConfiguration struct {

	// 预留存率
	// 注意：此字段可能返回 null，表示取不到有效值。
	TargetValue *uint64 `json:"TargetValue,omitempty" name:"TargetValue"`
}

type UpdateAliasRequest struct {
	*tchttp.BaseRequest

	// 要更新的别名的唯一标识符
	AliasId *string `json:"AliasId,omitempty" name:"AliasId"`

	// 名字，长度不小于1字符不超过1024字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 别名的可读说明，长度不小于1字符不超过1024字符
	Description *string `json:"Description,omitempty" name:"Description"`

	// 别名的路由配置
	RoutingStrategy *RoutingStrategy `json:"RoutingStrategy,omitempty" name:"RoutingStrategy"`
}

func (r *UpdateAliasRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAliasResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 别名对象
	// 注意：此字段可能返回 null，表示取不到有效值。
		Alias *Alias `json:"Alias,omitempty" name:"Alias"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAliasResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAliasResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetRequest struct {
	*tchttp.BaseRequest

	// 生成包ID
	AssetId *string `json:"AssetId,omitempty" name:"AssetId"`

	// 生成包名字，最小长度为1，最大长度为64
	AssetName *string `json:"AssetName,omitempty" name:"AssetName"`

	// 生成包版本，最小长度为1，最大长度为64
	AssetVersion *string `json:"AssetVersion,omitempty" name:"AssetVersion"`
}

func (r *UpdateAssetRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssetRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateAssetResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateAssetResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateAssetResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetAttributesRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队描述
	Description *string `json:"Description,omitempty" name:"Description"`

	// 服务器舰队名称
	Name *string `json:"Name,omitempty" name:"Name"`

	// 新建游戏会话保护策略
	NewGameSessionProtectionPolicy *string `json:"NewGameSessionProtectionPolicy,omitempty" name:"NewGameSessionProtectionPolicy"`

	// 资源创建限制策略
	ResourceCreationLimitPolicy *ResourceCreationLimitPolicy `json:"ResourceCreationLimitPolicy,omitempty" name:"ResourceCreationLimitPolicy"`

	// 时限保护超时时间，默认60分钟
	GameServerSessionProtectionTimeLimit *int64 `json:"GameServerSessionProtectionTimeLimit,omitempty" name:"GameServerSessionProtectionTimeLimit"`
}

func (r *UpdateFleetAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFleetAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetCapacityRequest struct {
	*tchttp.BaseRequest

	// 服务部署ID
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 期望的服务器实例数
	DesiredInstances *uint64 `json:"DesiredInstances,omitempty" name:"DesiredInstances"`

	// 服务器实例数最小限制
	MinSize *uint64 `json:"MinSize,omitempty" name:"MinSize"`

	// 服务器实例数最大限制
	MaxSize *uint64 `json:"MaxSize,omitempty" name:"MaxSize"`

	// 服务器伸缩容间隔
	ScalingInterval *uint64 `json:"ScalingInterval,omitempty" name:"ScalingInterval"`
}

func (r *UpdateFleetCapacityRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetCapacityRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetCapacityResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFleetCapacityResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetCapacityResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetPortSettingsRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队 Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 新增安全组
	InboundPermissionAuthorizations []*InboundPermissionAuthorization `json:"InboundPermissionAuthorizations,omitempty" name:"InboundPermissionAuthorizations" list`

	// 移除安全组
	InboundPermissionRevocations []*InboundPermissionRevocations `json:"InboundPermissionRevocations,omitempty" name:"InboundPermissionRevocations" list`
}

func (r *UpdateFleetPortSettingsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetPortSettingsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateFleetPortSettingsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务部署 Id
	// 注意：此字段可能返回 null，表示取不到有效值。
		FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateFleetPortSettingsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateFleetPortSettingsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionQueueRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话队列名字，长度1~128
	Name *string `json:"Name,omitempty" name:"Name"`

	// 目的服务部署（可为别名）列表
	Destinations []*GameServerSessionQueueDestination `json:"Destinations,omitempty" name:"Destinations" list`

	// 延迟策略集合
	PlayerLatencyPolicies []*PlayerLatencyPolicy `json:"PlayerLatencyPolicies,omitempty" name:"PlayerLatencyPolicies" list`

	// 超时时间
	TimeoutInSeconds *uint64 `json:"TimeoutInSeconds,omitempty" name:"TimeoutInSeconds"`
}

func (r *UpdateGameServerSessionQueueRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionQueueRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionQueueResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 部署服务组对象
		GameServerSessionQueue *GameServerSessionQueue `json:"GameServerSessionQueue,omitempty" name:"GameServerSessionQueue"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGameServerSessionQueueResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionQueueResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionRequest struct {
	*tchttp.BaseRequest

	// 游戏服务器会话ID，最小长度1个ASCII字符，最大长度不超过256个ASCII字符
	GameServerSessionId *string `json:"GameServerSessionId,omitempty" name:"GameServerSessionId"`

	// 最大玩家数量，最小值不小于0
	MaximumPlayerSessionCount *uint64 `json:"MaximumPlayerSessionCount,omitempty" name:"MaximumPlayerSessionCount"`

	// 游戏服务器会话名称，最小长度不小于1个ASCII字符，最大长度不超过1024个ASCII字符
	Name *string `json:"Name,omitempty" name:"Name"`

	// 玩家会话创建策略（ACCEPT_ALL,DENY_ALL）
	PlayerSessionCreationPolicy *string `json:"PlayerSessionCreationPolicy,omitempty" name:"PlayerSessionCreationPolicy"`

	// 保护策略(NoProtection,TimeLimitProtection,FullProtection)
	ProtectionPolicy *string `json:"ProtectionPolicy,omitempty" name:"ProtectionPolicy"`
}

func (r *UpdateGameServerSessionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateGameServerSessionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 更新后的游戏会话
		GameServerSession *GameServerSession `json:"GameServerSession,omitempty" name:"GameServerSession"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateGameServerSessionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateGameServerSessionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRuntimeConfigurationRequest struct {
	*tchttp.BaseRequest

	// 服务器舰队Id
	FleetId *string `json:"FleetId,omitempty" name:"FleetId"`

	// 服务器舰队配置
	RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`
}

func (r *UpdateRuntimeConfigurationRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRuntimeConfigurationRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type UpdateRuntimeConfigurationResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 服务器舰队配置
	// 注意：此字段可能返回 null，表示取不到有效值。
		RuntimeConfiguration *RuntimeConfiguration `json:"RuntimeConfiguration,omitempty" name:"RuntimeConfiguration"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *UpdateRuntimeConfigurationResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *UpdateRuntimeConfigurationResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}
