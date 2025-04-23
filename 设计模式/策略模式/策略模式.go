package main

import "fmt"

// 策略模式
// 策略模式定义了一系列算法,把它们一个个封装起来,并且使它们可以互相替换。
// 策略模式让算法的变化独立于使用算法的客户。

type Strategy interface {
	DoOperation(int, int) int
}

type AddOperation struct{}

func (a *AddOperation) DoOperation(num1, num2 int) int {
	return num1 + num2
}

type SubtractOperation struct{}

func (s *SubtractOperation) DoOperation(num1, num2 int) int {
	return num1 - num2
}

type MultiplyOperation struct{}

func (m *MultiplyOperation) DoOperation(num1, num2 int) int {
	return num1 * num2
}

type Context struct {
	strategy Strategy
}

func (c *Context) SetStrategy(strategy Strategy) {
	c.strategy = strategy
}

func (c *Context) ExecuteStrategy(num1, num2 int) int {
	return c.strategy.DoOperation(num1, num2)
}

func main() {
	num1 := 10
	num2 := 5

	context := &Context{}

	add := &AddOperation{}
	context.SetStrategy(add)
	fmt.Printf("10 + 5 = %d\n", context.ExecuteStrategy(num1, num2))

	subtract := &SubtractOperation{}
	context.SetStrategy(subtract)
	fmt.Printf("10 - 5 = %d\n", context.ExecuteStrategy(num1, num2))

	multiply := &MultiplyOperation{}
	context.SetStrategy(multiply)
	fmt.Printf("10 * 5 = %d\n", context.ExecuteStrategy(num1, num2))
}
