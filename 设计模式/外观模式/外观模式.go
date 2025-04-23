package main

import "fmt"

// 外观模式
// 外观模式（Facade Pattern）是一种结构型设计模式，它为复杂的子系统提供一个统一的接口，使得子系统更易于使用。
// 外观模式隐藏了子系统的复杂性，客户端只需要与外观接口交互，而不需要了解子系统的内部实现细节。
// 案例：
// 智能家居系统，用户通过一个统一的控制面板来启动多种居家模式，如“离家模式”、“回家模式”等，而不需要分别操作每个设备。

// Facade Pattern
// 定义一个外观接口
type Facade interface {
	Operation1()
	Operation2()
	Operation3()
}

// 定义一个外观结构体
type FacadeImpl struct {
	// 这里可以包含多个子系统的实例
	Subsystem1 *Subsystem1
	Subsystem2 *Subsystem2
	Subsystem3 *Subsystem3
}

// 实现外观接口的方法
func (f *FacadeImpl) Operation1() {
	fmt.Println("Facade: Operation1")
	f.Subsystem1.OperationA()
	f.Subsystem2.OperationB()
}
func (f *FacadeImpl) Operation2() {
	fmt.Println("Facade: Operation2")
	f.Subsystem2.OperationB()
	f.Subsystem3.OperationC()
}

func (f *FacadeImpl) Operation3() {
	fmt.Println("Facade: Operation3")
	f.Subsystem1.OperationA()
	f.Subsystem2.OperationB()
	f.Subsystem3.OperationC()
}

// 定义子系统1
type Subsystem1 struct{}

func (s *Subsystem1) OperationA() {
	fmt.Println("Subsystem1: OperationA")
}

// 定义子系统2
type Subsystem2 struct{}

func (s *Subsystem2) OperationB() {
	fmt.Println("Subsystem2: OperationB")
}

// 定义子系统3
type Subsystem3 struct{}

func (s *Subsystem3) OperationC() {
	fmt.Println("Subsystem3: OperationC")
}

// 客户端代码
func main() {
	facade := &FacadeImpl{
		Subsystem1: &Subsystem1{},
		Subsystem2: &Subsystem2{},
		Subsystem3: &Subsystem3{},
	}

	facade.Operation1()
	facade.Operation2()
	facade.Operation3()
}
