package main

import (
	"bufio"
	"fmt"
	"net"
)

func proc(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		buf := [128]byte{}
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf[:n]))
		conn.Write(buf[:n])
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go proc(conn)
	}
}
