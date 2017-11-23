package ssecret

import (
	"encoding/base64"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"os"
)

// encryptAndSign encrypt a secret with a new pair of RSA keys
func (a BuildBlock) encryptAndSign() {
	valid := false
	// Key pair creation
	fmt.Println("Phase 1/2: Key pair creation\n")
	pkey := a.genKeyPair()

	// test if public key is empty
	if pkey != nil {
		fmt.Println("Phase 1/2: Key pair creation success\n")
		fmt.Println("Phase 2/2: Secret encryption\n")

		// Secret encryption and saving by using returned public key
		valid = a.saveSecret(pkey)

		// valid secret encryption
		if valid {
			fmt.Println("Phase 2/2: Secret encryption success\n")
			a.createFileChecksum()

		} else {
			fmt.Println("Phase 2/2: Secret encryption failure!\n")
			os.Exit(1)
		}
	} else {
		fmt.Println("Phase 1/2: failure!\n")
		os.Exit(1)
	}
}

// signCheckandDecrypt open the .sec file containing the encrypted password, decrypt it and return the
// secret in base64 format
func (a BuildBlock) signCheckandDecrypt() *string {
	v := a.validateFileChecksum()
	if v == false {
		fmt.Println("p12 checksum failure!\n")
		return nil
		os.Exit(1)
	}

	ciphertext, err := ioutil.ReadFile(a.runpath + a.Filename + ".sec")
	if err != nil {
		fmt.Println("ciphertext failure!\n")
		return nil
		os.Exit(1)
	}
	secret, err := a.rsaDecryptOAEP(ciphertext)
	if err != nil {
		fmt.Println("Decrypt failure!\n")
		return nil
		os.Exit(1)
	}
	hash, err := ioutil.ReadFile(a.runpath + a.Filename + ".hash")
	if err != nil {
		fmt.Printf("Cannot open hash file: %v", err)
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(secret))
	sEnc := base64.StdEncoding.EncodeToString([]byte(secret))
	secret = []byte("fake")

	if err != nil {
		fmt.Println("Hash compare failure!\n")
		return nil
		os.Exit(1)
	}

	return &sEnc
}
