package main

import "fmt"

// 访问者模式
// 表示一个作用于某对象结构中的各元素的操作，它使你可以在不改变各元素的类的前提下定义作用于这些元素的新操作。

type Visitor interface {
	Visit(elem Element)
}

type ComputerVisitor struct {
}

func (cv *ComputerVisitor) Visit(elem Element) {
	switch elem.(type) {
	case *Computer:
		fmt.Println("visiting computer")
	case *Keyboard:
		fmt.Println("visiting keyboard")
	case *Monitor:
		fmt.Println("visiting monitor")
	case *Mouse:
		fmt.Println("visiting mouse")
	}
}

type Element interface {
	Accept(Visitor)
}

type Computer struct {
	parts []Element
}

func (c *Computer) Accept(v Visitor) {
	for _, p := range c.parts {
		v.Visit(p)
	}
	v.Visit(c)
}

func NewComputer() *Computer {
	return &Computer{parts: []Element{new(Mouse), new(Monitor), new(Keyboard)}}
}

type Mouse struct {
}

func (m *Mouse) Accept(v Visitor) {
	v.Visit(m)
}

type Monitor struct {
}

func (m *Monitor) Accept(v Visitor) {
	v.Visit(m)
}

type Keyboard struct {
}

func (k *Keyboard) Accept(v Visitor) {
	v.Visit(k)
}

func main() {
	computer := NewComputer()
	visitor := &ComputerVisitor{}

	computer.Accept(visitor)
}
