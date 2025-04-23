package main

import "fmt"

type Factory interface {
	Create(name string) Fruit
}

type FactoryImpl struct{}
func (f *FactoryImpl) Create(name string) Fruit {
	switch name {
	case "apple":
		return &Apple{}
	case "banana":
		return &Banana{}
	default:
	}
	return nil
}

type Fruit interface {
	Info()
}

type Apple struct{}
func(a *Apple) Info() {
	fmt.Println("Apple")
}

type Banana struct{}
func(b *Banana) Info() {
	fmt.Println("Banana")
}

// 简单工厂模式不符合开闭原则，
// 如果要添加新的水果类型，需要修改工厂类的代码

func main() {
	var f Factory
	f = &FactoryImpl{}

	Apple := f.Create("apple")
	Apple.Info()

	Banana := f.Create("banana")
	Banana.Info()
}