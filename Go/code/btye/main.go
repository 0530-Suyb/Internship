package main

import (
	"bytes"
	"fmt"
)

func main() {
	buf1 := new(bytes.Buffer)
	buf1.WriteString("hello")
	buf1.Write([]byte(" world"))
	fmt.Println(buf1.String())

	buf2 := bytes.NewBufferString("hello")
	buf2.WriteString(" world")
	fmt.Println(buf2.String())

	buf3 := bytes.NewBuffer([]byte("hello"))
	buf3.WriteString(" world")
	buf3.Truncate(5)
	fmt.Println(buf3.String())

	buf4 := bytes.NewBuffer([]byte("hello"))
	buf4.WriteString(" world")
	buf4.Reset()
	fmt.Println(buf4.String())

	buf5 := bytes.NewBuffer([]byte("hello"))
	buf5.WriteString(" world")
	buf5.Next(5)
	fmt.Println(buf5.String())

	buf6 := bytes.NewBuffer([]byte("hello"))
	buf6.WriteString(" world")
	buf6.Next(5)
	buf6.UnreadByte()
	fmt.Println(buf6.String())
}
