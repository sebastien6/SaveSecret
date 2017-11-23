package ssecret

import (
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
)

// open the c-share library
func (a BuildBlock) openCLib() *string {
	// open the c-shared library
	var err error
	var f []byte
	if runtime.GOOS == "windows" {
		f, err = ioutil.ReadFile(a.libpath + "c_hieroglyph.dll")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open program c-shared library: %s\n", err)
			return nil
		}
	} else if runtime.GOOS == "linux" {
		f, err = ioutil.ReadFile(a.libpath + "c_hieroglyph.so")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open program c-shared library: %s\n", err)
			return nil
		}
	}
	file := string(f)
	return &file
}

// open the python library
func (a BuildBlock) openPLib() *string {
	var err error
	var f []byte
	if a.Extension == ".py" {
		// open the python library
		f, err = ioutil.ReadFile(a.libpath + "hieroglyph.py")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Cannot open python library: %s\n", err)
			return nil
		}
	} else {
		return nil
	}
	file := string(f)
	return &file
}

// open the script file calling the library
func (a BuildBlock) openScFile() *string {
	var err error
	var f []byte
	f, err = ioutil.ReadFile(a.runpath + a.Filename + a.Extension)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open script file: %s\n", err)
		return nil
	}
	file := string(f)
	return &file
}

// open the RSA key file (p12)
func (a BuildBlock) openKeyFile() *string {
	var err error
	var f []byte
	f, err = ioutil.ReadFile(a.runpath + a.Filename + ".key")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open key file: %s\n", err)
		return nil
	}
	file := string(f)
	return &file
}

// open the RSA key checksum file (p12)
func (a BuildBlock) openSumFile() *string {
	var err error
	var f []byte
	f, err = ioutil.ReadFile(a.runpath + a.Filename + ".sig")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Cannot open key checksum file: %s\n", err)
		return nil
	}
	file := string(f)
	return &file
}
