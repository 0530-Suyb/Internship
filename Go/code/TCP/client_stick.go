package main

import (
	"fmt"
	"net"

	"tcp/proto"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `hello, world!`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		conn.Write(data)
	}
}
