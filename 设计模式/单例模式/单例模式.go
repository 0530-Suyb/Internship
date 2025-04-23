package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

func (s *Singleton) DoSomething() {
	// Implementation of the method
	fmt.Println("Doing something...")
}

var singletonInstance *Singleton
var once sync.Once

func NewSingleton() *Singleton {
	once.Do(func() {
		singletonInstance = &Singleton{}
	})
	return singletonInstance
}

func main() {
	singleton := NewSingleton()
	singleton.DoSomething()
}
