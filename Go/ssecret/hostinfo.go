package ssecret

import (
	"errors"
	"net"
	"os"
	"regexp"
)

// externalIP return the main ip address of the host
// found this one on the net, really good but don't recall
// on which website, sorry.
func extIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}

	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("No network connection!")
}

func getHostname() string {
	// extract hostname from FQDN
	h, _ := os.Hostname()
	re, _ := regexp.Compile(`([A-Za-z0-9_-]+).*`)
	match := re.FindAllStringSubmatch(h, -1)
	hostname := match[0][1]
	return hostname
}
