package main

import "fmt"

// 观察者模式
// 观察者模式是一种行为设计模式，它定义了一种一对多的依赖关系，让多个观察者对象同时监听某一个主题对象。当这个主题对象的状态发生变化时，它会通知所有观察者对象，使它们能够自动更新自己。
// 观察者模式通常用于实现事件处理系统、发布-订阅系统等场景。它的优点是降低了对象之间的耦合度，提高了系统的灵活性和可扩展性。
// 但是，观察者模式也有一些缺点，比如可能会导致过多的通知和更新操作，从而影响性能；另外，如果观察者对象过多，可能会导致系统的复杂性增加。因此，在使用观察者模式时，需要根据具体情况进行权衡和选择。

type Observer interface {
	Notify(data string)
}

type ConcreteObserver struct {
	name string
}

func (o *ConcreteObserver) Notify(data string) {
	fmt.Printf("Observer %s received data: %s\n", o.name, data)
}

type Subject interface {
	Register(observer Observer)
	Unregister(observer Observer)
	NotifyObservers(data string)
}

type ConcreteSubject struct {
	observers []Observer
}

func (s *ConcreteSubject) Register(observer Observer) {
	s.observers = append(s.observers, observer)
}

func (s *ConcreteSubject) Unregister(observer Observer) {
	for i, o := range s.observers {
		if o == observer {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
}

func (s *ConcreteSubject) NotifyObservers(data string) {
	for _, observer := range s.observers {
		observer.Notify(data)
	}
}

func main() {
	subject := &ConcreteSubject{}
	observer1 := &ConcreteObserver{name: "Observer 1"}
	observer2 := &ConcreteObserver{name: "Observer 2"}

	subject.Register(observer1)
	subject.Register(observer2)

	subject.NotifyObservers("Hello, Observers!")

	subject.Unregister(observer1)

	subject.NotifyObservers("Goodbye, Observers!")
}