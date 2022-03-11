# Signatures playground

This project is considered as a test for packing and signing messages that may be sent over to a *-third party- to validate the authenticity and trust the content inside.

> \* This test does not cover sending the data to the third party, both roles are assumed in this solution.

## Tools

- Go 1.17
- Docker
- Visual Studio Code
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started
- Protocol Buffers
    > For more about what Protocol Buffers is please refer to: https://developers.google.com/protocol-buffers/docs/overview

## Data structure

Message data will contains the following attributes:

```
sender (String)
payload (String)
timestamp (Timestamp)
```

A *Container contains:

```
message (Message)
signature (Bytes)
```

> \* `Container` is a structure that contains the message that will be checked against the signature to validate the authenticity.

## Development

This project contains a `Dockerfile` file with all required dependencies to run it using `Visual Studio Code` + `Remote - Containers` extension.
However, if you want to make it run locally in your development machine, please follow the instructions below.

### Install Go

Install it from https://go.dev/dl/

### Protocol Buffers for Go

To install the `Protocol Buffers` compiler select the right OS and architecture from https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.4 and follow the commands:

> NOTE: Replace `protoc-3.19.4-linux-x86_64.zip` by your right configuration, **For the purpose of this project, we are using a Linux machine inside a docker container**. To get more information about other options and platforms please refer to https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

    unzip protoc-3.19.4-linux-x86_64.zip -d /usr/local

    rm protoc-3.19.4-linux-x86_64.zip

Run the following command to install the Go protocol buffers plugin:

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

> NOTE: To get more information about Protocol Buffers in Go, please referes to: https://developers.google.com/protocol-buffers/docs/gotutorial and https://github.com/protocolbuffers/protobuf-go

### Data generation

Only required if you add/remove new fields.

    protoc --go_out=. proto/*.proto