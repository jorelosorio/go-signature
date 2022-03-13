# Signatures playground

This project is considered as a test for packing and signing messages that may be sent over to a *-third party- to validate the authenticity and trust the content inside.

> \* This test does not cover sending the data to the third party, both roles are assumed in this solution.

## Tools

- Go `1.17`
- Docker
- Visual Studio Code
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started
- Protocol Buffers `3.19.4`
    > For more about what Protocol Buffers is please refer to: https://developers.google.com/protocol-buffers/docs/overview

## Data structure

Message struct has the following attributes.

```
sender (String)
payload (String)
timestamp (Timestamp)
```

A * `Container` has the following attributes.

```
message (Message)
signature (Bytes)
```

> \* `Container` is the main structure to be sent, because it contains the message to be transported and the signature associated with it, **When validating the authenticity it shall be made with the message**.

## Development

This project contains a `Dockerfile` file with all required dependencies to run it using `Visual Studio Code` + `Remote - Containers` extension.
However, if you want to make it run locally in your development machine, please follow the instructions below.

### Install Go

Install it from https://go.dev/dl/

### Protocol Buffers for Go

To install the `Protocol Buffers` compiler select the right OS and architecture from https://github.com/protocolbuffers/protobuf/releases/tag/v3.19.4 and follow the commands.

> NOTE: Replace `protoc-3.19.4-linux-x86_64.zip` by your right configuration, **For the purpose of this project, we are using a Linux machine inside a docker container**. To get more information about other options and platforms please refer to https://github.com/protocolbuffers/protobuf/blob/master/src/README.md

    wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

    unzip protoc-3.19.4-linux-x86_64.zip -d /usr/local

    rm protoc-3.19.4-linux-x86_64.zip

Run the following command to install the Go protocol buffers plugin

    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

> NOTE: To get more information about Protocol Buffers in Go, please referes to: https://developers.google.com/protocol-buffers/docs/gotutorial and https://github.com/protocolbuffers/protobuf-go

### Data generation

Only required if you add/remove new fields

    protoc --go_out=. proto/*.proto¨


### Build the command CLI tool

In the main workspace path run the following command to generate a build executable

    go build -o bin/sp

### Using the CLI tool

> If you ran the build command, make sure to change the directory to the `bin` folder. **If you are using windows the `sp` will become `sp.exe` instead.**

To create a new pair of private and public keys and export them as `private_key.pem` and `public_key.pem`

    sp create-keys --export-path .

> If `sp` is not found, try `./sp` in the same `bin` folder. **You could move it to your local bin directory as well.**

To pack a new message

    sp pack-message --sender "Jorge Osorio" --payload "HOLA" --private-key-path ./private_key.pem

To pack a new message importing the data from a file

    sp pack-message --sender "Jorge Osorio" --payload-path ./payload.txt --private-key-path ./private_key.pem

To unpack a new message

    sp unpack-message --public-key-path ./public_key.pem --base64-message {BASE64_ENCODED_CONTAINER_DATA}

> If you want more details about a specific command usage use `--help` argument. For instance: `sp pack-message --help` . For general information run `sp --help`.