package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int = 1
	var p *int = &a
	var fp *float32
	// 1. 指针不可运算
	// [X] p++或p=p+3
	//
	// 2. 不同类型指针不可转换
	// [X] fp = (*float32)(p)
	//
	// 3. 不同类型指针不可比较
	// [X] if fp == p {}
	// 但可以和nil做==或!=比较
	if fp == nil {
	}
	
	// unsafe

	// 4.
	// uintptr可以和unsafe.Pointer相互转换
	// 任何类型指针可以和unsafe.Pointer相互转换
	var unp unsafe.Pointer = unsafe.Pointer(p)
	fmt.Printf("%T %v %v\n", unp, unp, *(*int)(unp))

	// 5.
	// uintptr没有指针的语义，存的地址的变量会被gc回收
	// uintptr可运算，unsafe.Pointer若需要运算先转uintptr
	// unsafe.Pointer有指针语义，不会被gc回收

	// 6.
	// unsafe包应用：修改私有成员
	type Programmer struct {
		name string
		age int
		language string
	}

	pger := Programmer{"su", 24, "go"}
	fmt.Println(pger)
	name := (*string)(unsafe.Pointer(&pger))
	*name = "suyb"
	lang := (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&pger))+unsafe.Offsetof(pger.language)))
	*lang = "golang"
	fmt.Println(pger)
	// 假若是私有成员，不能直接访问
	lang = (*string)(unsafe.Pointer(uintptr(unsafe.Pointer(&pger))+unsafe.Sizeof(int(0))+unsafe.Sizeof(string(""))))
	*lang = "Golang"
	fmt.Println(pger)

	// 7.
	// unsafe包应用：获取slice、map的长度
	s := make([]int, 9, 20)
	sLen := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s))+uintptr(8)))
	sCap := *(*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&s))+uintptr(16))) // unsafe.Pointer和int在当前系统都是8字节
	fmt.Println(sLen, len(s))
	fmt.Println(sCap, cap(s))

	m := make(map[string]int)
	m["su"] = 24
	m["gong"] = 84
	count := **(**int)(unsafe.Pointer(&m))
	fmt.Println(count, len(m))

	// 8.
	// 字符串与byte切片零拷贝
	str1 := "string"
	bFs := string2bytes(str1)
	fmt.Println(string(bFs), len(bFs), cap(bFs))
	b := []byte("string")
	sFb := bytes2string(b)
	fmt.Println(sFb, len(sFb))
}


// 字符串与byte切片零拷贝
func string2bytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&s))
}

func bytes2string(b []byte) string {
	return *(*string)(unsafe.Pointer(&b))
}