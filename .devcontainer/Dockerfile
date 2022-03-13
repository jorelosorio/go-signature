FROM mcr.microsoft.com/vscode/devcontainers/go:1.17

RUN apt-get update

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.19.4/protoc-3.19.4-linux-x86_64.zip

RUN unzip protoc-3.19.4-linux-x86_64.zip -d /usr/local

RUN rm protoc-3.19.4-linux-x86_64.zip

RUN go mod download