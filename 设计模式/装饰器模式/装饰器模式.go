package main

import "fmt"

// 装饰器模式
// 装饰器模式是一种结构型设计模式，它允许在不改变对象的情况下，动态地给对象添加额外的功能。
// 这种模式创建了一个装饰类，用于包装原始类，并在保持类方法签名的情况下提供额外的功能。
// 装饰器模式通常用于需要在运行时添加功能的场景，而不是在编译时就确定功能。

// 例如，在图形用户界面（GUI）中，可以使用装饰器模式来动态地添加滚动条、边框等功能，而不需要修改原始组件的代码。
// 这使得代码更加灵活和可扩展。

// GUI例子

type Component interface {
	Draw() string
}

// ConcreteComponent 是一个具体的组件
type ConcreteComponent struct{}

func (c *ConcreteComponent) Draw() string {
	return "Drawing a basic component"
}

// Decorator 是一个装饰器基类
type Decorator struct {
	component Component
}

func (d *Decorator) Draw() string {
	if d.component != nil {
		return d.component.Draw()
	}
	return ""
}

// BorderDecorator 是一个具体的装饰器，添加边框功能
type BorderDecorator struct {
	Decorator
}

func (b *BorderDecorator) Draw() string {
	return b.Decorator.Draw() + " with a border"
}

// ScrollDecorator 是另一个具体的装饰器，添加滚动条功能
type ScrollDecorator struct {
	Decorator
}

func (s *ScrollDecorator) Draw() string {
	return s.Decorator.Draw() + " with a scrollbar"
}

func main() {
	// 创建一个基本组件
	component := &ConcreteComponent{}

	// 添加边框装饰器
	border := &BorderDecorator{
		Decorator: Decorator{component: component},
	}

	// 添加滚动条装饰器
	scroll := &ScrollDecorator{
		Decorator: Decorator{component: border},
	}

	// 绘制最终组件
	fmt.Println(scroll.Draw())
}


