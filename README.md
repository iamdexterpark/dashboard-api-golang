# dashboard-api-golang

## Introduction
A Community developed Golang Library for the Cisco Meraki Dashboard API. 
This library was generated using [go-swagger](https://github.com/go-swagger/go-swagger) against the Cisco Meraki [OpenAPI Specification.](https://github.com/meraki/openapi)

Inspired by the vendor supported  [Meraki Dashboard API Python Library](https://github.com/meraki/dashboard-api-python).

## Requirements
[Go Programming Language](https://golang.org/doc/install) >= 1.16

## Getting Started

Please see [INSTALLATION.md](/docs/en/INSTALLATION.md) to get started.
Reference code can be found in the [examples folder](/examples).

## Disclaimer

Even with the Dashboard rate-limit, you have a theoretical potential to make:

| Number of API Calls  | 5            | 300          | 180,000    | 1,400,000   |
|----------------------|--------------|--------------|------------|-------------|
| **Time to Complete** | **1 second** | **1 minute** | **1 hour** | **8 hours** |

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

## Contributing

Contributions are always appreciated. Please use following guidelines:

- [Contributor Guidelines](.github/CONTRIBUTING.md)
- [Code of Conduct](.github/CODE_OF_CONDUCT.md)
- [Pull Request Template](.github/PULL_REQUEST_TEMPLATE.md)
- [Issue Template](.github/ISSUE_TEMPLATE.md)


Full list of [Contributors](https://github.com/ddexterpark/dashboard-api-golang/graphs/contributors) who participated in this project.