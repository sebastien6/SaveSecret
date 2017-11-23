package ssecret

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/shirou/gopsutil/cpu"
	"net"
	"runtime"
	"strconv"
)

// pemSecret is returning a sha256 checksum from a complex string based on files and device information
func (a BuildBlock) pemSecret() *[]byte {
	// Collect device information
	c, _ := cpu.Info()
	n, _ := net.Interfaces()

	// extract hostname without fqdn
	hostname := getHostname()

	// Collect CPU information
	family := c[0].Family
    vendid := c[0].VendorID

	// Collect network info
	iname := n[1].Name

	// open the script file protected
	f1 := a.openScFile()

	// open the c-shared library
	f3 := a.openCLib()

	// open python library
	f4 := a.openPLib()

	// create a complex string
	complexstring := "SOME_CHARACTERS" + string(*f1) + "SOME_CHARACTERS" + vendid + string(*f3) +  "SOME_CHARACTERS" + string(*f4) + "SOME_CHARACTERS" + hostname + iname + "SOME_CHARACTERS" + family + "SOME_CHARACTERS"

	// open a new sha256 hash.Hash
	sha := sha256.New()

	// send complex string to hash.Hash sha256
	sha.Write([]byte(complexstring))

	// extract hash checkcum of complex string
	hsum := hex.EncodeToString(sha.Sum(nil))

	// convert checksum to []byte
	cstringsha := []byte(hsum)

	// return the checksum
	return &cstringsha
}
