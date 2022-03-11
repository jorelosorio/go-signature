# Signatures playground

This project is considered as a test for packing and signing messages that may be sent over to a -third party- to validate the authenticity and trust the content inside.

> NOTE: This test does not cover sending the data to the third party, both roles are assumed in this solution.

## Tools

- Go 1.17
- Docker
- Visual Studio Code
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started
- Protocol Buffers
    > For more about what Protocol Buffers is please refer to: https://developers.google.com/protocol-buffers/docs/overview

## Installation

### Protocol Buffers for Go

To install the `Protocol Buffers` compiler using the Linux machine in `dev containers` please run the following commands:

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

    unzip protoc-3.19.4-linux-x86_64.zip -d /usr/local

    rm protoc-3.19.4-linux-x86_64.zip

> To get more details about the above commands please refer to https://developers.google.com/protocol-buffers/docs/gotutorial

## Data generation

    protoc --go_out=. proto/*.proto

## Data structure

Message data will contains the following attributes:

```
sender (String)
payload (String)
timestamp (Timestamp)
```