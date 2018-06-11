package SimpleCoin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

// Node :
type Node struct {
	Address string `json:"node"`
}

type Announce struct {
	BlockNumber uint64 `json:"blocknumber"`
	MinedNode   Node   `json:"node"`
}

// AnnouncedBlock
var AnnouncedBlock Announce

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

func AnnounceMakeBlock(blocknumber uint64) {
	for _, aNode := range NodeList {
		MyIP := GetMyIP()
		if MyIP == "" {
			return
		}

		var myNode Node
		myNode = Node{"http://" + MyIP + ":" + Port}
		aAnnounce := Announce{blocknumber, myNode}
		announcebyte, _ := json.Marshal(aAnnounce)
		buff := bytes.NewBuffer(announcebyte)

		resp, err := http.Post(aNode.Address+"/listenmakeblock", "application/json", buff)
		if err != nil {
			fmt.Println(err)
		}
		defer resp.Body.Close()
	}
}
