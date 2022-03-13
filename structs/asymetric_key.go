package structs

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
	"path/filepath"
	"signatures-playground/utilities"
)

type AsymmetricKey struct {
	PrivateKey *rsa.PrivateKey
	PublicKey  *rsa.PublicKey
}

func (aKey *AsymmetricKey) EncodedToPem() ([]byte, []byte) {
	return aKey.PrivateKeyEncodedToPem(), aKey.PublicKeyEncodedToPem()
}

func (aKey *AsymmetricKey) PrivateKeyEncodedToPem() []byte {
	privateKeyBytes := x509.MarshalPKCS1PrivateKey(aKey.PrivateKey)

	privateKeyBlock := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: privateKeyBytes,
	}

	return pem.EncodeToMemory(privateKeyBlock)
}

func (aKey *AsymmetricKey) PublicKeyEncodedToPem() []byte {
	publicKeyBytes, err := x509.MarshalPKIXPublicKey(aKey.PublicKey)
	if err != nil {
		fmt.Printf("Failed to dump public key bytes: %s", err)
		os.Exit(1)
	}

	publicKeyBlock := &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: publicKeyBytes,
	}

	return pem.EncodeToMemory(publicKeyBlock)
}

func (aKey *AsymmetricKey) ExportPrivateKeyToPem(path string) {
	utilities.WriteFile(filepath.Join(path, "private_key.pem"), aKey.PrivateKeyEncodedToPem())
}

func (aKey *AsymmetricKey) ExportPublicKeyToPem(path string) {
	utilities.WriteFile(filepath.Join(path, "public_key.pem"), aKey.PublicKeyEncodedToPem())
}

func (aKey *AsymmetricKey) ImportPrivateKey(filePath string) {
	keyPemBytes := readPemFile(filePath)
	key, err := x509.ParsePKCS1PrivateKey(keyPemBytes)
	if err != nil {
		fmt.Printf("Failed to import private key from %s: %s", filePath, err)
		os.Exit(1)
	}

	aKey.PrivateKey = key
}

func (aKey *AsymmetricKey) ImportPublicKey(filePath string) {
	keyPemBytes := readPemFile(filePath)
	pub, err := x509.ParsePKIXPublicKey(keyPemBytes)
	if err != nil {
		fmt.Printf("Failed to import public key from %s: %s", filePath, err)
		os.Exit(1)
	}

	switch pub := pub.(type) {
	case *rsa.PublicKey:
		aKey.PublicKey = pub
	default:
		fmt.Printf("Failed to import public key from %s: Wrong type", filePath)
		os.Exit(1)
	}
}

func NewAsymmetricKey() *AsymmetricKey {
	privatekey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Printf("Failed to create a private key: %s", err)
		os.Exit(1)
	}

	publickey := &privatekey.PublicKey

	return &AsymmetricKey{privatekey, publickey}
}

func readPemFile(filePath string) []byte {
	block, _ := pem.Decode(utilities.ReadFile(filePath))
	if block == nil {
		fmt.Printf("Failed to decode pem: %s", filePath)
		os.Exit(1)
	}

	return block.Bytes
}
