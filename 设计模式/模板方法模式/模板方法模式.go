package main

import "fmt"

// 模板方法模式
// 模板方法模式定义了一个操作中的算法的骨架，而将一些步骤延迟到子类中。模板方法使得子类可以在不改变算法结构的情况下重新定义该算法的某些特定步骤。
// 这种模式的主要优点是代码复用和代码的可读性。

type Methods interface {
	Step1()
	Step2()
	Step3()
}

type template struct {
	m Methods
}

func (t *template) Do() {
	t.m.Step1()
	t.m.Step2()
	t.m.Step3()
}

type MakeTee struct {
}

func (m *MakeTee) Step1() {
	fmt.Println("boiling water")
}

func (m *MakeTee) Step2() {
	fmt.Println("puting tee into cup")
}

func (m *MakeTee) Step3() {
	fmt.Println("brewing tee")
}

type MakeCoffee struct {
}

func (m *MakeCoffee) Step1() {
	fmt.Println("boiling water")
}

func (m *MakeCoffee) Step2() {
	fmt.Println("puting coffee into cup")
}

func (m *MakeCoffee) Step3() {
	fmt.Println("brewing coffee")
}

func main() {
	mt := &MakeTee{}
	mC := &MakeCoffee{}

	t1 := &template{m: mt}
	t2 := &template{m: mC}

	t1.Do()
	t2.Do()
}
