package ssecret

import (
	"crypto/sha256"
	"encoding/hex"
	"io/ioutil"
	"os"
)

func (a BuildBlock) createFileChecksum() {
	// open the script file protected
	f1 := a.openKeyFile()
	hn, _ := os.Hostname()

	s := *f1 + "SOME_CHARECTERS" + genNum(hn) + "SOME_CHARACTERS"

	// Create a new hash.Hash computing the SHA256 checksum
	h := sha256.New()

	// send the aggregate string to hash.Hash sha256
	h.Write([]byte(s))

	// return the checksum of the aggregate string
	hsum := hex.EncodeToString(h.Sum(nil))

	// save the sha256 checksum into a file
	ioutil.WriteFile(a.runpath+a.Filename+".sig", []byte(hsum), 0644)
}

func (a BuildBlock) validateFileChecksum() (r bool) {
	// open the key file
	f1 := a.openKeyFile()
	f2 := a.openSumFile()
	hn, _ := os.Hostname()

	s := *f1 + "SOME_CHARECTERS" + genNum(hn) + "SOME_CHARACTERS"

	// Create a new hash.Hash computing the SHA256 checksum
	h := sha256.New()

	// send the aggregate string to hash.Hash sha256
	h.Write([]byte(s))

	// return the checksum of the aggregate string
	hsum := hex.EncodeToString(h.Sum(nil))

	// Compare new checksum with saved one to validate the protected script file and private key were not changed since
	// the time the secret was saved in the encrypted file. Return true or false based on the result of the string comparison
	if *f2 == hsum {
		return true
	} else {
		return false
	}
}
