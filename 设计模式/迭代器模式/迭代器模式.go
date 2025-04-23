package main

import "fmt"

// 迭代器模式
// 迭代器模式是一个对象行为模式，它提供一种方法顺序访问一个集合对象中的各个元素，而又不暴露该对象的内部表示。
// 迭代器模式的主要优点是：
// 1. 封装迭代算法：迭代器模式将迭代算法封装在迭代器类中，客户端不需要关心迭代的细节，只需要使用迭代器提供的方法来访问集合中的元素。
// 2. 支持多种遍历方式：迭代器模式可以支持多种遍历方式，例如顺序遍历、逆序遍历等，客户端可以根据需要选择不同的迭代器来访问集合中的元素。
// 3. 提高代码的可读性和可维护性：迭代器模式将迭代算法与集合对象分离，使得代码更加清晰易懂，便于维护和扩展。
// 4. 支持并发访问：迭代器模式可以支持多个客户端同时访问同一个集合对象，而不会互相干扰。
// 5. 支持懒加载：迭代器模式可以支持懒加载，即在需要时才加载集合中的元素，从而提高性能。
// 迭代器模式的主要缺点是：
// 1. 增加了类的数量：迭代器模式需要定义迭代器类和集合类，增加了类的数量，可能会导致代码复杂度增加。
// 2. 可能会影响性能：迭代器模式需要维护一个迭代器对象，可能会增加内存开销和性能开销，尤其是在集合对象很大时。

// 例子：
// 假设我们有一个集合类 `Collection`，它包含多个元素，我们需要一个迭代器类 `Iterator` 来遍历这些元素。
// 迭代器类提供了 `Next` 方法来访问集合中的元素。
// 集合类提供了 `CreateIterator` 方法来创建迭代器对象。

type Iterator interface {
	Next() interface{} // 返回下一个元素
}

type ConcreteIterator struct {
	items []interface{} // 集合中的元素
	index int           // 当前索引
}

func (c *ConcreteIterator) Next() interface{} {
	if c.index < len(c.items) {
		item := c.items[c.index]
		c.index++
		return item
	}
	return nil // 返回 nil 表示没有更多元素
}

func NewConcreteIterator(items []interface{}) *ConcreteIterator {
	return &ConcreteIterator{
		items: items,
		index: 0,
	}
}

// 集合接口
// 定义一个集合接口，提供创建迭代器的方法和获取集合元素的方法
type Collection interface {
	CreateIterator() Iterator // 创建迭代器
	GetItems() []interface{}  // 获取集合中的元素
}

type ConcreteCollection struct {
	items []interface{} // 集合中的元素
}

func (c *ConcreteCollection) CreateIterator() Iterator {
	return NewConcreteIterator(c.items)
}

func (c *ConcreteCollection) GetItems() []interface{} {
	return c.items
}

func NewConcreteCollection(items []interface{}) *ConcreteCollection {
	return &ConcreteCollection{
		items: items,
	}
}

func main() {
	// 创建一个集合对象
	collection := NewConcreteCollection([]interface{}{"A", 'b', "C", "D", 1})

	// 创建一个迭代器对象
	iterator := collection.CreateIterator()

	// 使用迭代器遍历集合中的元素
	for item := iterator.Next(); item != nil; item = iterator.Next() {
		fmt.Println(item)
	}
}
