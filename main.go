package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println(IpsBetween("10.0.0.0", "10.0.0.50"))
}

func IpsBetween(start, end string) int {
	return int(ipToInt(net.ParseIP(end).To4()) - ipToInt(net.ParseIP(start).To4()))
}

func ipToInt(ip net.IP) uint32 {
	return uint32(ip[0])<<24 | uint32(ip[1])<<16 | uint32(ip[2])<<8 | uint32(ip[3])
}
