# Go Signature

This project is the result of learning `Signatures` for packing and signing messages that may be sent out to a **-third-party-server-** to validate the authenticity and trust of the content.

> NOTE: This project does not send the data to any server, but both roles are assumed in this solution.

This project makes a CLI tool to make it easy to test the encoding, decoding and signing features.

## Try it

> Use it now with a Docker instance. It will open an interactive shell where you can start to play with some commands. Please refer to `How to use the command CLI tool` below, for more details.

    docker pull jorelosorio/go-signature:latest

    docker run -i -t jorelosorio/go-signature

## How to use the command CLI tool

To create a new pair of private and public keys and export them as `private_key.pem` and `public_key.pem`

    sign create-keys --export-path .

To encode a new message

    sign encode-message --sender "Jorge Osorio" --payload "HOLA" --private-key-path ./private_key.pem

To encode a new message importing the `payload` from a file

    sign encode-message --sender "Jorge Osorio" --payload-path ./payload.txt --private-key-path ./private_key.pem

To decode a new message

    sign decode-message --public-key-path ./public_key.pem --base64-message {BASE64_ENCODED_CONTAINER_DATA}

> If you want more details about a specific command usage use `--help` argument. For instance: `sign encode-message --help` . For general information run `sign --help`. There are also shortcuts for the commands in the help description.

## Tools

- GoLang `1.17`
- Docker
- Visual Studio Code `Optional!`
    > It requires a `Remote - Containers` extension. for more information please refers to: https://code.visualstudio.com/docs/remote/containers#_getting-started
- Protocol Buffers `3.19.4`
    > For more about what Protocol Buffers is please refer to: https://developers.google.com/protocol-buffers/docs/overview

## Data structure

A `Message` struct has the following attributes.

```
Sender (String)
Payload (Bytes)
```

A * `Container` struct has the following attributes.

```
Signature (Bytes)
Message (Message)
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

    protoc --go_out=. proto/*.proto

### Build the command CLI tool

In the main workspace path run the following command to generate a build executable

    go build -o bin/sign

> **If you are using windows the `sign` will become `sign.exe` instead.**

### Build Docker

To build the docker image use `Dockerfile.deploy` and the command

    docker build -f Dockerfile.deploy -t jorelosorio/go-signature:latest .

To run the docker image as an interactive shell

    docker run -i -t jorelosorio/go-signature

If everything goes well run:

    sign --help