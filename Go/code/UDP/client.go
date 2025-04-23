package main

import (
	"fmt"
	"net"
)

func main() {
	socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 8000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer socket.Close()
	sendData := []byte("hello, world")
	_, err = socket.Write(sendData)
	if err != nil {
		fmt.Println(err)
		return
	}
	data := make([]byte, 4096)
	n, rAddr, err := socket.ReadFromUDP(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("recv: %s, addr: %v, count: %d", data[:n], rAddr, n)
}
