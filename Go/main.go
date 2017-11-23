/*
ssecret:
-----------

C-Shared library, for Linux and Windows, to encrypt/decrypt a "Secret" for a specific script/code that will need
the decrypt secret to automatically provision a password/secret field in a script/code function call,
compatible with all development languages able to call up a C dynamic library (GO, C/C++, Python, Ruby, Java, Node, ..)

Each time a "Secret" is encrypted for a designated script/code, a new pair of RSA key (size 4096) is created with
a generated password based on a mixed of file checksum and host information. For the specified script/code to receive the
decrypted password from the GetSecret function, the script/code need to pass a couple of validations:
	-> the script/code run on the same host as it was encrypted for
	-> the script/code run at the same location than when the secret was encrypted for it
	-> no modification were made to any of the files:
			* the script/code itself calling the function
			* the RSA key
			* the language specific library (hieroglyph.py for python)
			* the c-shared library itself
Any check failure end up with a decrypt reject from the library. The decrypted "Secret" will be validated against
the bcrypt hash before return by the function.

USAGE:
------

To use the library, prepare your script/code or code with the function call GetSecret into it. Then call the SaveSecret
from another script/code or core or REPL.
Python example using the library hieroglyph.py:		python -c "import hieroglyph; hieroglyph.SaveSecret()"

RESTRICTION:
------------

Any change required to the script/code will required a new call to the SaveSecret function to re-sign the new
script/code. The call will truncate and replace the existing .key, .sig, .hash, and .sec file.

FUNCTIONS:
----------

The library export two functions:
	func SaveSecret(libpath, runpath string)
	func GetSecret(libpath, runpath, scriptfile string) *C.char

SaveSecret:
-----------

SaveSecret will ask the user for the name of the script/code that will use the "Secret", and the secret messge. The
script/code will be "Signed" to prevent any modification after the secret encryption.
SaveSecret is taking two parameters:
	libpath		(string)	-> 	Location of the c-shared library and language specific library (both have to be in
								the same folder.
								example: libpath=/usr/lib/python2.7 with files hieroglyph.pyc and c_hieroglyph.so
										 present in that location
	runpath	(string)	->  location of the script/code to "Signe" and encrypt "Secret" for
								example: runpath=/home/user/script/ where myscript.py is located.

After the function call, four file will be present with the same name as the targeted script/code in the runpath:
	<script/code_name>.key	-> 	RSA key in PEM(p12) format
	<script/code_name>.sec	->  Encrypted "Secret"
	<script/code_name>.sig	->  script/code signature
	<script/code_name>.hash  ->  BCrypt hash

GetSecret:
----------

GetSecret will return the "Secret" to the script/code it was encrypted for in a base64 format.
GetSecret is taking three parameters:
	libpath		(string)	-> 	Location of the c-shared library and language specific library (both have to be in
								the same folder.
								example: libpath=/usr/lib/python2.7 with files hieroglyph.pyc and c_hieroglyph.so
										 present in that location
	runpath	(string)	->  location of the script/code to "Signe" and encrypt "Secret" for
								example: runpath=/home/user/script/ where myscript/code.py is located.
	scriptfile	(string)	->	name of the script/code itself with its extension
								example: myscript/code.py

*/
package main

import (
	"C"
	"SaveSecret/Go/ssecret"
	"os"
)

func main() {}

//export SaveSecret
func SaveSecret(libpath, runpath string) {
	ssecret.SaveSecret(libpath, runpath)
}

//export GetSecret
func GetSecret(libpath, runpath, scriptfile string) *C.char {
	s := ssecret.GetSecret(libpath, runpath, scriptfile)
	if *s != "" {
		// Return b64encoded value of plaintext
		return C.CString(*s)
	} else {
		os.Exit(1)
		return nil
	}
}
