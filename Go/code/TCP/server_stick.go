package main

import (
	"bufio"
	"fmt"
	"io"
	"net"

	"tcp/proto"
)

func proc(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("recv data: ", msg)
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
