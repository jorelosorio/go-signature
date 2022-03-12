package main

import (
	"log"
	"signatures-playground/helpers"
	"signatures-playground/structs"
	pb "signatures-playground/structspb"
)

func main() {
	asymmetricKey := structs.NewAsymmetricKey()
	privateKey, publicKey := asymmetricKey.EncodedToPem()
	log.Println(privateKey, publicKey)

	message := &pb.Message{Sender: "ok", Payload: "hola"}
	containerBytes := helpers.SignMessage(message, asymmetricKey)

	log.Println(helpers.EncodeBase64(containerBytes))

	container := helpers.DecodeContainer(containerBytes)

	log.Println("Message:", container.Message)

	isAuthentic := helpers.IsAuthentic(container, asymmetricKey)
	log.Printf("Is the message authentic? %t", isAuthentic)
}
