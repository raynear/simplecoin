package SimpleCoin

import (
	"fmt"
	"net"
)

// Node :
type Node struct {
	Address string `json:"node"`
}

// NodeList :
var NodeList []Node

// GetMyIP Get local IP
func GetMyIP() string {
	netInterfaceAddresses, err := net.InterfaceAddrs()

	if err != nil {
		return ""
	}

	for _, netInterfaceAddress := range netInterfaceAddresses {
		networkIP, ok := netInterfaceAddress.(*net.IPNet)

		if ok && !networkIP.IP.IsLoopback() && networkIP.IP.To4() != nil {
			ip := networkIP.IP.String()
			fmt.Println("Resolved Host IP: " + ip)
			return ip
		}
	}
	return ""
}
