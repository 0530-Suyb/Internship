package main

import (
	"fmt"
)

type Mover interface {
	move()
	bigger()
}

type dog struct{ weight int }

func (d dog) move() {
	fmt.Println("dog move")
}

func (d dog) bigger() {
	d.weight++
}

type cat struct{ weight int }

func (c *cat) move() {
	fmt.Println("cat move")
}

func (c *cat) bigger() {
	c.weight++
}

func main() {
	var m Mover
	m = dog{}
	m.move() // ok
	var n Mover
	n = &dog{} // n Type=*main.dog value=&{}
	// fmt.Printf("n Type=%T value=%v\n", n, n)
	n.move() // ok

	var p Mover
	p = &cat{}
	p.move() // ok
	// p = cat{} // error: cat do not implement method move()
	// 当使用指针接收者实现接口方法时，Go 语言要求必须使用指针类型来调用该方法，因为指针接收者的方法可能会修改接收者指向的对象。如果允许使用值类型来赋值给接口变量，那么就无法保证方法对接收者对象的修改能正确生效。

	// 值接收者、指针接收者
	d1 := dog{}
	d2 := &dog{}
	c1 := cat{}
	c2 := &cat{}
	// all can move
	d1.move()
	d2.move()
	c1.move()
	c2.move()
	// dog can't bigger, cat can bigger
	// 值接收者和指针接收者实现的方法区别：值接收者定义的方法调用时操作的是接收者的一个副本！
	// d1.bigger()
	var q Mover
	q = d1
	q.bigger()
	fmt.Println("d1 w: ", d1.weight) // 0
	q = d2
	q.bigger()                       // (*d2).bigger() 语法糖
	fmt.Println("d2 w: ", d2.weight) // 0
	c1.bigger()                      // (&c1).bigger() 语法糖
	fmt.Println("c1 w: ", c1.weight) // 1
	q = c2
	q.bigger()
	fmt.Println("c2 w: ", c2.weight) // 1

	// 因此
	// 1. 接收者的方法要能修改本身的值时，要用指针接收者。（接收者本质上是第一个参数，编译器会将方法转成普通函数）
	// 2. 由于值接收者的方法改不了本身值，所以go中不允许将值类型的变量赋给接口类型变量（指针接收者实现接口）

	// 空接口没有定义任何方法，因此任何类型都实现了空接口
	var x interface{}
	s := "string"
	x = s
	fmt.Printf("type:%T value:%v\n", x, x)
	i := 100
	x = i
	fmt.Printf("type:%T value:%v\n", x, x)
	// 空接口作为函数的参数
	fun1 := func(a interface{}) {
		fmt.Printf("type:%T value:%v\n", a, a)
	}
	fun1(struct{}{})
	// 空接口不能用短变量声明操作符这样interface{}{}初始化，像slice、map才{}，int、string等都是()
	k := interface{}("hello")
	fmt.Printf("type:%T value:%v\n", k, k)
	// 空接口做map值
	var myInfo = make(map[string]interface{})
	myInfo["name"] = "libai"
	myInfo["age"] = 18
	myInfo["married"] = false
	fmt.Println(myInfo)
	// x.(T) 类型断言（动态类型和动态值）
	x = "string"
	v, ok := x.(string)
	if ok {
		fmt.Println("value:", v)
	} else {
		fmt.Println("类型断言失败")
	}
	switch v := x.(type) {
	case string:
		fmt.Printf("x is a string, value is %v\n", v)
	case int:
		fmt.Printf("x is a int, value is %v\n", v)
	case bool:
		fmt.Printf("x is a bool, value is %v\n", v)
	default:
		fmt.Println("unsupport type!")
	}
}
