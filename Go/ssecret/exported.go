package ssecret

import (
	"fmt"
	"regexp"
	"runtime"
)

// fileExtenson receive a filename and split the name of the file form its extension
func FileExtension(scriptfile string) (filename, extension *string) {
	re, _ := regexp.Compile(`([A-Za-z0-9_-]+)(.*)`)
	matches := re.FindAllStringSubmatch(scriptfile, -1)
	filename = &matches[0][1]
	extension = &matches[0][2]
	return
}

// SaveSecret
func SaveSecret(libpath, runpath string) {
	fmt.Print("Enter file name: ")
	var file string
	fmt.Scanln(&file)

	filename, extension := FileExtension(file)
	if runtime.GOOS == "windows" {
		runpath = runpath + `\`
	}
	a := BuildBlock{runpath, libpath, *filename, *extension, 4096}

	a.encryptAndSign()
}

// GetSecret
func GetSecret(libpath, runpath, scriptfile string) *string {
	if runtime.GOOS == "windows" {
		runpath = runpath + `\`
	}
	filename, extension := FileExtension(scriptfile)
	a := BuildBlock{runpath, libpath, *filename, *extension, 4096}

	sEnc := a.signCheckandDecrypt()

	return sEnc
}
