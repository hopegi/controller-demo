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

package v20200304

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type ArtifactReduction struct {

	// 去毛刺方式：weak,,strong
	Type *string `json:"Type,omitempty" name:"Type"`

	// 去毛刺算法，可选项：
	// edaf,
	// wdaf，
	// 默认edaf。
	// 注意：edaf：速度快，去毛刺效果强，保护边缘效果较弱；
	// wdaf：速度慢，保护边缘效果好
	Algorithm *string `json:"Algorithm,omitempty" name:"Algorithm"`
}

type AudioInfo struct {

	// 音频码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 注意：当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频编码器，可选项：aac,mp3,ac3,flac,mp2。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 声道数，可选项：
	// 1：单声道，
	// 2：双声道，
	// 6：立体声。
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 采样率，单位：Hz。可选项：32000，44100,48000
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`
}

type AudioInfoResultItem struct {

	// 音频流的流id
	Stream *int64 `json:"Stream,omitempty" name:"Stream"`

	// 音频采样率 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Sample *int64 `json:"Sample,omitempty" name:"Sample"`

	// 音频声道数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Channel *int64 `json:"Channel,omitempty" name:"Channel"`

	// 编码格式，如aac, mp3等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频时长，单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`
}

type CallbackInfo struct {

	// 回调URL。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type ClassificationEditingInfo struct {

	// 是否开启视频分类识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type ClassificationTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 视频分类识别结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*ClassificationTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type ClassificationTaskResultItem struct {

	// 分类名称。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type ColorEnhance struct {

	// 颜色增强类型，可选项：weak,strong。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type CosAuthMode struct {

	// 授权类型，可选值： 
	// 0：bucket授权，需要将对应bucket授权给本服务帐号（3020447271），否则会读写cos失败； 
	// 1：key托管，把cos的账号id和key托管于本服务，本服务会提供一个托管id； 
	// 3：临时key授权。
	// 注意：目前智能编辑还不支持临时key授权；画质重生目前只支持bucket授权
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// cos账号托管id，Type等于1时必选。
	HostedId *string `json:"HostedId,omitempty" name:"HostedId"`

	// cos身份识别id，Type等于3时必选。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// cos身份秘钥，Type等于3时必选。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// 临时授权 token，Type等于3时必选。
	Token *string `json:"Token,omitempty" name:"Token"`
}

type CosInfo struct {

	// cos 区域值。例如：ap-beijing。
	Region *string `json:"Region,omitempty" name:"Region"`

	// cos 存储桶，格式为BuketName-AppId。例如：test-123456。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// cos 路径。 
	// 对于写表示目录，例如：/test； 
	// 对于读表示文件路径，例如：/test/test.mp4。
	Path *string `json:"Path,omitempty" name:"Path"`

	// cos 授权信息，不填默认为公有权限。
	CosAuthMode *CosAuthMode `json:"CosAuthMode,omitempty" name:"CosAuthMode"`
}

type CoverEditingInfo struct {

	// 是否开启智能封面。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type CoverTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能封面结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*CoverTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type CoverTaskResultItem struct {

	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type CreateEditingTaskRequest struct {
	*tchttp.BaseRequest

	// 智能编辑任务参数。
	EditingInfo *EditingInfo `json:"EditingInfo,omitempty" name:"EditingInfo"`

	// 视频源信息。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 结果存储信息。对于包含智能拆条、智能集锦或者智能封面的任务必选。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

func (r *CreateEditingTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEditingTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateEditingTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑任务 ID，可以通过该 ID 查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateEditingTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateEditingTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaQualityRestorationTaskRequest struct {
	*tchttp.BaseRequest

	// 源文件地址。
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 画质重生任务参数信息。
	TransInfo []*SubTaskTranscodeInfo `json:"TransInfo,omitempty" name:"TransInfo" list`

	// 任务结束后文件存储信息。
	SaveInfo *SaveInfo `json:"SaveInfo,omitempty" name:"SaveInfo"`

	// 任务结果回调地址信息。
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

func (r *CreateMediaQualityRestorationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaQualityRestorationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateMediaQualityRestorationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 画质重生任务ID，可以通过该ID查询任务状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateMediaQualityRestorationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateMediaQualityRestorationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateQualityControlTaskRequest struct {
	*tchttp.BaseRequest

	// 质检任务参数
	QualityControlInfo *QualityControlInfo `json:"QualityControlInfo,omitempty" name:"QualityControlInfo"`

	// 视频源信息
	DownInfo *DownInfo `json:"DownInfo,omitempty" name:"DownInfo"`

	// 任务结果回调地址信息
	CallbackInfo *CallbackInfo `json:"CallbackInfo,omitempty" name:"CallbackInfo"`
}

func (r *CreateQualityControlTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateQualityControlTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateQualityControlTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 质检任务 ID
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateQualityControlTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateQualityControlTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DarInfo struct {

	// 填充模式，可选值：
	// 1：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。
	// 2：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“。
	// 默认为2。
	FillMode *uint64 `json:"FillMode,omitempty" name:"FillMode"`
}

type Denoising struct {

	// 去噪方式，可选项：
	// templ：时域降噪；
	// spatial：空域降噪,
	// fast-spatial：快速空域降噪。
	// 注意：可选择组合方式：
	// 1.type:"templ,spatial" ;
	// 2.type:"templ,fast-spatial"。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 时域去噪强度，可选值：0.0-1.0 。小于0.0的默认为0.0，大于1.0的默认为1.0。
	TemplStrength *float64 `json:"TemplStrength,omitempty" name:"TemplStrength"`

	// 空域去噪强度，可选值：0.0-1.0 。小于0.0的默认为0.0，大于1.0的默认为1.0。
	SpatialStrength *float64 `json:"SpatialStrength,omitempty" name:"SpatialStrength"`
}

type DescribeEditingTaskResultRequest struct {
	*tchttp.BaseRequest

	// 编辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeEditingTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEditingTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeEditingTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑任务结果信息。
		TaskResult *EditingTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeEditingTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeEditingTaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaQualityRestorationTaskRusultRequest struct {
	*tchttp.BaseRequest

	// 画质重生任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeMediaQualityRestorationTaskRusultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaQualityRestorationTaskRusultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaQualityRestorationTaskRusultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 画质重生任务结果信息
		TaskResult *MediaQualityRestorationTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaQualityRestorationTaskRusultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaQualityRestorationTaskRusultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQualityControlTaskResultRequest struct {
	*tchttp.BaseRequest

	// 质检任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *DescribeQualityControlTaskResultRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQualityControlTaskResultRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeQualityControlTaskResultResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 质检任务结果信息
		TaskResult *QualityControlInfoTaskResult `json:"TaskResult,omitempty" name:"TaskResult"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeQualityControlTaskResultResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeQualityControlTaskResultResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DownInfo struct {

	// 下载类型，可选值： 
	// 0：UrlInfo； 
	// 1：CosInfo。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Url形式下载信息，当Type等于0时必选。
	UrlInfo *UrlInfo `json:"UrlInfo,omitempty" name:"UrlInfo"`

	// Cos形式下载信息，当Type等于1时必选。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type EditInfo struct {

	// 剪辑开始时间，单位：ms。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 剪辑结束时间，单位：ms
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type EditingInfo struct {

	// 视频标签识别任务参数，不填则不开启。
	TagEditingInfo *TagEditingInfo `json:"TagEditingInfo,omitempty" name:"TagEditingInfo"`

	// 视频分类识别任务参数，不填则不开启。
	ClassificationEditingInfo *ClassificationEditingInfo `json:"ClassificationEditingInfo,omitempty" name:"ClassificationEditingInfo"`

	// 智能拆条任务参数，不填则不开启。
	StripEditingInfo *StripEditingInfo `json:"StripEditingInfo,omitempty" name:"StripEditingInfo"`

	// 智能集锦任务参数，不填则不开启。
	HighlightsEditingInfo *HighlightsEditingInfo `json:"HighlightsEditingInfo,omitempty" name:"HighlightsEditingInfo"`

	// 智能封面任务参数，不填则不开启。
	CoverEditingInfo *CoverEditingInfo `json:"CoverEditingInfo,omitempty" name:"CoverEditingInfo"`

	// 片头片尾识别任务参数，不填则不开启。
	OpeningEndingEditingInfo *OpeningEndingEditingInfo `json:"OpeningEndingEditingInfo,omitempty" name:"OpeningEndingEditingInfo"`
}

type EditingTaskResult struct {

	// 编辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 编辑任务状态。 
	// 1：执行中；2：已完成。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 视频标签识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagTaskResult *TagTaskResult `json:"TagTaskResult,omitempty" name:"TagTaskResult"`

	// 视频分类识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationTaskResult *ClassificationTaskResult `json:"ClassificationTaskResult,omitempty" name:"ClassificationTaskResult"`

	// 智能拆条结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StripTaskResult *StripTaskResult `json:"StripTaskResult,omitempty" name:"StripTaskResult"`

	// 智能集锦结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighlightsTaskResult *HighlightsTaskResult `json:"HighlightsTaskResult,omitempty" name:"HighlightsTaskResult"`

	// 智能封面结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverTaskResult *CoverTaskResult `json:"CoverTaskResult,omitempty" name:"CoverTaskResult"`

	// 片头片尾识别结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OpeningEndingTaskResult *OpeningEndingTaskResult `json:"OpeningEndingTaskResult,omitempty" name:"OpeningEndingTaskResult"`
}

type FaceProtect struct {

	// 人脸区域增强强度，可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	FaceUsmRatio *float64 `json:"FaceUsmRatio,omitempty" name:"FaceUsmRatio"`
}

type FileInfo struct {

	// 任务结束后生成的文件大小。
	// 注意：此字段可能返回 null，表示取不到有效值 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileSize *int64 `json:"FileSize,omitempty" name:"FileSize"`

	// 任务结束后生成的文件格式，例如：mp4,flv等等。
	// 注意：此字段可能返回 null，表示取不到有效值 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 任务结束后生成的文件整体码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 任务结束后生成的文件时长，单位：ms。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 任务结束后生成的文件视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoInfoResult []*VideoInfoResultItem `json:"VideoInfoResult,omitempty" name:"VideoInfoResult" list`

	// 任务结束后生成的文件音频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioInfoResult []*AudioInfoResultItem `json:"AudioInfoResult,omitempty" name:"AudioInfoResult" list`
}

type HighlightsEditingInfo struct {

	// 是否开启智能集锦。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type HighlightsTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能集锦结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*HighlightsTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type HighlightsTaskResultItem struct {

	// 智能集锦地址。
	HighlightUrl *string `json:"HighlightUrl,omitempty" name:"HighlightUrl"`

	// 智能集锦封面地址。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 智能集锦持续时间，单位：秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 智能集锦子片段结果集，集锦片段由这些子片段拼接生成。
	SegmentSet []*HighlightsTaskResultItemSegment `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type HighlightsTaskResultItemSegment struct {

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 集锦片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 集锦片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type LowLightEnhance struct {

	// 低光照增强类型，可选项：normal。
	Type *string `json:"Type,omitempty" name:"Type"`
}

type MediaQualityRestorationTaskResult struct {

	// 画质重生任务ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 画质重生处理后文件的详细信息。
	SubTaskResult []*SubTaskResultItem `json:"SubTaskResult,omitempty" name:"SubTaskResult" list`
}

type MuxInfo struct {

	// 删除流，可选项：video,audio。
	DeleteStream *string `json:"DeleteStream,omitempty" name:"DeleteStream"`
}

type OpeningEndingEditingInfo struct {

	// 是否开启片头片尾识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type OpeningEndingTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 片头片尾识别结果项。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Item *OpeningEndingTaskResultItem `json:"Item,omitempty" name:"Item"`
}

type OpeningEndingTaskResultItem struct {

	// 视频片头的结束时间点，单位：秒。
	OpeningTimeOffset *float64 `json:"OpeningTimeOffset,omitempty" name:"OpeningTimeOffset"`

	// 片头识别置信度，取值范围是 0 到 100。
	OpeningConfidence *float64 `json:"OpeningConfidence,omitempty" name:"OpeningConfidence"`

	// 视频片尾的开始时间点，单位：秒。
	EndingTimeOffset *float64 `json:"EndingTimeOffset,omitempty" name:"EndingTimeOffset"`

	// 片尾识别置信度，取值范围是 0 到 100。
	EndingConfidence *float64 `json:"EndingConfidence,omitempty" name:"EndingConfidence"`
}

type PicMarkInfoItem struct {

	// 图片水印的X坐标。
	PosX *int64 `json:"PosX,omitempty" name:"PosX"`

	// 图片水印的Y坐标 。
	PosY *int64 `json:"PosY,omitempty" name:"PosY"`

	// 图片水印路径,，如果不从Cos拉取水印，则必填
	Path *string `json:"Path,omitempty" name:"Path"`

	// 图片水印的Cos 信息，从Cos上拉取图片水印时必填。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`

	// 图片水印宽度，不填为图片原始宽度。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 图片水印高度，不填为图片原始高度。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 添加图片水印的开始时间,单位：ms。
	StartTime *int64 `json:"StartTime,omitempty" name:"StartTime"`

	// 添加图片水印的结束时间,单位：ms。
	EndTime *int64 `json:"EndTime,omitempty" name:"EndTime"`
}

type QualityControlInfo struct {

	// 对流进行截图的间隔ms，默认1000ms
	Interval *uint64 `json:"Interval,omitempty" name:"Interval"`

	// 是否保存截图
	VideoShot *bool `json:"VideoShot,omitempty" name:"VideoShot"`

	// 是否检测抖动重影
	Jitter *bool `json:"Jitter,omitempty" name:"Jitter"`

	// 是否检测模糊
	Blur *bool `json:"Blur,omitempty" name:"Blur"`

	// 是否检测低光照、过曝
	AbnormalLighting *bool `json:"AbnormalLighting,omitempty" name:"AbnormalLighting"`

	// 是否检测花屏
	CrashScreen *bool `json:"CrashScreen,omitempty" name:"CrashScreen"`

	// 是否检测黑边、白边、黑屏、白屏、绿屏
	BlackWhiteEdge *bool `json:"BlackWhiteEdge,omitempty" name:"BlackWhiteEdge"`

	// 是否检测噪点
	Noise *bool `json:"Noise,omitempty" name:"Noise"`

	// 是否检测马赛克
	Mosaic *bool `json:"Mosaic,omitempty" name:"Mosaic"`

	// 是否检测二维码，包括小程序码、条形码
	QRCode *bool `json:"QRCode,omitempty" name:"QRCode"`

	// 是否开启画面质量评价
	QualityEvaluation *bool `json:"QualityEvaluation,omitempty" name:"QualityEvaluation"`

	// 画面质量评价过滤阈值，结果只返回低于阈值的时间段，默认60
	QualityEvalScore *uint64 `json:"QualityEvalScore,omitempty" name:"QualityEvalScore"`

	// 是否检测视频音频，包含静音、低音、爆音
	Voice *bool `json:"Voice,omitempty" name:"Voice"`
}

type QualityControlInfoTaskResult struct {

	// 质检任务 ID
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 质检任务状态。
	// 1：执行中；2：成功；3：失败
	Status *uint64 `json:"Status,omitempty" name:"Status"`

	// 表示处理进度百分比
	Progress *uint64 `json:"Progress,omitempty" name:"Progress"`

	// 处理时长(s)
	UsedTime *uint64 `json:"UsedTime,omitempty" name:"UsedTime"`

	// 计费时长(s)
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 为true时表示视频无音频轨
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoAudio *bool `json:"NoAudio,omitempty" name:"NoAudio"`

	// 为true时表示视频无视频轨
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoVideo *bool `json:"NoVideo,omitempty" name:"NoVideo"`

	// 视频无参考质量打分，百分制
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationScore *uint64 `json:"QualityEvaluationScore,omitempty" name:"QualityEvaluationScore"`

	// 视频画面无参考评分低于阈值的时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	QualityEvaluationResults []*QualityControlResultItems `json:"QualityEvaluationResults,omitempty" name:"QualityEvaluationResults" list`

	// 视频画面抖动时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	JitterResults []*QualityControlResultItems `json:"JitterResults,omitempty" name:"JitterResults" list`

	// 视频画面模糊时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlurResults []*QualityControlResultItems `json:"BlurResults,omitempty" name:"BlurResults" list`

	// 视频画面低光、过曝时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	AbnormalLightingResults []*QualityControlResultItems `json:"AbnormalLightingResults,omitempty" name:"AbnormalLightingResults" list`

	// 视频画面花屏时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	CrashScreenResults []*QualityControlResultItems `json:"CrashScreenResults,omitempty" name:"CrashScreenResults" list`

	// 视频画面黑边、白边、黑屏、白屏、纯色屏时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	BlackWhiteEdgeResults []*QualityControlResultItems `json:"BlackWhiteEdgeResults,omitempty" name:"BlackWhiteEdgeResults" list`

	// 视频画面有噪点时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	NoiseResults []*QualityControlResultItems `json:"NoiseResults,omitempty" name:"NoiseResults" list`

	// 视频画面有马赛克时间段
	// 注意：此字段可能返回 null，表示取不到有效值。
	MosaicResults []*QualityControlResultItems `json:"MosaicResults,omitempty" name:"MosaicResults" list`

	// 视频画面有二维码的时间段，包括小程序码、条形码
	// 注意：此字段可能返回 null，表示取不到有效值。
	QRCodeResults []*QualityControlResultItems `json:"QRCodeResults,omitempty" name:"QRCodeResults" list`

	// 视频音频异常时间段，包括静音、低音、爆音
	// 注意：此字段可能返回 null，表示取不到有效值。
	VoiceResults []*QualityControlResultItems `json:"VoiceResults,omitempty" name:"VoiceResults" list`

	// 任务错误码
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 任务错误信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`
}

type QualityControlItem struct {

	// 置信度，取值范围是 0 到 100
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *uint64 `json:"Confidence,omitempty" name:"Confidence"`

	// 出现的起始时间戳，秒
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 出现的结束时间戳，秒
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 区域坐标(px)，即左上角坐标、右下角坐标
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordsSet []*uint64 `json:"AreaCoordsSet,omitempty" name:"AreaCoordsSet" list`
}

type QualityControlResultItems struct {

	// 异常类型
	// 注意：此字段可能返回 null，表示取不到有效值。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 质检结果项
	QualityControlItems []*QualityControlItem `json:"QualityControlItems,omitempty" name:"QualityControlItems" list`
}

type SaveInfo struct {

	// 存储类型，可选值： 
	// 1：CosInfo。
	Type *int64 `json:"Type,omitempty" name:"Type"`

	// Cos形式存储信息，当Type等于1时必选。
	CosInfo *CosInfo `json:"CosInfo,omitempty" name:"CosInfo"`
}

type ScratchRepair struct {

	// 去划痕方式，取值：normal。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 去划痕强度， 可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type SegmentInfo struct {

	// 每个切片平均时长，默认10s。
	FragmentTime *int64 `json:"FragmentTime,omitempty" name:"FragmentTime"`

	// 切片类型，可选项：hls，不填时默认hls。
	SegmentType *string `json:"SegmentType,omitempty" name:"SegmentType"`

	// 切片文件名字。注意：
	// 1.不填切片文件名时，默认按照按照如下格式命名：m3u8文件名{order}。
	// 2.若填了切片文件名字，则会按照如下格式命名：用户指定文件名{order}。
	FragmentName *string `json:"FragmentName,omitempty" name:"FragmentName"`
}

type Sharp struct {

	// 细节增强方式,取值：normal。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 细节增强强度，可选项：0.0-1.0。小于0.0的默认为0.0，大于1.0的默认为1.0。
	Ratio *float64 `json:"Ratio,omitempty" name:"Ratio"`
}

type StopMediaQualityRestorationTaskRequest struct {
	*tchttp.BaseRequest

	// 要删除的画质重生任务ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`
}

func (r *StopMediaQualityRestorationTaskRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMediaQualityRestorationTaskRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StopMediaQualityRestorationTaskResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *StopMediaQualityRestorationTaskResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *StopMediaQualityRestorationTaskResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type StripEditingInfo struct {

	// 是否开启智能拆条。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type StripTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 智能拆条结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*StripTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type StripTaskResultItem struct {

	// 视频拆条片段地址。
	SegmentUrl *string `json:"SegmentUrl,omitempty" name:"SegmentUrl"`

	// 拆条封面图片地址。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 拆条片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 拆条片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type SubTaskResultItem struct {

	// 子任务名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 子任务状态。
	// 0：成功；
	// 1：执行中；
	// 其他值：失败。
	StatusCode *int64 `json:"StatusCode,omitempty" name:"StatusCode"`

	// 子任务状态描述。
	StatusMsg *string `json:"StatusMsg,omitempty" name:"StatusMsg"`

	// 子任务进度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProgressRate *int64 `json:"ProgressRate,omitempty" name:"ProgressRate"`

	// 画质重生处理后文件的下载地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	DownloadUrl *string `json:"DownloadUrl,omitempty" name:"DownloadUrl"`

	// 画质重生处理后文件的MD5。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 画质重生处理后文件的详细信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfo *FileInfo `json:"FileInfo,omitempty" name:"FileInfo"`
}

type SubTaskTranscodeInfo struct {

	// 子任务名称。
	TaskName *string `json:"TaskName,omitempty" name:"TaskName"`

	// 目标文件信息。
	TargetInfo *TargetInfo `json:"TargetInfo,omitempty" name:"TargetInfo"`

	// 视频剪辑信息。注意：如果填写了EditInfo，则VideoInfo和AudioInfo必填
	EditInfo *EditInfo `json:"EditInfo,omitempty" name:"EditInfo"`

	// 视频转码信息，不填保持和源文件一致。
	VideoInfo *VideoInfo `json:"VideoInfo,omitempty" name:"VideoInfo"`

	// 音频转码信息，不填保持和源文件一致。
	AudioInfo *AudioInfo `json:"AudioInfo,omitempty" name:"AudioInfo"`

	// 指定封装信息。
	MuxInfo *MuxInfo `json:"MuxInfo,omitempty" name:"MuxInfo"`
}

type TagEditingInfo struct {

	// 是否开启视频标签识别。0为关闭，1为开启。其他非0非1值默认为0。
	Switch *int64 `json:"Switch,omitempty" name:"Switch"`

	// 额外定制化服务参数。参数为序列化的Json字符串，例如：{"k1":"v1"}。
	CustomInfo *string `json:"CustomInfo,omitempty" name:"CustomInfo"`
}

type TagTaskResult struct {

	// 编辑任务状态。 
	// 1：执行中；2：成功；3：失败。
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 编辑任务失败错误码。 
	// 0：成功；其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 编辑任务失败错误描述。
	ErrMsg *string `json:"ErrMsg,omitempty" name:"ErrMsg"`

	// 视频标签识别结果集。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ItemSet []*TagTaskResultItem `json:"ItemSet,omitempty" name:"ItemSet" list`
}

type TagTaskResultItem struct {

	// 标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 置信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type TargetInfo struct {

	// 目标文件名
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 目标文件切片信息
	SegmentInfo *SegmentInfo `json:"SegmentInfo,omitempty" name:"SegmentInfo"`
}

type UrlInfo struct {

	// 视频 URL。音视频支持mp4、ts等格式；直播流支持flv、rtmp格式。
	// 注意：目前智能编辑还不支持直播流场景。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 视频地址格式，可选值： 
	// 0：音视频 ;
	// 1：直播流。 
	// 默认为0。其他非0非1值默认为0。画质重生任务只支持0。
	Format *int64 `json:"Format,omitempty" name:"Format"`

	// 指定请求资源时，HTTP头部host的值。
	Host *string `json:"Host,omitempty" name:"Host"`
}

type VideoEnhance struct {

	// 去编码毛刺、伪影参数。
	ArtifactReduction *ArtifactReduction `json:"ArtifactReduction,omitempty" name:"ArtifactReduction"`

	// 去噪声参数。
	Denoising *Denoising `json:"Denoising,omitempty" name:"Denoising"`

	// 颜色增强参数。
	ColorEnhance *ColorEnhance `json:"ColorEnhance,omitempty" name:"ColorEnhance"`

	// 细节增强参数。
	Sharp *Sharp `json:"Sharp,omitempty" name:"Sharp"`

	// 超分参数，可选项：2，目前仅支持2倍超分。
	WdSuperResolution *int64 `json:"WdSuperResolution,omitempty" name:"WdSuperResolution"`

	// 人脸保护信息。
	FaceProtect *FaceProtect `json:"FaceProtect,omitempty" name:"FaceProtect"`

	// 插帧，取值范围：[0, 60]，单位：Hz。
	// 注意：当取值为 0，表示帧率和原始视频保持一致。
	WdFps *int64 `json:"WdFps,omitempty" name:"WdFps"`

	// 去划痕参数
	ScratchRepair *ScratchRepair `json:"ScratchRepair,omitempty" name:"ScratchRepair"`

	// 低光照增强参数
	LowLightEnhance *LowLightEnhance `json:"LowLightEnhance,omitempty" name:"LowLightEnhance"`
}

type VideoInfo struct {

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 注意：当取值为 0，表示帧率和原始视频保持一致。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`

	// 宽度，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 Width、Height 均为 0，则分辨率同源；
	// 当 Width 为 0，Height 非 0，则 Width 按比例缩放；
	// 当 Width 非 0，Height 为 0，则 Height 按比例缩放；
	// 当 Width、Height 均非 0，则分辨率按用户指定。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 高度，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 Width、Height 均为 0，则分辨率同源；
	// 当 Width 为 0，Height 非 0，则 Width 按比例缩放；
	// 当 Width 非 0，Height 为 0，则 Height 按比例缩放；
	// 当 Width、Height 均非 0，则分辨率按用户指定。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 长边分辨率，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 LongSide、ShortSide 均为 0，则分辨率按照Width，Height；
	// 当 LongSide 为 0，ShortSide 非 0，则 LongSide 按比例缩放；
	// 当 LongSide非 0，ShortSide为 0，则 ShortSide 按比例缩放；
	// 当 LongSide、ShortSide 均非 0，则分辨率按用户指定。
	// 长短边优先级高于Weight,Height,设置长短边则忽略宽高。
	LongSide *int64 `json:"LongSide,omitempty" name:"LongSide"`

	// 短边分辨率，取值范围：0 和 [128, 4096]
	// 注意：
	// 当 LongSide、ShortSide 均为 0，则分辨率按照Width，Height；
	// 当 LongSide 为 0，ShortSide 非 0，则 LongSide 按比例缩放；
	// 当 LongSide非 0，ShortSide为 0，则 ShortSide 按比例缩放；
	// 当 LongSide、ShortSide 均非 0，则分辨率按用户指定。
	// 长短边优先级高于Weight,Height,设置长短边则忽略宽高。
	ShortSide *int64 `json:"ShortSide,omitempty" name:"ShortSide"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 固定I帧之间，视频帧数量，取值范围： [25, 2500]，如果不填，使用编码默认最优序列。
	Gop *int64 `json:"Gop,omitempty" name:"Gop"`

	// 编码器支持选项，可选值：
	// h264,
	// h265,
	// av1
	// 。
	// 不填默认h264。
	VideoCodec *string `json:"VideoCodec,omitempty" name:"VideoCodec"`

	// 图片水印。
	PicMarkInfo []*PicMarkInfoItem `json:"PicMarkInfo,omitempty" name:"PicMarkInfo" list`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。
	DarInfo *DarInfo `json:"DarInfo,omitempty" name:"DarInfo"`

	// 支持hdr,可选项：
	// hdr10,
	// hlg。
	// 此时，VideoCodec会强制设置为h265, 编码位深为10
	Hdr *string `json:"Hdr,omitempty" name:"Hdr"`

	// 画质增强参数信息。
	VideoEnhance *VideoEnhance `json:"VideoEnhance,omitempty" name:"VideoEnhance"`
}

type VideoInfoResultItem struct {

	// 视频流的流id。
	Stream *int64 `json:"Stream,omitempty" name:"Stream"`

	// 视频宽度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频高度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频帧率，用分数格式表示，如：25/1, 99/32等等。
	// 注意：此字段可能返回 null，表示取不到有效值 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *string `json:"Fps,omitempty" name:"Fps"`

	// 编码格式，如h264,h265等等 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 播放旋转角度，可选值0-360。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频时长，单位：ms 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *int64 `json:"Duration,omitempty" name:"Duration"`

	// 颜色空间，如yuv420p，yuv444p等等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PixFormat *string `json:"PixFormat,omitempty" name:"PixFormat"`
}
