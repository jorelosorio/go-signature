package helpers

import (
	"encoding/base64"
	"log"
	pb "signatures-playground/src/structspb"

	"google.golang.org/protobuf/proto"
)

func EncodeBase64(data []byte) string {
	return base64.StdEncoding.Strict().Strict().EncodeToString(data)
}

func DecodeBase64(data string) []byte {
	decoded, err := base64.StdEncoding.Strict().DecodeString(data)
	if err == nil {
		log.Fatalln("Failed to encode base64 data:", err)
	}

	return decoded
}

func EncodeMessage(message *pb.Message) []byte {
	messageBytes, err := proto.Marshal(message)
	if err != nil {
		log.Fatalln("Failed to serialize message:", err)
	}

	return messageBytes
}

func EncodeContainer(container *pb.Container) []byte {
	containerBytes, err := proto.Marshal(container)
	if err != nil {
		log.Fatalln("Failed to serialize container:", err)
	}

	return containerBytes
}

func DecodeContainer(data []byte) *pb.Container {
	container := &pb.Container{}
	err := proto.Unmarshal(data, container)
	if err != nil {
		log.Fatalln("Failed to decode container:", err)
	}

	return container
}
