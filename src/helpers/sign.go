package helpers

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"log"
	"signatures-playground/src/structs"
	pb "signatures-playground/src/structspb"
)

func SignMessage(message *pb.Message, key structs.AsymmetricKey) []byte {
	// Before signing, we need to hash our message
	// The hash is what we actually sign
	hash := sha256.New()

	messageBytes := EncodeMessage(message)
	_, err := hash.Write(messageBytes)
	if err != nil {
		log.Fatalln("Failed hashing message", err)
	}

	hashSum := hash.Sum(nil)

	signatureBytes, err := rsa.SignPKCS1v15(rand.Reader, key.PrivateKey, crypto.SHA256, hashSum)

	if err != nil {
		log.Fatalln("Failed to sign message")
	}

	container := &pb.Container{Message: message, Signature: signatureBytes}
	containerBytes := EncodeContainer(container)

	return containerBytes
}

func IsAuthentic(container *pb.Container, key structs.AsymmetricKey) bool {
	hash := sha256.New()

	messageBytes := EncodeMessage(container.Message)
	_, err := hash.Write(messageBytes)
	if err != nil {
		log.Fatalln("Failed hashing message", err)
	}

	hashSum := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(key.PublicKey, crypto.SHA256, hashSum, container.Signature)

	return err == nil
}
