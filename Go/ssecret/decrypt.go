package ssecret

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// rsaDecryptOAEP receive the content of an encrypted file, decrypt it, and return the plaintext value
func (a BuildBlock) rsaDecryptOAEP(encText []byte) ([]byte, error) {
	pKey, err := ioutil.ReadFile(a.runpath + a.Filename + ".key")
	block, _ := pem.Decode(pKey)
	if block == nil {
		return nil, errors.New("private key error!")
	}

	rsaSecret := a.pemSecret()

	nblock, _ := x509.DecryptPEMBlock(block, *rsaSecret)
	if nblock == nil {
		return nil, errors.New("nBlock decrypt error!")
	}

	priv, err := x509.ParsePKCS1PrivateKey(nblock)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// crypto/rand.Reader is a good source of entropy for blinding the RSA
	// operation.
	rng := rand.Reader

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, priv, encText, label())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return nil, err
	}

	return plaintext, nil
}
