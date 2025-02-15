# 简介

欢迎使用腾讯云开发者工具套件（SDK）3.0，SDK3.0 是云 API3.0 平台的配套工具。后续所有的云服务产品都会接入进来。新版 SDK实 现了统一化，具有各个语言版本的 SDK 使用方法相同，接口调用方式相同，统一的错误码和返回包格式这些优点。

为方便 GO 开发者调试和接入腾讯云产品 API，这里向您介绍适用于 GO 的腾讯云开发工具包，并提供首次使用开发工具包的简单示例。让您快速获取腾讯云 GO SDK 并开始调用。

# 依赖环境

1. Go 1.9 版本及以上，并设置好 GOPATH 等必须的环境变量。
2. 使用相关产品前需要在腾讯云控制台已开通相应产品。
3. 在腾讯云控制台[访问管理](https://console.cloud.tencent.com/cam/capi)页面获取 SecretID 和 SecretKey 。

# 获取安装

安装 Go SDK 前，先获取安全凭证。在第一次使用云 API 之前，用户首先需要在腾讯云控制台上申请安全凭证，安全凭证包括 SecretID 和 SecretKey, SecretID 是用于标识 API 调用者的身份，SecretKey 是用于加密签名字符串和服务器端验证签名字符串的密钥。SecretKey 必须严格保管，避免泄露。

## 通过go get安装（推荐）

推荐使用语言自带的工具安装 SDK ：

    go get -u github.com/Hyzhou/tencentcloud-sdk-go


## 通过源码安装

前往 [Github 代码托管地址](https://github.com/Hyzhou/tencentcloud-sdk-go) 下载最新代码，解压后安装到 $GOPATH/src/github.com/tencentcloud 目录下。

# 示例

每个接口都有一个对应的 Request 结构和一个 Response 结构。例如云服务器的查询实例列表接口 DescribeInstances 有对应的请求结构体 DescribeInstancesRequest 和返回结构体 DescribeInstancesResponse 。

下面以云服务器查询实例列表接口为例，介绍 SDK 的基础用法。出于演示的目的，有一些非必要的内容也加上去了，以尽量展示 SDK 常用的功能，但也显得臃肿。在实际编写代码使用 SDK 的时候，应尽量简化。

```
package main

import (
        "fmt"
        "os"

        "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common"
        "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/errors"
        "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/profile"
        "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/common/regions"
        cvm "github.com/Hyzhou/tencentcloud-sdk-go/tencentcloud/cvm/v20170312"
)

func main() {
        // 必要步骤：
        // 实例化一个认证对象，入参需要传入腾讯云账户密钥对secretId，secretKey。
        // 这里采用的是从环境变量读取的方式，需要在环境变量中先设置这两个值。
        // 你也可以直接在代码中写死密钥对，但是小心不要将代码复制、上传或者分享给他人，
        // 以免泄露密钥对危及你的财产安全。
        credential := common.NewCredential(
                os.Getenv("TENCENTCLOUD_SECRET_ID"),
                os.Getenv("TENCENTCLOUD_SECRET_KEY"),
        )

        // 非必要步骤
        // 实例化一个客户端配置对象，可以指定超时时间等配置
        cpf := profile.NewClientProfile()
        // SDK默认使用POST方法。
        // 如果你一定要使用GET方法，可以在这里设置。GET方法无法处理一些较大的请求。
        // 如非必要请不要修改默认设置。
        //cpf.HttpProfile.ReqMethod = "GET"
        // SDK有默认的超时时间，如非必要请不要修改默认设置。
        // 如有需要请在代码中查阅以获取最新的默认值。
        //cpf.HttpProfile.ReqTimeout = 10
        // SDK会自动指定域名。通常是不需要特地指定域名的，但是如果你访问的是金融区的服务，
        // 则必须手动指定域名，例如云服务器的上海金融区域名： cvm.ap-shanghai-fsi.tencentcloudapi.com
        //cpf.HttpProfile.Endpoint = "cvm.tencentcloudapi.com"
        // SDK默认用HmacSHA256进行签名，它更安全但是会轻微降低性能。
        // 如非必要请不要修改默认设置。
        //cpf.SignMethod = "HmacSHA1"
        // SDK 默认用 zh-CN 调用返回中文。此外还可以设置 en-US 返回全英文。
        // 但大部分产品或接口并不支持全英文的返回。
        // 如非必要请不要修改默认设置。
        //cpf.Language = "en-US"

        // 实例化要请求产品(以cvm为例)的client对象
        // 第二个参数是地域信息，可以直接填写字符串ap-guangzhou，或者引用预设的常量
        client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
        // 实例化一个请求对象，根据调用的接口和实际情况，可以进一步设置请求参数
        // 你可以直接查询SDK源码确定DescribeInstancesRequest有哪些属性可以设置，
        // 属性可能是基本类型，也可能引用了另一个数据结构。
        // 推荐使用IDE进行开发，可以方便的跳转查阅各个接口和数据结构的文档说明。
        request := cvm.NewDescribeInstancesRequest()

        // 基本类型的设置。
        // 此接口允许设置返回的实例数量。此处指定为只返回一个。
        // SDK采用的是指针风格指定参数，即使对于基本类型你也需要用指针来对参数赋值。
        // SDK提供对基本类型的指针引用封装函数
        request.Limit = common.Int64Ptr(1)

        // 数组类型的设置。
        // 此接口允许指定实例 ID 进行过滤，但是由于和接下来要演示的 Filter 参数冲突，先注释掉。
        // request.InstanceIds = common.StringPtrs([]string{"ins-r8hr2upy"})

        // 复杂对象的设置。
        // 在这个接口中，Filters是数组，数组的元素是复杂对象Filter，Filter的成员Values是string数组。
        request.Filters = []*cvm.Filter{
            &cvm.Filter{
                Name: common.StringPtr("zone"),
                Values: common.StringPtrs([]string{"ap-guangzhou-1"}),
            },
        }

        // 使用json字符串设置一个request，注意这里实际是更新request，即Limit=1将会被保留，
        // 而过滤条件的zone将会变为ap-guangzhou-2。
        // 如果需要一个全新的request，则需要用cvm.NewDescribeInstancesRequest()创建。
        err := request.FromJsonString(`{"Filters":[{"Name":"zone","Values":["ap-guangzhou-2"]}]}`)
        if err != nil {
                panic(err)
        }
        // 通过client对象调用想要访问的接口，需要传入请求对象
        response, err := client.DescribeInstances(request)
        // 处理异常
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
```

更多示例参见 [examples](https://github.com/Hyzhou/tencentcloud-sdk-go/tree/master/examples) 目录。对于复杂接口的 Request 初始化例子，可以参考 examples/cvm/v20170312/run_instances.go 。对于使用json字符串初始化 Request 的例子，可以参考 examples/cvm/v20170312/describe_instances.go 。

# 相关配置

## 代理

如果是有代理的环境下，需要设置系统环境变量 `https_proxy` ，否则可能无法正常调用，抛出连接超时的异常。

## 开启 DNS 缓存

当前 GO SDK 总是会去请求 DNS 服务器，而没有使用到 nscd 的缓存，可以通过导出环境变量`GODEBUG=netdns=cgo`，或者`go build`编译时指定参数`-tags 'netcgo'`控制读取 nscd 缓存。

## 忽略服务器证书校验

虽然使用 SDK 调用公有云服务时，必须校验服务器证书，以识破他人伪装的服务器，确保请求的安全。
但是某些极端情况下，例如测试时，你可能会需要忽略自签名的服务器证书。
以下是其中一种可能的方法：

```
import "crypto/tls"
...
    client, _ := cvm.NewClient(credential, regions.Guangzhou, cpf)
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    }
    client.WithHttpTransport(tr)
...
```

再次强调，除非你知道自己在做什么，并明白由此带来的风险，否则不要尝试关闭服务器证书校验。

# 支持产品列表

| 包名 | 产品中文名 |
|------------|------|
| aai | 智能语音服务 |
| as | 弹性伸缩 |
| batch | 批量计算 |
| billing | 计费 |
| bm | 黑石物理服务器 |
| bmeip | 黑石弹性 IP |
| cbs | 云硬盘 |
| cdb | 数据库 MySQL |
| cdn | 内容分发网络 |
| cis | 容器实例服务 |
| clb | 负载均衡 |
| cr | 催收机器人 |
| cvm | 云服务器 |
| cws | Web 漏洞扫描 |
| dc | 专线接入 |
| dcdb | 分布式数据库 DCDB |
| drm | 数字版权管理 |
| ds | 电子合同服务 |
| dts | 数据传输服务 DTS |
| es | Elasticsearch Service |
| facefusion | 人脸融合 |
| faceid | 人脸核身·云智慧眼 |
| habo | 样本智能分析平台 |
| hcm | 数学作业批改 |
| iai | 人脸识别 |
| iot | 加速物联网套件 |
| iotcloud | 物联网通信 |
| live | 云直播 |
| mariadb | 数据库 MariaDB(TDSQL) |
| mongodb | 云数据库 MongoDB |
| monitor | 云监控 |
| ms | 应用安全 |
| msp | 迁移服务平台 |
| partners | 渠道合作伙伴 |
| postgres | 数据库 PostgreSQL |
| redis | 云数据库 Redis |
| scf | 无服务器云函数 |
| soe | 智聆口语评测 |
| sqlserver | 云数据库 SQL Server |
| sts | 访问管理 |
| tbaas | TBaaS |
| tbm | 腾讯优评 |
| tmt | 腾讯机器翻译 |
| vod | 云点播 |
| vpc | 私有网络 |
| youmall | 腾讯优Mall |
| yunjing | 主机安全 |
