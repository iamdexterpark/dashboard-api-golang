＃dashboard-api-golang
Meraki仪表板API Go语言库

*多语言文档：[英文](README.md)，[日本语](docs / README.ja.md)，[简体中文](docs / README.zh-cn.md)。*

社区为Meraki Dashboard API开发了Golang库。
有关供应商支持的API接口，请参见精彩内容：[Meraki仪表板API Python库](https://github.com/meraki/dashboard-api-python)

##从源代码运行

####安装
安装[Go](http://golang.org)编程语言。

####设置路径

```
导出GOPATH = $ HOME / go
导出PATH = $ PATH：$ GOPATH / bin
```

####下载项目

``
去获取github.com/ddexterpark/dashboard-api-meraki
```

####编译(可选)
外壳脚本
    ＃Linux / MacOS
    去建立main.go

    ＃Windows 64位
    env GOOS = windows GOARCH = amd64编译main.go
    
    ＃Windows 32位
    env GOOS = windows GOARCH = 386去建立main.go
```
    
＃＃ 环境变量

至少，要使用此工具，您将需要为API密钥设置一个环境变量。
您还可以设置一些可选的env var来自定义API调用：

####必填

** MERAKI_API_TOKEN **
外壳脚本
      重击-
      出口MERAKI_API_TOKEN = 1234567890987654321
      回声$ MERAKI_API_TOKEN
      
      电源外壳 -
            setx MERAKI_API_TOKEN“ 1234567890987654321”
            回声％MERAKI_API_TOKEN％
```
＃＃＃＃ 可选的
 
** MERAKI_API_URL **
外壳脚本
        默认='https://api.meraki.com/api/'
        中国='https://api.meraki.cn/api/'
```

** MERAKI_API_VERSION **

默认版本为v1，此工具仅对v0提供有限支持，因为该版本将于2022年停用。
并非所有端点都可以在v0中工作。
 
外壳脚本
    默认='v1'
```
    
＃＃ 句法

外壳脚本
    merakictl [命令] [子命令] [目标] [标志]
```


###免责声明

User-Agent非常强大。即使使用仪表板速率限制，您也有理论上的潜力：


API调用数| 5 | 300 | 180,000 | 1,400,000 |
--- | --- | --- | --- | --- |
**完成时间** | ** 1秒** | ** 1分钟** | ** 1小时** | ** 8小时** |


进行生产更改时请小心。如果不确定该工具，请查看代码。

切勿运行您不理解的程序。在接触生产之前，请在测试环境中练习使用此工具。

创建生产变更计划。在测试环境中实施计划的各个方面，以确保对更改有最高的信心。

####出色的生产变更计划的要素包括：
-**同行审查**要求其他人审查您的测试计划，请他们在您的测试环境中运行它。
-**预先检查**捕获更改前的网络状态。
-**事后检查**捕获更改后的网络状态。
-**备份配置**复制配置，以便在发生回滚时可以重新应用它。
-**回滚过程**请勿轻易采取此步骤，否则会出错。
最糟糕的情况是更改失败并且没有经过测试的可靠回滚计划。
-**指数变更时间表**不要一次做所有事情。从单个网络开始，
监视它，给它时间正常运行，然后，如果没有错误，则调度下5个网络，然后是10、25、50、100等。
-**故障阈值**一批网络可接受的故障更改百分比
在取消所有计划的更改之前？通常，根据您的规模，可接受1-5％的费用。
一切都需要根本原因分析和计划修改。