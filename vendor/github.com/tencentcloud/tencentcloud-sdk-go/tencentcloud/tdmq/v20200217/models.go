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

package v20200217

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type Connection struct {

	// 生产者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Address *string `json:"Address,omitempty" name:"Address"`

	// 主题分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`

	// 生产者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`

	// 生产者ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerId *string `json:"ProducerId,omitempty" name:"ProducerId"`

	// 消息平均大小(byte)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 生成速率(byte/秒)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`
}

type Consumer struct {

	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`

	// 消费者版本。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClientVersion *string `json:"ClientVersion,omitempty" name:"ClientVersion"`
}

type ConsumersSchedule struct {

	// 当前分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *uint64 `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 消息积压数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *uint64 `json:"MsgBacklog,omitempty" name:"MsgBacklog"`

	// 消费者每秒分发消息的数量之和。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 超时丢弃比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`
}

type CreateEnvironmentRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称，不支持中字以及除了短线和下划线外的特殊字符且不超过16个字符。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，最小60，最大1296000，（15天）。
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 说明，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateEnvironmentRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEnvironmentRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEnvironmentResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 未消费消息过期时间，单位：秒。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEnvironmentResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEnvironmentResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅者名称，不支持中字以及除了短线和下划线外的特殊字符且不超过150个字符。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 是否幂等创建，若否不允许创建同名的订阅关系。
	IsIdempotent *bool `json:"IsIdempotent,omitempty" name:"IsIdempotent"`

	// 备注，128个字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateSubscriptionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscriptionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateSubscriptionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 创建结果。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateSubscriptionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateSubscriptionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名，不支持中字以及除了短线和下划线外的特殊字符且不超过32个字符。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 0：非分区topic，无分区；非0：具体分区topic的分区数，最大不允许超过128。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列；
	// 5 ：事务消息。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *CreateTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 主题名。
		TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

		// 0：非分区topic，无分区；非0：具体分区topic的分区数。
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// 备注，128字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 0： 普通消息；
	// 1 ：全局顺序消息；
	// 2 ：局部顺序消息；
	// 3 ：重试队列；
	// 4 ：死信队列；
	// 5 ：事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）数组，每次最多删除20个。
	EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds" list`
}

func (r *DeleteEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEnvironmentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功删除的环境（命名空间）数组。
		EnvironmentIds []*string `json:"EnvironmentIds,omitempty" name:"EnvironmentIds" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteEnvironmentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 订阅关系集合，每次最多删除20个。
	SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets" list`
}

func (r *DeleteSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubscriptionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 成功删除的订阅关系数组。
		SubscriptionTopicSets []*SubscriptionTopic `json:"SubscriptionTopicSets,omitempty" name:"SubscriptionTopicSets" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteSubscriptionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsRequest struct {
	*tchttp.BaseRequest

	// 主题集合，每次最多删除20个。
	TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets" list`
}

func (r *DeleteTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 被删除的主题数组。
		TopicSets []*TopicRecord `json:"TopicSets,omitempty" name:"TopicSets" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`
}

func (r *DescribeEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvironmentAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 未消费消息过期时间，单位：秒，最大1296000（15天）。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 消费速率限制，单位：byte/秒，0：不限速。
		RateInByte *uint64 `json:"RateInByte,omitempty" name:"RateInByte"`

		// 消费速率限制，单位：个数/秒，0：不限速。
		RateInSize *uint64 `json:"RateInSize,omitempty" name:"RateInSize"`

		// 已消费消息保存策略，单位：小时，0：消费完马上删除。
		RetentionHours *uint64 `json:"RetentionHours,omitempty" name:"RetentionHours"`

		// 已消费消息保存策略，单位：G，0：消费完马上删除。
		RetentionSize *uint64 `json:"RetentionSize,omitempty" name:"RetentionSize"`

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 副本数。
		Replicas *uint64 `json:"Replicas,omitempty" name:"Replicas"`

		// 备注。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvironmentAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称，模糊搜索。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`
}

func (r *DescribeEnvironmentsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvironmentsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEnvironmentsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）记录数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 环境（命名空间）集合数组。
		EnvironmentSet []*Environment `json:"EnvironmentSet,omitempty" name:"EnvironmentSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEnvironmentsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEnvironmentsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProducersRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 生产者名称，模糊匹配。
	ProducerName *string `json:"ProducerName,omitempty" name:"ProducerName"`
}

func (r *DescribeProducersRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProducersRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProducersResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 生产者集合数组。
		ProducerSets []*Producer `json:"ProducerSets,omitempty" name:"ProducerSets" list`

		// 记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProducersResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProducersResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 订阅者名称，模糊匹配。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 数据过滤条件。
	Filters []*FilterSubscription `json:"Filters,omitempty" name:"Filters" list`
}

func (r *DescribeSubscriptionsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscriptionsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubscriptionsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 订阅者集合数组。
		SubscriptionSets []*Subscription `json:"SubscriptionSets,omitempty" name:"SubscriptionSets" list`

		// 数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubscriptionsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubscriptionsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名模糊匹配。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 起始下标，不填默认为0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回数量，不填则默认为10，最大值为20。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
}

func (r *DescribeTopicsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTopicsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 主题集合数组。
		TopicSets []*Topic `json:"TopicSets,omitempty" name:"TopicSets" list`

		// 主题数量。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTopicsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTopicsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Environment struct {

	// 环境（命名空间）名称
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 说明
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 未消费消息过期时间，单位：秒，最大1296000（15天）
	MsgTTL *int64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 创建时间
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type FilterSubscription struct {

	// 是否仅展示包含真实消费者的订阅。
	ConsumerHasCount *bool `json:"ConsumerHasCount,omitempty" name:"ConsumerHasCount"`

	// 是否仅展示消息堆积的订阅。
	ConsumerHasBacklog *bool `json:"ConsumerHasBacklog,omitempty" name:"ConsumerHasBacklog"`

	// 是否仅展示存在消息超期丢弃的订阅。
	ConsumerHasExpired *bool `json:"ConsumerHasExpired,omitempty" name:"ConsumerHasExpired"`
}

type ModifyEnvironmentAttributesRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 未消费消息过期时间，单位：秒，最大1296000。
	MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

	// 备注，字符串最长不超过128。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyEnvironmentAttributesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvironmentAttributesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyEnvironmentAttributesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 环境（命名空间）名称。
		EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

		// 未消费消息过期时间，单位：秒。
		MsgTTL *uint64 `json:"MsgTTL,omitempty" name:"MsgTTL"`

		// 备注，字符串最长不超过128。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyEnvironmentAttributesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyEnvironmentAttributesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 分区数，必须大于或者等于原分区数，若想维持原分区数请输入原数目，修改分区数仅对非全局顺序消息起效果，不允许超过128个分区。
	Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

	// 备注，128字符以内。
	Remark *string `json:"Remark,omitempty" name:"Remark"`
}

func (r *ModifyTopicRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTopicResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分区数
		Partitions *uint64 `json:"Partitions,omitempty" name:"Partitions"`

		// 备注，128字符以内。
		Remark *string `json:"Remark,omitempty" name:"Remark"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTopicResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTopicResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PartitionsTopic struct {

	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 子分区id。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`

	// topic类型描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`
}

type Producer struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 连接数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CountConnect *int64 `json:"CountConnect,omitempty" name:"CountConnect"`

	// 连接集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectionSets []*Connection `json:"ConnectionSets,omitempty" name:"ConnectionSets" list`
}

type ResetMsgSubOffsetByTimestampRequest struct {
	*tchttp.BaseRequest

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅者名称。
	Subscription *string `json:"Subscription,omitempty" name:"Subscription"`

	// 时间戳，精确到毫秒。
	ToTimestamp *uint64 `json:"ToTimestamp,omitempty" name:"ToTimestamp"`
}

func (r *ResetMsgSubOffsetByTimestampRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetMsgSubOffsetByTimestampRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetMsgSubOffsetByTimestampResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *bool `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetMsgSubOffsetByTimestampResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetMsgSubOffsetByTimestampResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type Subscription struct {

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 消费者开始连接的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConnectedSince *string `json:"ConnectedSince,omitempty" name:"ConnectedSince"`

	// 消费者地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerAddr *string `json:"ConsumerAddr,omitempty" name:"ConsumerAddr"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 消费者名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerName *string `json:"ConsumerName,omitempty" name:"ConsumerName"`

	// 堆积的消息数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgBacklog *string `json:"MsgBacklog,omitempty" name:"MsgBacklog"`

	// 于TTL，此订阅下没有被发送而是被丢弃的比例。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateExpired *string `json:"MsgRateExpired,omitempty" name:"MsgRateExpired"`

	// 消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 消费者每秒消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 订阅名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`

	// 消费者集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerSets []*Consumer `json:"ConsumerSets,omitempty" name:"ConsumerSets" list`

	// 是否在线。
	// 注意：此字段可能返回 null，表示取不到有效值。
	IsOnline *bool `json:"IsOnline,omitempty" name:"IsOnline"`

	// 消费进度集合。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumersScheduleSets []*ConsumersSchedule `json:"ConsumersScheduleSets,omitempty" name:"ConsumersScheduleSets" list`

	// 备注。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type SubscriptionTopic struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 订阅名称。
	SubscriptionName *string `json:"SubscriptionName,omitempty" name:"SubscriptionName"`
}

type Topic struct {

	// 最后一次间隔内发布消息的平均byte大小。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AverageMsgSize *string `json:"AverageMsgSize,omitempty" name:"AverageMsgSize"`

	// 消费者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConsumerCount *string `json:"ConsumerCount,omitempty" name:"ConsumerCount"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastConfirmedEntry *string `json:"LastConfirmedEntry,omitempty" name:"LastConfirmedEntry"`

	// 最后一个ledger创建的时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	LastLedgerCreatedTimestamp *string `json:"LastLedgerCreatedTimestamp,omitempty" name:"LastLedgerCreatedTimestamp"`

	// 本地和复制的发布者每秒发布消息的速率。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateIn *string `json:"MsgRateIn,omitempty" name:"MsgRateIn"`

	// 本地和复制的消费者每秒分发消息的数量之和。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgRateOut *string `json:"MsgRateOut,omitempty" name:"MsgRateOut"`

	// 本地和复制的发布者每秒发布消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputIn *string `json:"MsgThroughputIn,omitempty" name:"MsgThroughputIn"`

	// 本地和复制的消费者每秒分发消息的byte。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MsgThroughputOut *string `json:"MsgThroughputOut,omitempty" name:"MsgThroughputOut"`

	// 被记录下来的消息总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	NumberOfEntries *string `json:"NumberOfEntries,omitempty" name:"NumberOfEntries"`

	// 分区数<=0：topic下无子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Partitions *int64 `json:"Partitions,omitempty" name:"Partitions"`

	// 生产者数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProducerCount *string `json:"ProducerCount,omitempty" name:"ProducerCount"`

	// 以byte计算的所有消息存储总量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalSize *string `json:"TotalSize,omitempty" name:"TotalSize"`

	// 分区topic里面的子分区。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubTopicSets []*PartitionsTopic `json:"SubTopicSets,omitempty" name:"SubTopicSets" list`

	// topic类型描述：
	// 0：普通消息；
	// 1：全局顺序消息；
	// 2：局部顺序消息；
	// 3：重试队列；
	// 4：死信队列；
	// 5：事务消息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicType *uint64 `json:"TopicType,omitempty" name:"TopicType"`

	// 环境（命名空间）名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`

	// 说明，128个字符以内。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Remark *string `json:"Remark,omitempty" name:"Remark"`

	// 创建时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最近修改时间。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TopicRecord struct {

	// 环境（命名空间）名称。
	EnvironmentId *string `json:"EnvironmentId,omitempty" name:"EnvironmentId"`

	// 主题名称。
	TopicName *string `json:"TopicName,omitempty" name:"TopicName"`
}
