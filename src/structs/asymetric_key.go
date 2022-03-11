package structs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"log"
)

type AsymmetricKey struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (aKey *AsymmetricKey) PrivateKeyEncodedToPem() string {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(aKey.PrivateKey)

	privateKeyBlock := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	privateKeyPemBytes := pem.EncodeToMemory(privateKeyBlock)

	return base64.StdEncoding.EncodeToString(privateKeyPemBytes)
}

func (aKey *AsymmetricKey) EncodedToPem() (string, string) {
	return aKey.PrivateKeyEncodedToPem(), aKey.PublicKeyEncodedToPem()
}

func (aKey *AsymmetricKey) PublicKeyEncodedToPem() string {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(aKey.PublicKey)
	if err != nil {
		log.Printf("Failed to dump publickey bytes: %s \n", err)
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	publicKeyPemBytes := pem.EncodeToMemory(publicKeyBlock)

	return base64.StdEncoding.EncodeToString(publicKeyPemBytes)
}

func NewAsymmetricKey() AsymmetricKey {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalln("Failed to create a private key")
	}

	publickey := &privatekey.PublicKey

	return AsymmetricKey{privatekey, publickey}
}
