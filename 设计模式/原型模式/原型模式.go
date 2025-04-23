package main

import (
	"fmt"
	"time"
)

// 原型模式
// 定义一个原型接口，包含一个克隆方法
type Prototype interface {
	GetName() string
	GetAge() int
	SetName(name string)
	SetAge(age int)
	// 克隆方法，返回一个新的原型对象
	Clone() Prototype
}

// 具体的原型类，实现克隆方法
type ConcretePrototype struct {
	Name string
	Age  int
}

func NewConcretePrototype(name string, age int) *ConcretePrototype {
	// 直接New一个原型对象需要做大量的初始化工作
	// 这里我们假设有一个复杂的初始化过程
	// 例如从数据库中加载数据、读取配置文件等
	// 这里我们sleep后直接返回一个新的实例
	// 以简化示例
	time.Sleep(1 * time.Second) // 模拟复杂的初始化过程
	return &ConcretePrototype{
		Name: name,
		Age:  age,
	}
}

func (p *ConcretePrototype) GetName() string {
	return p.Name
}

func (p *ConcretePrototype) GetAge() int {
	return p.Age
}

func (p *ConcretePrototype) SetName(name string) {
	p.Name = name
}

func (p *ConcretePrototype) SetAge(age int) {
	p.Age = age
}

func (p *ConcretePrototype) Clone() Prototype {
	// 创建一个新的实例，并复制当前对象的属性
	return &ConcretePrototype{
		Name: p.Name,
		Age:  p.Age,
	}
}

func main() {
	// 创建一个原型对象
	t1 := time.Now()
	prototype := NewConcretePrototype("John", 30)
	t2 := time.Now()
	fmt.Println("Time taken to create prototype:", t2.Sub(t1))

	// 克隆原型对象
	t1 = time.Now()
	clonedPrototype := prototype.Clone()
	t2 = time.Now()
	fmt.Println("Time taken to clone prototype:", t2.Sub(t1))

	// 修改克隆对象的属性
	clonedPrototype.SetName("Jane")
	clonedPrototype.SetAge(25)

	// 打印原型对象和克隆对象的属性
	fmt.Println("Original Prototype:", prototype.Name, prototype.Age)
	fmt.Println("Cloned Prototype:", clonedPrototype.(*ConcretePrototype).Name, clonedPrototype.(*ConcretePrototype).Age)
}
