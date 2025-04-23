package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now().UnixMicro()
	fmt.Println(t)
}
