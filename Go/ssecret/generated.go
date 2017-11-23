package ssecret

import (
	"github.com/shirou/gopsutil/cpu"
	"math"
	"net"
	"os"
	"strconv"
)

// label generate a label for the rsa encrypt and decrypt function based on some collected system information
func label() []byte {
	ip, _ := extIP()
	hostname, _ := os.Hostname()

	// Collect device information
	c, _ := cpu.Info()
	modelname := c[0].ModelName

	label := []byte(modelname + "SOME_CHARACTERS" + ip + "SOME_CHARACTERS" + hostname)
	return label
}

// genNum generate a string based on calculated value from a received string, host CPU and network specific parameters
func genNum(m string) string {
	// convert string to []byte
	mySlice := []byte(m)

	// convert []byte to int
	d := len(mySlice)

	// extract CPU(s) info
	c, _ := cpu.Info()

	// extract network interface info
	n, _ := net.Interfaces()

	// calculate Log of first network interface MTU size
	v := int(math.Sqrt(float64(n[1].MTU)))

	// calculate Sqrt of number of CPU cores
	s := int(math.Sqrt(float64(len(c))))

	// multiply string convert of int of bytes values with sqrt of number of cpu cores and divide by Log of first
	// network interface MTU size
	r := float64(d * s * v)

	// power by 5
	r = math.Pow(r, float64(5))

	// convert float64 to int
	tot := int(r)

	// return result as a string
	res := strconv.Itoa(d) + strconv.Itoa(s) + strconv.Itoa(v) + strconv.Itoa(tot)
	return res
}
