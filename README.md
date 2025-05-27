# go-wsman-messages

[![CodeQL](https://img.shields.io/github/actions/workflow/status/open-amt-cloud-toolkit/go-wsman-messages/codeql-analysis.yml?style=for-the-badge&label=CodeQL&logo=github)](https://github.com/device-management-toolkit/go-wsman-messages/actions/workflows/codeql-analysis.yml)
[![Build](https://img.shields.io/github/actions/workflow/status/open-amt-cloud-toolkit/go-wsman-messages/ci.yml?style=for-the-badge&logo=github)](https://github.com/device-management-toolkit/go-wsman-messages/actions/workflows/ci.yml)
[![Codecov](https://img.shields.io/codecov/c/github/open-amt-cloud-toolkit/go-wsman-messages?style=for-the-badge&logo=codecov)](https://app.codecov.io/gh/open-amt-cloud-toolkit/go-wsman-messages)
[![OSSF-Scorecard Score](https://img.shields.io/ossf-scorecard/github.com/device-management-toolkit/go-wsman-messages?style=for-the-badge&label=OSSF%20Score)](https://api.securityscorecards.dev/projects/github.com/device-management-toolkit/go-wsman-messages)
[![Discord](https://img.shields.io/discord/1063200098680582154?style=for-the-badge&label=Discord&logo=discord&logoColor=white&labelColor=%235865F2&link=https%3A%2F%2Fdiscord.gg%2FDKHeUNEWVH)](https://discord.gg/DKHeUNEWVH)

> Disclaimer: Production viable releases are tagged and listed under 'Releases'.  All other check-ins should be considered 'in-development' and should not be used in production

This repository contains a Go library that implements APIs for communicating with Intel® Active Management Technology (AMT) capable platforms. These APIs are based on the AMT SDK documentation, which can be found [here](https://www.intel.com/content/www/us/en/developer/tools/active-management-technology-sdk/overview.html).

## How to use it

A few steps are required to use the library in your code.

1. Import the library in your project,
1. Provide the connection parameters,
1. Create the desired messages to act on the AMT device.

See the commented source code as follows:

``` go
package main

import (
    "fmt"
    // 1. Import the library
    "github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman"
    "github.com/device-management-toolkit/go-wsman-messages/v2/pkg/wsman/client"
)

func main() {
    // 2. Provide the connection parameters
    clientParams := client.Parameters{
        Target:            "192.168.1.164",
        Username:          "admin",          
        Password:          "Your_AMT_Password",
        UseDigest:         true,
        UseTLS:            true,
        SelfSignedAllowed: true,
        LogAMTMessages:    true,
    }

    // 3.a Create an instance to send the messages and defer the close of the connection.
    // NewMessages instantiates a new Messages class with client connection parameters.
    // Messages implements client.WSMan, amt.Messages, cim.Messages, and ips.Messages.
    amtClass := wsman.NewMessages(clientParams)
    defer amtClass.Client.CloseConnection()

    // 3.b Query the AMT General Settings
    gset, err := amtClass.AMT.GeneralSettings.Get()
    if err != nil {
        fmt.Println("Error getting AMT General Settings:", err)
        return
    } else {
        fmt.Println("")
        fmt.Println("AMT General Settings: ")
        fmt.Println(string(gset.JSON()))
    }

    // 3.c Recover the Audit Log in chronological order, starting with the oldest one (i.e., 1)
    alog, err := amtClass.AMT.AuditLog.ReadRecords(1)
    if err != nil {
        fmt.Println("Error getting AMT Audit Log:", err)
        return
    } else {
        fmt.Println("")
        fmt.Println("AMT Audit Log: ")
        fmt.Println(string(alog.JSON()))
    }

    // 3.d Get the Processor Information
    aproc, err := amtClass.CIM.Processor.Get()
    if err != nil {
        fmt.Println("Error getting Processor Info:", err)
        return
    } else {
        fmt.Println("")
        fmt.Println("Processor Info: ")
        fmt.Println(string(aproc.JSON()))
    }

    // 3.d Get the Power Management Service Information
    power, err := amtClass.CIM.PowerManagementService.Get()
    if err != nil {
        fmt.Println("Error getting Power Info:", err)
        return
    } else {
        fmt.Println("")
        fmt.Println("Power Info: ")
        fmt.Println(string(power.JSON()))
    }
}
```

As shown, you can call the various methods of the wsman.Messages struct. go-wsman-messages authenticates with AMT using the client parameters provided, sends messages to the Intel® AMT device, and handles responses, returning a package-specific Response struct or error message.

## Dev tips for passing CI Checks

- Install gofumpt `go install mvdan.cc/gofumpt@latest` (replaces gofmt)
- Install gci `go install github.com/daixiang0/gci@latest` (organizes imports)
- Ensure code is formatted correctly with `gofumpt -l -w -extra ./`
- Ensure code is gci'd with `gci.exe write --skip-generated -s standard -s default .`
- Ensure all unit tests pass with `go test ./...`
- Ensure code has been linted with `docker run --rm -v ${pwd}:/app -w /app golangci/golangci-lint:latest golangci-lint run -v`
