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

package v20180717

import (
    "encoding/json"

    tchttp "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/http"
)

type AIAnalysisTemplateItem struct {

	// 智能分析模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 智能分析模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 智能分析模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AIRecognitionTemplateItem struct {

	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频内容识别模板名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频内容识别模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 头尾识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// 拆条识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// 物体识别控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// 截图时间间隔，单位：秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AdaptiveDynamicStreamingInfoItem struct {

	// 转自适应码流规格。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 打包格式，可能为 hls 和 dash 两种。
	Package *string `json:"Package,omitempty" name:"Package"`

	// 加密类型。
	DrmType *string `json:"DrmType,omitempty" name:"DrmType"`

	// 播放地址。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AdaptiveDynamicStreamingTaskInput struct {

	// 转自适应码流模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type AiAnalysisResult struct {

	// 任务的类型，可以取的值有：
	// <li>Classification：智能分类</li>
	// <li>Cover：智能封面</li>
	// <li>Tag：智能标签</li>
	// <li>FrameTag：智能按帧标签</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频内容分析智能分类任务的查询结果，当任务类型为 Classification 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassificationTask *AiAnalysisTaskClassificationResult `json:"ClassificationTask,omitempty" name:"ClassificationTask"`

	// 视频内容分析智能封面任务的查询结果，当任务类型为 Cover 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverTask *AiAnalysisTaskCoverResult `json:"CoverTask,omitempty" name:"CoverTask"`

	// 视频内容分析智能标签任务的查询结果，当任务类型为 Tag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagTask *AiAnalysisTaskTagResult `json:"TagTask,omitempty" name:"TagTask"`

	// 视频内容分析智能按帧标签任务的查询结果，当任务类型为 FrameTag 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FrameTagTask *AiAnalysisTaskFrameTagResult `json:"FrameTagTask,omitempty" name:"FrameTagTask"`

	// 视频内容分析智能按集锦任务的查询结果，当任务类型为 Highlight 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HighlightsTask []*AiAnalysisTaskHighlightResult `json:"HighlightsTask,omitempty" name:"HighlightsTask" list`
}

type AiAnalysisTaskClassificationInput struct {

	// 视频智能分类模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskClassificationOutput struct {

	// 视频智能分类列表。
	ClassificationSet []*MediaAiAnalysisClassificationItem `json:"ClassificationSet,omitempty" name:"ClassificationSet" list`
}

type AiAnalysisTaskClassificationResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能分类任务输入。
	Input *AiAnalysisTaskClassificationInput `json:"Input,omitempty" name:"Input"`

	// 智能分类任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskClassificationOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskCoverInput struct {

	// 视频智能封面模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskCoverOutput struct {

	// 智能封面列表。
	CoverSet []*MediaAiAnalysisCoverItem `json:"CoverSet,omitempty" name:"CoverSet" list`
}

type AiAnalysisTaskCoverResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能封面任务输入。
	Input *AiAnalysisTaskCoverInput `json:"Input,omitempty" name:"Input"`

	// 智能封面任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskCoverOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskFrameTagInput struct {

	// 视频智能按帧标签模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskFrameTagOutput struct {

	// 视频按帧标签列表。
	SegmentSet []*MediaAiAnalysisFrameTagSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiAnalysisTaskFrameTagResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能按帧标签任务输入。
	Input *AiAnalysisTaskFrameTagInput `json:"Input,omitempty" name:"Input"`

	// 智能按帧标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskFrameTagOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskHighlightInput struct {

	// 视频智能精彩片段模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskHighlightOutput struct {

	// 视频智能精彩片段列表。
	HighlightSet []*MediaAiAnalysisHighlightItem `json:"HighlightSet,omitempty" name:"HighlightSet" list`
}

type AiAnalysisTaskHighlightResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能精彩片段任务输入。
	Input *AiAnalysisTaskHighlightInput `json:"Input,omitempty" name:"Input"`

	// 智能精彩片段任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskHighlightOutput `json:"Output,omitempty" name:"Output"`
}

type AiAnalysisTaskInput struct {

	// 视频内容分析模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagInput struct {

	// 视频智能标签模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiAnalysisTaskTagOutput struct {

	// 视频智能标签列表。
	TagSet []*MediaAiAnalysisTagItem `json:"TagSet,omitempty" name:"TagSet" list`
}

type AiAnalysisTaskTagResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 智能标签任务输入。
	Input *AiAnalysisTaskTagInput `json:"Input,omitempty" name:"Input"`

	// 智能标签任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiAnalysisTaskTagOutput `json:"Output,omitempty" name:"Output"`
}

type AiContentReviewResult struct {

	// 任务的类型，可以取的值有：
	// <li>Porn：图片鉴黄</li>
	// <li>Terrorism：图片鉴恐</li>
	// <li>Political：图片鉴政</li>
	// <li>Porn.Asr：Asr 文字鉴黄</li>
	// <li>Porn.Ocr：Ocr 文字鉴黄</li>
	// <li>Political.Asr：Asr 文字鉴政</li>
	// <li>Political.Ocr：Ocr 文字鉴政</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频内容审核智能画面鉴黄任务的查询结果，当任务类型为 Porn 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornTask *AiReviewTaskPornResult `json:"PornTask,omitempty" name:"PornTask"`

	// 视频内容审核智能画面鉴恐任务的查询结果，当任务类型为 Terrorism 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismTask *AiReviewTaskTerrorismResult `json:"TerrorismTask,omitempty" name:"TerrorismTask"`

	// 视频内容审核智能画面鉴政任务的查询结果，当任务类型为 Political 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalTask *AiReviewTaskPoliticalResult `json:"PoliticalTask,omitempty" name:"PoliticalTask"`

	// 视频内容审核 Asr 文字鉴黄任务的查询结果，当任务类型为 Porn.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornAsrTask *AiReviewTaskPornAsrResult `json:"PornAsrTask,omitempty" name:"PornAsrTask"`

	// 视频内容审核 Ocr 文字鉴黄任务的查询结果，当任务类型为 Porn.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornOcrTask *AiReviewTaskPornOcrResult `json:"PornOcrTask,omitempty" name:"PornOcrTask"`

	// 视频内容审核 Asr 文字鉴政任务的查询结果，当任务类型为 Political.Asr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalAsrTask *AiReviewTaskPoliticalAsrResult `json:"PoliticalAsrTask,omitempty" name:"PoliticalAsrTask"`

	// 视频内容审核 Ocr 文字鉴政任务的查询结果，当任务类型为 Political.Ocr 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalOcrTask *AiReviewTaskPoliticalOcrResult `json:"PoliticalOcrTask,omitempty" name:"PoliticalOcrTask"`
}

type AiContentReviewTaskInput struct {

	// 视频内容审核模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionResult struct {

	// 任务的类型，取值范围：
	// <li>FaceRecognition：人脸识别，</li>
	// <li>AsrWordsRecognition：语音关键词识别，</li>
	// <li>OcrWordsRecognition：文本关键词识别，</li>
	// <li>AsrFullTextRecognition：语音全文识别，</li>
	// <li>OcrFullTextRecognition：文本全文识别，</li>
	// <li>HeadTailRecognition：视频片头片尾识别，</li>
	// <li>ObjectRecognition：物体识别。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频片头片尾识别结果，当 Type 为
	//  HeadTailRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	HeadTailTask *AiRecognitionTaskHeadTailResult `json:"HeadTailTask,omitempty" name:"HeadTailTask"`

	// 视频拆条识别结果，当 Type 为
	//  SegmentRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentTask *AiRecognitionTaskSegmentResult `json:"SegmentTask,omitempty" name:"SegmentTask"`

	// 人脸识别结果，当 Type 为 
	//  FaceRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceTask *AiRecognitionTaskFaceResult `json:"FaceTask,omitempty" name:"FaceTask"`

	// 语音关键词识别结果，当 Type 为
	//  AsrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrWordsTask *AiRecognitionTaskAsrWordsResult `json:"AsrWordsTask,omitempty" name:"AsrWordsTask"`

	// 语音全文识别结果，当 Type 为
	//  AsrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrFullTextTask *AiRecognitionTaskAsrFullTextResult `json:"AsrFullTextTask,omitempty" name:"AsrFullTextTask"`

	// 文本关键词识别结果，当 Type 为
	//  OcrWordsRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrWordsTask *AiRecognitionTaskOcrWordsResult `json:"OcrWordsTask,omitempty" name:"OcrWordsTask"`

	// 文本全文识别结果，当 Type 为
	//  OcrFullTextRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrFullTextTask *AiRecognitionTaskOcrFullTextResult `json:"OcrFullTextTask,omitempty" name:"OcrFullTextTask"`

	// 物体识别结果，当 Type 为
	//  ObjectRecognition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ObjectTask *AiRecognitionTaskObjectResult `json:"ObjectTask,omitempty" name:"ObjectTask"`
}

type AiRecognitionTaskAsrFullTextResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 语音全文识别任务输入信息。
	Input *AiRecognitionTaskAsrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// 语音全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrFullTextResultInput struct {

	// 语音全文识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrFullTextResultOutput struct {

	// 语音全文识别片段列表。
	SegmentSet []*AiRecognitionTaskAsrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`

	// 字幕文件 Url。
	SubtitleUrl *string `json:"SubtitleUrl,omitempty" name:"SubtitleUrl"`
}

type AiRecognitionTaskAsrFullTextSegmentItem struct {

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别文本。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskAsrWordsResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 语音关键词识别任务输入信息。
	Input *AiRecognitionTaskAsrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// 语音关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskAsrWordsResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskAsrWordsResultInput struct {

	// 语音关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskAsrWordsResultItem struct {

	// 语音关键词。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 语音关键词出现的时间片段列表。
	SegmentSet []*AiRecognitionTaskAsrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskAsrWordsResultOutput struct {

	// 语音关键词识别结果集。
	ResultSet []*AiRecognitionTaskAsrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskAsrWordsSegmentItem struct {

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type AiRecognitionTaskFaceResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 人脸识别任务输入信息。
	Input *AiRecognitionTaskFaceResultInput `json:"Input,omitempty" name:"Input"`

	// 人脸识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskFaceResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskFaceResultInput struct {

	// 人脸识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskFaceResultItem struct {

	// 人物唯一标识 ID。
	Id *string `json:"Id,omitempty" name:"Id"`

	// 人物库类型，表示识别出的人物来自哪个人物库：
	// <li>Default：默认人物库；</li>
	// <li>UserDefine：用户自定义人物库。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 人物名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人物出现的片段结果集。
	SegmentSet []*AiRecognitionTaskFaceSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskFaceResultOutput struct {

	// 智能人脸识别结果集。
	ResultSet []*AiRecognitionTaskFaceResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskFaceSegmentItem struct {

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskHeadTailResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频片头片尾识别任务输入信息。
	Input *AiRecognitionTaskHeadTailResultInput `json:"Input,omitempty" name:"Input"`

	// 视频片头片尾识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskHeadTailResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskHeadTailResultInput struct {

	// 视频片头片尾识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskHeadTailResultOutput struct {

	// 片头识别置信度。取值：0~100。
	HeadConfidence *float64 `json:"HeadConfidence,omitempty" name:"HeadConfidence"`

	// 视频片头的结束时间点，单位：秒。
	HeadTimeOffset *float64 `json:"HeadTimeOffset,omitempty" name:"HeadTimeOffset"`

	// 片尾识别置信度。取值：0~100。
	TailConfidence *float64 `json:"TailConfidence,omitempty" name:"TailConfidence"`

	// 视频片尾的开始时间点，单位：秒。
	TailTimeOffset *float64 `json:"TailTimeOffset,omitempty" name:"TailTimeOffset"`
}

type AiRecognitionTaskInput struct {

	// 视频智能识别模板 ID 。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 物体识别任务输入信息。
	Input *AiRecognitionTaskObjectResultInput `json:"Input,omitempty" name:"Input"`

	// 物体识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskObjectResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskObjectResultInput struct {

	// 物体识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskObjectResultItem struct {

	// 识别的物体名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 物体出现的片段列表。
	SegmentSet []*AiRecognitionTaskObjectSeqmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskObjectResultOutput struct {

	// 智能物体识别结果集。
	ResultSet []*AiRecognitionTaskObjectResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskObjectSeqmentItem struct {

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskOcrFullTextResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 文本全文识别任务输入信息。
	Input *AiRecognitionTaskOcrFullTextResultInput `json:"Input,omitempty" name:"Input"`

	// 文本全文识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrFullTextResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrFullTextResultInput struct {

	// 文本全文识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrFullTextResultOutput struct {

	// 文本全文识别结果集。
	SegmentSet []*AiRecognitionTaskOcrFullTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskOcrFullTextSegmentItem struct {

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别片段结果集。
	TextSet []*AiRecognitionTaskOcrFullTextSegmentTextItem `json:"TextSet,omitempty" name:"TextSet" list`
}

type AiRecognitionTaskOcrFullTextSegmentTextItem struct {

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// 识别文本。
	Text *string `json:"Text,omitempty" name:"Text"`
}

type AiRecognitionTaskOcrWordsResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 文本关键词识别任务输入信息。
	Input *AiRecognitionTaskOcrWordsResultInput `json:"Input,omitempty" name:"Input"`

	// 文本关键词识别任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskOcrWordsResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskOcrWordsResultInput struct {

	// 文本关键词识别模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskOcrWordsResultItem struct {

	// 文本关键词。
	Word *string `json:"Word,omitempty" name:"Word"`

	// 文本关键出现的片段列表。
	SegmentSet []*AiRecognitionTaskOcrWordsSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskOcrWordsResultOutput struct {

	// 文本关键词识别结果集。
	ResultSet []*AiRecognitionTaskOcrWordsResultItem `json:"ResultSet,omitempty" name:"ResultSet" list`
}

type AiRecognitionTaskOcrWordsSegmentItem struct {

	// 识别片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 识别片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 识别片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 识别结果的区域坐标。数组包含 4 个元素 [x1,y1,x2,y2]，依次表示区域左上点、右下点的横纵坐标。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type AiRecognitionTaskSegmentResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频拆条任务输入信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *AiRecognitionTaskSegmentResultInput `json:"Input,omitempty" name:"Input"`

	// 视频拆条任务输出信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiRecognitionTaskSegmentResultOutput `json:"Output,omitempty" name:"Output"`
}

type AiRecognitionTaskSegmentResultInput struct {

	// 视频拆条模板 ID。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type AiRecognitionTaskSegmentResultOutput struct {

	// 视频拆条片段列表。
	SegmentSet []*AiRecognitionTaskSegmentSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiRecognitionTaskSegmentSegmentItem struct {

	// 视频拆条片段 Url。
	SegmentUrl *string `json:"SegmentUrl,omitempty" name:"SegmentUrl"`

	// 拆条片段置信度。取值：0~100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 拆条片段起始的偏移时间，单位：秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 拆条片段终止的偏移时间，单位：秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 拆条封面图片 Url。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 特殊字段，请忽略。
	SpecialInfo *string `json:"SpecialInfo,omitempty" name:"SpecialInfo"`
}

type AiReviewPoliticalAsrTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalAsrTaskOutput struct {

	// Asr 文字涉政、敏感评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉政、敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉政、敏感嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalOcrTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalOcrTaskOutput struct {

	// Ocr 文字涉政、敏感评分，分值为0到100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉政、敏感结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉政、敏感嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPoliticalTaskInput struct {

	// 鉴政模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPoliticalTaskOutput struct {

	// 视频涉政评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 涉政结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴政结果标签，取值范围：
	// <li>politician：政治人物。</li>
	// <li>violation_photo：违规图标。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉政嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewPoliticalSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornAsrTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornAsrTaskOutput struct {

	// Asr 文字涉黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Asr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Asr 文字有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewAsrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornOcrTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornOcrTaskOutput struct {

	// Ocr 文字涉黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// Ocr 文字涉黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// Ocr 文字有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewOcrTextSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewPornTaskInput struct {

	// 鉴黄模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewPornTaskOutput struct {

	// 视频鉴黄评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频鉴黄结果标签，取值范围：
	// <li>porn：色情。</li>
	// <li>sexy：性感。</li>
	// <li>vulgar：低俗。</li>
	// <li>intimacy：亲密行为。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有涉黄嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiReviewTaskPoliticalAsrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Asr 文字鉴政任务输入。
	Input *AiReviewPoliticalAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Asr 文字鉴政任务输出。
	Output *AiReviewPoliticalAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalOcrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Ocr 文字鉴政任务输入。
	Input *AiReviewPoliticalOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Ocr 文字鉴政任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPoliticalResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴政任务输入。
	Input *AiReviewPoliticalTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴政任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPoliticalTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornAsrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Asr 文字鉴黄任务输入。
	Input *AiReviewPornAsrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Asr 文字鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornAsrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornOcrResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核 Ocr 文字鉴黄任务输入。
	Input *AiReviewPornOcrTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核 Ocr 文字鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornOcrTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskPornResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴黄任务输入。
	Input *AiReviewPornTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴黄任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewPornTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTaskTerrorismResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 内容审核鉴恐任务输入。
	Input *AiReviewTerrorismTaskInput `json:"Input,omitempty" name:"Input"`

	// 内容审核鉴恐任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *AiReviewTerrorismTaskOutput `json:"Output,omitempty" name:"Output"`
}

type AiReviewTerrorismTaskInput struct {

	// 鉴恐模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type AiReviewTerrorismTaskOutput struct {

	// 视频暴恐评分，分值为0到100。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 暴恐结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 视频暴恐结果标签，取值范围：
	// <li>guns：武器枪支。</li>
	// <li>crowd：人群聚集。</li>
	// <li>police：警察部队。</li>
	// <li>bloody：血腥画面。</li>
	// <li>banners：暴恐旗帜。</li>
	// <li>militant：武装分子。</li>
	// <li>explosion：爆炸火灾。</li>
	// <li>terrorists：暴恐人物。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 有暴恐嫌疑的视频片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SegmentSet []*MediaContentReviewSegmentItem `json:"SegmentSet,omitempty" name:"SegmentSet" list`
}

type AiSampleFaceInfo struct {

	// 人脸图片 ID。
	FaceId *string `json:"FaceId,omitempty" name:"FaceId"`

	// 人脸图片地址。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type AiSampleFaceOperation struct {

	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。重置操作将清空该人物已有人脸数据，并添加 FaceContents 指定人脸数据。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 人脸 ID 集合，当 Type为delete 时，该字段必填。
	FaceIds []*string `json:"FaceIds,omitempty" name:"FaceIds" list`

	// 人脸图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串集合。
	// <li>当 Type为add 或 reset 时，该字段必填；</li>
	// <li>数组长度限制：5 张图片。</li>
	// 注意：图片必须是单人像正面人脸较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents" list`
}

type AiSampleFailFaceInfo struct {

	// 对应入参 FaceContents 中错误图片下标，从 0 开始。
	Index *uint64 `json:"Index,omitempty" name:"Index"`

	// 错误码，取值：
	// <li>0：成功；</li>
	// <li>其他：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误描述。
	Message *string `json:"Message,omitempty" name:"Message"`
}

type AiSamplePerson struct {

	// 人物 ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 人物名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人物描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 人脸信息。
	FaceInfoSet []*AiSampleFaceInfo `json:"FaceInfoSet,omitempty" name:"FaceInfoSet" list`

	// 人物标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 应用场景。
	UsageSet []*string `json:"UsageSet,omitempty" name:"UsageSet" list`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AiSampleTagOperation struct {

	// 操作类型，可选值：add（添加）、delete（删除）、reset（重置）。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 标签，长度限制：128 个字符。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`
}

type AiSampleWord struct {

	// 关键词。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 关键词标签。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 关键词应用场景。
	UsageSet []*string `json:"UsageSet,omitempty" name:"UsageSet" list`

	// 创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type AiSampleWordInfo struct {

	// 关键词，长度限制：20 个字符。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// 关键词标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`
}

type AnimatedGraphicTaskInput struct {

	// 视频转动图模板 ID
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 动图在视频中的开始时间，单位为秒。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间，单位为秒。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type ApplyUploadRequest struct {
	*tchttp.BaseRequest

	// 媒体类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	MediaType *string `json:"MediaType,omitempty" name:"MediaType"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// 封面类型，可选值请参考 [上传能力综述](/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。
	CoverType *string `json:"CoverType,omitempty" name:"CoverType"`

	// 媒体后续任务处理操作，即完成媒体上传后，可自动发起任务流操作。参数值为任务流模板名，云点播支持 [创建任务流模板](/document/product/266/33819) 并为模板命名。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户。
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过 [创建分类](/document/product/266/7812) 接口，创建分类，获得分类 ID。
	// <li>默认值：0，表示其他分类。</li>
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，[上传完成回调](/document/product/266/7830) 将返回该字段值，最长 250 个字符。
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`

	// 会话上下文，用于透传用户请求信息，当指定 Procedure 参数后，[任务流状态变更回调](/document/product/266/9636) 将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 点播 [子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ApplyUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ApplyUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 存储桶，用于上传接口 URL 的 bucket_name。
		StorageBucket *string `json:"StorageBucket,omitempty" name:"StorageBucket"`

		// 存储园区，用于上传接口 Host 的 Region。
		StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

		// 点播会话，用于确认上传接口的参数 VodSessionKey。
		VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

		// 媒体存储路径，用于上传接口存储媒体的对象键（Key）。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaStoragePath *string `json:"MediaStoragePath,omitempty" name:"MediaStoragePath"`

		// 封面存储路径，用于上传接口存储封面的对象键（Key）。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CoverStoragePath *string `json:"CoverStoragePath,omitempty" name:"CoverStoragePath"`

		// 临时凭证，用于上传接口的权限验证。
		TempCertificate *TempCertificate `json:"TempCertificate,omitempty" name:"TempCertificate"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ApplyUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ApplyUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type AsrFullTextConfigureInfo struct {

	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 生成的字幕文件格式，不填或者填空字符串表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件。</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitempty" name:"SubtitleFormat"`
}

type AsrFullTextConfigureInfoForUpdate struct {

	// 语音全文识别任务开关，可选值：
	// <li>ON：开启智能语音全文识别任务；</li>
	// <li>OFF：关闭智能语音全文识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 生成的字幕文件格式，填空字符串表示不生成字幕文件，可选值：
	// <li>vtt：生成 WebVTT 字幕文件。</li>
	SubtitleFormat *string `json:"SubtitleFormat,omitempty" name:"SubtitleFormat"`
}

type AsrWordsConfigureInfo struct {

	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type AsrWordsConfigureInfoForUpdate struct {

	// 语音关键词识别任务开关，可选值：
	// <li>ON：开启语音关键词识别任务；</li>
	// <li>OFF：关闭语音关键词识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type AudioTemplateInfo struct {

	// 音频流的编码格式。
	// 当外层参数 Container 为 mp3 时，可选值为：
	// <li>libmp3lame。</li>
	// 当外层参数 Container 为 ogg 或 flac 时，可选值为：
	// <li>flac。</li>
	// 当外层参数 Container 为 m4a 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame；</li>
	// <li>ac3。</li>
	// 当外层参数 Container 为 mp4 或 flv 时，可选值为：
	// <li>libfdk_aac：更适合 mp4；</li>
	// <li>libmp3lame：更适合 flv；</li>
	// <li>mp2。</li>
	// 当外层参数 Container 为 hls 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame。</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。
	// 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTemplateInfoForUpdate struct {

	// 音频流的编码格式。
	// 当外层参数 Container 为 mp3 时，可选值为：
	// <li>libmp3lame。</li>
	// 当外层参数 Container 为 ogg 或 flac 时，可选值为：
	// <li>flac。</li>
	// 当外层参数 Container 为 m4a 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame；</li>
	// <li>ac3。</li>
	// 当外层参数 Container 为 mp4 或 flv 时，可选值为：
	// <li>libfdk_aac：更适合 mp4；</li>
	// <li>libmp3lame：更适合 flv；</li>
	// <li>mp2。</li>
	// 当外层参数 Container 为 hls 时，可选值为：
	// <li>libfdk_aac；</li>
	// <li>libmp3lame。</li>
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频流的码率，取值范围：0 和 [26, 256]，单位：kbps。 当取值为 0，表示音频码率和原始音频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，可选值：
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	SampleRate *uint64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频通道方式，可选值：
	// <li>1：单通道</li>
	// <li>2：双通道</li>
	// <li>6：立体声</li>
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type AudioTrackItem struct {

	// 音频素材的媒体文件来源。可以是点播的文件 ID，也可以是其它文件的 URL。
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// 音频片段取自素材文件的起始时间，单位为秒。0 表示从素材开始位置截取。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// 音频片段的时长，单位为秒。默认和素材本身长度一致，表示截取全部素材。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 对音频片段进行的操作，如音量调节等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations" list`
}

type AudioTransform struct {

	// 音频操作类型，取值有：
	// <li>Volume：音量调节。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 音量调节参数， 当 Type = Volume 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VolumeParam *AudioVolumeParam `json:"VolumeParam,omitempty" name:"VolumeParam"`
}

type AudioVolumeParam struct {

	// 音频增益，取值范围0~10。仅在Mute=0时生效。
	// <li>大于1表示增加音量。</li>
	// <li>小于1表示降低音量。</li>
	// <li>1：表示不改变。</li>
	// 默认是1。
	Gain *float64 `json:"Gain,omitempty" name:"Gain"`

	// 是否静音，取值范围0或1。
	// <li>0表示不静音。</li>
	// <li>1表示静音。</li>
	// 默认是0。
	Mute *int64 `json:"Mute,omitempty" name:"Mute"`
}

type Canvas struct {

	// 背景颜色，取值有：
	// <li>Black：黑色背景</li>
	// <li>White：白色背景</li>
	// 默认值：Black。
	Color *string `json:"Color,omitempty" name:"Color"`

	// 画布宽度，即输出视频的宽度，取值范围：0~ 4096，单位：px。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频宽度一致。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 画布高度，即输出视频的高度（或长边），取值范围：0~ 4096，单位：px。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频高度一致。
	Height *int64 `json:"Height,omitempty" name:"Height"`
}

type ClassificationConfigureInfo struct {

	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ClassificationConfigureInfoForUpdate struct {

	// 智能分类任务开关，可选值：
	// <li>ON：开启智能分类任务；</li>
	// <li>OFF：关闭智能分类任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ClipFileInfo2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 输出目标文件的文件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 输出目标文件的文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 输出目标文件的文件类型。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ClipTask2017 struct {

	// 视频剪辑任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 视频剪辑任务源文件 ID。
	SrcFileId *string `json:"SrcFileId,omitempty" name:"SrcFileId"`

	// 视频剪辑输出的文件信息。
	FileInfo *ClipFileInfo2017 `json:"FileInfo,omitempty" name:"FileInfo"`
}

type CommitUploadRequest struct {
	*tchttp.BaseRequest

	// 点播会话，取申请上传接口的返回值 VodSessionKey。
	VodSessionKey *string `json:"VodSessionKey,omitempty" name:"VodSessionKey"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CommitUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CommitUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件的唯一标识。
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// 媒体播放地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

		// 媒体封面地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CommitUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CommitUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaOutput struct {

	// 文件名称，最长 64 个字符。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 描述信息，最长 128 个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 封装格式，可选值：mp4、mp3。其中，mp3 为纯音频文件。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 输出的视频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStream *OutputVideoStream `json:"VideoStream,omitempty" name:"VideoStream"`

	// 输出的音频信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStream *OutputAudioStream `json:"AudioStream,omitempty" name:"AudioStream"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`
}

type ComposeMediaRequest struct {
	*tchttp.BaseRequest

	// 输入的媒体轨道列表，包括视频、音频、图片等素材组成的多个轨道信息。输入的多个轨道在时间轴上和输出媒体文件的时间轴对齐，时间轴上相同时间点的各个轨道的素材进行重叠，视频或者图片按轨道顺序进行图像的叠加，轨道顺序高的素材叠加在上面；音频素材进行混音。
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks" list`

	// 输出的媒体文件信息。
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`

	// 制作视频文件时使用的画布。
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ComposeMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ComposeMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 制作媒体文件的任务 ID，可以通过该 ID 查询制作任务（任务类型为 MakeMedia）的状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ComposeMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ComposeMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ComposeMediaTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 制作媒体文件任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *ComposeMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// 制作媒体文件任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *ComposeMediaTaskOutput `json:"Output,omitempty" name:"Output"`
}

type ComposeMediaTaskInput struct {

	// 输入的媒体轨道列表，包括视频、音频、图片等素材组成的多个轨道信息。
	Tracks []*MediaTrack `json:"Tracks,omitempty" name:"Tracks" list`

	// 制作视频文件时使用的画布。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Canvas *Canvas `json:"Canvas,omitempty" name:"Canvas"`

	// 输出的媒体文件信息。
	Output *ComposeMediaOutput `json:"Output,omitempty" name:"Output"`
}

type ComposeMediaTaskOutput struct {

	// 文件类型，例如 mp4、mp3 等。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件播放地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

type ConcatFileInfo2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频拼接源文件的 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频拼接源文件的地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 视频拼接源文件的格式。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`
}

type ConcatTask2017 struct {

	// 视频拼接任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 视频拼接源文件信息。
	FileInfoSet []*ConcatFileInfo2017 `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`
}

type ConfirmEventsRequest struct {
	*tchttp.BaseRequest

	// 事件句柄，数组长度限制：16。
	EventHandles []*string `json:"EventHandles,omitempty" name:"EventHandles" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ConfirmEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ConfirmEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ConfirmEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ConfirmEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ContentReviewTemplateItem struct {

	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 鉴黄控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// 鉴恐控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// 鉴政控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// 用户自定义内容审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// 审核结果是否进入审核墙（对审核结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type CoverBySnapshotTaskInput struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图方式。包含：
	// <li>Time：依照时间点截图</li>
	// <li>Percent：依照百分比截图</li>
	PositionType *string `json:"PositionType,omitempty" name:"PositionType"`

	// 截图位置：
	// <li>对于依照时间点截图，该值表示指定视频第几秒的截图作为封面</li>
	// <li>对于依照百分比截图，该值表示使用视频百分之多少的截图作为封面</li>
	PositionValue *float64 `json:"PositionValue,omitempty" name:"PositionValue"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type CoverBySnapshotTaskOutput struct {

	// 封面 URL。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`
}

type CoverConfigureInfo struct {

	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CoverConfigureInfoForUpdate struct {

	// 智能封面任务开关，可选值：
	// <li>ON：开启智能封面任务；</li>
	// <li>OFF：关闭智能封面任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type CreateAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfo `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfo `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfo `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfo `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频内容分析模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfo `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfo `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfo `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfo `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfo `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfo `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfo `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfo `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIRecognitionTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 视频内容识别模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateAIRecognitionTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassRequest struct {
	*tchttp.BaseRequest

	// 父类 ID，一级分类填写 -1。
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 分类名称，长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类 ID
		ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// 审核结果是否进入审核墙（对审核结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 鉴黄控制参数。
	PornConfigure *PornConfigureInfo `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// 鉴恐控制参数。
	TerrorismConfigure *TerrorismConfigureInfo `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// 鉴政控制参数。
	PoliticalConfigure *PoliticalConfigureInfo `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfo `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒。当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContentReviewTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 内容审核模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateContentReviewTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateImageSpriteTask2017 struct {

	// 截图雪碧图任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 截取雪碧图文件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 雪碧图规格，参见[雪碧图截图模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图小图总数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 截取雪碧图输出的地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteUrlSet []*string `json:"ImageSpriteUrlSet,omitempty" name:"ImageSpriteUrlSet" list`

	// 雪碧图子图位置与时间关系 WebVtt 文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type CreatePersonSampleRequest struct {
	*tchttp.BaseRequest

	// 人物名称，长度限制：20 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 人脸图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 jpeg、png 图片格式。数组长度限制：5 张图片。
	// 注意：图片必须是单人像正面人脸较清晰的照片，像素不低于 200*200。
	FaceContents []*string `json:"FaceContents,omitempty" name:"FaceContents" list`

	// 人物应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于内容审核，等价于 Review.Face。
	// 3. All：用于内容识别、内容审核，等价于 1+2。
	Usages []*string `json:"Usages,omitempty" name:"Usages" list`

	// 人物描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 人物标签
	// <li>数组长度限制：20 个标签；</li>
	// <li>单个标签长度限制：128 个字符。</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreatePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonSampleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreatePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人物信息。
		Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

		// 处理失败的人脸信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreatePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreatePersonSampleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// 任务流名字（支持中文，不超过20个字）。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// AI 智能内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	// 默认值：0。
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，当 RemoveVideo 为 0，该字段必填。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，当 RemoveAudio 为 0，该字段必填。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 转码模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印类型，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 默认值：TopLeft。目前，当 Type 为 image，该字段仅支持 TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，当 Type 为 image，该字段必填。当 Type 为 text，该字段无效。
	ImageTemplate *ImageWatermarkInput `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，当 Type 为 text，该字段必填。当 Type 为 image，该字段无效。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG水印模板，当 Type 为 svg，该字段必填。当 Type 为 image 或 text，该字段无效。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 水印模板唯一标识。
		Definition *int64 `json:"Definition,omitempty" name:"Definition"`

		// 水印图片地址，仅当 Type 为 image，该字段有效。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWordSamplesRequest struct {
	*tchttp.BaseRequest

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过语音识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行内容审核；
	// 4. Review.Asr：通过语音识别技术，进行内容审核；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、语音识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、语音识别技术，进行内容审核，等价于 3+4；
	// 7. All：通过光学字符识别技术、语音识别技术，进行内容识别、内容审核，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitempty" name:"Usages" list`

	// 关键词，数组长度限制：100。
	Words []*AiSampleWordInfo `json:"Words,omitempty" name:"Words" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *CreateWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWordSamplesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type CreateWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *CreateWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *CreateWordSamplesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIRecognitionTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteAIRecognitionTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContentReviewTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteContentReviewTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaRequest struct {
	*tchttp.BaseRequest

	// 媒体文件的唯一标识。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 指定本次需要删除的部分。默认值为 "[]", 表示删除媒体及其对应的全部视频处理文件。
	DeleteParts []*MediaDeleteItem `json:"DeleteParts,omitempty" name:"DeleteParts" list`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonSampleRequest struct {
	*tchttp.BaseRequest

	// 人物 ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeletePersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonSampleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeletePersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeletePersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeletePersonSampleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// 任务流名字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWordSamplesRequest struct {
	*tchttp.BaseRequest

	// 关键词，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DeleteWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWordSamplesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DeleteWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DeleteWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DeleteWordSamplesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesRequest struct {
	*tchttp.BaseRequest

	// 视频内容分析模板唯一标识过滤条件，数组长度最大值：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAIAnalysisTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIAnalysisTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIAnalysisTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 视频内容分析模板详情列表。
		AIAnalysisTemplateSet []*AIAnalysisTemplateItem `json:"AIAnalysisTemplateSet,omitempty" name:"AIAnalysisTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIAnalysisTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIAnalysisTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesRequest struct {
	*tchttp.BaseRequest

	// 视频内容识别模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAIRecognitionTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIRecognitionTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAIRecognitionTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 视频内容识别模板详情列表。
		AIRecognitionTemplateSet []*AIRecognitionTemplateItem `json:"AIRecognitionTemplateSet,omitempty" name:"AIRecognitionTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAIRecognitionTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAIRecognitionTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassRequest struct {
	*tchttp.BaseRequest

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeAllClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeAllClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 分类信息集合
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClassInfoSet []*MediaClassInfo `json:"ClassInfoSet,omitempty" name:"ClassInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeAllClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeAllClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesRequest struct {
	*tchttp.BaseRequest

	// 内容审核模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeContentReviewTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentReviewTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeContentReviewTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 内容审核模板详情列表。
		ContentReviewTemplateSet []*ContentReviewTemplateItem `json:"ContentReviewTemplateSet,omitempty" name:"ContentReviewTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeContentReviewTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeContentReviewTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID 列表，N 从 0 开始取值，最大 19。
	FileIds []*string `json:"FileIds,omitempty" name:"FileIds" list`

	// 指定所有媒体文件需要返回的信息，可同时指定多个信息，N 从 0 开始递增。如果未填写该字段，默认返回所有信息。选项有：
	// <li>basicInfo（视频基础信息）。</li>
	// <li>metaData（视频元信息）。</li>
	// <li>transcodeInfo（视频转码结果信息）。</li>
	// <li>animatedGraphicsInfo（视频转动图结果信息）。</li>
	// <li>imageSpriteInfo（视频雪碧图信息）。</li>
	// <li>snapshotByTimeOffsetInfo（视频指定时间点截图信息）。</li>
	// <li>sampleSnapshotInfo（采样截图信息）。</li>
	// <li>keyFrameDescInfo（打点信息）。</li>
	// <li>adaptiveDynamicStreamingInfo（转自适应码流信息）。</li>
	// <li>miniProgramReviewInfo（小程序审核信息）。</li>
	Filters []*string `json:"Filters,omitempty" name:"Filters" list`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeMediaInfosRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeMediaInfosResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 媒体文件信息列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// 不存在的文件 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		NotExistFileIdSet []*string `json:"NotExistFileIdSet,omitempty" name:"NotExistFileIdSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeMediaInfosResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeMediaInfosResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonSamplesRequest struct {
	*tchttp.BaseRequest

	// 拉取的人物类型，可选值：
	// <li>UserDefine：用户自定义人物库；</li>
	// <li>Default：系统默认人物库。</li>
	// 
	// 默认值：UserDefine，拉取用户自定义人物库人物。
	// 说明：如果是拉取系统默认人物库，只能使用人物名字或者人物 ID + 人物名字的方式进行拉取，且人脸图片只返回一张。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 人物 ID，数组长度限制：100。
	PersonIds []*string `json:"PersonIds,omitempty" name:"PersonIds" list`

	// 人物名称，数组长度限制：20。
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// 人物标签，数组长度限制：20。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribePersonSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonSamplesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribePersonSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 人物信息。
		PersonSet []*AiSamplePerson `json:"PersonSet,omitempty" name:"PersonSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribePersonSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribePersonSamplesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcedureTemplatesRequest struct {
	*tchttp.BaseRequest

	// 任务流模板名字过滤条件，数组长度限制：100。
	Names []*string `json:"Names,omitempty" name:"Names" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeProcedureTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcedureTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeProcedureTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 任务流模板详情列表。
		ProcedureTemplateSet []*ProcedureTemplate `json:"ProcedureTemplateSet,omitempty" name:"ProcedureTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeProcedureTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeProcedureTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReviewDetailsRequest struct {
	*tchttp.BaseRequest

	// 起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 结束日期，需大于起始日期。使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeReviewDetailsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReviewDetailsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeReviewDetailsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 发起内容审核次数。
		TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 内容审核总时长。
		TotalDuration *int64 `json:"TotalDuration,omitempty" name:"TotalDuration"`

		// 内容审核时长统计数据，每天一个数据。
		Data []*StatDataItem `json:"Data,omitempty" name:"Data" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeReviewDetailsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeReviewDetailsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAppIdsRequest struct {
	*tchttp.BaseRequest
}

func (r *DescribeSubAppIdsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubAppIdsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeSubAppIdsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 子应用信息集合。
		SubAppIdInfoSet []*SubAppIdInfo `json:"SubAppIdInfoSet,omitempty" name:"SubAppIdInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeSubAppIdsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeSubAppIdsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailRequest struct {
	*tchttp.BaseRequest

	// 视频处理任务的任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTaskDetailRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTaskDetailResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务类型，取值：
	// <li>Procedure：视频处理任务；</li>
	// <li>EditMedia：视频编辑任务；</li>
	// <li>WechatPublish：微信发布任务；</li>
	// <li>WechatMiniProgramPublish：微信小程序视频发布任务；</li>
	// <li>ComposeMedia：制作媒体文件任务；</li>
	// <li>PullUpload：拉取上传媒体文件任务。</li>
	// 
	// 兼容 2017 版的任务类型：
	// <li>Transcode：视频转码任务；</li>
	// <li>SnapshotByTimeOffset：视频截图任务；</li>
	// <li>Concat：视频拼接任务；</li>
	// <li>Clip：视频剪辑任务；</li>
	// <li>ImageSprites：截取雪碧图任务。</li>
		TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

		// 任务状态，取值：
	// <li>WAITING：等待中；</li>
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
		Status *string `json:"Status,omitempty" name:"Status"`

		// 任务的创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

		// 任务开始执行的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

		// 任务执行完毕的时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
		FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`

		// 视频处理任务信息，仅当 TaskType 为 Procedure，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ProcedureTask *ProcedureTask `json:"ProcedureTask,omitempty" name:"ProcedureTask"`

		// 视频编辑任务信息，仅当 TaskType 为 EditMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		EditMediaTask *EditMediaTask `json:"EditMediaTask,omitempty" name:"EditMediaTask"`

		// 微信发布任务信息，仅当 TaskType 为 WechatPublish，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		WechatPublishTask *WechatPublishTask `json:"WechatPublishTask,omitempty" name:"WechatPublishTask"`

		// 制作媒体文件任务信息，仅当 TaskType 为 ComposeMedia，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ComposeMediaTask *ComposeMediaTask `json:"ComposeMediaTask,omitempty" name:"ComposeMediaTask"`

		// 拉取上传媒体文件任务信息，仅当 TaskType 为 PullUpload，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		PullUploadTask *PullUploadTask `json:"PullUploadTask,omitempty" name:"PullUploadTask"`

		// 视频转码任务信息，仅当 TaskType 为 Transcode，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranscodeTask *TranscodeTask2017 `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

		// 视频指定时间点截图任务信息，仅当 TaskType 为 SnapshotByTimeOffset，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		SnapshotByTimeOffsetTask *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

		// 视频拼接任务信息，仅当 TaskType 为 Concat，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ConcatTask *ConcatTask2017 `json:"ConcatTask,omitempty" name:"ConcatTask"`

		// 视频剪辑任务信息，仅当 TaskType 为 Clip，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ClipTask *ClipTask2017 `json:"ClipTask,omitempty" name:"ClipTask"`

		// 截取雪碧图任务信息，仅当 TaskType 为 ImageSprite，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		CreateImageSpriteTask *CreateImageSpriteTask2017 `json:"CreateImageSpriteTask,omitempty" name:"CreateImageSpriteTask"`

		// 微信小程序发布任务信息，仅当 TaskType 为 WechatMiniProgramPublish，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		WechatMiniProgramPublishTask *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishTask,omitempty" name:"WechatMiniProgramPublishTask"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTaskDetailResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTaskDetailResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksRequest struct {
	*tchttp.BaseRequest

	// 过滤条件：任务状态，可选值：WAITING（等待中）、PROCESSING（处理中）、FINISH（已完成）。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 过滤条件：文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 翻页标识，分批拉取时使用：当单次请求无法拉取所有数据，接口将会返回 ScrollToken，下一次请求携带该 Token，将会从下一条记录开始获取。
	ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTasksRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTasksResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务概要列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TaskSet []*TaskSimpleInfo `json:"TaskSet,omitempty" name:"TaskSet" list`

		// 翻页标识，当请求未返回所有数据，该字段表示下一条记录的 ID。当该字段为空，说明已无更多数据。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ScrollToken *string `json:"ScrollToken,omitempty" name:"ScrollToken"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTasksResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTasksResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 模板类型过滤条件，可选值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式板；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式。</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeTranscodeTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeTranscodeTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 转码模板详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TranscodeTemplateSet []*TranscodeTemplate `json:"TranscodeTemplateSet,omitempty" name:"TranscodeTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeTranscodeTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeTranscodeTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识过滤条件，数组长度限制：100。
	Definitions []*int64 `json:"Definitions,omitempty" name:"Definitions" list`

	// 水印类型过滤条件，可选值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数
	// <li>默认值：10；</li>
	// <li>最大值：100。</li>
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeWatermarkTemplatesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWatermarkTemplatesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合过滤条件的记录总数。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 水印模板详情列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		WatermarkTemplateSet []*WatermarkTemplate `json:"WatermarkTemplateSet,omitempty" name:"WatermarkTemplateSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWatermarkTemplatesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWatermarkTemplatesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWordSamplesRequest struct {
	*tchttp.BaseRequest

	// <b>关键词应用场景过滤条件，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过语音识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行内容审核；
	// 4. Review.Asr：通过语音识别技术，进行内容审核；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、语音识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、语音识别技术，进行内容审核，等价于 3+4；
	// 可多选，元素间关系为 or，即关键词的应用场景包含该字段集合中任意元素的记录，均符合该条件。
	Usages []*string `json:"Usages,omitempty" name:"Usages" list`

	// 关键词过滤条件，数组长度限制：100 个词。
	Keywords []*string `json:"Keywords,omitempty" name:"Keywords" list`

	// 标签过滤条件，数组长度限制：20 个词。
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 分页偏移量，默认值：0。
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：100，最大值：100。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *DescribeWordSamplesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWordSamplesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeWordSamplesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合条件的记录总数。
	// 注意：此字段可能返回 null，表示取不到有效值。
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 关键词信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		WordSet []*AiSampleWord `json:"WordSet,omitempty" name:"WordSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeWordSamplesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeWordSamplesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaFileInfo struct {

	// 视频的 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频剪辑的起始偏移时间偏移，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 视频剪辑的起始结束时间偏移，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type EditMediaOutputConfig struct {

	// 输出文件格式，可选值：mp4、hls。默认是 mp4。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 输出文件的过期时间，超过该时间文件将被删除，默认为永久不过期，格式按照 ISO 8601标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`
}

type EditMediaRequest struct {
	*tchttp.BaseRequest

	// 输入视频的类型，可以取的值为  File，Stream 两种。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入的视频文件信息，当 InputType 为 File 时必填。
	FileInfos []*EditMediaFileInfo `json:"FileInfos,omitempty" name:"FileInfos" list`

	// 输入的流信息，当 InputType 为 Stream 时必填。
	StreamInfos []*EditMediaStreamInfo `json:"StreamInfos,omitempty" name:"StreamInfos" list`

	// 编辑模板 ID，取值有 10，20，不填代表使用 10 模板。
	// <li>10：拼接时，以分辨率最高的输入为基准；</li>
	// <li>20：拼接时，以码率最高的输入为基准；</li>
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字，如果要对生成的新视频执行任务流时填写。
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// 编辑后生成的文件配置。
	OutputConfig *EditMediaOutputConfig `json:"OutputConfig,omitempty" name:"OutputConfig"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *EditMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 编辑视频的任务 ID，可以通过该 ID 查询编辑任务（任务类型为 EditMedia）的状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *EditMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *EditMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type EditMediaStreamInfo struct {

	// 录制的流 ID
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 流剪辑的起始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`
}

type EditMediaTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 视频编辑任务的输入。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Input *EditMediaTaskInput `json:"Input,omitempty" name:"Input"`

	// 视频编辑任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *EditMediaTaskOutput `json:"Output,omitempty" name:"Output"`

	// 若发起视频编辑任务时指定了视频处理流程，则该字段为流程任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type EditMediaTaskInput struct {

	// 输入视频的来源类型，可以取的值为 File，Stream 两种。
	InputType *string `json:"InputType,omitempty" name:"InputType"`

	// 输入的视频文件信息，当 InputType 为 File 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileInfoSet []*EditMediaFileInfo `json:"FileInfoSet,omitempty" name:"FileInfoSet" list`

	// 输入的流信息，当 InputType 为 Stream 时，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StreamInfoSet []*EditMediaStreamInfo `json:"StreamInfoSet,omitempty" name:"StreamInfoSet" list`
}

type EditMediaTaskOutput struct {

	// 文件类型，例如 mp4、flv 等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileType *string `json:"FileType,omitempty" name:"FileType"`

	// 媒体文件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件播放地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`
}

type EmptyTrackItem struct {

	// 持续时间，单位为秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type EventContent struct {

	// 事件句柄，调用方必须调用 ConfirmEvents 来确认消息已经收到，确认有效时间 30 秒。失效后，事件可重新被获取。
	EventHandle *string `json:"EventHandle,omitempty" name:"EventHandle"`

	// <b>支持事件类型：</b>
	// <li>NewFileUpload：视频上传完成；</li>
	// <li>ProcedureStateChanged：任务流状态变更；</li>
	// <li>FileDeleted：视频删除完成；</li>
	// <li>PullComplete：视频转拉完成；</li>
	// <li>EditMediaComplete：视频编辑完成；</li>
	// <li>WechatPublishComplete：微信发布完成；</li>
	// <li>ComposeMediaComplete：制作媒体文件完成；</li>
	// <li>WechatMiniProgramPublishComplete：微信小程序发布完成。</li>
	// <b>兼容 2017 版的事件类型：</b>
	// <li>TranscodeComplete：视频转码完成；</li>
	// <li>ConcatComplete：视频拼接完成；</li>
	// <li>ClipComplete：视频剪辑完成；</li>
	// <li>CreateImageSpriteComplete：视频截取雪碧图完成；</li>
	// <li>CreateSnapshotByTimeOffsetComplete：视频按时间点截图完成。</li>
	EventType *string `json:"EventType,omitempty" name:"EventType"`

	// 视频上传完成事件，当事件类型为 NewFileUpload 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileUploadEvent *FileUploadTask `json:"FileUploadEvent,omitempty" name:"FileUploadEvent"`

	// 任务流状态变更事件，当事件类型为 ProcedureStateChanged 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureStateChangeEvent *ProcedureTask `json:"ProcedureStateChangeEvent,omitempty" name:"ProcedureStateChangeEvent"`

	// 文件删除事件，当事件类型为 FileDeleted 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileDeleteEvent *FileDeleteTask `json:"FileDeleteEvent,omitempty" name:"FileDeleteEvent"`

	// 视频转拉完成事件，当事件类型为 PullComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PullCompleteEvent *PullUploadTask `json:"PullCompleteEvent,omitempty" name:"PullCompleteEvent"`

	// 视频编辑完成事件，当事件类型为 EditMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EditMediaCompleteEvent *EditMediaTask `json:"EditMediaCompleteEvent,omitempty" name:"EditMediaCompleteEvent"`

	// 微信发布完成事件，当事件类型为 WechatPublishComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatPublishCompleteEvent *WechatPublishTask `json:"WechatPublishCompleteEvent,omitempty" name:"WechatPublishCompleteEvent"`

	// 视频转码完成事件，当事件类型为 TranscodeComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeCompleteEvent *TranscodeTask2017 `json:"TranscodeCompleteEvent,omitempty" name:"TranscodeCompleteEvent"`

	// 视频拼接完成事件，当事件类型为 ConcatComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ConcatCompleteEvent *ConcatTask2017 `json:"ConcatCompleteEvent,omitempty" name:"ConcatCompleteEvent"`

	// 视频剪辑完成事件，当事件类型为 ClipComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClipCompleteEvent *ClipTask2017 `json:"ClipCompleteEvent,omitempty" name:"ClipCompleteEvent"`

	// 视频截取雪碧图完成事件，当事件类型为 CreateImageSpriteComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateImageSpriteCompleteEvent *CreateImageSpriteTask2017 `json:"CreateImageSpriteCompleteEvent,omitempty" name:"CreateImageSpriteCompleteEvent"`

	// 视频按时间点截图完成事件，当事件类型为 CreateSnapshotByTimeOffsetComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetCompleteEvent *SnapshotByTimeOffsetTask2017 `json:"SnapshotByTimeOffsetCompleteEvent,omitempty" name:"SnapshotByTimeOffsetCompleteEvent"`

	// 制作媒体文件任务完成事件，当事件类型为 ComposeMediaComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ComposeMediaCompleteEvent *ComposeMediaTask `json:"ComposeMediaCompleteEvent,omitempty" name:"ComposeMediaCompleteEvent"`

	// 微信小程序发布任务完成事件，当事件类型为 WechatMiniProgramPublishComplete 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatMiniProgramPublishCompleteEvent *WechatMiniProgramPublishTask `json:"WechatMiniProgramPublishCompleteEvent,omitempty" name:"WechatMiniProgramPublishCompleteEvent"`
}

type ExecuteFunctionRequest struct {
	*tchttp.BaseRequest

	// 调用后端接口名称。
	FunctionName *string `json:"FunctionName,omitempty" name:"FunctionName"`

	// 接口参数，具体参数格式调用时与后端协调。
	FunctionArg *string `json:"FunctionArg,omitempty" name:"FunctionArg"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ExecuteFunctionRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteFunctionRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ExecuteFunctionResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 处理结果打包后的字符串，具体与后台一同协调。
	// 注意：此字段可能返回 null，表示取不到有效值。
		Result *string `json:"Result,omitempty" name:"Result"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ExecuteFunctionResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ExecuteFunctionResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type FaceConfigureInfo struct {

	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。默认 95 分。取值范围：0 - 100。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：政治人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet" list`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet" list`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	// 默认值：All，使用系统默认人物库及用户自定义人物库。
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FaceConfigureInfoForUpdate struct {

	// 人脸识别任务开关，可选值：
	// <li>ON：开启智能人脸识别任务；</li>
	// <li>OFF：关闭智能人脸识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 人脸识别过滤分数，当识别结果达到该分数以上，返回识别结果。取值范围：0-100。
	Score *float64 `json:"Score,omitempty" name:"Score"`

	// 默认人物过滤标签，指定需要返回的默认人物的标签。如果未填或者为空，则全部默认人物结果都返回。标签可选值：
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星；</li>
	// <li>politician：政治人物。</li>
	DefaultLibraryLabelSet []*string `json:"DefaultLibraryLabelSet,omitempty" name:"DefaultLibraryLabelSet" list`

	// 用户自定义人物过滤标签，指定需要返回的用户自定义人物的标签。如果未填或者为空，则全部自定义人物结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	UserDefineLibraryLabelSet []*string `json:"UserDefineLibraryLabelSet,omitempty" name:"UserDefineLibraryLabelSet" list`

	// 人物库选择，可选值：
	// <li>Default：使用默认人物库；</li>
	// <li>UserDefine：使用用户自定义人物库。</li>
	// <li>All：同时使用默认人物库和用户自定义人物库。</li>
	FaceLibrary *string `json:"FaceLibrary,omitempty" name:"FaceLibrary"`
}

type FileDeleteTask struct {

	// 删除文件 ID 列表。
	FileIdSet []*string `json:"FileIdSet,omitempty" name:"FileIdSet" list`
}

type FileUploadTask struct {

	// 文件唯一 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 上传完成后生成的媒体文件基础信息。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`

	// 若视频上传时指定了视频处理流程，则该字段为流程任务 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`
}

type FrameTagConfigureInfo struct {

	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 截帧间隔，单位为秒，当不填时，默认截帧间隔为 1 秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type FrameTagConfigureInfoForUpdate struct {

	// 智能按帧标签任务开关，可选值：
	// <li>ON：开启智能按帧标签任务；</li>
	// <li>OFF：关闭智能按帧标签任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`
}

type HeadTailConfigureInfo struct {

	// 视频片头片尾识别任务开关，可选值：
	// <li>ON：开启智能视频片头片尾识别任务；</li>
	// <li>OFF：关闭智能视频片头片尾识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type HeadTailConfigureInfoForUpdate struct {

	// 视频片头片尾识别任务开关，可选值：
	// <li>ON：开启智能视频片头片尾识别任务；</li>
	// <li>OFF：关闭智能视频片头片尾识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type ImageSpriteTaskInput struct {

	// 雪碧图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`
}

type ImageTransform struct {

	// 类型，取值有：
	// <li> Rotate：图像旋转。</li>
	// <li> Flip：图像翻转。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 图像以中心点为原点进行旋转的角度，取值范围0~360。当 Type = Rotate 时有效。
	RotateAngle *float64 `json:"RotateAngle,omitempty" name:"RotateAngle"`

	// 图像翻转动作，取值有：
	// <li>Horizental：水平翻转，即左右镜像。</li>
	// <li>Vertical：垂直翻转，即上下镜像。</li>
	// 当 Type = Flip 时有效。
	Flip *string `json:"Flip,omitempty" name:"Flip"`
}

type ImageWatermarkInput struct {

	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：10%。
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 默认值：0px，表示 Height 按照原始水印图片的宽高比缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkInputForUpdate struct {

	// 水印图片 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串。支持 jpeg、png 图片格式。
	ImageContent *string `json:"ImageContent,omitempty" name:"ImageContent"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// 0px 表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type ImageWatermarkTemplate struct {

	// 水印图片地址。
	ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

	// 水印的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Width 为视频宽度的百分比大小，如 10% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 Height 为视频高度的百分比大小，如 10% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；</li>
	// 0px：表示 Height 按照 Width 对视频宽度的比例缩放。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type LiveRealTimeClipRequest struct {
	*tchttp.BaseRequest

	// 推流[直播码](https://cloud.tencent.com/document/product/267/5959)。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 流剪辑的开始时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 流剪辑的结束时间，格式参照 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 是否固化。0 不固化，1 固化。默认不固化。
	IsPersistence *int64 `json:"IsPersistence,omitempty" name:"IsPersistence"`

	// 剪辑固化后的视频存储过期时间。格式参照 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。仅 IsPersistence 为 1 时有效，默认剪辑固化的视频永不过期。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 剪辑固化后的视频点播任务流处理，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。仅 IsPersistence 为 1 时有效。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// 是否需要返回剪辑后的视频元信息。0 不需要，1 需要。默认不需要。
	MetaDataRequired *uint64 `json:"MetaDataRequired,omitempty" name:"MetaDataRequired"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *LiveRealTimeClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type LiveRealTimeClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 剪辑后的视频播放 URL。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 剪辑固化后的视频的媒体文件的唯一标识。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FileId *string `json:"FileId,omitempty" name:"FileId"`

		// 剪辑固化后的视频任务流 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
		VodTaskId *string `json:"VodTaskId,omitempty" name:"VodTaskId"`

		// 剪辑后的视频元信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *LiveRealTimeClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *LiveRealTimeClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type MediaAdaptiveDynamicStreamingInfo struct {

	// 转自适应码流信息数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingSet []*AdaptiveDynamicStreamingInfoItem `json:"AdaptiveDynamicStreamingSet,omitempty" name:"AdaptiveDynamicStreamingSet" list`
}

type MediaAiAnalysisClassificationItem struct {

	// 智能分类的类别名称。
	Classification *string `json:"Classification,omitempty" name:"Classification"`

	// 智能分类的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisCoverItem struct {

	// 智能封面地址。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 智能封面的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagItem struct {

	// 按帧标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 按帧标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAiAnalysisFrameTagSegmentItem struct {

	// 按帧标签起始的偏移时间。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 按帧标签结束的偏移时间。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 时间片段内的标签列表。
	TagSet []*MediaAiAnalysisFrameTagItem `json:"TagSet,omitempty" name:"TagSet" list`
}

type MediaAiAnalysisHighlightItem struct {

	// 智能精彩片段地址。
	HighlightUrl *string `json:"HighlightUrl,omitempty" name:"HighlightUrl"`

	// 智能精彩片段封面地址。
	CovImgUrl *string `json:"CovImgUrl,omitempty" name:"CovImgUrl"`

	// 智能精彩片段的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 智能精彩片段持续时间。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`
}

type MediaAiAnalysisTagItem struct {

	// 标签名称。
	Tag *string `json:"Tag,omitempty" name:"Tag"`

	// 标签的可信度，取值范围是 0 到 100。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaAnimatedGraphicsInfo struct {

	// 视频转动图结果信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicsSet []*MediaAnimatedGraphicsItem `json:"AnimatedGraphicsSet,omitempty" name:"AnimatedGraphicsSet" list`
}

type MediaAnimatedGraphicsItem struct {

	// 转动图的文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转动图模板 ID，参见[转动图参数模板](https://cloud.tencent.com/document/product/266/33481#.E8.BD.AC.E5.8A.A8.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 动图格式，如 gif。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 动图的高度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 动图的宽度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 动图码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 动图大小，单位：字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 动图的md5值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 动图在视频中的起始时间偏移，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 动图在视频中的结束时间偏移，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

type MediaAudioStreamItem struct {

	// 音频流的码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 音频流的采样率，单位：hz。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SamplingRate *int64 `json:"SamplingRate,omitempty" name:"SamplingRate"`

	// 音频流的编码格式，例如 aac。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`
}

type MediaBasicInfo struct {

	// 媒体文件名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体文件描述。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 媒体文件的创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 媒体文件的最近更新时间（如修改视频属性、发起视频处理等会触发更新媒体文件信息的操作），使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 媒体文件的过期时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。“9999-12-31T23:59:59Z”表示永不过期。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 媒体文件的分类 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 媒体文件的分类名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 媒体文件的分类路径，分类间以“-”分隔，如“新的一级分类 - 新的二级分类”。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ClassPath *string `json:"ClassPath,omitempty" name:"ClassPath"`

	// 媒体文件的封面图片地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 媒体文件的封装格式，例如 mp4、flv 等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 原始媒体文件的 URL 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// 该媒体文件的来源信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceInfo *MediaSourceData `json:"SourceInfo,omitempty" name:"SourceInfo"`

	// 媒体文件存储地区，如 ap-guangzhou，参见[地域列表](https://cloud.tencent.com/document/api/213/15692#.E5.9C.B0.E5.9F.9F.E5.88.97.E8.A1.A8)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 媒体文件的标签信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TagSet []*string `json:"TagSet,omitempty" name:"TagSet" list`

	// 直播录制文件的唯一标识
	// 注意：此字段可能返回 null，表示取不到有效值。
	Vid *string `json:"Vid,omitempty" name:"Vid"`
}

type MediaClassInfo struct {

	// 分类 ID
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 父类 ID，一级分类的父类 ID 为 -1。
	ParentId *int64 `json:"ParentId,omitempty" name:"ParentId"`

	// 分类名称
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 分类级别，一级分类为 0，最大值为 3，即最多允许 4 级分类层次。
	Level *uint64 `json:"Level,omitempty" name:"Level"`

	// 当前分类的第一级子类 ID 集合
	// 注意：此字段可能返回 null，表示取不到有效值。
	SubClassIdSet []*int64 `json:"SubClassIdSet,omitempty" name:"SubClassIdSet" list`
}

type MediaContentReviewAsrTextSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`
}

type MediaContentReviewOcrTextSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段置信度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段审核结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑关键词列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeywordSet []*string `json:"KeywordSet,omitempty" name:"KeywordSet" list`

	// 嫌疑文字出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`
}

type MediaContentReviewPoliticalSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉政分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴政结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 涉政人物、违规图标名字。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 嫌疑片段鉴政结果标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 涉政人物、违规图标出现的区域坐标 (像素级)，[x1, y1, x2, y2]，即左上角坐标、右下角坐标。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AreaCoordSet []*int64 `json:"AreaCoordSet,omitempty" name:"AreaCoordSet" list`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`
}

type MediaContentReviewSegmentItem struct {

	// 嫌疑片段起始的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 嫌疑片段结束的偏移时间，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`

	// 嫌疑片段涉黄分数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Confidence *float64 `json:"Confidence,omitempty" name:"Confidence"`

	// 嫌疑片段鉴黄结果标签。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Label *string `json:"Label,omitempty" name:"Label"`

	// 嫌疑片段鉴黄结果建议，取值范围：
	// <li>pass。</li>
	// <li>review。</li>
	// <li>block。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 嫌疑图片 URL （图片不会永久存储，到达
	//  PicUrlExpireTime 时间点后图片将被删除）。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 嫌疑图片 URL 失效时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicUrlExpireTimeStamp *int64 `json:"PicUrlExpireTimeStamp,omitempty" name:"PicUrlExpireTimeStamp"`
}

type MediaDeleteItem struct {

	// 所指定的删除部分。如果未填写该字段则参数无效。可选值有：
	// <li>TranscodeFiles（删除转码文件）。</li>
	// <li>WechatPublishFiles（删除微信发布文件）。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 删除由Type参数指定的种类下的视频模板号，模板定义参见[转码模板](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	// 默认值为0，表示删除参数Type指定种类下所有的视频。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`
}

type MediaImageSpriteInfo struct {

	// 特定规格的雪碧图信息集合，每个元素代表一套相同规格的雪碧图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteSet []*MediaImageSpriteItem `json:"ImageSpriteSet,omitempty" name:"ImageSpriteSet" list`
}

type MediaImageSpriteItem struct {

	// 雪碧图规格，参见[雪碧图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.9B.AA.E7.A2.A7.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 雪碧图小图的高度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 雪碧图小图的宽度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 每一张雪碧图大图里小图的数量。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TotalCount *int64 `json:"TotalCount,omitempty" name:"TotalCount"`

	// 每一张雪碧图大图的地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// 雪碧图子图位置与时间关系的 WebVtt 文件地址。WebVtt 文件表明了各个雪碧图小图对应的时间点，以及在在雪碧大图里的坐标位置，一般被播放器用于实现预览。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WebVttUrl *string `json:"WebVttUrl,omitempty" name:"WebVttUrl"`
}

type MediaInfo struct {

	// 基础信息。包括视频名称、分类、播放地址、封面图片等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BasicInfo *MediaBasicInfo `json:"BasicInfo,omitempty" name:"BasicInfo"`

	// 元信息。包括大小、时长、视频流信息、音频流信息等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 转码结果信息。包括该视频转码生成的各种码率的视频的地址、规格、码率、分辨率等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeInfo *MediaTranscodeInfo `json:"TranscodeInfo,omitempty" name:"TranscodeInfo"`

	// 转动图结果信息。对视频转动图（如 gif）后，动图相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicsInfo *MediaAnimatedGraphicsInfo `json:"AnimatedGraphicsInfo,omitempty" name:"AnimatedGraphicsInfo"`

	// 采样截图信息。对视频采样截图后，相关截图信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotInfo *MediaSampleSnapshotInfo `json:"SampleSnapshotInfo,omitempty" name:"SampleSnapshotInfo"`

	// 雪碧图信息。对视频截取雪碧图之后，雪碧的相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteInfo *MediaImageSpriteInfo `json:"ImageSpriteInfo,omitempty" name:"ImageSpriteInfo"`

	// 指定时间点截图信息。对视频依照指定时间点截图后，各个截图的信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetInfo *MediaSnapshotByTimeOffsetInfo `json:"SnapshotByTimeOffsetInfo,omitempty" name:"SnapshotByTimeOffsetInfo"`

	// 视频打点信息。对视频设置的各个打点信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyFrameDescInfo *MediaKeyFrameDescInfo `json:"KeyFrameDescInfo,omitempty" name:"KeyFrameDescInfo"`

	// 转自适应码流信息。包括规格、加密类型、打包格式等相关信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingInfo *MediaAdaptiveDynamicStreamingInfo `json:"AdaptiveDynamicStreamingInfo,omitempty" name:"AdaptiveDynamicStreamingInfo"`

	// 小程序审核信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MiniProgramReviewInfo *MediaMiniProgramReviewInfo `json:"MiniProgramReviewInfo,omitempty" name:"MiniProgramReviewInfo"`

	// 媒体文件唯一标识 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`
}

type MediaInputInfo struct {

	// 视频 URL。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 视频名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频自定义 ID。
	Id *string `json:"Id,omitempty" name:"Id"`
}

type MediaKeyFrameDescInfo struct {

	// 视频打点信息数组。
	// 注意：此字段可能返回 null，表示取不到有效值。
	KeyFrameDescSet []*MediaKeyFrameDescItem `json:"KeyFrameDescSet,omitempty" name:"KeyFrameDescSet" list`
}

type MediaKeyFrameDescItem struct {

	// 打点的视频偏移时间，单位：秒。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 打点的内容字符串，限制 1-128 个字符。
	Content *string `json:"Content,omitempty" name:"Content"`
}

type MediaMetaData struct {

	// 上传的媒体文件大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 容器类型，例如 m4a，mp4 等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频流码率平均值与音频流码率平均值之和，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频拍摄时的选择角度，单位：度。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Rotate *int64 `json:"Rotate,omitempty" name:"Rotate"`

	// 视频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`

	// 音频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoDuration *float64 `json:"VideoDuration,omitempty" name:"VideoDuration"`

	// 音频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioDuration *float64 `json:"AudioDuration,omitempty" name:"AudioDuration"`
}

type MediaMiniProgramReviewElem struct {

	// 审核类型。 
	// <li>Porn：画面涉黄，</li>
	// <li>Porn.Ocr：文字涉黄，</li>
	// <li>Porn.Asr：声音涉黄，</li>
	// <li>Terrorism：画面涉暴恐，</li>
	// <li>Political：画面涉政，</li>
	// <li>Political.Ocr：文字涉政，</li>
	// <li>Political.Asr：声音涉政。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 审核意见。
	// <li>pass：确认正常，</li>
	// <li>block：确认违规，</li>
	// <li>review：疑似违规。</li>
	Suggestion *string `json:"Suggestion,omitempty" name:"Suggestion"`

	// 审核结果置信度。取值 0~100。
	Confidence *string `json:"Confidence,omitempty" name:"Confidence"`
}

type MediaMiniProgramReviewInfo struct {

	// 审核信息列表。
	MiniProgramReivewList []*MediaMiniProgramReviewInfoItem `json:"MiniProgramReivewList,omitempty" name:"MiniProgramReivewList" list`
}

type MediaMiniProgramReviewInfoItem struct {

	// 模板id。小程序视频发布的视频所对应的转码模板ID，为0代表原始视频。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 小程序审核视频播放地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 小程序视频发布状态：
	// <li>Pass：成功。</li>
	// <li>Rejected：未通过。</li>
	ReviewResult *string `json:"ReviewResult,omitempty" name:"ReviewResult"`

	// 小程序审核元素。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ReviewSummery *MediaMiniProgramReviewElem `json:"ReviewSummery,omitempty" name:"ReviewSummery"`
}

type MediaOutputInfo struct {

	// 输出文件 Bucket 所属地域，如 ap-guangzhou  。
	Region *string `json:"Region,omitempty" name:"Region"`

	// 输出文件 Bucket 。
	Bucket *string `json:"Bucket,omitempty" name:"Bucket"`

	// 输出文件目录，目录名必须以 "/" 结尾。
	Dir *string `json:"Dir,omitempty" name:"Dir"`
}

type MediaProcessTaskAdaptiveDynamicStreamingResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频转自适应码流任务的输入。
	Input *AdaptiveDynamicStreamingTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频转自适应码流任务的输出。
	Output *AdaptiveDynamicStreamingInfoItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskAnimatedGraphicResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转动图任务的输入。
	Input *AnimatedGraphicTaskInput `json:"Input,omitempty" name:"Input"`

	// 转动图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaAnimatedGraphicsItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskCoverBySnapshotResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频截图做封面任务的输入。
	Input *CoverBySnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频截图做封面任务的输出。
	Output *CoverBySnapshotTaskOutput `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskImageSpriteResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频截雪碧图任务的输入。
	Input *ImageSpriteTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频截雪碧图任务的输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaImageSpriteItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskInput struct {

	// 视频转码任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTaskSet []*TranscodeTaskInput `json:"TranscodeTaskSet,omitempty" name:"TranscodeTaskSet" list`

	// 视频转动图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTaskSet []*AnimatedGraphicTaskInput `json:"AnimatedGraphicTaskSet,omitempty" name:"AnimatedGraphicTaskSet" list`

	// 对视频按时间点截图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTaskSet []*SnapshotByTimeOffsetTaskInput `json:"SnapshotByTimeOffsetTaskSet,omitempty" name:"SnapshotByTimeOffsetTaskSet" list`

	// 对视频采样截图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTaskSet []*SampleSnapshotTaskInput `json:"SampleSnapshotTaskSet,omitempty" name:"SampleSnapshotTaskSet" list`

	// 对视频截雪碧图任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTaskSet []*ImageSpriteTaskInput `json:"ImageSpriteTaskSet,omitempty" name:"ImageSpriteTaskSet" list`

	// 对视频截图做封面任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverBySnapshotTaskSet []*CoverBySnapshotTaskInput `json:"CoverBySnapshotTaskSet,omitempty" name:"CoverBySnapshotTaskSet" list`

	// 对视频转自适应码流任务列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTaskSet []*AdaptiveDynamicStreamingTaskInput `json:"AdaptiveDynamicStreamingTaskSet,omitempty" name:"AdaptiveDynamicStreamingTaskSet" list`
}

type MediaProcessTaskResult struct {

	// 任务的类型，可以取的值有：
	// <li>Transcode：转码</li>
	// <li>AnimatedGraphics：转动图</li>
	// <li>SnapshotByTimeOffset：时间点截图</li>
	// <li>SampleSnapshot：采样截图</li>
	// <li>ImageSprites：雪碧图</li>
	// <li>CoverBySnapshot：截图做封面</li>
	// <li>AdaptiveDynamicStreaming：自适应码流</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频转码任务的查询结果，当任务类型为 Transcode 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeTask *MediaProcessTaskTranscodeResult `json:"TranscodeTask,omitempty" name:"TranscodeTask"`

	// 视频转动图任务的查询结果，当任务类型为 AnimatedGraphics 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AnimatedGraphicTask *MediaProcessTaskAnimatedGraphicResult `json:"AnimatedGraphicTask,omitempty" name:"AnimatedGraphicTask"`

	// 对视频按时间点截图任务的查询结果，当任务类型为 SnapshotByTimeOffset 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetTask *MediaProcessTaskSnapshotByTimeOffsetResult `json:"SnapshotByTimeOffsetTask,omitempty" name:"SnapshotByTimeOffsetTask"`

	// 对视频采样截图任务的查询结果，当任务类型为 SampleSnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotTask *MediaProcessTaskSampleSnapshotResult `json:"SampleSnapshotTask,omitempty" name:"SampleSnapshotTask"`

	// 对视频截雪碧图任务的查询结果，当任务类型为 ImageSprite 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageSpriteTask *MediaProcessTaskImageSpriteResult `json:"ImageSpriteTask,omitempty" name:"ImageSpriteTask"`

	// 对视频截图做封面任务的查询结果，当任务类型为 CoverBySnapshot 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverBySnapshotTask *MediaProcessTaskCoverBySnapshotResult `json:"CoverBySnapshotTask,omitempty" name:"CoverBySnapshotTask"`

	// 对视频转自适应码流任务的查询结果，当任务类型为 AdaptiveDynamicStreaming 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AdaptiveDynamicStreamingTask *MediaProcessTaskAdaptiveDynamicStreamingResult `json:"AdaptiveDynamicStreamingTask,omitempty" name:"AdaptiveDynamicStreamingTask"`
}

type MediaProcessTaskSampleSnapshotResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频做采样截图任务输入。
	Input *SampleSnapshotTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频做采样截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSampleSnapshotItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskSnapshotByTimeOffsetResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 对视频按指定时间点截图任务输入。
	Input *SnapshotByTimeOffsetTaskInput `json:"Input,omitempty" name:"Input"`

	// 对视频按指定时间点截图任务输出。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Output *MediaSnapshotByTimeOffsetItem `json:"Output,omitempty" name:"Output"`
}

type MediaProcessTaskTranscodeResult struct {

	// 任务状态，有 PROCESSING，SUCCESS 和 FAIL 三种。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码，0：成功，其他值：失败。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转码任务的输入。
	Input *TranscodeTaskInput `json:"Input,omitempty" name:"Input"`

	// 转码任务的输出。
	Output *MediaTranscodeItem `json:"Output,omitempty" name:"Output"`
}

type MediaSampleSnapshotInfo struct {

	// 特定规格的采样截图信息集合，每个元素代表一套相同规格的采样截图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleSnapshotSet []*MediaSampleSnapshotItem `json:"SampleSnapshotSet,omitempty" name:"SampleSnapshotSet" list`
}

type MediaSampleSnapshotItem struct {

	// 采样截图规格 ID，参见[采样截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E9.87.87.E6.A0.B7.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 采样方式，取值范围：
	// <li>Percent：根据百分比间隔采样。</li>
	// <li>Time：根据时间间隔采样。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SampleType *string `json:"SampleType,omitempty" name:"SampleType"`

	// 采样间隔
	// <li>当 SampleType 为 Percent 时，该值表示多少百分比一张图。</li>
	// <li>当 SampleType 为 Time 时，该值表示多少时间间隔一张图，单位秒， 第一张图均为视频首帧。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	Interval *int64 `json:"Interval,omitempty" name:"Interval"`

	// 生成的截图 url 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageUrlSet []*string `json:"ImageUrlSet,omitempty" name:"ImageUrlSet" list`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSnapshotByTimeOffsetInfo struct {

	// 特定规格的指定时间点截图信息集合。目前每种规格只能有一套截图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SnapshotByTimeOffsetSet []*MediaSnapshotByTimeOffsetItem `json:"SnapshotByTimeOffsetSet,omitempty" name:"SnapshotByTimeOffsetSet" list`
}

type MediaSnapshotByTimeOffsetItem struct {

	// 指定时间点截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 同一规格的截图信息集合，每个元素代表一张截图。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PicInfoSet []*MediaSnapshotByTimePicInfoItem `json:"PicInfoSet,omitempty" name:"PicInfoSet" list`
}

type MediaSnapshotByTimePicInfoItem struct {

	// 该张截图对应视频文件中的时间偏移，单位为<font color=red>毫秒</font>。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOffset *float64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 该张截图的 URL 地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 截图如果被打上了水印，被打水印的模板 ID 列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WaterMarkDefinition []*int64 `json:"WaterMarkDefinition,omitempty" name:"WaterMarkDefinition" list`
}

type MediaSourceData struct {

	// 媒体文件的来源类别：
	// <li>Record：来自录制。如直播录制、直播时移录制等。</li>
	// <li>Upload：来自上传。如拉取上传、服务端上传、客户端 UGC 上传等。</li>
	// <li>VideoProcessing：来自视频处理。如视频拼接、视频剪辑等。</li>
	// <li>Unknown：未知来源。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 用户创建文件时透传的字段
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceContext *string `json:"SourceContext,omitempty" name:"SourceContext"`
}

type MediaTrack struct {

	// 轨道类型，取值有：
	// <ul>
	// <li>Video ：视频轨道。视频轨道由以下 Item 组成：<ul><li>VideoTrackItem</li><li>MediaTransitionItem</li> <li>EmptyTrackItem</li></ul> </li>
	// <li>Audio ：音频轨道。音频轨道由以下 Item 组成：<ul><li>AudioTrackItem</li><li>MediaTransitionItem</li><li>EmptyTrackItem</li></ul></li>
	// <li>Sticker ：贴图轨道。贴图轨道以下 Item 组成：<ul><li> StickerTrackItem</li><li>EmptyTrackItem</li></ul></li>	
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 轨道上的媒体片段列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TrackItems []*MediaTrackItem `json:"TrackItems,omitempty" name:"TrackItems" list`
}

type MediaTrackItem struct {

	// 片段类型。取值有：
	// <li>Video：视频片段。</li>
	// <li>Audio：音频片段。</li>
	// <li>Sticker：贴图片段。</li>
	// <li>Transition：转场。</li>
	// <li>Empty：空白片段。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频片段，当 Type = Video 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoItem *VideoTrackItem `json:"VideoItem,omitempty" name:"VideoItem"`

	// 音频片段，当 Type = Audio 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioItem *AudioTrackItem `json:"AudioItem,omitempty" name:"AudioItem"`

	// 贴图片段，当 Type = Sticker 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	StickerItem *StickerTrackItem `json:"StickerItem,omitempty" name:"StickerItem"`

	// 转场，当 Type = Transition 时有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TransitionItem *MediaTransitionItem `json:"TransitionItem,omitempty" name:"TransitionItem"`

	// 空白片段，当 Type = Empty 时有效。空片段用于时间轴的占位。<li>如需要两个音频片段之间有一段时间的静音，可以用 EmptyTrackItem 来进行占位。</li>
	// <li>使用 EmptyTrackItem 进行占位，来定位某个Item。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	EmptyItem *EmptyTrackItem `json:"EmptyItem,omitempty" name:"EmptyItem"`
}

type MediaTranscodeInfo struct {

	// 各规格的转码信息集合，每个元素代表一个规格的转码结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TranscodeSet []*MediaTranscodeItem `json:"TranscodeSet,omitempty" name:"TranscodeSet" list`
}

type MediaTranscodeItem struct {

	// 转码后的视频文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 媒体文件总大小（视频为 HLS 时，大小是 m3u8 和 ts 文件大小的总和），单位：字节。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Size *int64 `json:"Size,omitempty" name:"Size"`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 容器类型，例如 m4a，mp4 等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 视频的 md5 值。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Md5 *string `json:"Md5,omitempty" name:"Md5"`

	// 音频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioStreamSet []*MediaAudioStreamItem `json:"AudioStreamSet,omitempty" name:"AudioStreamSet" list`

	// 视频流信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoStreamSet []*MediaVideoStreamItem `json:"VideoStreamSet,omitempty" name:"VideoStreamSet" list`
}

type MediaTransitionItem struct {

	// 转场持续时间，单位为秒。进行转场处理的两个媒体片段，第二个片段在轨道上的起始时间会自动进行调整，设置为前面一个片段的结束时间减去转场的持续时间。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 转场操作列表。图像转场操作和音频转场操作各自最多支持一个。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Transitions []*TransitionOpertion `json:"Transitions,omitempty" name:"Transitions" list`
}

type MediaVideoStreamItem struct {

	// 视频流的码率，单位：bps。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流的高度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流的宽度，单位：px。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *int64 `json:"Width,omitempty" name:"Width"`

	// 视频流的编码格式，例如 h264。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 帧率，单位：hz。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type ModifyAIAnalysisTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容分析模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频内容分析模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频内容分析模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 智能分类任务控制参数。
	ClassificationConfigure *ClassificationConfigureInfoForUpdate `json:"ClassificationConfigure,omitempty" name:"ClassificationConfigure"`

	// 智能标签任务控制参数。
	TagConfigure *TagConfigureInfoForUpdate `json:"TagConfigure,omitempty" name:"TagConfigure"`

	// 智能封面任务控制参数。
	CoverConfigure *CoverConfigureInfoForUpdate `json:"CoverConfigure,omitempty" name:"CoverConfigure"`

	// 智能按帧标签任务控制参数。
	FrameTagConfigure *FrameTagConfigureInfoForUpdate `json:"FrameTagConfigure,omitempty" name:"FrameTagConfigure"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyAIAnalysisTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIAnalysisTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAIAnalysisTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAIAnalysisTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIAnalysisTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAIRecognitionTemplateRequest struct {
	*tchttp.BaseRequest

	// 视频内容识别模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频内容识别模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频内容识别模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 视频片头片尾识别控制参数。
	HeadTailConfigure *HeadTailConfigureInfoForUpdate `json:"HeadTailConfigure,omitempty" name:"HeadTailConfigure"`

	// 视频拆条识别控制参数。
	SegmentConfigure *SegmentConfigureInfoForUpdate `json:"SegmentConfigure,omitempty" name:"SegmentConfigure"`

	// 人脸识别控制参数。
	FaceConfigure *FaceConfigureInfoForUpdate `json:"FaceConfigure,omitempty" name:"FaceConfigure"`

	// 文本全文识别控制参数。
	OcrFullTextConfigure *OcrFullTextConfigureInfoForUpdate `json:"OcrFullTextConfigure,omitempty" name:"OcrFullTextConfigure"`

	// 文本关键词识别控制参数。
	OcrWordsConfigure *OcrWordsConfigureInfoForUpdate `json:"OcrWordsConfigure,omitempty" name:"OcrWordsConfigure"`

	// 语音全文识别控制参数。
	AsrFullTextConfigure *AsrFullTextConfigureInfoForUpdate `json:"AsrFullTextConfigure,omitempty" name:"AsrFullTextConfigure"`

	// 语音关键词识别控制参数。
	AsrWordsConfigure *AsrWordsConfigureInfoForUpdate `json:"AsrWordsConfigure,omitempty" name:"AsrWordsConfigure"`

	// 物体识别控制参数。
	ObjectConfigure *ObjectConfigureInfoForUpdate `json:"ObjectConfigure,omitempty" name:"ObjectConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyAIRecognitionTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIRecognitionTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyAIRecognitionTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyAIRecognitionTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyAIRecognitionTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClassRequest struct {
	*tchttp.BaseRequest

	// 分类 ID
	ClassId *uint64 `json:"ClassId,omitempty" name:"ClassId"`

	// 分类名称。长度限制：1-64 个字符。
	ClassName *string `json:"ClassName,omitempty" name:"ClassName"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyClassRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyClassResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyClassResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyClassResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContentReviewTemplateRequest struct {
	*tchttp.BaseRequest

	// 内容审核模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 内容审核模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 内容审核模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 鉴黄控制参数。
	PornConfigure *PornConfigureInfoForUpdate `json:"PornConfigure,omitempty" name:"PornConfigure"`

	// 鉴恐控制参数。
	TerrorismConfigure *TerrorismConfigureInfoForUpdate `json:"TerrorismConfigure,omitempty" name:"TerrorismConfigure"`

	// 鉴政控制参数。
	PoliticalConfigure *PoliticalConfigureInfoForUpdate `json:"PoliticalConfigure,omitempty" name:"PoliticalConfigure"`

	// 用户自定义内容审核控制参数。
	UserDefineConfigure *UserDefineConfigureInfoForUpdate `json:"UserDefineConfigure,omitempty" name:"UserDefineConfigure"`

	// 截帧间隔，单位为秒，最小值为 0.5 秒。
	ScreenshotInterval *float64 `json:"ScreenshotInterval,omitempty" name:"ScreenshotInterval"`

	// 审核结果是否进入审核墙（对审核结果进行人工复核）的开关。
	// <li>ON：是；</li>
	// <li>OFF：否。</li>
	ReviewWallSwitch *string `json:"ReviewWallSwitch,omitempty" name:"ReviewWallSwitch"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyContentReviewTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContentReviewTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyContentReviewTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyContentReviewTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyContentReviewTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoRequest struct {
	*tchttp.BaseRequest

	// 媒体文件唯一标识。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件名称，最长 64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 媒体文件描述，最长 128 个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 媒体文件分类 ID。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 媒体文件过期时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。填“9999-12-31T23:59:59Z”表示永不过期。过期后该媒体文件及其相关资源（转码结果、雪碧图等）将被永久删除。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 视频封面图片文件（如 jpeg, png 等）进行 [Base64](https://tools.ietf.org/html/rfc4648) 编码后的字符串，仅支持 gif、jpeg、png 三种图片格式。
	CoverData *string `json:"CoverData,omitempty" name:"CoverData"`

	// 新增的一组视频打点信息，如果某个偏移时间已存在打点，则会进行覆盖操作，单个媒体文件最多 100 个打点信息。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	AddKeyFrameDescs []*MediaKeyFrameDescItem `json:"AddKeyFrameDescs,omitempty" name:"AddKeyFrameDescs" list`

	// 要删除的一组视频打点信息的时间偏移，单位：秒。同一个请求里，AddKeyFrameDescs 的时间偏移参数必须与 DeleteKeyFrameDescs 都不同。
	DeleteKeyFrameDescs []*float64 `json:"DeleteKeyFrameDescs,omitempty" name:"DeleteKeyFrameDescs" list`

	// 取值 1 表示清空视频打点信息，其他值无意义。
	// 同一个请求里，ClearKeyFrameDescs 与 AddKeyFrameDescs 不能同时出现。
	ClearKeyFrameDescs *int64 `json:"ClearKeyFrameDescs,omitempty" name:"ClearKeyFrameDescs"`

	// 新增的一组标签，单个媒体文件最多 16 个标签，单个标签最多 16 个字符。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	AddTags []*string `json:"AddTags,omitempty" name:"AddTags" list`

	// 要删除的一组标签。同一个请求里，AddTags 参数必须与 DeleteTags 都不同。
	DeleteTags []*string `json:"DeleteTags,omitempty" name:"DeleteTags" list`

	// 取值 1 表示清空媒体文件所有标签，其他值无意义。
	// 同一个请求里，ClearTags 与 AddTags 不能同时出现。
	ClearTags *int64 `json:"ClearTags,omitempty" name:"ClearTags"`

	// 点播[子应用](/document/product/266/14574) ID 。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyMediaInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyMediaInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 新的视频封面 URL。
	// * 注意：仅当请求携带 CoverData 时此返回值有效。 *
	// 注意：此字段可能返回 null，表示取不到有效值。
		CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyMediaInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyMediaInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonSampleRequest struct {
	*tchttp.BaseRequest

	// 人物 ID。
	PersonId *string `json:"PersonId,omitempty" name:"PersonId"`

	// 名称，长度限制：128 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 描述，长度限制：1024 个字符。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 人物应用场景，可选值：
	// 1. Recognition：用于内容识别，等价于 Recognition.Face。
	// 2. Review：用于内容审核，等价于 Review.Face。
	// 3. All：用于内容识别、内容审核，等价于 1+2。
	Usages []*string `json:"Usages,omitempty" name:"Usages" list`

	// 人脸操作信息。
	FaceOperationInfo *AiSampleFaceOperation `json:"FaceOperationInfo,omitempty" name:"FaceOperationInfo"`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitempty" name:"TagOperationInfo"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyPersonSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonSampleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyPersonSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 人物信息。
		Person *AiSamplePerson `json:"Person,omitempty" name:"Person"`

		// 处理失败的人脸信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
		FailFaceInfoSet []*AiSampleFailFaceInfo `json:"FailFaceInfoSet,omitempty" name:"FailFaceInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyPersonSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyPersonSampleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdInfoRequest struct {
	*tchttp.BaseRequest

	// 子应用 ID。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子应用名称，长度限制：40个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子应用简介，长度限制： 300个字符。
	Description *string `json:"Description,omitempty" name:"Description"`
}

func (r *ModifySubAppIdInfoRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdInfoRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdInfoResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubAppIdInfoResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdInfoResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdStatusRequest struct {
	*tchttp.BaseRequest

	// 子应用 ID。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子应用状态，取值范围：
	// <li>On：启用</li>
	// <li>Off：停用</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

func (r *ModifySubAppIdStatusRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdStatusRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifySubAppIdStatusResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifySubAppIdStatusResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifySubAppIdStatusResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateRequest struct {
	*tchttp.BaseRequest

	// 转码模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 封装格式，可选值：mp4、flv、hls、mp3、flac、ogg、m4a。其中，mp3、flac、ogg、m4a 为纯音频文件。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字节。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 是否去除视频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，可选值：
	// <li>0：保留</li>
	// <li>1：去除</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数。
	VideoTemplate *VideoTemplateInfoForUpdate `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数。
	AudioTemplate *AudioTemplateInfoForUpdate `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyTranscodeTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyTranscodeTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyTranscodeTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyTranscodeTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateRequest struct {
	*tchttp.BaseRequest

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 水印模板名称，长度限制：64 个字符。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息，长度限制：256 个字符。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 原点位置，可选值：
	// <li>TopLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>TopRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>BottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>BottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下角。</li>
	// 目前，当 Type 为 image，该字段仅支持 TopLeft。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 水印原点距离视频图像坐标原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 XPos 为视频宽度指定百分比，如 10% 表示 XPos 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 XPos 为指定像素，如 100px 表示 XPos 为 100 像素。</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印原点距离视频图像坐标原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示水印 YPos 为视频高度指定百分比，如 10% 表示 YPos 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 YPos 为指定像素，如 100px 表示 YPos 为 100 像素。</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，该字段仅对图片水印模板有效。
	ImageTemplate *ImageWatermarkInputForUpdate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，该字段仅对文字水印模板有效。
	TextTemplate *TextWatermarkTemplateInputForUpdate `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG 水印模板，该字段仅对 SVG 水印模板有效。
	SvgTemplate *SvgWatermarkInputForUpdate `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyWatermarkTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWatermarkTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 图片水印地址，仅当 ImageTemplate.ImageContent 非空，该字段有值。
	// 注意：此字段可能返回 null，表示取不到有效值。
		ImageUrl *string `json:"ImageUrl,omitempty" name:"ImageUrl"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWatermarkTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWatermarkTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWordSampleRequest struct {
	*tchttp.BaseRequest

	// 关键词，长度限制：128 个字符。
	Keyword *string `json:"Keyword,omitempty" name:"Keyword"`

	// <b>关键词应用场景，可选值：</b>
	// 1. Recognition.Ocr：通过光学字符识别技术，进行内容识别；
	// 2. Recognition.Asr：通过语音识别技术，进行内容识别；
	// 3. Review.Ocr：通过光学字符识别技术，进行内容审核；
	// 4. Review.Asr：通过语音识别技术，进行内容审核；
	// <b>可合并简写为：</b>
	// 5. Recognition：通过光学字符识别技术、语音识别技术，进行内容识别，等价于 1+2；
	// 6. Review：通过光学字符识别技术、语音识别技术，进行内容审核，等价于 3+4；
	// 7. All：通过光学字符识别技术、语音识别技术，进行内容识别、内容审核，等价于 1+2+3+4。
	Usages []*string `json:"Usages,omitempty" name:"Usages" list`

	// 标签操作信息。
	TagOperationInfo *AiSampleTagOperation `json:"TagOperationInfo,omitempty" name:"TagOperationInfo"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ModifyWordSampleRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWordSampleRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ModifyWordSampleResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ModifyWordSampleResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ModifyWordSampleResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ObjectConfigureInfo struct {

	// 物体识别任务开关，可选值：
	// <li>ON：开启智能物体识别任务；</li>
	// <li>OFF：关闭智能物体识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 物体库选择，可选值：
	// <li>Default：使用默认物体库；</li>
	// <li>UserDefine：使用用户自定义物体库。</li>
	// <li>All：同时使用默认物体库和用户自定义物体库。</li>
	// 默认值： All，同时使用默认物体库和用户自定义物体库。
	ObjectLibrary *string `json:"ObjectLibrary,omitempty" name:"ObjectLibrary"`
}

type ObjectConfigureInfoForUpdate struct {

	// 物体识别任务开关，可选值：
	// <li>ON：开启智能物体识别任务；</li>
	// <li>OFF：关闭智能物体识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 物体库选择，可选值：
	// <li>Default：使用默认物体库；</li>
	// <li>UserDefine：使用用户自定义物体库。</li>
	// <li>All：同时使用默认物体库和用户自定义物体库。</li>
	ObjectLibrary *string `json:"ObjectLibrary,omitempty" name:"ObjectLibrary"`
}

type OcrFullTextConfigureInfo struct {

	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrFullTextConfigureInfoForUpdate struct {

	// 文本全文识别任务开关，可选值：
	// <li>ON：开启智能文本全文识别任务；</li>
	// <li>OFF：关闭智能文本全文识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type OcrWordsConfigureInfo struct {

	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type OcrWordsConfigureInfoForUpdate struct {

	// 文本关键词识别任务开关，可选值：
	// <li>ON：开启文本关键词识别任务；</li>
	// <li>OFF：关闭文本关键词识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 关键词过滤标签，指定需要返回的关键词的标签。如果未填或者为空，则全部结果都返回。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`
}

type OutputAudioStream struct {

	// 音频流的编码格式，可选值：
	// <li>libfdk_aac：适合 mp4 文件。</li>
	// 默认值：libfdk_aac。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 音频流的采样率，可选值：
	// <li>16000</li>
	// <li>32000</li>
	// <li>44100</li>
	// <li>48000</li>
	// 单位：Hz。
	// 默认值：16000。
	SampleRate *int64 `json:"SampleRate,omitempty" name:"SampleRate"`

	// 音频声道数，可选值：
	// <li>1：单声道 。</li>
	// <li>2：双声道</li>
	// 默认值：2。
	AudioChannel *int64 `json:"AudioChannel,omitempty" name:"AudioChannel"`
}

type OutputVideoStream struct {

	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码 </li>
	// 默认值：libx264。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 默认值：0，表示和第一个视频轨的第一个视频片段的视频帧率一致。
	Fps *int64 `json:"Fps,omitempty" name:"Fps"`
}

type PoliticalAsrReviewTemplateInfo struct {

	// 语音鉴政任务开关，可选值：
	// <li>ON：开启语音鉴政任务；</li>
	// <li>OFF：关闭语音鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalAsrReviewTemplateInfoForUpdate struct {

	// 语音鉴政任务开关，可选值：
	// <li>ON：开启语音鉴政任务；</li>
	// <li>OFF：关闭语音鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalConfigureInfo struct {

	// 画面鉴政控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *PoliticalImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// 语音鉴政控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 文本鉴政控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalConfigureInfoForUpdate struct {

	// 画面鉴政控制参数。
	ImgReviewInfo *PoliticalImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// 语音鉴政控制参数。
	AsrReviewInfo *PoliticalAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 文本鉴政控制参数。
	OcrReviewInfo *PoliticalOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PoliticalImgReviewTemplateInfo struct {

	// 画面鉴政任务开关，可选值：
	// <li>ON：开启画面鉴政任务；</li>
	// <li>OFF：关闭画面鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴政过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：政治人物；</li>
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalImgReviewTemplateInfoForUpdate struct {

	// 画面鉴政任务开关，可选值：
	// <li>ON：开启画面鉴政任务；</li>
	// <li>OFF：关闭画面鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴政过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>violation_photo：违规图标；</li>
	// <li>politician：政治人物；</li>
	// <li>entertainment：娱乐明星；</li>
	// <li>sport：体育明星。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfo struct {

	// 文本鉴政任务开关，可选值：
	// <li>ON：开启文本鉴政任务；</li>
	// <li>OFF：关闭文本鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PoliticalOcrReviewTemplateInfoForUpdate struct {

	// 文本鉴政任务开关，可选值：
	// <li>ON：开启文本鉴政任务；</li>
	// <li>OFF：关闭文本鉴政任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfo struct {

	// 语音鉴黄任务开关，可选值：
	// <li>ON：开启语音鉴黄任务；</li>
	// <li>OFF：关闭语音鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornAsrReviewTemplateInfoForUpdate struct {

	// 语音鉴黄任务开关，可选值：
	// <li>ON：开启语音鉴黄任务；</li>
	// <li>OFF：关闭语音鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornConfigureInfo struct {

	// 画面鉴黄控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *PornImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// 语音鉴黄控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *PornAsrReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 文本鉴黄控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *PornOcrReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornConfigureInfoForUpdate struct {

	// 画面鉴黄控制参数。
	ImgReviewInfo *PornImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`

	// 语音鉴黄控制参数。
	AsrReviewInfo *PornAsrReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 文本鉴黄控制参数。
	OcrReviewInfo *PornOcrReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type PornImgReviewTemplateInfo struct {

	// 画面鉴黄任务开关，可选值：
	// <li>ON：开启画面鉴黄任务；</li>
	// <li>OFF：关闭画面鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴黄过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 0 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornImgReviewTemplateInfoForUpdate struct {

	// 画面鉴黄任务开关，可选值：
	// <li>ON：开启画面鉴黄任务；</li>
	// <li>OFF：关闭画面鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴黄过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>porn：色情；</li>
	// <li>vulgar：低俗；</li>
	// <li>intimacy：亲密行为；</li>
	// <li>sexy：性感。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfo struct {

	// 文本鉴黄任务开关，可选值：
	// <li>ON：开启文本鉴黄任务；</li>
	// <li>OFF：关闭文本鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type PornOcrReviewTemplateInfoForUpdate struct {

	// 文本鉴黄任务开关，可选值：
	// <li>ON：开启文本鉴黄任务；</li>
	// <li>OFF：关闭文本鉴黄任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type ProcedureTask struct {

	// 视频处理任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 媒体文件 ID
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 FileId；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Id。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 媒体文件名称
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 BasicInfo.Name；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Name。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 媒体文件地址
	// <li>若流程由 [ProcessMedia](https://cloud.tencent.com/document/product/266/33427) 发起，该字段表示 [MediaInfo](https://cloud.tencent.com/document/product/266/31773#MediaInfo) 的 BasicInfo.MediaUrl；</li>
	// <li>若流程由 [ProcessMediaByUrl](https://cloud.tencent.com/document/product/266/33426) 发起，该字段表示 [MediaInputInfo](https://cloud.tencent.com/document/product/266/31773#MediaInputInfo) 的 Url。</li>
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 原始视频的元信息。
	MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

	// 视频处理任务的执行状态与结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaProcessResultSet []*MediaProcessTaskResult `json:"MediaProcessResultSet,omitempty" name:"MediaProcessResultSet" list`

	// 视频内容审核任务的执行状态与结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewResultSet []*AiContentReviewResult `json:"AiContentReviewResultSet,omitempty" name:"AiContentReviewResultSet" list`

	// 视频内容分析任务的执行状态与结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisResultSet []*AiAnalysisResult `json:"AiAnalysisResultSet,omitempty" name:"AiAnalysisResultSet" list`

	// 视频内容识别任务的执行状态与结果。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionResultSet []*AiRecognitionResult `json:"AiRecognitionResultSet,omitempty" name:"AiRecognitionResultSet" list`

	// 任务流的优先级，取值范围为 [-10, 10]。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式。
	// <li>Finish：只有当任务流全部执行完毕时，才发起一次事件通知；</li>
	// <li>Change：只要任务流中每个子任务的状态发生变化，都进行事件通知；</li>
	// <li>None：不接受该任务流回调。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type ProcedureTemplate struct {

	// 任务流名字。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频处理类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// AI 智能内容审核类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type ProcessMediaByProcedureRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// [任务流模板](/document/product/266/11700#.E4.BB.BB.E5.8A.A1.E6.B5.81.E6.A8.A1.E6.9D.BF)名字。
	ProcedureName *string `json:"ProcedureName,omitempty" name:"ProcedureName"`

	// 任务流的优先级，数值越大优先级越高，取值范围是-10到10，不填代表0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果一天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByProcedureRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByProcedureRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByProcedureResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaByProcedureResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByProcedureResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByUrlRequest struct {
	*tchttp.BaseRequest

	// 输入视频信息，包括视频 URL ， 名称、视频自定义 ID。
	InputInfo *MediaInputInfo `json:"InputInfo,omitempty" name:"InputInfo"`

	// 输出文件 COS 路径信息。
	OutputInfo *MediaOutputInfo `json:"OutputInfo,omitempty" name:"OutputInfo"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaByUrlRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaByUrlResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaByUrlResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaByUrlResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// 视频内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// 视频内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// 视频内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 任务流的优先级，数值越大优先级越高，取值范围是 -10 到 10，不填代表 0。
	TasksPriority *int64 `json:"TasksPriority,omitempty" name:"TasksPriority"`

	// 任务流状态变更通知模式，可取值有 Finish，Change 和 None，不填代表 Finish。
	TasksNotifyMode *string `json:"TasksNotifyMode,omitempty" name:"TasksNotifyMode"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ProcessMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ProcessMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ProcessMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ProcessMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullEventsRequest struct {
	*tchttp.BaseRequest

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullEventsRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullEventsResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 事件列表。
	// 注意：此字段可能返回 null，表示取不到有效值。
		EventSet []*EventContent `json:"EventSet,omitempty" name:"EventSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullEventsResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullEventsResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadRequest struct {
	*tchttp.BaseRequest

	// 要拉取的媒体 URL，暂不支持拉取 HLS 和 Dash 格式。
	// <li>URL 里文件名需要包括扩展名, 比如 ```https://xxxx.mp4``` ，扩展名为 mp4，支持的扩展名详见[文件类型](https://cloud.tencent.com/document/product/266/9760#.E6.96.87.E4.BB.B6.E7.B1.BB.E5.9E.8B)。</li>
	MediaUrl *string `json:"MediaUrl,omitempty" name:"MediaUrl"`

	// 媒体名称。
	MediaName *string `json:"MediaName,omitempty" name:"MediaName"`

	// 要拉取的视频封面 URL。仅支持 gif、jpeg、png 三种图片格式。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 媒体后续任务操作，详见[上传指定任务流](https://cloud.tencent.com/document/product/266/9759)。
	Procedure *string `json:"Procedure,omitempty" name:"Procedure"`

	// 媒体文件过期时间，格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	ExpireTime *string `json:"ExpireTime,omitempty" name:"ExpireTime"`

	// 指定上传园区，仅适用于对上传地域有特殊需求的用户（目前仅支持北京、上海和重庆园区）。
	StorageRegion *string `json:"StorageRegion,omitempty" name:"StorageRegion"`

	// 分类ID，用于对媒体进行分类管理，可通过[创建分类](https://cloud.tencent.com/document/product/266/7812)接口，创建分类，获得分类 ID。
	ClassId *int64 `json:"ClassId,omitempty" name:"ClassId"`

	// 来源上下文，用于透传用户请求信息，当指定 Procedure 任务后，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PullUploadRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullUploadRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 拉取上传视频的任务 ID，可以通过该 ID 查询拉取上传任务的状态。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PullUploadResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PullUploadResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PullUploadTask struct {

	// 转拉上传任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务流状态，取值：
	// <li>PROCESSING：处理中；</li>
	// <li>FINISH：已完成。</li>
	// 
	// 注意：此字段可能返回 null，表示取不到有效值。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 转拉上传完成后生成的视频 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 转拉完成后生成的媒体文件基础信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	MediaBasicInfo *MediaBasicInfo `json:"MediaBasicInfo,omitempty" name:"MediaBasicInfo"`

	// 转拉上传完成后生成的播放地址。
	FileUrl *string `json:"FileUrl,omitempty" name:"FileUrl"`

	// 若转拉上传时指定了视频处理流程，则该参数为流程任务 ID。
	ProcedureTaskId *string `json:"ProcedureTaskId,omitempty" name:"ProcedureTaskId"`

	// 来源上下文，用于透传用户请求信息，任务流状态变更回调将返回该字段值，最长 1000 个字符。
	SessionContext *string `json:"SessionContext,omitempty" name:"SessionContext"`

	// 用于去重的识别码，如果七天内曾有过相同的识别码的请求，则本次的请求会返回错误。最长 50 个字符，不带或者带空字符串表示不做去重。
	SessionId *string `json:"SessionId,omitempty" name:"SessionId"`
}

type PushUrlCacheRequest struct {
	*tchttp.BaseRequest

	// 预热的 URL 列表，单次最多指定20个 URL。
	Urls []*string `json:"Urls,omitempty" name:"Urls" list`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *PushUrlCacheRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlCacheRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type PushUrlCacheResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *PushUrlCacheResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *PushUrlCacheResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetProcedureTemplateRequest struct {
	*tchttp.BaseRequest

	// 任务流名字
	Name *string `json:"Name,omitempty" name:"Name"`

	// 视频处理类型任务参数。
	MediaProcessTask *MediaProcessTaskInput `json:"MediaProcessTask,omitempty" name:"MediaProcessTask"`

	// AI 智能内容审核类型任务参数。
	AiContentReviewTask *AiContentReviewTaskInput `json:"AiContentReviewTask,omitempty" name:"AiContentReviewTask"`

	// AI 智能内容分析类型任务参数。
	AiAnalysisTask *AiAnalysisTaskInput `json:"AiAnalysisTask,omitempty" name:"AiAnalysisTask"`

	// AI 内容识别类型任务参数。
	AiRecognitionTask *AiRecognitionTaskInput `json:"AiRecognitionTask,omitempty" name:"AiRecognitionTask"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *ResetProcedureTemplateRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetProcedureTemplateRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type ResetProcedureTemplateResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *ResetProcedureTemplateResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *ResetProcedureTemplateResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SampleSnapshotTaskInput struct {

	// 采样截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SearchMediaRequest struct {
	*tchttp.BaseRequest

	// 搜索文本，模糊匹配媒体文件名称或描述信息，匹配项越多，匹配度越高，排序越优先。长度限制：64 个字符。
	Text *string `json:"Text,omitempty" name:"Text"`

	// 标签集合，匹配集合中任意元素。
	// <li>单个标签长度限制：8 个字符。</li>
	// <li>数组长度限制：10。</li>
	Tags []*string `json:"Tags,omitempty" name:"Tags" list`

	// 分类 ID 集合，匹配集合指定 ID 的分类及其所有子类。数组长度限制：10。
	ClassIds []*int64 `json:"ClassIds,omitempty" name:"ClassIds" list`

	// 创建时间的开始时间。
	// <li>大于等于开始时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	StartTime *string `json:"StartTime,omitempty" name:"StartTime"`

	// 创建时间的结束时间。
	// <li>小于结束时间。</li>
	// <li>格式按照 ISO 8601 标准表示，详见 [ISO 日期格式说明](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。</li>
	EndTime *string `json:"EndTime,omitempty" name:"EndTime"`

	// 媒体文件来源，来源取值参见 [SourceType](https://cloud.tencent.com/document/product/266/31773#MediaSourceData)。
	SourceType *string `json:"SourceType,omitempty" name:"SourceType"`

	// 推流[直播码](https://cloud.tencent.com/document/product/267/5959)。
	StreamId *string `json:"StreamId,omitempty" name:"StreamId"`

	// 直播录制文件的唯一标识。
	Vid *string `json:"Vid,omitempty" name:"Vid"`

	// 排序方式。
	// <li>Sort.Field 可选值：CreateTime</li>
	// <li>指定 Text 搜索时，将根据匹配度排序，该字段无效</li>
	Sort *SortBy `json:"Sort,omitempty" name:"Sort"`

	// 偏移量。
	// <li>默认值：0。</li>
	// <li>取值范围：Offset + Limit 不超过 5000。</li>
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 返回记录条数，默认值：10。
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *SearchMediaRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SearchMediaResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 符合搜索条件的记录总数
	// <li>最大值：5000，即，当命中记录数超过 5000，该字段将返回 5000，而非实际命中总数。</li>
		TotalCount *uint64 `json:"TotalCount,omitempty" name:"TotalCount"`

		// 媒体文件信息列表，只包含基础信息（BasicInfo）
	// 注意：此字段可能返回 null，表示取不到有效值。
		MediaInfoSet []*MediaInfo `json:"MediaInfoSet,omitempty" name:"MediaInfoSet" list`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SearchMediaResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SearchMediaResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SegmentConfigureInfo struct {

	// 视频拆条识别任务开关，可选值：
	// <li>ON：开启智能视频拆条识别任务；</li>
	// <li>OFF：关闭智能视频拆条识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SegmentConfigureInfoForUpdate struct {

	// 视频拆条识别任务开关，可选值：
	// <li>ON：开启智能视频拆条识别任务；</li>
	// <li>OFF：关闭智能视频拆条识别任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type SimpleHlsClipRequest struct {
	*tchttp.BaseRequest

	// 需要裁剪的腾讯云点播 HLS 视频 URL。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 裁剪的开始偏移时间，单位秒。默认 0，即从视频开头开始裁剪。负数表示距离视频结束多少秒开始裁剪。比如 -10 表示从倒数第 10 秒开始裁剪。
	StartTimeOffset *float64 `json:"StartTimeOffset,omitempty" name:"StartTimeOffset"`

	// 裁剪的结束偏移时间，单位秒。默认 0，即裁剪到视频尾部。负数表示距离视频结束多少秒结束裁剪。比如 -10 表示到倒数第 10 秒结束裁剪。
	EndTimeOffset *float64 `json:"EndTimeOffset,omitempty" name:"EndTimeOffset"`
}

func (r *SimpleHlsClipRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SimpleHlsClipResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 裁剪后的视频地址。
		Url *string `json:"Url,omitempty" name:"Url"`

		// 裁剪后的视频元信息。目前`Size`，`Rotate`，`VideoDuration`，`AudioDuration` 几个字段暂时缺省，没有真实数据。
		MetaData *MediaMetaData `json:"MetaData,omitempty" name:"MetaData"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *SimpleHlsClipResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *SimpleHlsClipResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type SnapshotByTimeOffset2017 struct {

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 截图的具体时间点，单位：毫秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOffset *uint64 `json:"TimeOffset,omitempty" name:"TimeOffset"`

	// 截图输出文件地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Url *string `json:"Url,omitempty" name:"Url"`
}

type SnapshotByTimeOffsetTask2017 struct {

	// 截图任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 截图文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 截图规格，参见[指定时间点截图参数模板](https://cloud.tencent.com/document/product/266/33480#.E6.97.B6.E9.97.B4.E7.82.B9.E6.88.AA.E5.9B.BE.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 截图结果信息。
	SnapshotInfoSet []*SnapshotByTimeOffset2017 `json:"SnapshotInfoSet,omitempty" name:"SnapshotInfoSet" list`
}

type SnapshotByTimeOffsetTaskInput struct {

	// 指定时间点截图模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 截图时间点列表，单位为<font color=red>毫秒</font>。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TimeOffsetSet []*float64 `json:"TimeOffsetSet,omitempty" name:"TimeOffsetSet" list`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type SortBy struct {

	// 排序字段
	Field *string `json:"Field,omitempty" name:"Field"`

	// 排序方式，可选值：Asc（升序）、Desc（降序）
	Order *string `json:"Order,omitempty" name:"Order"`
}

type StatDataItem struct {

	// 数据所在时间区间的开始时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。如：当时间粒度为天，2018-12-01T00:00:00+08:00，表示2018年12月1日（含）到2018年12月2日（不含）区间。
	Time *string `json:"Time,omitempty" name:"Time"`

	// 数据大小。
	// <li>存储空间的数据，单位是字节。</li>
	// <li>转码时长的数据，单位是秒。</li>
	// <li>流量数据，单位是字节。</li>
	// <li>带宽数据，单位是比特每秒。</li>
	Value *int64 `json:"Value,omitempty" name:"Value"`
}

type StickerTrackItem struct {

	// 贴图素材的媒体文件来源。可以是点播的文件 ID，也可以是其它文件的 URL。
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// 贴图的持续时间，单位为秒。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 贴图在轨道上的起始时间，单位为秒。
	StartTime *float64 `json:"StartTime,omitempty" name:"StartTime"`

	// 原点位置，取值有：
	// <li>Center：坐标原点为中心位置，如画布中心。</li>
	// 默认值：Center。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 贴图原点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布宽度的 10%。</li><li>当字符串以 px 结尾，表示贴图 XPos 单位为像素，如 100px 表示 XPos 为 100 像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 贴图原点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 YPos 单位为像素，如 100px 表示 YPos 为 100 像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 贴图的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 Width 单位为像素，如 100px 表示 Width 为 100 像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取贴图素材本身的 Width、Height。</li>
	// <li>当 Width 为空0，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 贴图的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示贴图 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示贴图 Height 单位为像素，如 100px 表示 Hieght 为 100 像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取贴图素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Height *string `json:"Height,omitempty" name:"Height"`

	// 对贴图进行的操作，如图像旋转等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations" list`
}

type SubAppIdInfo struct {

	// 子应用 ID。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`

	// 子应用名称。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 子应用简介。
	Description *string `json:"Description,omitempty" name:"Description"`

	// 子应用创建时间，采用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 子应用状态，有效值：
	// <li>On：启用；</li>
	// <li>Off：停用。</li>
	Status *string `json:"Status,omitempty" name:"Status"`
}

type SvgWatermarkInput struct {

	// 水印的宽度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；当填 0px 且
	//  Height 不为 0px 时，表示水印的宽度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的宽度取原始 SVG 图像的宽度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Width 为视频宽度的百分比大小，如 10W% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Width 为视频高度的百分比大小，如 10H% 表示 Width 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Width 为视频短边的百分比大小，如 10S% 表示 Width 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Width 为视频长边的百分比大小，如 10L% 表示 Width 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 W%。</li>
	// 默认值为 10W%。
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度，支持 px，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。</li>
	// 默认值为 0px。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type SvgWatermarkInputForUpdate struct {

	// 水印的宽度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Width 单位为像素，如 100px 表示 Width 为 100 像素；当填 0px 且
	//  Height 不为 0px 时，表示水印的宽度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的宽度取原始 SVG 图像的宽度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Width 为视频宽度的百分比大小，如 10W% 表示 Width 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Width 为视频高度的百分比大小，如 10H% 表示 Width 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Width 为视频短边的百分比大小，如 10S% 表示 Width 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Width 为视频长边的百分比大小，如 10L% 表示 Width 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 W%。</li>
	// 默认值为 10W%。
	Width *string `json:"Width,omitempty" name:"Width"`

	// 水印的高度，支持 px，%，W%，H%，S%，L% 六种格式：
	// <li>当字符串以 px 结尾，表示水印 Height 单位为像素，如 100px 表示 Height 为 100 像素；当填 0px 且
	//  Width 不为 0px 时，表示水印的高度按原始 SVG 图像等比缩放；当 Width、Height 都填 0px 时，表示水印的高度取原始 SVG 图像的高度；</li>
	// <li>当字符串以 W% 结尾，表示水印 Height 为视频宽度的百分比大小，如 10W% 表示 Height 为视频宽度的 10%；</li>
	// <li>当字符串以 H% 结尾，表示水印 Height 为视频高度的百分比大小，如 10H% 表示 Height 为视频高度的 10%；</li>
	// <li>当字符串以 S% 结尾，表示水印 Height 为视频短边的百分比大小，如 10S% 表示 Height 为视频短边的 10%；</li>
	// <li>当字符串以 L% 结尾，表示水印 Height 为视频长边的百分比大小，如 10L% 表示 Height 为视频长边的 10%；</li>
	// <li>当字符串以 % 结尾时，含义同 H%。
	// 默认值为 0px。
	Height *string `json:"Height,omitempty" name:"Height"`
}

type TEHDConfig struct {

	// 极速高清类型，可选值：
	// <li>TEHD-100：极速高清-100。</li>
	// 不填代表不启用极速高清。
	Type *string `json:"Type,omitempty" name:"Type"`

	// 视频码率上限，当 Type 指定了极速高清类型时有效。
	// 不填或填0表示不设视频码率上限。
	MaxVideoBitrate *uint64 `json:"MaxVideoBitrate,omitempty" name:"MaxVideoBitrate"`
}

type TagConfigureInfo struct {

	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type TagConfigureInfoForUpdate struct {

	// 智能标签任务开关，可选值：
	// <li>ON：开启智能标签任务；</li>
	// <li>OFF：关闭智能标签任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`
}

type TaskSimpleInfo struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务类型，取值：
	// <li>Procedure：视频处理任务；</li>
	// <li>EditMedia：视频编辑任务</li>
	// <li>WechatDistribute：微信发布任务。</li>
	// 兼容 2017 版的任务类型：
	// <li>Transcode：视频转码任务；</li>
	// <li>SnapshotByTimeOffset：视频截图任务；</li>
	// <li>Concat：视频拼接任务；</li>
	// <li>Clip：视频剪辑任务；</li>
	// <li>ImageSprites：截取雪碧图任务。</li>
	TaskType *string `json:"TaskType,omitempty" name:"TaskType"`

	// 任务创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreatTime *string `json:"CreatTime,omitempty" name:"CreatTime"`

	// 任务开始执行时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未开始，该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	BeginProcessTime *string `json:"BeginProcessTime,omitempty" name:"BeginProcessTime"`

	// 任务结束时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。若任务尚未完成，该字段为空。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FinishTime *string `json:"FinishTime,omitempty" name:"FinishTime"`
}

type TempCertificate struct {

	// 临时安全证书 Id。
	SecretId *string `json:"SecretId,omitempty" name:"SecretId"`

	// 临时安全证书 Key。
	SecretKey *string `json:"SecretKey,omitempty" name:"SecretKey"`

	// Token 值。
	Token *string `json:"Token,omitempty" name:"Token"`

	// 证书无效的时间，返回 Unix 时间戳，精确到秒。
	ExpiredTime *uint64 `json:"ExpiredTime,omitempty" name:"ExpiredTime"`
}

type TerrorismConfigureInfo struct {

	// 画面鉴恐任务控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImgReviewInfo *TerrorismImgReviewTemplateInfo `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`
}

type TerrorismConfigureInfoForUpdate struct {

	// 画面鉴恐任务控制参数。
	ImgReviewInfo *TerrorismImgReviewTemplateInfoForUpdate `json:"ImgReviewInfo,omitempty" name:"ImgReviewInfo"`
}

type TerrorismImgReviewTemplateInfo struct {

	// 画面鉴恐任务开关，可选值：
	// <li>ON：开启画面鉴恐任务；</li>
	// <li>OFF：关闭画面鉴恐任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴恐过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：暴恐旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：暴恐人物。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 90 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 80 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TerrorismImgReviewTemplateInfoForUpdate struct {

	// 画面鉴恐任务开关，可选值：
	// <li>ON：开启画面鉴恐任务；</li>
	// <li>OFF：关闭画面鉴恐任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 画面鉴恐过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回，可选值为：
	// <li>guns：武器枪支；</li>
	// <li>crowd：人群聚集；</li>
	// <li>bloody：血腥画面；</li>
	// <li>police：警察部队；</li>
	// <li>banners：暴恐旗帜；</li>
	// <li>militant：武装分子；</li>
	// <li>explosion：爆炸火灾；</li>
	// <li>terrorists：暴恐人物。</li>
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type TextWatermarkTemplateInput struct {

	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	// 默认值：1。
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TextWatermarkTemplateInputForUpdate struct {

	// 字体类型，目前可以支持两种：
	// <li>simkai.ttf：可以支持中文和英文；</li>
	// <li>arial.ttf：仅支持英文。</li>
	FontType *string `json:"FontType,omitempty" name:"FontType"`

	// 字体大小，格式：Npx，N 为数值。
	FontSize *string `json:"FontSize,omitempty" name:"FontSize"`

	// 字体颜色，格式：0xRRGGBB，默认值：0xFFFFFF（白色）。
	FontColor *string `json:"FontColor,omitempty" name:"FontColor"`

	// 文字透明度，取值范围：(0, 1]
	// <li>0：完全透明</li>
	// <li>1：完全不透明</li>
	FontAlpha *float64 `json:"FontAlpha,omitempty" name:"FontAlpha"`
}

type TranscodePlayInfo2017 struct {

	// 播放地址。
	Url *string `json:"Url,omitempty" name:"Url"`

	// 转码规格 ID，参见[转码参数模板](https://cloud.tencent.com/document/product/266/33478#.E8.BD.AC.E7.A0.81.E6.A8.A1.E6.9D.BF)。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 视频流码率平均值与音频流码率平均值之和， 单位：bps。
	Bitrate *int64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 视频流高度的最大值，单位：px。
	Height *int64 `json:"Height,omitempty" name:"Height"`

	// 视频流宽度的最大值，单位：px。
	Width *int64 `json:"Width,omitempty" name:"Width"`
}

type TranscodeTask2017 struct {

	// 转码任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 被转码文件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 被转码文件名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileName *string `json:"FileName,omitempty" name:"FileName"`

	// 视频时长，单位：秒。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Duration *uint64 `json:"Duration,omitempty" name:"Duration"`

	// 封面地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	CoverUrl *string `json:"CoverUrl,omitempty" name:"CoverUrl"`

	// 视频转码后生成的播放信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	PlayInfoSet []*TranscodePlayInfo2017 `json:"PlayInfoSet,omitempty" name:"PlayInfoSet" list`
}

type TranscodeTaskInput struct {

	// 视频转码模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 水印列表，支持多张图片或文字水印，最大可支持 10 张。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WatermarkSet []*WatermarkInput `json:"WatermarkSet,omitempty" name:"WatermarkSet" list`
}

type TranscodeTemplate struct {

	// 转码模板唯一标识。
	Definition *string `json:"Definition,omitempty" name:"Definition"`

	// 封装格式，取值：mp4、flv、hls、mp3、flac、ogg。
	Container *string `json:"Container,omitempty" name:"Container"`

	// 转码模板名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 模板类型，取值：
	// <li>Preset：系统预置模板；</li>
	// <li>Custom：用户自定义模板。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 是否去除视频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveVideo *int64 `json:"RemoveVideo,omitempty" name:"RemoveVideo"`

	// 是否去除音频数据，取值：
	// <li>0：保留；</li>
	// <li>1：去除。</li>
	RemoveAudio *int64 `json:"RemoveAudio,omitempty" name:"RemoveAudio"`

	// 视频流配置参数，仅当 RemoveVideo 为 0，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	VideoTemplate *VideoTemplateInfo `json:"VideoTemplate,omitempty" name:"VideoTemplate"`

	// 音频流配置参数，仅当 RemoveAudio 为 0，该字段有效 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioTemplate *AudioTemplateInfo `json:"AudioTemplate,omitempty" name:"AudioTemplate"`

	// 极速高清转码参数，需联系商务架构师开通后才能使用。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TEHDConfig *TEHDConfig `json:"TEHDConfig,omitempty" name:"TEHDConfig"`

	// 封装格式过滤条件，可选值：
	// <li>Video：视频格式，可以同时包含视频流和音频流的封装格式；</li>
	// <li>PureAudio：纯音频格式，只能包含音频流的封装格式板。</li>
	ContainerType *string `json:"ContainerType,omitempty" name:"ContainerType"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`
}

type TransitionOpertion struct {

	// 转场类型，取值有：
	// <ul>
	// <li>图像的转场操作，用于两个视频片段图像间的转场处理：
	// <ul>
	// <li>ImageFadeInFadeOut：图像淡入淡出。 </li>
	// <li>BowTieHorizontal：水平蝴蝶结。 </li>
	// <li>BowTieVertical：垂直蝴蝶结。 </li>
	// <li>ButterflyWaveScrawler：晃动。 </li>
	// <li>Cannabisleaf：枫叶。 </li>
	// <li>Circle：弧形收放。 </li>
	// <li>CircleCrop：圆环聚拢。 </li>
	// <li>Circleopen：椭圆聚拢。 </li>
	// <li>Crosswarp：横向翘曲。 </li>
	// <li>Cube：立方体。 </li>
	// <li>DoomScreenTransition：幕布。 </li>
	// <li>Doorway：门廊。 </li>
	// <li>Dreamy：波浪。 </li>
	// <li>DreamyZoom：水平聚拢。 </li>
	// <li>FilmBurn：火烧云。 </li>
	// <li>GlitchMemories：抖动。 </li>
	// <li>Heart：心形。 </li>
	// <li>InvertedPageCurl：翻页。 </li>
	// <li>Luma：腐蚀。 </li>
	// <li>Mosaic：九宫格。 </li>
	// <li>Pinwheel：风车。 </li>
	// <li>PolarFunction：椭圆扩散。 </li>
	// <li>PolkaDotsCurtain：弧形扩散。 </li>
	// <li>Radial：雷达扫描 </li>
	// <li>RotateScaleFade：上下收放。 </li>
	// <li>Squeeze：上下聚拢。 </li>
	// <li>Swap：放大切换。 </li>
	// <li>Swirl：螺旋。 </li>
	// <li>UndulatingBurnOutSwirl：水流蔓延。 </li>
	// <li>Windowblinds：百叶窗。 </li>
	// <li>WipeDown：向下收起。 </li>
	// <li>WipeLeft：向左收起。 </li>
	// <li>WipeRight：向右收起。 </li>
	// <li>WipeUp：向上收起。 </li>
	// <li>ZoomInCircles：水波纹。 </li>
	// </ul>
	// </li>
	// <li>音频的转场操作，用于两个音频片段间的转场处理：
	// <ul>
	// <li>AudioFadeInFadeOut：声音淡入淡出。 </li>
	// </ul>
	// </li>
	// </ul>
	Type *string `json:"Type,omitempty" name:"Type"`
}

type UserDefineAsrTextReviewTemplateInfo struct {

	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineAsrTextReviewTemplateInfoForUpdate struct {

	// 用户自定语音审核任务开关，可选值：
	// <li>ON：开启自定义语音审核任务；</li>
	// <li>OFF：关闭自定义语音审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义语音过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义语音关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineConfigureInfo struct {

	// 用户自定义人物审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfo `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// 用户自定义语音审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfo `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 用户自定义文本审核控制参数。
	// 注意：此字段可能返回 null，表示取不到有效值。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfo `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineConfigureInfoForUpdate struct {

	// 用户自定义人物审核控制参数。
	FaceReviewInfo *UserDefineFaceReviewTemplateInfoForUpdate `json:"FaceReviewInfo,omitempty" name:"FaceReviewInfo"`

	// 用户自定义语音审核控制参数。
	AsrReviewInfo *UserDefineAsrTextReviewTemplateInfoForUpdate `json:"AsrReviewInfo,omitempty" name:"AsrReviewInfo"`

	// 用户自定义文本审核控制参数。
	OcrReviewInfo *UserDefineOcrTextReviewTemplateInfoForUpdate `json:"OcrReviewInfo,omitempty" name:"OcrReviewInfo"`
}

type UserDefineFaceReviewTemplateInfo struct {

	// 用户自定义人物审核任务开关，可选值：
	// <li>ON：开启自定义人物审核任务；</li>
	// <li>OFF：关闭自定义人物审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义人物过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 97 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 95 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineFaceReviewTemplateInfoForUpdate struct {

	// 用户自定义人物审核任务开关，可选值：
	// <li>ON：开启自定义人物审核任务；</li>
	// <li>OFF：关闭自定义人物审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义人物过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义人物库的时，需要添加对应人物标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfo struct {

	// 用户自定文本审核任务开关，可选值：
	// <li>ON：开启自定义文本审核任务；</li>
	// <li>OFF：关闭自定义文本审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义文本过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规，不填默认为 100 分。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核，不填默认为 75 分。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type UserDefineOcrTextReviewTemplateInfoForUpdate struct {

	// 用户自定文本审核任务开关，可选值：
	// <li>ON：开启自定义文本审核任务；</li>
	// <li>OFF：关闭自定义文本审核任务。</li>
	Switch *string `json:"Switch,omitempty" name:"Switch"`

	// 用户自定义文本过滤标签，审核结果包含选择的标签则返回结果，如果过滤标签为空，则审核结果全部返回。如果要使用标签过滤功能，添加自定义文本关键词素材时需要添加对应标签。
	// 标签个数最多 10 个，每个标签长度最多 16 个字符。
	LabelSet []*string `json:"LabelSet,omitempty" name:"LabelSet" list`

	// 判定涉嫌违规的分数阈值，当智能审核达到该分数以上，认为涉嫌违规。取值范围：0~100。
	BlockConfidence *int64 `json:"BlockConfidence,omitempty" name:"BlockConfidence"`

	// 判定需人工复核是否违规的分数阈值，当智能审核达到该分数以上，认为需人工复核。取值范围：0~100。
	ReviewConfidence *int64 `json:"ReviewConfidence,omitempty" name:"ReviewConfidence"`
}

type VideoTemplateInfo struct {

	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// 目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	// 默认值：open。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	// 默认值：0。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type VideoTemplateInfoForUpdate struct {

	// 视频流的编码格式，可选值：
	// <li>libx264：H.264 编码</li>
	// <li>libx265：H.265 编码</li>
	// 目前 H.265 编码必须指定分辨率，并且需要在 640*480 以内。
	Codec *string `json:"Codec,omitempty" name:"Codec"`

	// 视频帧率，取值范围：[0, 60]，单位：Hz。
	// 当取值为 0，表示帧率和原始视频保持一致。
	Fps *uint64 `json:"Fps,omitempty" name:"Fps"`

	// 视频流的码率，取值范围：0 和 [128, 35000]，单位：kbps。
	// 当取值为 0，表示视频码率和原始视频保持一致。
	Bitrate *uint64 `json:"Bitrate,omitempty" name:"Bitrate"`

	// 分辨率自适应，可选值：
	// <li>open：开启，此时，Width 代表视频的长边，Height 表示视频的短边；</li>
	// <li>close：关闭，此时，Width 代表视频的宽度，Height 表示视频的高度。</li>
	ResolutionAdaptive *string `json:"ResolutionAdaptive,omitempty" name:"ResolutionAdaptive"`

	// 视频流宽度（或长边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	// <li>当 Width、Height 均为 0，则分辨率同源；</li>
	// <li>当 Width 为 0，Height 非 0，则 Width 按比例缩放；</li>
	// <li>当 Width 非 0，Height 为 0，则 Height 按比例缩放；</li>
	// <li>当 Width、Height 均非 0，则分辨率按用户指定。</li>
	Width *uint64 `json:"Width,omitempty" name:"Width"`

	// 视频流高度（或短边）的最大值，取值范围：0 和 [128, 4096]，单位：px。
	Height *uint64 `json:"Height,omitempty" name:"Height"`

	// 填充方式，当视频流配置宽高参数与原始视频的宽高比不一致时，对转码的处理方式，即为“填充”。可选填充方式：
	// <li> stretch：拉伸，对每一帧进行拉伸，填满整个画面，可能导致转码后的视频被“压扁“或者“拉长“；</li>
	// <li>black：留黑，保持视频宽高比不变，边缘剩余部分使用黑色填充。</li>
	// 默认值：black 。
	FillType *string `json:"FillType,omitempty" name:"FillType"`
}

type VideoTrackItem struct {

	// 视频片段的媒体素材来源，可以是点播的文件 ID，或者是其它文件的 URL。
	SourceMedia *string `json:"SourceMedia,omitempty" name:"SourceMedia"`

	// 视频片段取自素材文件的起始时间，单位为秒。默认为0。
	SourceMediaStartTime *float64 `json:"SourceMediaStartTime,omitempty" name:"SourceMediaStartTime"`

	// 视频片段时长，单位为秒。默认取视频素材本身长度，表示截取全部素材。如果源文件是图片，Duration需要大于0。
	Duration *float64 `json:"Duration,omitempty" name:"Duration"`

	// 视频原点位置，取值有：
	// <li>Center：坐标原点为中心位置，如画布中心。</li>
	// 默认值 ：Center。
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`

	// 视频片段原点距离画布原点的水平位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 XPos 为画布宽度指定百分比的位置，如 10% 表示 XPos 为画布口宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 XPos 单位为像素，如 100px 表示 XPos 为100像素。</li>
	// 默认值：0px。
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 视频片段原点距离画布原点的垂直位置。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 YPos 为画布高度指定百分比的位置，如 10% 表示 YPos 为画布高度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 YPos 单位为像素，如 100px 表示 YPos 为100像素。</li>
	// 默认值：0px。
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 视频片段的宽度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Width 为画布宽度的百分比大小，如 10% 表示 Width 为画布宽度的 10%。</li>
	// <li>当字符串以 px 结尾，表示视频片段 Width 单位为像素，如 100px 表示 Width 为100像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Width *string `json:"Width,omitempty" name:"Width"`

	// 视频片段的高度。支持 %、px 两种格式：
	// <li>当字符串以 % 结尾，表示视频片段 Height 为画布高度的百分比大小，如 10% 表示 Height 为画布高度的 10%；
	// </li><li>当字符串以 px 结尾，表示视频片段 Height 单位为像素，如 100px 表示 Height 为100像素。</li>
	// <li>当 Width、Height 均为空，则 Width 和 Height 取视频素材本身的 Width、Height。</li>
	// <li>当 Width 为空，Height 非空，则 Width 按比例缩放</li>
	// <li>当 Width 非空，Height 为空，则 Height 按比例缩放。</li>
	Height *string `json:"Height,omitempty" name:"Height"`

	// 对图像进行的操作，如图像旋转等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageOperations []*ImageTransform `json:"ImageOperations,omitempty" name:"ImageOperations" list`

	// 对音频进行操作，如静音等。
	// 注意：此字段可能返回 null，表示取不到有效值。
	AudioOperations []*AudioTransform `json:"AudioOperations,omitempty" name:"AudioOperations" list`
}

type WatermarkInput struct {

	// 水印模板 ID。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 文字内容，长度不超过100个字符。仅当水印类型为文字水印时填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextContent *string `json:"TextContent,omitempty" name:"TextContent"`

	// SVG 内容。长度不超过 2000000 个字符。仅当水印类型为 SVG 水印时填写。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvgContent *string `json:"SvgContent,omitempty" name:"SvgContent"`
}

type WatermarkTemplate struct {

	// 水印模板唯一标识。
	Definition *int64 `json:"Definition,omitempty" name:"Definition"`

	// 水印类型，取值：
	// <li>image：图片水印；</li>
	// <li>text：文字水印。</li>
	Type *string `json:"Type,omitempty" name:"Type"`

	// 水印模板名称。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Name *string `json:"Name,omitempty" name:"Name"`

	// 模板描述信息。
	Comment *string `json:"Comment,omitempty" name:"Comment"`

	// 水印图片原点距离视频图像原点的水平位置。
	// <li>当字符串以 % 结尾，表示水印 Left 为视频宽度指定百分比的位置，如 10% 表示 Left 为视频宽度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Left 为视频宽度指定像素的位置，如 100px 表示 Left 为 100 像素。</li>
	XPos *string `json:"XPos,omitempty" name:"XPos"`

	// 水印图片原点距离视频图像原点的垂直位置。
	// <li>当字符串以 % 结尾，表示水印 Top 为视频高度指定百分比的位置，如 10% 表示 Top 为视频高度的 10%；</li>
	// <li>当字符串以 px 结尾，表示水印 Top 为视频高度指定像素的位置，如 100px 表示 Top 为 100 像素。</li>
	YPos *string `json:"YPos,omitempty" name:"YPos"`

	// 图片水印模板，仅当 Type 为 image，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	ImageTemplate *ImageWatermarkTemplate `json:"ImageTemplate,omitempty" name:"ImageTemplate"`

	// 文字水印模板，仅当 Type 为 text，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	TextTemplate *TextWatermarkTemplateInput `json:"TextTemplate,omitempty" name:"TextTemplate"`

	// SVG 水印模板，当 Type 为 svg，该字段有效。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SvgTemplate *SvgWatermarkInput `json:"SvgTemplate,omitempty" name:"SvgTemplate"`

	// 模板创建时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	CreateTime *string `json:"CreateTime,omitempty" name:"CreateTime"`

	// 模板最后修改时间，使用 [ISO 日期格式](https://cloud.tencent.com/document/product/266/11732#iso-.E6.97.A5.E6.9C.9F.E6.A0.BC.E5.BC.8F)。
	UpdateTime *string `json:"UpdateTime,omitempty" name:"UpdateTime"`

	// 原点位置，可选值：
	// <li>topLeft：表示坐标原点位于视频图像左上角，水印原点为图片或文字的左上角；</li>
	// <li>topRight：表示坐标原点位于视频图像的右上角，水印原点为图片或文字的右上角；</li>
	// <li>bottomLeft：表示坐标原点位于视频图像的左下角，水印原点为图片或文字的左下角；</li>
	// <li>bottomRight：表示坐标原点位于视频图像的右下角，水印原点为图片或文字的右下。；</li>
	CoordinateOrigin *string `json:"CoordinateOrigin,omitempty" name:"CoordinateOrigin"`
}

type WeChatMiniProgramPublishRequest struct {
	*tchttp.BaseRequest

	// 媒体文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 发布视频所对应的转码模板 ID，为0代表原始视频。
	SourceDefinition *int64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// 点播[子应用](/document/product/266/14574) ID。如果要访问子应用中的资源，则将该字段填写为子应用 ID；否则无需填写该字段。
	SubAppId *uint64 `json:"SubAppId,omitempty" name:"SubAppId"`
}

func (r *WeChatMiniProgramPublishRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WeChatMiniProgramPublishRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WeChatMiniProgramPublishResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 任务 ID。
		TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *WeChatMiniProgramPublishResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *WeChatMiniProgramPublishResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type WechatMiniProgramPublishTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 发布视频文件 ID。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// 微信小程序视频发布状态，取值：
	// <li>Pass：发布成功；</li>
	// <li>Failed：发布失败；</li>
	// <li>Rejected：审核未通过。</li>
	PublishResult *string `json:"PublishResult,omitempty" name:"PublishResult"`
}

type WechatPublishTask struct {

	// 任务 ID。
	TaskId *string `json:"TaskId,omitempty" name:"TaskId"`

	// 任务状态，取值：
	// WAITING：等待中；
	// PROCESSING：处理中；
	// FINISH：已完成。
	Status *string `json:"Status,omitempty" name:"Status"`

	// 错误码
	// <li>0：成功；</li>
	// <li>其他值：失败。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	ErrCode *int64 `json:"ErrCode,omitempty" name:"ErrCode"`

	// 错误信息。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Message *string `json:"Message,omitempty" name:"Message"`

	// 发布视频文件 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	FileId *string `json:"FileId,omitempty" name:"FileId"`

	// 微信发布模板 ID。
	// 注意：此字段可能返回 null，表示取不到有效值。
	Definition *uint64 `json:"Definition,omitempty" name:"Definition"`

	// 发布视频所对应的转码模板 ID，为 0 代表原始视频。
	// 注意：此字段可能返回 null，表示取不到有效值。
	SourceDefinition *uint64 `json:"SourceDefinition,omitempty" name:"SourceDefinition"`

	// 微信发布状态，取值：
	// <li>FAIL：失败；</li>
	// <li>SUCCESS：成功；</li>
	// <li>AUDITNOTPASS：审核未通过；</li>
	// <li>NOTTRIGGERED：尚未发起微信发布。</li>
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatStatus *string `json:"WechatStatus,omitempty" name:"WechatStatus"`

	// 微信 Vid。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatVid *string `json:"WechatVid,omitempty" name:"WechatVid"`

	// 微信地址。
	// 注意：此字段可能返回 null，表示取不到有效值。
	WechatUrl *string `json:"WechatUrl,omitempty" name:"WechatUrl"`
}
