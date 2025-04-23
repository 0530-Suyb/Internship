package main

import (
	"fmt"
	"sync"
	"time"
)

// ### 避免关闭后发送 ###

// 单发单收，发者关
func chClose() {
	ch := make(chan int)
	after := time.After(10 * time.Millisecond)
	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		for {
			select {
			case <-after:
				wg.Done()
				close(ch)
				return
			default:
				ch <- 1
			}
		}
	}()

	go func() {
		defer wg.Done()
		for v := range ch {
			fmt.Println(v)
		}
		return
	}()

	wg.Wait()
}

// 多发单收，收者关
func chClose2() {
	numCh := make(chan int)
	stopCh := make(chan struct{})
	after := time.After(10 * time.Millisecond)
	wg := sync.WaitGroup{}
	wg.Add(1)

	sender := 3
	for i := 0; i < sender; i++ {
		go func() {
			for {
				select {
				case <-stopCh:
					fmt.Println("退出通道")
					return
				case numCh <- 1:

				}
			}
		}()
	}

	go func() {
		for {
			select {
			case v := <-numCh:
				fmt.Printf("v: %d\n", v)
			case <-after:
				close(stopCh)
				time.Sleep(1 * time.Second)
				wg.Done()
				return
			}
		}
	}()

	wg.Wait()
}

// 多发多收，引入中间者，关时通知中间者，中间者收到信号后关闭stopCh（中间者也结束），发者和收者检测到stopCh关闭了就退出

// ### 避免重复关闭 ###

// sync.once
type myCh1 struct {
	ch   chan interface{}
	once sync.Once
}

func NewMyCh1() *myCh1 {
	return &myCh1{ch: make(chan interface{})}
}

func (ch *myCh1) CloseCh() {
	ch.once.Do(func() {
		close(ch.ch)
	})
}

// sync.Mutex
type myCh2 struct {
	ch     chan interface{}
	closed bool
	mu     sync.Mutex
}

func NewMyCh2() *myCh2 {
	return &myCh2{ch: make(chan interface{})}
}

func (ch *myCh2) CloseCh() {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	if !ch.closed {
		close(ch.ch)
		ch.closed = true
	}
}

func (ch *myCh2) IsClosed() bool {
	ch.mu.Lock()
	defer ch.mu.Unlock()
	return ch.closed
}

func stringType() {
	var str string = "hello"
	fmt.Printf("%T\n", str[0]) // uint8
	for _, v := range str {
		fmt.Printf("%T\n", v) // int32
	}
}

func slice_nil_spare() {
	var s []int // nil切片
	fmt.Println("var s []int, nil?", s == nil) // true
	s1 := new([]int) // 空切片
	fmt.Println("s1 := new([]int), nil?", s1 == nil) // false
	s11 := *new([]int) // nil切片
	fmt.Println("s11 := *new([]int), nil?", s11 == nil) // true
	s2 := make([]int, 0) // 空切片
	fmt.Println("s2 := make([]int, 0), nil?", s2 == nil) // false
	s3 := []int{} // 空切片
	fmt.Println("s3 := []int{}, nil?", s3 == nil) // false
}


func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	ch := make(chan int)
	go func() {
		ch <- 1
		close(ch)
		wg.Done()
	}()
	i, ok := <-ch
	fmt.Println(i, ok)
	
	wg.Wait()
	i, ok = <-ch
	fmt.Println(i, ok)
}
