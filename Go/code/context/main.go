package main

import (
	"context"
	"fmt"
	"time"
)

type UserInfo struct {
	Name string
}

func main() {
	// var ctx1 context.Context = context.Background()
	// ctx1 = context.WithValue(ctx1, "name", UserInfo{Name: "su"})
	// GetUser(ctx1)

	// ctx2, _ := context.WithTimeout(context.Background(), 2*time.Second)
	// go GetIp(ctx2)
	// time.Sleep(5 * time.Second)

	// WithDeadline内部将deadline-now()得到间隔，然后调用time.AfterFunc(间隔, func(){//取消context})
	// ctx3, _ := context.WithDeadline(context.Background(), time.Now().Add(2*time.Second))
	// go GetIp(ctx3)
	// time.Sleep(5*time.Second)

	ctx4, cancel := context.WithCancel(context.Background())
	go GetIp(ctx4)
	select {
	case <-time.After(2 * time.Second):
		fmt.Println("cancel!")
		cancel()
	}
	time.Sleep(3*time.Second)
}

func GetUser(ctx context.Context) {
	fmt.Println(ctx.Value("name").(UserInfo).Name)
}

func GetIp(ctx context.Context) {
	startT := time.Now()
	workCh := time.After(3000 * time.Millisecond)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("timeout!")
			fmt.Printf("run %v\n", time.Since(startT))
			return
		case <-workCh:
			ip := "127.0.0.1"
			fmt.Printf("get ip %s\n", ip)
			fmt.Printf("run %v\n", time.Since(startT))
			return
		default:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("...")
		}
	}
}
