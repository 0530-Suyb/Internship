package main

import "fmt"

// 工厂方法模式

type Fruit interface {
	Info()
}

type Apple struct {}
func(a *Apple) Info() {
	fmt.Println("Apple")
}

type Banana struct {}
func(b *Banana) Info() {
	fmt.Println("Banana")
}

type Factory interface {
	Create() Fruit
}

type AppleFactory struct {}
func(a *AppleFactory) Create() Fruit {
	return &Apple{}
}

type BananaFactory struct {}
func(b *BananaFactory) Create() Fruit {
	return &Banana{}
}

type FactoryMethod struct {
	Factory Factory
}

func(f *FactoryMethod) Create() Fruit {
	return f.Factory.Create()
}

func NewFactory(Factory Factory) *FactoryMethod {
	return &FactoryMethod{Factory: Factory}
}

func main() {
	appleFactory := &AppleFactory{}
	bananaFactory := &BananaFactory{}

	factoryMethod1 := NewFactory(appleFactory)
	factoryMethod2 := NewFactory(bananaFactory)

	fruit1 := factoryMethod1.Create()
	fruit2 := factoryMethod2.Create()

	fruit1.Info() // Apple
	fruit2.Info() // Banana
}