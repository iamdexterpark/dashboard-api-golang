# dashboard-api-golang
Meraki Dashboard API Go Lang Library

*Multi-language Documentation: [English](README.md), [日本語](docs/README.ja.md), [简体中文](docs/README.zh-cn.md).*

A Community developed Golang Library for the Meraki Dashboard API. 
For a Vendor supported API interface please see the wonderful: [Meraki Dashboard API Python Library](https://github.com/meraki/dashboard-api-python)

## Run From Source Code 

#### Installation
Install the [Go](http://golang.org) programming language.

#### Set PATH
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### Download Project

```bash
go get github.com/ddexterpark/dashboard-api-meraki
```

#### Compile (Optional) 
```shell script
    # Linux/MacOS
    go build main.go

    # Windows 64-bit
    env GOOS=windows GOARCH=amd64 go build main.go
    
    # Windows 32-bit
    env GOOS=windows GOARCH=386 go build main.go
```
    
## Environment Variables

At minimum, to use this tool you will need to set an environmental variable for the API key. 
There are also some optional env vars that you can set to customize your API calls:

#### Required

**MERAKI_API_TOKEN**
```shell script
      Bash -
      export MERAKI_API_TOKEN=1234567890987654321
      echo $MERAKI_API_TOKEN 
      
      PowerShell -
            setx MERAKI_API_TOKEN "1234567890987654321"
            echo %MERAKI_API_TOKEN%
```
#### Optional
 
**MERAKI_API_URL**
```shell script
        Default = 'https://api.meraki.com/api/'
        China = 'https://api.meraki.cn/api/' 
```

**MERAKI_API_VERSION**

The default version is v1, this tool has limited support for v0 as it is being sunset in 2022. 
Not all endpoints will work in v0.
 
```shell script
    Default = 'v1'
```
    
## Syntax

```shell script
    merakictl [COMMAND] [SUBCOMMAND] [TARGET]  [flags]
```


### Disclaimer

The User-Agent is extremely powerful. Even with the Dashboard rate-limit, you have a theoretical potential to make:


Number of API Calls | 5 | 300 | 180,000 | 1,400,000 |
--- | --- | --- | --- | --- |
**Time to Complete** | **1 second** | **1 minute** | **1 hour** | **8 hours** |


Please be cautious when making production changes. If you are unsure about the tool, look at the code. 

Never run a program that you do not understand. Practice using this tool in a test environment before touching production.

Create a production change plan. Implement every aspect of your plan in a test environment to ensure the highest levels of confidence in your change.   

#### Elements of a great production change plan include:
- **Peer Review** Have someone else review your test plan, ask them to run it in your test environment.
- **Pre-checks**  Capture the state of the network before the change.
- **Post-checks** Capture the state of the network after the change. 
- **Backup Config** Copy the config so that you can re-apply it in the event of a rollback.
- **Rollback Procedure** Do not take this step lightly, things go wrong. 
The worst possible position is to have a change fail and not have a tested, reliable rollback plan.
- **Exponential Change Schedule** Don't do everything at once. Start with a single network, 
monitor it, give it time to operate normally, then, if nothing is wrong schedule the next 5 networks, then 10, 25, 50, 100, etc..
- **Failure threshold** What percentage of failed changes are acceptable in a batch of networks 
before all scheduled changes are canceled? Typically, 1-5% is acceptable depending on your scale. 
Anything over that needs a root cause analysis, and modification of your plan. 
