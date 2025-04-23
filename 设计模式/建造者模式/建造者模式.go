package main

import "fmt"

// 建造者模式
// 1. 定义一个产品类，包含多个属性
// 2. 定义一个建造者接口，包含创建产品的方法
// 3. 定义多个具体的建造者类，实现建造者接口，创建产品的具体实现，分别实现不同的属性
// 4. 定义一个指挥者类，使用建造者接口创建产品
// 5. 客户端使用指挥者类创建产品
// 6. 通过建造者模式，可以将产品的创建和使用分离，方便扩展和维护
// 7. 适用于复杂对象的创建，尤其是当对象的创建过程需要多个步骤时

type Building struct {
	Rooms   int
	Windows int
	Doors   int
}

type Builder interface {
	Build()
}

type RoomBuilder struct {
	Building *Building
}

func (b *RoomBuilder) Build() {
	b.Building.Rooms++
	fmt.Println("Building a room")
}

func NewRoomBuilder(b *Building) *RoomBuilder {
	return &RoomBuilder{Building: b}
}

type WindowBuilder struct {
	Building *Building
}

func (b *WindowBuilder) Build() {
	b.Building.Windows++
	fmt.Println("Building a window")
}

func NewWindowBuilder(b *Building) *WindowBuilder {
	return &WindowBuilder{Building: b}
}

type DoorBuilder struct {
	Building *Building
}

func (b *DoorBuilder) Build() {
	b.Building.Doors++
	fmt.Println("Building a door")
}

func NewDoorBuilder(b *Building) *DoorBuilder {
	return &DoorBuilder{Building: b}
}

type Director struct {
	Builder Builder
}

func (d *Director) SetBuilder(builder Builder) {
	d.Builder = builder
}

func (d *Director) Construct() {
	d.Builder.Build()
}

func main() {
	// 创建一个建筑
	building := Building{}

	// 创建建造者
	roomBuilder := NewRoomBuilder(&building)
	windowBuilder := NewWindowBuilder(&building)
	doorBuilder := NewDoorBuilder(&building)

	// 创建指挥者
	director := &Director{}
	director.SetBuilder(roomBuilder)
	director.Construct()
	director.SetBuilder(windowBuilder)
	director.Construct()
	director.SetBuilder(doorBuilder)
	director.Construct()

	fmt.Printf("Building has %d rooms, %d windows, and %d doors\n", building.Rooms, building.Windows, building.Doors)
}
