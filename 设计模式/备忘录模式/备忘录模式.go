package main

import "fmt"

// 备忘录模式
// 备忘录模式是用于保存对象的状态，以便在需要时可以恢复到之前的状态。
// 备忘录模式通常用于实现撤销操作、历史记录等功能。
// 备忘录模式包含三个角色：
// 1. 发起人（Originator）：创建一个备忘录，用于保存当前状态，并可以使用备忘录恢复状态。
// 2. 备忘录（Memento）：用于保存发起人的状态，可以是一个简单的对象，也可以是一个复杂的对象。
// 3. 管理者（Caretaker）：负责保存和管理备忘录对象，不能访问备忘录的内容，只能通过发起人来恢复状态。
// 备忘录模式的优点是可以将对象的状态保存到备忘录中，而不需要暴露对象的内部结构。
// 备忘录模式的缺点是会增加系统的复杂性，因为需要创建和管理备忘录对象。
// 备忘录模式的使用场景包括：
// 1. 实现撤销操作：在文本编辑器中，可以使用备忘录模式来保存文本的历史记录，以便在需要时可以撤销操作。
// 2. 实现历史记录：在游戏中，可以使用备忘录模式来保存游戏的历史记录，以便在需要时可以恢复到之前的状态。
// 3. 实现版本控制：在文件管理系统中，可以使用备忘录模式来保存文件的历史版本，以便在需要时可以恢复到之前的版本。

// 例子：版本控制

type Memento struct {
	// 备忘录对象，保存发起人的状态
	State string
}

type Originator struct {
	// 发起人对象，包含需要保存的状态
	State string
}

type Caretaker struct {
	// 管理者对象，保存备忘录对象
	Memento *Memento
}

func (o *Originator) CreateMemento() *Memento {
	// 创建备忘录对象，保存当前状态
	return &Memento{State: o.State}
}

func (o *Originator) Restore(m *Memento) {
	// 恢复发起人的状态
	o.State = m.State
}

func (o *Originator) SetState(state string) {
	// 设置发起人的状态
	o.State = state
}

func (o *Originator) GetState() string {
	// 获取发起人的状态
	return o.State
}

func main() {
	// 创建发起人对象
	originator := &Originator{}
	originator.SetState("Version 1.0")
	fmt.Println("Current State:", originator.GetState())

	// 创建管理者对象
	caretaker := &Caretaker{}

	// 创建备忘录对象，保存当前状态
	caretaker.Memento = originator.CreateMemento()

	// 修改发起人的状态
	originator.SetState("Version 2.0")
	fmt.Println("Current State:", originator.GetState())

	// 恢复发起人的状态
	originator.Restore(caretaker.Memento)
	fmt.Println("Restored State:", originator.GetState())
}
