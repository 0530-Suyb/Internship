package main

import "fmt"

type Power interface {
	Open()
	GetVoltage() int
}

// 5V
type V5 struct {
	Voltage int
}

func (v *V5) Open() {
	fmt.Println("开启5V电源")
}

func (v *V5) GetVoltage() int {
	return v.Voltage
}

func (v *V5) Charge() {
	fmt.Println("充电5V电源")
}

// 220V
type V220 struct {
	Voltage int
}

func (v *V220) Open() {
	fmt.Println("开启220V电源")
}

func (v *V220) GetVoltage() int {
	return v.Voltage
}

// 适配器
type Adaptor interface {
	Charge()
}

// 5V适配器
type AdaptorV5 struct {
	Power
}

func (a *AdaptorV5) Charge() {
	a.Power.Open()
	if a.Power.GetVoltage() < 5 {
		fmt.Println("电压过低，无法充电")
		return
	} else if a.Power.GetVoltage() > 5 {
		fmt.Printf("将%d电压转换为5V\n", a.Power.GetVoltage())
	}
	fmt.Println("充电5V电源")
}

func NewAdaptorV5(p Power) *AdaptorV5 {
	return &AdaptorV5{Power: p}
}

func main() {
	// 220V电源适配器
	adaptor1 := NewAdaptorV5(&V220{Voltage: 220})
	adaptor1.Charge()

	adaptor2 := NewAdaptorV5(&V5{Voltage: 5})
	adaptor2.Charge()
}
