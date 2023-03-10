package service

import "net"

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ":("
	}
	for i := range addrs {
		ipNet, isIpNet := addrs[i].(*net.IPNet)
		if isIpNet && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				return ipNet.IP.String()
			}
		}
	}
	return ":("
}
