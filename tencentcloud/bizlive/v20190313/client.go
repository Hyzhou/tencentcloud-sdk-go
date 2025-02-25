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

package v20190313

import (
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-03-13"

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


func NewDescribeStreamPlayInfoListRequest() (request *DescribeStreamPlayInfoListRequest) {
    request = &DescribeStreamPlayInfoListRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bizlive", APIVersion, "DescribeStreamPlayInfoList")
    return
}

func NewDescribeStreamPlayInfoListResponse() (response *DescribeStreamPlayInfoListResponse) {
    response = &DescribeStreamPlayInfoListResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查询播放数据，支持按流名称查询详细播放数据，也可按播放域名查询详细总数据。
func (c *Client) DescribeStreamPlayInfoList(request *DescribeStreamPlayInfoListRequest) (response *DescribeStreamPlayInfoListResponse, err error) {
    if request == nil {
        request = NewDescribeStreamPlayInfoListRequest()
    }
    response = NewDescribeStreamPlayInfoListResponse()
    err = c.Send(request, response)
    return
}

func NewForbidLiveStreamRequest() (request *ForbidLiveStreamRequest) {
    request = &ForbidLiveStreamRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bizlive", APIVersion, "ForbidLiveStream")
    return
}

func NewForbidLiveStreamResponse() (response *ForbidLiveStreamResponse) {
    response = &ForbidLiveStreamResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 禁止某条流的推送，可以预设某个时刻将流恢复。
func (c *Client) ForbidLiveStream(request *ForbidLiveStreamRequest) (response *ForbidLiveStreamResponse, err error) {
    if request == nil {
        request = NewForbidLiveStreamRequest()
    }
    response = NewForbidLiveStreamResponse()
    err = c.Send(request, response)
    return
}

func NewRegisterIMRequest() (request *RegisterIMRequest) {
    request = &RegisterIMRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("bizlive", APIVersion, "RegisterIM")
    return
}

func NewRegisterIMResponse() (response *RegisterIMResponse) {
    response = &RegisterIMResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 注册聊天室
func (c *Client) RegisterIM(request *RegisterIMRequest) (response *RegisterIMResponse, err error) {
    if request == nil {
        request = NewRegisterIMRequest()
    }
    response = NewRegisterIMResponse()
    err = c.Send(request, response)
    return
}
