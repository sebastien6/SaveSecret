package ssecret

import (
	"crypto/rsa"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
	"io/ioutil"
	"os"
)

// saveSecret ask user for Secret, encrypt it and save the secret and its hash in files
func (a BuildBlock) saveSecret(publicKey *rsa.PublicKey) bool {
	i := true
	var secretMessage []byte
	var fakesecret []byte
	fakesecret = []byte("fake")

	for i {
		fmt.Printf("\nEnter password to encrypt: ")
		pwd1, _ := terminal.ReadPassword(int(os.Stdin.Fd()))

		fmt.Printf("\nRe-Enter password to encrypt: ")
		pwd2, _ := terminal.ReadPassword(int(os.Stdin.Fd()))

		if pwd1 != nil && pwd2 != nil {
			if string(pwd1) == string(pwd2) {
				secretMessage = []byte(pwd1)
				fmt.Println("")
				i = false
			} else {
				fmt.Println("\nSecretmessages do not match, try again!")
			}
		} else {
			fmt.Println("\nYou entered empty values. Do it again!")
		}
	}
	fmt.Println("")

	enc, err := rsaEncryptOAEP(secretMessage, publicKey)
	if err != nil {
		fmt.Printf("Unable to encrypt: %v", err)
	}

	// Generate hash of SecretMessage
	bCryptHash, err := bcrypt.GenerateFromPassword([]byte(secretMessage), 15)
	err = bcrypt.CompareHashAndPassword(bCryptHash, []byte(secretMessage))
	if err != nil {
		fmt.Printf("Hashing issue: %v", err)
		os.Exit(1)
	}

	//clear secretMessage plaintext value
	secretMessage = fakesecret

	var x int = 0
	valid := false

	// Save the encrypted secretMessage in a binary file
	err = ioutil.WriteFile(a.runpath+a.Filename+".sec", enc, 0644)
	if err != nil {
		fmt.Printf("Error saving secret: %v", err)
	} else {
		x = 1
	}
	// Save hash to file
	err = ioutil.WriteFile(a.runpath+a.Filename+".hash", bCryptHash, 0644)
	if err != nil {
		fmt.Printf("Error saving secret hash: %v", err)
	} else {
		x = 2
	}

	if x == 2 {
		valid = true
	}

	return valid
}
