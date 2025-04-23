package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func main() {
	// 获取类型和值
	var x int = 2
	xT := reflect.TypeOf(x)
	xV := reflect.ValueOf(x)
	fmt.Println(xT, xV)

	// 动态修改值，需要传入指针，否则不能修改，因为value存的是a的一个拷贝
	a := 1
	v := reflect.ValueOf(&a)
	// v is unsettable, 只有存储了原变量本身才是可设置的，
	realA := v.Elem()
	realA.SetInt(2)
	fmt.Println(a)

	// 获取结构体变量
	p := person{"Alice", 20}
	pV := reflect.ValueOf(p)
	if pV.Kind() == reflect.Struct {
		for i:=0; i<pV.NumField(); i++ {
			field := pV.Field(i)
			fmt.Println(pV.Type().Field(i).Name, field)
		}
	}

	// 调用方法
	p2 := &person{"Bob", 21}
	p2V := reflect.ValueOf(p2)
	method := p2V.MethodByName("SayHello")
	method.Call(nil)

	

	var r io.Reader
	f, err := os.Open("file.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	r = f
	rv := reflect.ValueOf(r)
	fmt.Println(reflect.TypeOf(r), rv.Type())

	switch r.(type) {
	case *os.File:
		fmt.Println("yes, it's *os.File")
	}

	fmt.Println(add(1.1, 2.2))
}

func add(a, b any) (res any) {
	switch t := reflect.TypeOf(a); t.Kind() {
	case reflect.Int:
		fmt.Printf("add %s\n", t.Name())
		res = addInt(a.(int), b.(int))
	case reflect.Float64:
		fmt.Printf("add %s\n", t.Name())
		res = addFloat(a.(float64), b.(float64))
	}
	return
}

func addInt(a, b int) int {
	return a + b
}

func addFloat(a, b float64) float64 {
	return a + b
}


type person struct {
	Name string
	Age  int
}

func(p *person) SayHello() {
	fmt.Println("hello, I am", p.Name)
}
