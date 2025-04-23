package main

import (
	"fmt"
	"time"
)

var m map[int]int = map[int]int{1: 1}

func setM(v int) {
	// 需要加锁
	m[1] = v
}

func main() {
	go setM(1)
	go setM(2)
	time.Sleep(5 * time.Second)
	fmt.Println(m[1])
}
