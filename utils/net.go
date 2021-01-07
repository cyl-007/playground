package utils

import (
	"apodemakeles/playground/log"
	"net"
)

var ipString = "unknown"

func init() {
	resolveIp()
}

func GetIp() string {
	return ipString
}

func resolveIp() {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Logger.Errorf("fail to get ip address, error=%v", err)
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ipString = ipnet.IP.String()
			}
		}
	}
}
