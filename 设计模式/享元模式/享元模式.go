package main

import "fmt"

// 享元模式
// Flyweight Pattern
// 享元模式是一种结构型设计模式，旨在通过共享对象来减少内存使用和提高性能。
// 享元模式的核心思想是将对象的状态分为内部状态和外部状态。
// 内部状态是对象共享的部分，而外部状态是对象独有的部分。
// 享元模式适用于需要大量相似对象的场景，例如文本编辑器中的字符对象、图形编辑器中的形状对象等。

// 图形形状例子

type Circle struct {
	Color  string // 内部状态
	Radius int    // 内部状态
	X      int    // 外部状态
	Y      int    // 外部状态
}

func (c *Circle) Draw() string {
	return fmt.Sprintf("Circle: Color=%s, Radius=%d, X=%d, Y=%d", c.Color, c.Radius, c.X, c.Y)
}

type ShapeFactory struct {
	circleMap map[string]*Circle
}

func NewShapeFactory() *ShapeFactory {
	return &ShapeFactory{
		circleMap: make(map[string]*Circle),
	}
}

func (sf *ShapeFactory) GetCircle(color string) *Circle {
	if circle, exists := sf.circleMap[color]; exists {
		return circle
	}
	circle := &Circle{Color: color}
	sf.circleMap[color] = circle
	return circle
}

func main() {
	factory := NewShapeFactory()

	colors := []string{"Red", "Green", "Blue", "Red", "Green"}
	for i, color := range colors {
		circle := factory.GetCircle(color) // 内部共享了Red和Green的Circle对象
		circle.X = i * 10
		circle.Y = i * 20
		fmt.Println(circle.Draw())
	}
}
