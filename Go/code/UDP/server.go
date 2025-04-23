package main

import (
	"fmt"
	"net"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		data := [1024]byte{}
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Printf("data: %s, addr: %v, count: %d\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
}
