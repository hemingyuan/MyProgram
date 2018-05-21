package getLocalAddr

import (
	"fmt"
	"net"
)

// GetLocalAddr get local ip address list
func GetLocalAddr() (result []string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		fmt.Println("get local ip failed", err)
		return
	}

	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && ipnet.IP.IsGlobalUnicast() && ipnet.IP.To4() != nil {
			// fmt.Println(ipnet.IP.String())
			result = append(result, ipnet.IP.String())
		}
	}
	return
}
