package utils

import (
	"crypto/ed25519"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"os"
)

func ReadEd25519Keypair(privateKeyPath string, publicKeyPath string) (ed25519.PrivateKey, ed25519.PublicKey, error) {
	privateKeyPEM, err := readPEM(privateKeyPath)
	if err != nil {
		return nil, nil, err
	}

	publicKeyPEM, err := readPEM(publicKeyPath)
	if err != nil {
		return nil, nil, err
	}

	privateKey, err := ed25519PrivateKeyFromPEM(privateKeyPEM)
	if err != nil {
		return nil, nil, err
	}

	publicKey, err := ed25519PublicKeyFromPEM(publicKeyPEM)
	if err != nil {
		return nil, nil, err
	}

	return privateKey, publicKey, nil
}

func readPEM(path string) (*pem.Block, error) {
	pemData, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(pemData)
	if block == nil {
		return nil, errors.New("failed to decode PEM block")
	}

	return block, nil
}

func ed25519PrivateKeyFromPEM(block *pem.Block) (ed25519.PrivateKey, error) {
	if block.Type != "PRIVATE KEY" {
		return nil, fmt.Errorf("invalid PEM block type, expected PRIVATE KEY, got %s", block.Type)
	}

	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ed25519PrivateKey, ok := privateKey.(ed25519.PrivateKey)
	if !ok {
		return nil, errors.New("private key is not of type ed25519.PrivateKey")
	}

	return ed25519PrivateKey, nil
}

func ed25519PublicKeyFromPEM(block *pem.Block) (ed25519.PublicKey, error) {
	if block.Type != "PUBLIC KEY" {
		return nil, fmt.Errorf("invalid PEM block type, expected PUBLIC KEY, got %s", block.Type)
	}

	publicKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	ed25519PublicKey, ok := publicKey.(ed25519.PublicKey)
	if !ok {
		return nil, errors.New("public key is not of type ed25519.PublicKey")
	}

	return ed25519PublicKey, nil
}
