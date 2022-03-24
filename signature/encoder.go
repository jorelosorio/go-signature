package signature

import (
	"encoding/base64"
	"fmt"
	"os"

	"google.golang.org/protobuf/proto"
)

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.Strict().EncodeToString(data)
}

func DecodeBase64(data string) []byte {
	decoded, err := base64.StdEncoding.Strict().DecodeString(data)
	if err != nil {
		fmt.Printf("Failed to encode base64 data: %s", err)
		os.Exit(1)
	}

	return decoded
}

func EncodeMessage(message *Message) []byte {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		fmt.Printf("Failed to serialize message: %s", err)
		os.Exit(1)
	}

	return messageBytes
}

func EncodeContainer(container *Container) []byte {
	containerBytes, err := proto.Marshal(container)
	if err != nil {
		fmt.Printf("Failed to serialize container: %s", err)
		os.Exit(1)
	}

	return containerBytes
}

func DecodeContainer(data []byte) *Container {
	container := &Container{}
	err := proto.Unmarshal(data, container)
	if err != nil {
		fmt.Printf("Failed to decode container: %s", err)
		os.Exit(1)
	}

	return container
}
