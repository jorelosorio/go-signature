package helpers

import (
	"encoding/base64"
	"fmt"
	"os"
	pb "signatures-playground/structspb"

	"google.golang.org/protobuf/proto"
)

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.Strict().Strict().EncodeToString(data)
}

func DecodeBase64(data string) []byte {
	decoded, err := base64.StdEncoding.Strict().DecodeString(data)
	if err != nil {
		fmt.Printf("Failed to encode base64 data: %s", err)
		os.Exit(1)
	}

	return decoded
}

func EncodeMessage(message *pb.Message) []byte {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		fmt.Printf("Failed to serialize message: %s", err)
		os.Exit(1)
	}

	return messageBytes
}

func EncodeContainer(container *pb.Container) []byte {
	containerBytes, err := proto.Marshal(container)
	if err != nil {
		fmt.Printf("Failed to serialize container: %s", err)
		os.Exit(1)
	}

	return containerBytes
}

func DecodeContainer(data []byte) *pb.Container {
	container := &pb.Container{}
	err := proto.Unmarshal(data, container)
	if err != nil {
		fmt.Printf("Failed to decode container: %s", err)
		os.Exit(1)
	}

	return container
}
