package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(IpsBetween("10.0.0.0", "10.0.0.50"))
}

func IpsBetween(start, end string) int {
	startIP := net.ParseIP(start).To4()
	endIP := net.ParseIP(end).To4()

	if startIP == nil || endIP == nil {
		panic(fmt.Sprintf("Invalid IPv4 address: start=%s, end=%s", start, end))
	}

	startNum := ipToInt(startIP)
	endNum := ipToInt(endIP)

	return int(endNum - startNum)
}

func ipToInt(ip net.IP) uint32 {
	result := uint32(ip[0]) << 24
	result |= uint32(ip[1]) << 16
	result |= uint32(ip[2]) << 8
	result |= uint32(ip[3])
	return result
}
