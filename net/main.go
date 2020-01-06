package main

import (
	"net"

	"golang.org/x/net/ipv4"
)

func main() {
	//"udp://@239.255.4.236:1234"
	c1, err := net.ListenPacket("udp4", "239.255.4.236:1234")
	if err != nil {
		return
	}
	defer c1.Close()
	p1 := ipv4.NewPacketConn(c1)
	en0, err := net.InterfaceByName("eth0")
	addr, _ := net.ResolveUDPAddr("udp4", "239.255.4.236:1234")
	if err := p1.JoinGroup(en0, addr); err != nil {
		return
	}
}
