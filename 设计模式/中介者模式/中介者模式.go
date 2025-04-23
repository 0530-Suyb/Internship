package main

import "fmt"

// 中介者模式
// 中介者模式是用来降低多个对象和类之间的通信复杂性。
// 中介者模式通过引入一个中介者对象来封装一系列的对象交互。
// 中介者使得对象不需要显示地引用彼此，从而使得松耦合的系统更容易实现。
// 中介者可以是一个类，也可以是一个接口。

type Mediator interface {
	Notify(sender string, event string)
}

type ConcreteMediator struct {
	Colleague1 *ConcreteColleague1
	Colleague2 *ConcreteColleague2
}

type ConcreteColleague1 struct {
	Mediator Mediator
	Name     string
}

type ConcreteColleague2 struct {
	Mediator Mediator
	Name     string
}

func (c *ConcreteColleague1) Send(event string) {
	c.Mediator.Notify(c.Name, event)
}

func (c *ConcreteColleague2) Send(event string) {
	c.Mediator.Notify(c.Name, event)
}

func (m *ConcreteMediator) Notify(sender string, event string) {
	if sender == m.Colleague1.Name {
		fmt.Printf("%s received event from %s: %s\n", m.Colleague2.Name, sender, event)
	} else if sender == m.Colleague2.Name {
		fmt.Printf("%s received event from %s: %s\n", m.Colleague1.Name, sender, event)
	}
}

func main() {
	mediator := &ConcreteMediator{}
	colleague1 := &ConcreteColleague1{Name: "Colleague1", Mediator: mediator}
	colleague2 := &ConcreteColleague2{Name: "Colleague2", Mediator: mediator}
	mediator.Colleague1 = colleague1
	mediator.Colleague2 = colleague2

	colleague1.Send("Hello from Colleague1")
	colleague2.Send("Hello from Colleague2")
}
