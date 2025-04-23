package main

import "fmt"

// 职责链模式
// 责任链模式是一种行为设计模式，它允许将请求沿着处理者链传递，直到有一个处理者处理它为止。
// 责任链模式的主要优点是可以将请求的发送者和接收者解耦。
// 这样，发送者不需要知道哪个处理者会处理请求，只需要将请求传递给链中的第一个处理者即可。
// 责任链模式的主要缺点是可能会导致请求处理的延迟，因为请求需要沿着链传递，直到有一个处理者处理它为止。
// 责任链模式的实现通常涉及到创建一个处理者接口和一个具体的处理者类。
// 处理者接口定义了一个处理请求的方法，具体的处理者类实现了这个方法，并在其中执行具体的处理逻辑。
// 责任链模式通常用于需要将请求沿着处理者链传递的场景，例如日志记录、权限验证等。

// HTTP中间件案例
type Handler interface {
	Handle(request string) string
}

type FinalHandler struct{}

func (h *FinalHandler) Handle(request string) string {
	// 处理请求
	fmt.Println("FinalHandler: Handling request:", request)
	return "Request handled by FinalHandler"
}

type Middleware1 struct {
	next Handler
}

func (m *Middleware1) Handle(request string) string {
	// 处理请求
	fmt.Println("Middleware1: Handling request:", request)

	// 调用下一个处理者
	if m.next != nil {
		return m.next.Handle(request)
	}
	return "Request handled by Middleware1"
}

type Middleware2 struct {
	next Handler
}

func (m *Middleware2) Handle(request string) string {
	// 处理请求
	fmt.Println("Middleware2: Handling request:", request)

	// 调用下一个处理者
	if m.next != nil {
		return m.next.Handle(request)
	}
	return "Request handled by Middleware2"
}


func main() {
	// 创建责任链
	finalHandler := &FinalHandler{}
	middleware1 := &Middleware1{next: finalHandler}
	middleware2 := &Middleware2{next: middleware1}

	// 处理请求
	response := middleware2.Handle("Request data")
	fmt.Println("Response:", response)
}
