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

package v20190718

import (
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common"
    tchttp "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/http"
    "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/profile"
)

const APIVersion = "2019-07-18"

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


func NewCreateAccessGroupRequest() (request *CreateAccessGroupRequest) {
    request = &CreateAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessGroup")
    return
}

func NewCreateAccessGroupResponse() (response *CreateAccessGroupResponse) {
    response = &CreateAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建权限组
func (c *Client) CreateAccessGroup(request *CreateAccessGroupRequest) (response *CreateAccessGroupResponse, err error) {
    if request == nil {
        request = NewCreateAccessGroupRequest()
    }
    response = NewCreateAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewCreateAccessRulesRequest() (request *CreateAccessRulesRequest) {
    request = &CreateAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateAccessRules")
    return
}

func NewCreateAccessRulesResponse() (response *CreateAccessRulesResponse) {
    response = &CreateAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量创建权限规则
func (c *Client) CreateAccessRules(request *CreateAccessRulesRequest) (response *CreateAccessRulesResponse, err error) {
    if request == nil {
        request = NewCreateAccessRulesRequest()
    }
    response = NewCreateAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewCreateFileSystemRequest() (request *CreateFileSystemRequest) {
    request = &CreateFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateFileSystem")
    return
}

func NewCreateFileSystemResponse() (response *CreateFileSystemResponse) {
    response = &CreateFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建文件系统（异步创建）
func (c *Client) CreateFileSystem(request *CreateFileSystemRequest) (response *CreateFileSystemResponse, err error) {
    if request == nil {
        request = NewCreateFileSystemRequest()
    }
    response = NewCreateFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewCreateMountPointRequest() (request *CreateMountPointRequest) {
    request = &CreateMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "CreateMountPoint")
    return
}

func NewCreateMountPointResponse() (response *CreateMountPointResponse) {
    response = &CreateMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 创建挂载点
func (c *Client) CreateMountPoint(request *CreateMountPointRequest) (response *CreateMountPointResponse, err error) {
    if request == nil {
        request = NewCreateMountPointRequest()
    }
    response = NewCreateMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessGroupRequest() (request *DeleteAccessGroupRequest) {
    request = &DeleteAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessGroup")
    return
}

func NewDeleteAccessGroupResponse() (response *DeleteAccessGroupResponse) {
    response = &DeleteAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除权限组
func (c *Client) DeleteAccessGroup(request *DeleteAccessGroupRequest) (response *DeleteAccessGroupResponse, err error) {
    if request == nil {
        request = NewDeleteAccessGroupRequest()
    }
    response = NewDeleteAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteAccessRulesRequest() (request *DeleteAccessRulesRequest) {
    request = &DeleteAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteAccessRules")
    return
}

func NewDeleteAccessRulesResponse() (response *DeleteAccessRulesResponse) {
    response = &DeleteAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量删除权限规则
func (c *Client) DeleteAccessRules(request *DeleteAccessRulesRequest) (response *DeleteAccessRulesResponse, err error) {
    if request == nil {
        request = NewDeleteAccessRulesRequest()
    }
    response = NewDeleteAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteFileSystemRequest() (request *DeleteFileSystemRequest) {
    request = &DeleteFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteFileSystem")
    return
}

func NewDeleteFileSystemResponse() (response *DeleteFileSystemResponse) {
    response = &DeleteFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除文件系统
func (c *Client) DeleteFileSystem(request *DeleteFileSystemRequest) (response *DeleteFileSystemResponse, err error) {
    if request == nil {
        request = NewDeleteFileSystemRequest()
    }
    response = NewDeleteFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDeleteMountPointRequest() (request *DeleteMountPointRequest) {
    request = &DeleteMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DeleteMountPoint")
    return
}

func NewDeleteMountPointResponse() (response *DeleteMountPointResponse) {
    response = &DeleteMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 删除挂载点
func (c *Client) DeleteMountPoint(request *DeleteMountPointRequest) (response *DeleteMountPointResponse, err error) {
    if request == nil {
        request = NewDeleteMountPointRequest()
    }
    response = NewDeleteMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessGroupsRequest() (request *DescribeAccessGroupsRequest) {
    request = &DescribeAccessGroupsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessGroups")
    return
}

func NewDescribeAccessGroupsResponse() (response *DescribeAccessGroupsResponse) {
    response = &DescribeAccessGroupsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看权限组列表
func (c *Client) DescribeAccessGroups(request *DescribeAccessGroupsRequest) (response *DescribeAccessGroupsResponse, err error) {
    if request == nil {
        request = NewDescribeAccessGroupsRequest()
    }
    response = NewDescribeAccessGroupsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeAccessRulesRequest() (request *DescribeAccessRulesRequest) {
    request = &DescribeAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeAccessRules")
    return
}

func NewDescribeAccessRulesResponse() (response *DescribeAccessRulesResponse) {
    response = &DescribeAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看权限规则列表
func (c *Client) DescribeAccessRules(request *DescribeAccessRulesRequest) (response *DescribeAccessRulesResponse, err error) {
    if request == nil {
        request = NewDescribeAccessRulesRequest()
    }
    response = NewDescribeAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemRequest() (request *DescribeFileSystemRequest) {
    request = &DescribeFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystem")
    return
}

func NewDescribeFileSystemResponse() (response *DescribeFileSystemResponse) {
    response = &DescribeFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看文件系统详细信息
func (c *Client) DescribeFileSystem(request *DescribeFileSystemRequest) (response *DescribeFileSystemResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemRequest()
    }
    response = NewDescribeFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeFileSystemsRequest() (request *DescribeFileSystemsRequest) {
    request = &DescribeFileSystemsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeFileSystems")
    return
}

func NewDescribeFileSystemsResponse() (response *DescribeFileSystemsResponse) {
    response = &DescribeFileSystemsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看文件系统列表
func (c *Client) DescribeFileSystems(request *DescribeFileSystemsRequest) (response *DescribeFileSystemsResponse, err error) {
    if request == nil {
        request = NewDescribeFileSystemsRequest()
    }
    response = NewDescribeFileSystemsResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointRequest() (request *DescribeMountPointRequest) {
    request = &DescribeMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoint")
    return
}

func NewDescribeMountPointResponse() (response *DescribeMountPointResponse) {
    response = &DescribeMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看挂载点详细信息
func (c *Client) DescribeMountPoint(request *DescribeMountPointRequest) (response *DescribeMountPointResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointRequest()
    }
    response = NewDescribeMountPointResponse()
    err = c.Send(request, response)
    return
}

func NewDescribeMountPointsRequest() (request *DescribeMountPointsRequest) {
    request = &DescribeMountPointsRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "DescribeMountPoints")
    return
}

func NewDescribeMountPointsResponse() (response *DescribeMountPointsResponse) {
    response = &DescribeMountPointsResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 查看挂载点列表
func (c *Client) DescribeMountPoints(request *DescribeMountPointsRequest) (response *DescribeMountPointsResponse, err error) {
    if request == nil {
        request = NewDescribeMountPointsRequest()
    }
    response = NewDescribeMountPointsResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessGroupRequest() (request *ModifyAccessGroupRequest) {
    request = &ModifyAccessGroupRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessGroup")
    return
}

func NewModifyAccessGroupResponse() (response *ModifyAccessGroupResponse) {
    response = &ModifyAccessGroupResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改权限组属性
func (c *Client) ModifyAccessGroup(request *ModifyAccessGroupRequest) (response *ModifyAccessGroupResponse, err error) {
    if request == nil {
        request = NewModifyAccessGroupRequest()
    }
    response = NewModifyAccessGroupResponse()
    err = c.Send(request, response)
    return
}

func NewModifyAccessRulesRequest() (request *ModifyAccessRulesRequest) {
    request = &ModifyAccessRulesRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyAccessRules")
    return
}

func NewModifyAccessRulesResponse() (response *ModifyAccessRulesResponse) {
    response = &ModifyAccessRulesResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 批量修改权限规则属性
func (c *Client) ModifyAccessRules(request *ModifyAccessRulesRequest) (response *ModifyAccessRulesResponse, err error) {
    if request == nil {
        request = NewModifyAccessRulesRequest()
    }
    response = NewModifyAccessRulesResponse()
    err = c.Send(request, response)
    return
}

func NewModifyFileSystemRequest() (request *ModifyFileSystemRequest) {
    request = &ModifyFileSystemRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyFileSystem")
    return
}

func NewModifyFileSystemResponse() (response *ModifyFileSystemResponse) {
    response = &ModifyFileSystemResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改文件系统属性
func (c *Client) ModifyFileSystem(request *ModifyFileSystemRequest) (response *ModifyFileSystemResponse, err error) {
    if request == nil {
        request = NewModifyFileSystemRequest()
    }
    response = NewModifyFileSystemResponse()
    err = c.Send(request, response)
    return
}

func NewModifyMountPointRequest() (request *ModifyMountPointRequest) {
    request = &ModifyMountPointRequest{
        BaseRequest: &tchttp.BaseRequest{},
    }
    request.Init().WithApiInfo("chdfs", APIVersion, "ModifyMountPoint")
    return
}

func NewModifyMountPointResponse() (response *ModifyMountPointResponse) {
    response = &ModifyMountPointResponse{
        BaseResponse: &tchttp.BaseResponse{},
    }
    return
}

// 修改挂载点属性
func (c *Client) ModifyMountPoint(request *ModifyMountPointRequest) (response *ModifyMountPointResponse, err error) {
    if request == nil {
        request = NewModifyMountPointRequest()
    }
    response = NewModifyMountPointResponse()
    err = c.Send(request, response)
    return
}
