# dashboard-api-golang 

## Installation

Install the [Go](http://golang.org) programming language.

#### Set PATH
```bash
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

#### Download Project

```bash
go get github.com/ddexterpark/dashboard-api-golang
```

## Getting Started

At minimum, to use this tool you will need to 
[Enable API Access](https://documentation.meraki.com/General_Administration/Other_Topics/Cisco_Meraki_Dashboard_API) 
for the target organization and generate an API key.

I highly recommend you store this key as an environmental variable.

**MERAKI_DASHBOARD_API_KEY**
```shell script
      export MERAKI_DASHBOARD_API_KEY=1234567890987654321
```

## Shard URL

It is possible to use set our client to use a specific URL.

Reference shards:

| name       | URI                              | 
|------------|----------------------------------|
| Default    | https://api.meraki.com/api/v1    |
| China      | https://api.meraki.cn/api/v1     |
| mega-proxy | https://api-mp.meraki.com/api/v1 |


## Reference Code 
The [examples](/examples) folder contains scripts for common configuration tasks and is an excellent reference point.
