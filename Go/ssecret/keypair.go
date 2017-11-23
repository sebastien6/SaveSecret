package ssecret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

// genKeyPair generate a new RSA key pair
func (a BuildBlock) genKeyPair() *rsa.PublicKey {
	reader := rand.Reader

	key, err := rsa.GenerateKey(reader, a.Keysize)
	if err != nil {
		fmt.Printf("RSA key generation failure: %v", err)
	}

	publicKey := key.PublicKey

	r := a.savePEMKey(key)
	if r {
		return &publicKey
	} else {
		return nil
	}
}

// savePEMKey save in a file the generated key pair in PEM encoded format (p12)
func (a BuildBlock) savePEMKey(key *rsa.PrivateKey) bool {
	// Create a new file to save the RSA key PEM
	outFile, err := os.Create(a.Filename + ".key")
	if err != nil {
		fmt.Printf("PEM file initial creation failure: %v", err)
		return false
	}
	defer outFile.Close()

	// Convert it to pem
	block := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(key),
	}

	// Generate a private key password
	rsaSecret := a.pemSecret()

	// Encrypt the pem
	if string(*rsaSecret) != "" {
		// create the encrypted Secret block
		block, err = x509.EncryptPEMBlock(rand.Reader, block.Type, block.Bytes, *rsaSecret, x509.PEMCipherAES256)
		if err != nil {
			fmt.Printf("Failure encrypting secret: %v", err)
			return false
		}
	} else {
		fmt.Printf("Empty RSA private key password: %v", err)
	}

	// Encode block to PEM file
	err = pem.Encode(outFile, block)
	if err != nil {
		fmt.Printf("RSA private key PEM encoding to file failure: %v", err)
		return false
	}
	return true
}
