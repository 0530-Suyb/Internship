package main

import "fmt"

// 抽象工厂模式
// 需要先定义产品族、产品等级结构
// 抽象工厂模式是一个创建型设计模式，它提供了一个接口，用于创建一系列相关或相互依赖的对象，而无需指定它们的具体类。
// 这种模式通常用于需要创建多个相关对象的场景，例如在图形用户界面（GUI）库中创建不同平台的按钮、文本框等组件。
// 通过使用抽象工厂模式，可以确保创建的对象之间具有一致性和兼容性，从而提高代码的可维护性和可扩展性。
// 抽象工厂模式的主要优点是可以在不修改现有代码的情况下轻松添加新产品族，从而提高了系统的灵活性和可扩展性。
// 但是，它也有一些缺点，例如增加了系统的复杂性和难以理解的代码结构。

type ProductA interface {
	Use()
}

type ProductB interface {
	Use()
}

type Factory interface {
	CreateProductA() ProductA
	CreateProductB() ProductB
}

type ChinaProductA struct{}

func (p *ChinaProductA) Use() {
	fmt.Println("使用中国产品A")
}

type ChinaProductB struct{}

func (p *ChinaProductB) Use() {
	fmt.Println("使用中国产品B")
}

type ChinaFactory struct{}

func (f *ChinaFactory) CreateProductA() ProductA {
	return &ChinaProductA{}
}
func (f *ChinaFactory) CreateProductB() ProductB {
	return &ChinaProductB{}
}

type USProductA struct{}

func (p *USProductA) Use() {
	fmt.Println("使用美国产品A")
}

type USProductB struct{}

func (p *USProductB) Use() {
	fmt.Println("使用美国产品B")
}

type USFactory struct{}

func (f *USFactory) CreateProductA() ProductA {
	return &USProductA{}
}
func (f *USFactory) CreateProductB() ProductB {
	return &USProductB{}
}

type abstractFactory struct {
	Factory Factory
}

func (af *abstractFactory) CreateProductA() ProductA {
	return af.Factory.CreateProductA()
}

func (af *abstractFactory) CreateProductB() ProductB {
	return af.Factory.CreateProductB()
}

func NewAbstractFactory(factory Factory) *abstractFactory {
	return &abstractFactory{Factory: factory}
}

func main() {
	chinaFactory := NewAbstractFactory(&ChinaFactory{})
	productA := chinaFactory.CreateProductA()
	productB := chinaFactory.CreateProductB()
	productA.Use()
	productB.Use()

	usFactory := NewAbstractFactory(&USFactory{})
	productA = usFactory.CreateProductA()
	productB = usFactory.CreateProductB()
	productA.Use()
	productB.Use()
}
