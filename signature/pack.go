package signature

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

func PackAndSignMessage(message *Message, key *AsymmetricKey) (signature []byte, encodedContainer []byte) {
	// Before signing, we need to hash our message
	// The hash is what we actually sign
	hash := sha256.New()

	messageBytes := EncodeMessage(message)
	_, err := hash.Write(messageBytes)
	if err != nil {
		fmt.Printf("Failed hashing message: %s", err)
		os.Exit(1)
	}

	hashSum := hash.Sum(nil)

	signature, err = rsa.SignPKCS1v15(rand.Reader, key.PrivateKey, crypto.SHA256, hashSum)

	if err != nil {
		fmt.Printf("Failed to sign message: %s", err)
		os.Exit(1)
	}

	container := &Container{Message: message, Signature: signature}
	encodedContainer = EncodeContainer(container)

	return
}

func IsAuthentic(container *Container, key *AsymmetricKey) bool {
	hash := sha256.New()

	messageBytes := EncodeMessage(container.Message)

	_, err := hash.Write(messageBytes)
	if err != nil {
		fmt.Printf("Failed hashing message: %s", err)
		os.Exit(1)
	}

	hashSum := hash.Sum(nil)

	err = rsa.VerifyPKCS1v15(key.PublicKey, crypto.SHA256, hashSum, container.Signature)

	return err == nil
}
