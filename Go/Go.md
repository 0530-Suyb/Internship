[toc]



## 面向对象

### 匿名字段

即嵌入字段，可以不用写字段名只写类型

```go
type Person struct {
    name string
    sex  string
    age  int
}

type Student struct {
    Person	// 1. 匿名字段
    id   int
    addr string
    name string	// 2. 同名字段
}

// 匿名字段初始化
// Student{Person{"5lmh", "man", 20}, 1, "bj"}
// Student{Person: Person{"5lmh", "man", 20}}
// Student{Person: Person{name: "5lmh"}}

// 自定义类型
type mystr sting

type Teacher struct {
    *Person	// 3. 指针类型匿名字段 
    int
    mystr	// 4. 所有内置类型和自定义类型都可以作为匿名字段使用
}

func main() {
    var s Student
    
    // 同名字段
    s.name = "Student.name"
    s.Person.name = "Student.Person.name"
}
```

### 接口

接口interface定义对象的行为规范，只定义不实现。是一组方法的集合，是duck-type编程的一种体现。

**接口也是一种类型**，一种抽象的类型，区别于其他具体类型。





## 网络编程

### 互联网协议

![image-20250301164903646](C:\Users\hp-pc\Desktop\实习备战记\Go\img\OSI七层模型.png)

![image-20250301164730271](C:\Users\hp-pc\Desktop\实习备战记\Go\img\互联网五层模型.png)

### socket编程

**概述**

socket是一种进程通信机制，称“套接字”，用于描述IP地址和端口，是一个通信链的句柄。可以理解为TCP/IP网络的API，定义了许多函数，可利用socket开发应用程序。socket抽象层作为应用层与TCP/IP协议簇通信的中间软件抽象层，隐藏了很多底层细节。

常见Socket分为流式Socket（面向连接，TCP）和数据报式Socket（无连接，UDP）

![socket图解](C:\Users\hp-pc\Desktop\实习备战记\Go\img\socket图解.png)

**TCP编程**

TCP传输层控制协议，面向连接、可靠、基于字节流，存在粘包。

TCP服务端需要监听端口、接收客户端请求建立连接、创建协程处理连接。

TCP客户端需要建立与服务端的连接、进行数据收发、关闭连接。

```go
// server
package main

import (
	"bufio"
	"fmt"
	"net"
)

func proc(conn net.Conn) {
	defer conn.Close()
	for {
		reader := bufio.NewReader(conn)
		buf := [128]byte{}
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println(string(buf[:n]))
		conn.Write(buf[:n])
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go proc(conn)
	}
}
```

```go
// client
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	inputReader := bufio.NewReader(os.Stdin)
	for {
		input, _ := inputReader.ReadString('\n')
		if strings.ToUpper(input) == "Q" {
			return
		}
		_, err := conn.Write([]byte(input))
		if err != nil {
			fmt.Println(err)
			return
		}
		buf := [512]byte{}
		n, err := conn.Read(buf[:])
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
```

**UDP编程**

UDP协议，用户数据报协议，无连接，不可靠，但实时性号。

```go
// UDP server端
func main() {
    listen, err := net.ListenUDP("udp", &net.UDPAddr{
        IP:   net.IPv4(0, 0, 0, 0),
        Port: 30000,
    })
    if err != nil {
        fmt.Println("listen failed, err:", err)
        return
    }
    defer listen.Close()
    for {
        var data [1024]byte
        n, addr, err := listen.ReadFromUDP(data[:]) // 接收数据
        if err != nil {
            fmt.Println("read udp failed, err:", err)
            continue
        }
        fmt.Printf("data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
        _, err = listen.WriteToUDP(data[:n], addr) // 发送数据
        if err != nil {
            fmt.Println("write to udp failed, err:", err)
            continue
        }
    }
}
```

```go
// UDP 客户端
func main() {
    socket, err := net.DialUDP("udp", nil, &net.UDPAddr{
        IP:   net.IPv4(0, 0, 0, 0),
        Port: 30000,
    })
    if err != nil {
        fmt.Println("连接服务端失败，err:", err)
        return
    }
    defer socket.Close()
    sendData := []byte("Hello server")
    _, err = socket.Write(sendData) // 发送数据
    if err != nil {
        fmt.Println("发送数据失败，err:", err)
        return
    }
    data := make([]byte, 4096)
    n, remoteAddr, err := socket.ReadFromUDP(data) // 接收数据
    if err != nil {
        fmt.Println("接收数据失败，err:", err)
        return
    }
    fmt.Printf("recv:%v addr:%v count:%v\n", string(data[:n]), remoteAddr, n)
}
```

**TCP粘包**

tcp基于字节流，保持长连接时可以多次收发。粘包可能发生发送端也可能发生在接收端。

1. Nagle算法造成的发送端粘包：Nagle算法使得发送端不会立即发送数据，而是等待一段时间凑够数据或被ACK包触发而发送数据。
2. 接收端接收不及时造成的接收端粘包

粘包是由于接收端不确定传输的数据包的大小，通过在包头加入包体长度，读取包头后提取出包体长度，再往后读取出包体数据。

### http编程

http超文本传输协议，运行于TCP协议之上。

```go
// 服务端
package main

import (
    "fmt"
    "net/http"
)

func main() {
    //http://127.0.0.1:8000/go
    // 单独写回调函数
    http.HandleFunc("/go", myHandler)
    //http.HandleFunc("/ungo",myHandler2 )
    // addr：监听的地址
    // handler：回调函数
    http.ListenAndServe("127.0.0.1:8000", nil)
}

// handler函数
func myHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Println(r.RemoteAddr, "连接成功")
    // 请求方式：GET POST DELETE PUT UPDATE
    fmt.Println("method:", r.Method)
    // /go
    fmt.Println("url:", r.URL.Path)
    fmt.Println("header:", r.Header)
    fmt.Println("body:", r.Body)
    // 回复
    w.Write([]byte("www.5lmh.com"))
}
```

```go
// 客户端
package main

import (
    "fmt"
    "io"
    "net/http"
)

func main() {
    //resp, _ := http.Get("http://www.baidu.com")
    //fmt.Println(resp)
    resp, _ := http.Get("http://127.0.0.1:8000/go")
    defer resp.Body.Close()
    // 200 OK
    fmt.Println(resp.Status)
    fmt.Println(resp.Header)

    buf := make([]byte, 1024)
    for {
        // 接收服务端信息
        n, err := resp.Body.Read(buf)
        if err != nil && err != io.EOF {
            fmt.Println(err)
            return
        } else {
            fmt.Println("读取完毕")
            res := string(buf[:n])
            fmt.Println(res)
            break
        }
    }
}
```

### WebSocket编程

一种在单个TCP连接上进行全双工通信的协议，只需一次握手就能创建持久性连接，双向通信。



## gin框架

### 简介

- Gin是一个golang的微框架，封装比较优雅，API友好，源码注释比较明确，具有快速灵活，容错方便等特点
- 对于golang而言，web框架的依赖要远比Python，Java之类的要小。自身的`net/http`足够简单，性能也非常不错
- 借助框架开发，不仅可以省去很多常用的封装带来的时间，也有助于团队的编码风格和形成规范

### gin路由

**路由原理**

- httprouter将所有路由规则构造成一颗前缀树（Trie Tree，字典树）

  ![img](C:\Users\hp-pc\Desktop\实习备战记\Go\img\httprouter压缩字典树.png)

### gin数据解析和绑定

json、表单、URI

### gin渲染

JSON、结构体、XML、YAML、ProtoBuf数据响应

HTML模板渲染

重定向

同步异步 goroutine

### gin中间件

所有请求经过中间件

### 会话控制

**Cookie介绍**

HTTP无状态协议，服务器不能记录浏览器的访问状态，无法区分两个请求来自同一个客户端。

Cookie是服务器保留在浏览器上的一段信息，客户端每次请求带上Cookie

**Cookie缺点**

- 不安全，明文
- 增加带宽消耗
- 可以被禁用
- cookie有上限

**Sessions**





## 微服务

### 认识微服务

### 微服务生态

### 微服务详解

### RPC

远程过程调用RPC（Remote Procedure Call）是一个计算机协议，允许一台计算机的程序调用另一台计算机的子程序。在面向对象编程，也可称远程调用或远程方法调用。

![img](C:\Users\hp-pc\Desktop\实习备战记\Go\img\流行RPC框架对比.jpg)





## Goleveldb

### 简介

基于Go语言实现的LevelDB键值存储数据库，数据以键值对形式存储，模型简单直接，适合存储简单的数据映射关系。

优势：

- 高性能：**数据结构高效**，采用日志结构化合并树（LSM Tree）的数据结构，适合写密集型应用场景，数据写入时先追加到日志文件，再批量合并到磁盘上的SSTable（Sorted String Table）文件中，减少磁盘寻道时间，显著提升写入性能；**缓存机制优化**，支持对打开文件和数据块进行缓存，可调整缓存大小，减少磁盘操作。
- 易于使用：简单的API设计，Put、Get、Delete。

### 使用

安装`go get github.com/syndtr/goleveldb/leveldb`

```go
// 创建或打开数据库
{
	db, err := leveldb.OpenFile("path/to/db", nil)
	defer db.Close()
}

// 读取或修改数据库内容
{
    data, err := db.Get([]byte("key"), nil)
    err = db.Put([]byte("key"), []byte("value"), nil)
    err = db.Delete([]byte("key"), nil)
}

// 迭代数据库内容
{
    iter := db.NewIterator(nil, nil)
    for iter.Next() {
        key := iter.Key()
        value := iter.Value()
    }
	iter.Release()
	err = iter.Error()
}

// Seek-Then-Iterate
{
    iter := db.NewIterator(nil, nil)
    for ok := iter.Seek(key); ok; ok = iter.Next() {
        key := iter.Key()
        value := iter.Value()
    }
	iter.Release()
	err = iter.Error()
}

// 迭代数据库内容的子集
{
    iter := db.NewIterator(&util.Range{Start: []byte("foo"), Limit: []byte("xoo")}, nil)
    for iter.Next() {
        key := iter.Key()
        value := iter.Value()
    }
	iter.Release()
	err = iter.Error()
}

// 使用特定前缀迭代数据库内容的子集
{
    iter := db.NewIterator(util.BytesPrefix([]byte("foo-")), nil)
	for iter.Next() {
		// Use key/value.
		...
	}
	iter.Release()
	err = iter.Error()
}

// Batch writes
{
    batch := new(leveldb.Batch)
    batch.Put([]byte("foo"), []byte("value"))
	batch.Put([]byte("bar"), []byte("another value"))
	batch.Delete([]byte("baz"))
    err = db.Write(batch, nil)
}

// bloom过滤器
{
    o := &opt.Options{
        Filter: filter.NewBloomFilter(10),
    }
    db, err := leveldb.OpenFile("path/to/db", o)
    ...
    defer db.Close()
}
```

