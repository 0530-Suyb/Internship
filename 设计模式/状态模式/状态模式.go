package main

import "fmt"

// 状态模式
// 状态模式允许一个对象在其内部状态改变时改变它的行为。对象看起来好像修改了它的类。
// 状态模式的主要优点是可以将与特定状态相关的行为局部化，并且将不同状态的行为分离开来。
// 这样可以避免使用大量的条件语句来处理状态转换。
// 状态模式的缺点是会引入很多状态类，增加了系统的复杂性。
// 但是如果状态的数量不多，或者状态之间的转换关系比较复杂，使用状态模式是非常有用的。

type State interface {
	Handle(context *Context)
}

type ConcreteStateA struct{}

func (s *ConcreteStateA) Handle(context *Context) {
	fmt.Println("ConcreteStateA: Handling request.")
	context.SetState(&ConcreteStateB{})
}

type ConcreteStateB struct{}

func (s *ConcreteStateB) Handle(context *Context) {
	fmt.Println("ConcreteStateB: Handling request.")
	context.SetState(&ConcreteStateA{})
}

type Context struct {
	state State
}

func (c *Context) SetState(state State) {	
	c.state = state
}

func (c *Context) Request() {
	c.state.Handle(c)
}

func NewContext() *Context {
	return &Context{state: &ConcreteStateA{}}
}

func main() {
	context := NewContext()
	for i := 0; i < 5; i++ {
		context.Request()
	}
	fmt.Println("")
}
