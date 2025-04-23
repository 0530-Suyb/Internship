package main

import "fmt"

// 桥接模式
// 桥接模式是一种结构型设计模式，它通过将抽象部分与实现部分分离，使它们可以独立地变化。
// 这种模式使用了组合和继承的方式来实现桥接。
// 在桥接模式中，抽象类和实现类之间通过一个接口进行连接，从而实现了它们之间的解耦。
// 这种模式的主要优点是可以在不修改抽象类和实现类的情况下，增加新的实现类或抽象类，从而提高了系统的灵活性和可扩展性。
// 桥接模式的应用场景包括：
// 1. 当一个类存在两个或多个维度的变化时，可以使用桥接模式来分离这两个维度，从而减少类的数量。
// 2. 当一个类的实现部分可能会有多个变化时，可以使用桥接模式来将抽象部分与实现部分分离，从而提高系统的灵活性和可扩展性。
// 3. 当一个类的实现部分可能会被多个类共享时，可以使用桥接模式来将实现部分提取到一个独立的类中，从而减少代码的重复性。

// 定义一个抽象类，表示桥接的接口
type Bridge interface {
	Operation() string
}

// 定义一个具体的实现类，实现桥接接口
type ConcreteBridge struct {
	implementation string
}

func (c *ConcreteBridge) Operation() string {
	return fmt.Sprintf("ConcreteBridge: %s", c.implementation)
}

// 定义一个抽象类，表示桥接的抽象类
type Abstraction struct {
	bridge Bridge
}

func (a *Abstraction) SetBridge(bridge Bridge) {
	a.bridge = bridge
}

func (a *Abstraction) Operation() string {
	return fmt.Sprintf("Abstraction: %s", a.bridge.Operation())
}

// 定义一个具体的实现类，继承自抽象类
type RefinedAbstraction struct {
	Abstraction
}

func (r *RefinedAbstraction) Operation() string {
	return fmt.Sprintf("RefinedAbstraction: %s", r.Abstraction.Operation())
}

func main() {
	// 创建一个具体的实现类
	concreteBridge := &ConcreteBridge{implementation: "Concrete Implementation"}

	// 创建一个抽象类，并设置桥接的实现类
	abstraction := &Abstraction{}
	abstraction.SetBridge(concreteBridge)

	// 创建一个具体的实现类，继承自抽象类
	refinedAbstraction := &RefinedAbstraction{Abstraction: *abstraction}

	// 调用操作方法
	fmt.Println(refinedAbstraction.Operation())
}