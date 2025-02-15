package main

import (
	"fmt"

	"github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/errors"
	"github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/profile"
	tci "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/tci/v20190318"
)

func main() {
	credential := common.NewCredential(
		// os.Getenv("TENCENTCLOUD_SECRET_ID"),
		// os.Getenv("TENCENTCLOUD_SECRET_KEY"),
		"", "",
	)
	cpf := profile.NewClientProfile()
	cpf.HttpProfile.ReqMethod = "POST"
	cpf.HttpProfile.ReqTimeout = 30
	cpf.HttpProfile.Endpoint = "tci.tencentcloudapi.com"
	client, _ := tci.NewClient(credential, "ap-guangzhou", cpf)
	request := tci.NewDescribeAITaskResultRequest()
	request.TaskId = common.Int64Ptr(2716561495)
	// 通过client对象调用想要访问的接口，需要传入请求对象
	response, err := client.DescribeAITaskResult(request)
	// 处理异常
	fmt.Println(err)
	if _, ok := err.(*errors.TencentCloudSDKError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	// 非SDK异常，直接失败。实际代码中可以加入其他的处理。
	if err != nil {
		panic(err)
	}
	// 打印返回的json字符串
	fmt.Printf("%s", response.ToJsonString())
}
