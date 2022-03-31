package tf

import (
	"net"
)

func NetIps() []string {
	ips := []string{}
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ips
	}
	for _, addr := range addrs {
		ipNet, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		ips = append(ips, ipNet.IP.String())
	}
	return ips
}

func NetMacs() []string {
	macs := []string{}
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return macs
	}
	for _, netInterface := range netInterfaces {
		mac := netInterface.HardwareAddr.String()
		if len(mac) == 0 {
			continue
		}
		macs = append(macs, mac)
	}
	return macs
}
