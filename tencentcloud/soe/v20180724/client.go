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

package v20180724

import (
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2018-07-24"

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


func NewInitOralProcessRequest() (request *InitOralProcessRequest) {
    request = &InitOralProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("soe", APIVersion, "InitOralProcess")
    return
}

func NewInitOralProcessResponse() (response *InitOralProcessResponse) {
    response = &InitOralProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化发音评估过程，每一轮评估前进行调用。语音输入模式分为流式模式和非流式模式，流式模式支持数据分片传输，可以加快评估响应速度。评估模式分为词模式和句子模式，词模式会标注每个音节的详细信息；句子模式会有完整度和流利度的评估。
func (c *Client) InitOralProcess(request *InitOralProcessRequest) (response *InitOralProcessResponse, err error) {
    if request == nil {
        request = NewInitOralProcessRequest()
    }
    response = NewInitOralProcessResponse()
    err = c.Send(request, response)
    return
}

func NewKeywordEvaluateRequest() (request *KeywordEvaluateRequest) {
    request = &KeywordEvaluateRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("soe", APIVersion, "KeywordEvaluate")
    return
}

func NewKeywordEvaluateResponse() (response *KeywordEvaluateResponse) {
    response = &KeywordEvaluateResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 指定主题关键词词汇评估，分析语音与关键词的切合程度，可指定多个关键词，支持中文英文同时评测。分片传输时，尽量保证纯异步调用，即不等待上一个分片的传输结果边录边传，这样可以尽可能早的提供音频数据。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
func (c *Client) KeywordEvaluate(request *KeywordEvaluateRequest) (response *KeywordEvaluateResponse, err error) {
    if request == nil {
        request = NewKeywordEvaluateRequest()
    }
    response = NewKeywordEvaluateResponse()
    err = c.Send(request, response)
    return
}

func NewTransmitOralProcessRequest() (request *TransmitOralProcessRequest) {
    request = &TransmitOralProcessRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("soe", APIVersion, "TransmitOralProcess")
    return
}

func NewTransmitOralProcessResponse() (response *TransmitOralProcessResponse) {
    response = &TransmitOralProcessResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 传输音频数据，必须在完成发音评估初始化接口之后调用，且SessonId要与初始化接口保持一致。分片传输时，尽量保证SeqId顺序传输。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
func (c *Client) TransmitOralProcess(request *TransmitOralProcessRequest) (response *TransmitOralProcessResponse, err error) {
    if request == nil {
        request = NewTransmitOralProcessRequest()
    }
    response = NewTransmitOralProcessResponse()
    err = c.Send(request, response)
    return
}

func NewTransmitOralProcessWithInitRequest() (request *TransmitOralProcessWithInitRequest) {
    request = &TransmitOralProcessWithInitRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("soe", APIVersion, "TransmitOralProcessWithInit")
    return
}

func NewTransmitOralProcessWithInitResponse() (response *TransmitOralProcessWithInitResponse) {
    response = &TransmitOralProcessWithInitResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 初始化并传输音频数据，分片传输时，尽量保证SeqId顺序传输。音频源目前仅支持16k采样率16bit单声道编码方式，如有不一致可能导致评估不准确或失败。
func (c *Client) TransmitOralProcessWithInit(request *TransmitOralProcessWithInitRequest) (response *TransmitOralProcessWithInitResponse, err error) {
    if request == nil {
        request = NewTransmitOralProcessWithInitRequest()
    }
    response = NewTransmitOralProcessWithInitResponse()
    err = c.Send(request, response)
    return
}
