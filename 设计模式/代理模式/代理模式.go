package main

import "fmt"

// 代理模式
// 代理模式是指一个类代表另一个类的功能。代理模式可以用来控制对某个对象的访问，或者在访问对象之前或之后执行一些操作。
// 代理模式通常用于以下场景：
// 1. 惰性加载：在需要使用某个对象时才创建它，而不是在程序启动时就创建它。
// 2. 权限控制：在访问某个对象之前检查权限，只有在权限足够时才允许访问。
// 3. 日志记录：在访问某个对象之前或之后记录日志，方便调试和分析。
// 4. 缓存：在访问某个对象之前检查缓存，如果缓存中有数据，则直接返回缓存中的数据，而不是访问对象。
// 5. 远程代理：在访问远程对象时，使用代理对象来处理网络通信和数据传输。

// 权限控制案例
type User struct {
	Name string
	Role string
}

type Document struct {
	Title   string
	Content string
}

type Proxy struct {
	User     *User
	Document *Document
}

func (p *Proxy) View() {
	if p.User.Role == "admin" {
		fmt.Printf("%s is viewing document: %s\n", p.User.Name, p.Document.Title)
	} else {
		fmt.Printf("%s does not have permission to view document: %s\n", p.User.Name, p.Document.Title)
	}
}
func (p *Proxy) Edit(content string) {
	if p.User.Role == "admin" {
		p.Document.Content = content
		fmt.Printf("%s edited document: %s\n", p.User.Name, p.Document.Title)
	} else {
		fmt.Printf("%s does not have permission to edit document: %s\n", p.User.Name, p.Document.Title)
	}
}

func (p *Proxy) Delete() {
	if p.User.Role == "admin" {
		fmt.Printf("%s deleted document: %s\n", p.User.Name, p.Document.Title)
	} else {
		fmt.Printf("%s does not have permission to delete document: %s\n", p.User.Name, p.Document.Title)
	}
}

func main() {
	user1 := &User{Name: "Alice", Role: "admin"}
	user2 := &User{Name: "Bob", Role: "user"}
	document := &Document{Title: "Design Patterns", Content: "Proxy Pattern"}

	proxy1 := &Proxy{User: user1, Document: document}
	proxy2 := &Proxy{User: user2, Document: document}

	proxy1.View()
	proxy1.Edit("Updated Content")
	proxy1.Delete()

	proxy2.View()
	proxy2.Edit("Updated Content")
	proxy2.Delete()
}
