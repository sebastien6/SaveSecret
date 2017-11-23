package ssecret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
)

// rsaEncryptOAEP encrypt a plaintext message with the passed public key
func rsaEncryptOAEP(secretMessage []byte, pub *rsa.PublicKey) ([]byte, error) {
	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng := rand.Reader

	ciphertext, err := rsa.EncryptOAEP(sha256.New(), rng, pub, secretMessage, label())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return nil, err
	}

	return ciphertext, nil
}
