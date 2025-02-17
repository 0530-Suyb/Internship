[toc]

# 《图解网络》

## 一、网络基础篇

### 1.1 TCP/IP网络模型有哪几层？

同设备内进程通信可通过管道、消息队列、共享内存、信号等方式，而不同设备间进程通信需要通过网络通信，而设备多种多样，要兼容多种设备，就需要有一套通用的网络协议。

网络协议是分层的，各层有各自的功能职责。TCP/IP网络模型分层：**应用层、传输层、网络层、网络接口层**。

**1. 应用层**

**应用层（Application Layer）**只专注于为用户提供应用功能，如HTTP、FTP、Telnet、DNS、SMTP等，不关心数据的传输，数据交给下一层去处理。

应用层工作于操作系统的**用户态**，传输层及以下则工作于**内核态**。

**2. 传输层**

**传输层（Transport Layer）**接收应用层传输来的数据包，或将收到的数据包发送给应用层，为应用层提供网络支持。

传输层有**TCP（传输控制协议Transmission Control Protocol）**和**UDP（用户数据报协议User Datagram Protocol）**两种传输协议。

TCP协议保证数据包能**可靠传输**到接收端，相比UDP协议多了**流量控制、超时重传、拥塞控制**等特性。大部分应用都使用TCP，如HTTP。传输数据较大时，直接传输并不好控制，因此当传输层数据包大小超过**MSS（TCP最大报文段长度，）**就将数据包分块传输，某个块损坏或丢失时只需要重传这一个块。TCP中每个分块称一个**TCP段（TCP Segment）**。

UDP协议相对简单，只负责发送数据包，不保证数据包抵达接收端，但实时性更好，传输效率也高。UDP当然也可以实现可靠传输，将TCP的特性在应用层上实现就可以，不过实现一个商用的可靠UDP传输协议并不简单。

一台设备可能会有多个应用在接收或传输数据，需要通过**端口**来区分应用，如Web服务器用80端口、远程登录服务器用22端口等。浏览器中每个标签栏都是一个独立的进程，操作系统会为这些进程临时分配端口。传输层报文中携带了端口号。

传输层数据单元PDU（Protocol Data Unit）为段（Segment）或数据报（Datagram），网络层为报文（Packet），网络接口层为帧（frame）。应用层为数据（Data），HTTP传输单元是消息或报文（message）。其实没本质区别，都可以统称数据包。

**3. 网络层**

网络中存在各样的线路和分叉路口，数据的传输需要对各种路径和节点进行选择（路由），而传输层的设计理念是简单、高效、专注，并不负责这部分的功能。传输层只需服务好应用层，实际的传输工作交给下一层网络层。

**网络层（Internet Layer）**常用**IP协议（Internet Protocol）**，IP协议将传输层报文作为数据部分，加上IP包头组成**IP报文**，IP报文大小超过**MTU（最大传输单元Maximum Transmission Unit，以太网中一般1500字节）**则再次分片（分的是包含传输层TCP头部的数据包）。

<img src=".\img\基础篇\IP报文组成.png" style="zoom:50%;" />

网络中使用**IP地址**区分设备，对**IPv4协议**，IP地址共32位，分四段，每段8位，如192.168.100.1。

为便于寻址，将IP地址划分为两段，前n位作**网络号**，后32-n位作**主机号**。网络号标识IP地址属于哪个子网，主机号标识同一子网内不同主机。寻址中先匹配网络号，后找对应主机。

网络号和主机号通过**子网掩码**来划分，如192.168.100.1/24，/24就是子网掩码255.255.255.0，即前24位为网络号，后8位为主机号。通过IP地址与子网掩码按位与运算可以得到网络号，子网掩码的取反后与IP地址按位与运算则是主机号。24位掩码长度，可用IP数有254个，0和255分别用于网络号和广播。

CIDR无类别域间路由，`IP地址/前缀长度`，采用网络掩码“位格式”但不局限于固定的`/8`、`/16`、`/24`。

**寻址通过IP地址确定目标设备的位置，而数据包从源到目的地的过程则由路由负责**。实际中，设备间需要通过许多网关、路由器、交换机等众多网络设备连接，会形成很多条网络路径，需要路由算法来决定最佳路径。

**寻址**为网络设备分配唯一标识以定位目标，**路由**负责数据包传输。

**4. 网络接口层**

**网络接口层（Link Layer）**从网络层收到IP报文后再IP头部加上**MAC**（媒体存取控制位址Media Access Control Address）头部，并封装成**数据帧（Data frame）**发送到网络上。

IP头部中接收方IP地址表示网络包的目的地，但在以太网内便行不通了。以太网是一种在局域网内的，将周围设备连接起来，使它们之间进行通信的技术，包括了以太网接口、Wi-Fi接口、以太网交换机、路由器的千兆万兆以太网口、网线。

以太网中通过MAC地址来识别设备，因此数据帧的MAC头部包含了接收方和发送方的MAC地址，通过ARP协议可获取对方的MAC地址。

网络接口层位网络层提供**「链路级别」**传输的服务，负责在以太网、WiFi 这样的底层网络上发送原始数据包，工作在网卡这个层次，使用 MAC 地址来标识网络上的设备。

**5. 总结**

TCP/IP 网络通常是由上到下分成 4 层，分别是**应用层，传输层，网络层和网络接口层**。

<img src=".\img\基础篇\tcpip参考模型.drawio.png" alt="img" style="zoom:50%;" />

每一层封装格式：

![img](.\img\基础篇\封装.png)

网络接口层的传输单位是帧（frame），IP 层的传输单位是包（packet），TCP 层的传输单位是段（segment），HTTP 的传输单位则是消息或报文（message）。但这些名词并没有什么本质的区分，可以统称为数据包。

- 比特流bit：物理层
- 数据帧frame：数据链路层
- 数据包/报文分组packet：网络层
- 数据报datagram：传输层UDP
- 数据段segment：传输层TCP
- 消息/报文message：应用层

### 1.2 键入网址到网页显示，期间发生了什么？

**1. HTTP**

**解析URL**

![URL 解析](.\img\基础篇\URL组成.png)

**生产HTTP请求信息**

![HTTP 的消息格式](.\img\基础篇\HTTP请求报文和响应报文.png)



**2. DNS**

**域名层级关系**

如www.server.com，实际上最后还有一个点，代表根域名。

- 根DNS服务器（.）
- 顶级域DNS服务器（.com）
- 权威DNS服务器（server.com）

根域的DNS服务器信息保存在任一台DNS服务器中，通过根域DNS服务器可以知道目标DNS服务器。

**DNS域名解析**



![域名解析的工作流程](.\img\基础篇\DNS域名解析.png)

**缓存**

浏览器看自身有无对域名的缓存，没有再看操作系统是否有对域名的缓存，没有再看hosts文件，没有再问本地DNS服务器。本地DNS服务器也有缓存。

**3. 协议栈**

![img](.\img\基础篇\协议栈.png)

- ICMP：Internet控制报文协议，用于告知网络包传送过程中产生的错误以及各种控制信息。
- ARP：地址解析协议，用于根据IP地址查询相应的以太网MAC地址。

**4. 可靠传输——TCP**

**TCP包头格式**

![TCP 包头格式](.\img\基础篇\TCP报文头格式.png)

- 源端口号、目标端口号：标识发送给的应用。

- 序号：包的序号，解决包乱序问题。

- 确认序号：用于确认是否收到包，没收到就重发，解决丢包问题。

- 状态位：SYN发起连接同步序列号，ACK用于确认已接收的数据，RST重置连接，FIN结束连接，URG标志是否包含紧急数据，PSH推送功能以要求接收方立即处理数据。TCP面向连接，要维护连接状态，带状态位的包会引起双方状态变更。

  ![](.\img\基础篇\TCP包头状态位.png)

- 窗口大小：做流量控制，标识当前处理能力。

- 拥塞控制：控制自己发送速度。

**三次握手**

![TCP 三次握手](.\img\基础篇\TCP三次握手.png)

三次握手中客户端和服务端都在经过一发一收后进入ESTABLISHED状态，**三次握手目的就是保证双方都有发送和接收的能力**。

**查看TCP连接状态**

`netstat -napt`

![TCP 连接状态查看](.\img\基础篇\TCP连接状态查看.jpg)

**TCP分割数据**

![MTU 与 MSS](.\img\基础篇\MTU与MSS.jpg)

- MTU：网络包最大长度，以太网中一般1500字节
- MSS：除去IP和TCP头部，网络包能容纳的TCP数据最大长度。



![数据包分割](.\img\基础篇\HTTP数据拆分.jpg)

**TCP报文生成**

![TCP 层报文](.\img\基础篇\TCP报文格式.jpg)

源端口号为浏览器随机生成，目标端口号为web服务器监听的端口，HTTP默认80，HTTPS默认443。

**5. 远程定位——IP**

**IP包头格式**

<img src=".\img\基础篇\IP包头格式.jpg" alt="IP 包头格式" style="zoom:50%;" />

HTTP经过TCP传输，IP包头的协议号为0x06。

如果客户端有多个网卡，即有多个IP地址，需要根据路由表规则来判断使用哪一个网卡作为源地址IP。使用目标IP和路由条目的子网掩码做与运算，如果和条目的目的地址`Destination`匹配则选该网卡，否则继续与下一条路由判断。`0.0.0.0`为默认网关，如果其他条目都无法匹配则将包发送给路由器，`Gateway`为路由器的IP地址。

`route -n`查看路由表。

![路由规则判断](.\img\基础篇\路由规则判断.jpg)

**IP报文生成**

![IP 层报文](.\img\基础篇\IP层报文格式.jpg)



**6. 两点传输——MAC**

**MAC包头格式**

![MAC 包头格式](.\img\基础篇\MAC包头格式.jpg)

一般TCP/IP通信里，MAC包头协议类型只使用：

- `0800`：IP协议
- `0806`：ARP协议

发送方MAC地址为网卡ROM中写死的MAC地址，接收方MAC地址需要通过ARP协议来获得。![ARP 广播](.\img\基础篇\ARP广播.jpg)

ARP缓存时间几分钟。

发包时先查询ARP缓存，有就直接使用，不存在则发送ARP广播查询。

查看ARP缓存内容，`arp -a`。

如果源和目标不是在同一子网，则MAC地址可以使用广播地址（全1，FF-FF-FF-FF-FF-FF）或组播地址（第8bit为1）。

**MAC报文生成**

![MAC 层报文](.\img\基础篇\MAC报文格式.jpg)



**7. 出口——网卡**

网卡将内存中的二进制网络包从数字信号转换为电信号发送出去。

**网卡驱动程序**用于控制网卡，在获取网络包后将其复制到网卡的缓冲区，接着在其开头加上**报头和起始帧分界符**，在末尾加上**用于检测错误的帧校验序列**。

![数据包](.\img\基础篇\网卡数据包发送.png)

- 起始帧分界符：表示包起始位置的标记。
- FCS帧校验序列：检查包传输中是否有损坏。

**8. 送别者——交换机**

交换机工作在**MAC层**，二层网络设备。

交换机端口不具有MAC地址，直接接收所有包，经过FCS校验后放入缓冲区。

随后查询包的接收方MAC地址是否在MAC地址表中有记录，并将信号从对应端口发送出。

![交换机的 MAC 地址表](.\img\基础篇\交换机的MAC地址表.jpg)

当MAC地址表找不到接收方MAC地址时，可能是该设备还没向交换机发送过包，或者该设备一段时间没工作导致被从表中删除。这时交换机会将包转发到除了源端口外的所有端口上，只有相应的接收者会接收包，其他则忽略包。

通过接收方MAC地址是广播地址，也会发送到除源端口外的所有端口。

广播地址

- MAC地址中的`FF:FF:FF:FF:FF:FF`
- IP地址中的`255.255.255.255`

**9. 出境大门——路由器**

路由器基于IP设计，三层网络设备，每个端口都有MAC地址和IP地址。

转发包时，路由器端口接收发给自己的以太网包，接着在路由表中查询转发目标，从相应端口作为发送方将以太网包发送出去。

接收的以太网包的MAC头部作用是将包送达路由器。

![路由器转发](.\img\基础篇\路由器转发.jpg)

**路由器的发送操作**

根据路由表网关列判断对方地址

- 如果网关为一个IP地址，说明还未抵达终点，需要转发到这个目标地址上，通过其继续转发。
- 如果网关为空，说明到达终点，IP头部的接收方IP地址就是要转发到的目标地址。

确认下一步要转发的地址后，通过ARP协议获得接收方MAC地址，从ARP缓存取或发送ARP查询请求。

网络包传输中，源IP和目标IP始终不变，变得一直是MAC地址。

**10. 互相扒皮——服务器与客户端**

![网络分层模型](.\img\基础篇\网络分层模型.jpg)

链路层检查MAC地址，网络层检查IP地址，传输层检查序列号和回复ACK，并将包给到监听对应端口的HTTP服务器



### 1.3 Linux系统是如何收发网络包的？

**1. 网络模型**

OSI模型7层（Open System Interconnection Reference Model）

- 应用层：给应用程序提供统一的接口
- 表示层：将数据转换成兼容另一个系统能识别的格式
- 会话层：建立、管理和终止表示层实体之间的通信会话
- 传输层：端到端数据传输
- 网络层：数据的路由、转发、分片
- 数据链路层：数据的封帧和差错检测，以及MAC寻址
- 物理层：物理网络中传输数据帧

TCP/IP网络模型，四层模型，Linux网络协议栈依此实现。

- 应用层
- 传输层
- 网络层
- 网络接口层

常说的七层和四层负载均衡，是用OSI网络模型描述，七层是应用层，四层是传输层。

![](.\img\基础篇\四层和七层负载均衡对比.png)



**2. Linux网络协议栈**

<img src=".\img\基础篇\Linux网络协议栈.png" alt="img" style="zoom:50%;" />

应用程序通过系统调用和Socket层数据交互。



**3. Linux接收网络包的流程**

![img](.\img\基础篇\Linux网络包收发流程.png)

**接收通知**

网卡接收网络包后通过DMA技术将网络包写入指定内存地址（Ring Buffer，环形缓冲区），接着通知操作系统网络包到达。

通知操作系统最简单方式就是触发中断，但网络包众多，频繁中断CPU会影响效率。因此Linux内核2.6版本引入**NAPI机制**，混合**中断和轮询**的方式接收网络包，首先采用中断唤醒数据接收的服务程序，在poll方式轮询数据。

即DMA技术将网络包写入指定内存地址后，网卡向CPU发起中断，CPU收到硬件中断请求后根据中断表，调用注册的中断处理函数。硬件中断处理函数中，先**暂时屏蔽中断**（表示已知内存有数据，网卡下次收到数据包直接写内存不用通知CPU，以提高效率），接着发起**软中断**，然后恢复刚才屏蔽的中断。

**软中断处理**

内核中ksoftirqd线程负责软中断处理，收到中断后轮询处理数据。线程从Ring Buffer获取一个数据帧sk_buff，作为一个网络包交给网络协议栈逐层处理。

**网络协议栈处理**

进入网络接口层，检查报文合法性，合法则确认网络包的上层协议类型，IPv4还是IPv6，然后去掉帧头帧尾，交给网络层。

网络层取出IP包，判断包下一步走向是交给上层处理还是转发出去。确认发给本机则从IP头看上一层协议类型（TCP/UDP），去掉IP头交给传输层。

传输层取出TCP或UDP头，根据四元组[源IP、源端口、目的IP、目的端口]作为标识找出对应Socket，把数据放到Socket的接收缓冲区。

应用程序调用Socket接口，将内核的Socket接收缓冲区的数据**拷贝**到应用层缓冲区，再唤醒用户进程。



**4. Linux发送网络包的流程**

发送时，应用程序调用（系统调用）Socket发送数据包的接口，从用户态转入内核态的Socket层，内核申请一个内核态的sk_buff内存，将用户待发送数据拷贝到sk_buff内存，并将其加入发送缓冲区。

网络协议栈从Socket发送缓冲区取出sk_buff，按TCP/IP协议栈从上至下逐层处理。

TCP传输数据时会拷贝一个新sk_buff副本，因为TCP支持丢失重传，而sk_buff数据在网卡发送后会被释放。

发送报文时sk_buff结构体给数据预留足够空间来填充各层首部，加过各层时通过减少skb-\>data来增加协议首部。而在接收报文时则是增加skb-\>data来逐步剥离协议首部。

![img](.\img\基础篇\网络包收发使用sk_buff.jpg)

网络层选取下一跳路由、填充IP头、netfilter过滤、对超过MTU大小数据包分片，再交给网络接口层。

网络接口层通过ARP协议获取下一跳MAC地址，对sk_buff填充帧头帧尾，将sk_buff放入网卡发送队列。

触发软中断告诉网卡驱动程序，网卡驱动程序将sk_buff取出挂入Ring Buffer，接着将sk_buff数据映射到网卡可访问的内存DMA区域，触发真实发送。

数据发完，网卡触发硬中断释放sk_buff和RingBuffer等内存。对TCP协议，收到ACK应答传输层再释放原始sk_buff。

**发送网络数据涉及几次内存拷贝？**

- 发送数据时系统调用Socket接口，内核申请内核态sk_buff内存将用户待发送数据拷贝进来，并将其加入发送缓冲区。
- 使用TCP传输协议，会将原始sk_buff保留在传输层，拷贝一个副本往下传。
- IP层发现sk_buff大于MTU时会申请额外sk_buff，将原sk_buff拷贝为多个小的sk_buff。



## 二、HTTP篇

### 2.1 HTTP常见面试题

![提纲](.\img\HTTP篇\HTTP提纲.png)



**1. HTTP基本概念**

**HTTP是什么？**

HTTP（HyperText Transfer Protocol，超文本传输协议），是一个在计算机世界里专门在两点间传输文字、图片、音频、视频等超文本数据的约定和规范。



**HTTP常见的状态码有哪些？**

![ 五大类 HTTP 状态码 ](.\img\HTTP篇\五大类HTTP状态码.png)

2xx

- 200 OK：一切正常
- 204 No Content：与200基本相同，但响应头没有body数据
- 206 Partial Content：用于HTTP分块下载或断点续传，body数据只是资源的一部分

3xx

- 301 Moved Permanently：永久重定向，资源不存在了，需改用新URL访问
- 302 Found：临时重定向，请求资源还在但暂时用另一个URL访问
- 304 Not Modified：缓存重定向，用于服务器说明资源未修改，客户端可以继续使用本地缓存资源

301和302会在响应头里使用Location字段指明跳转的URL。

4xx

- 400 Bad Request：客户端请求报文有误，笼统
- 403 Forbidden：服务器禁止访问资源，不是客户端请求出错
- 404 Not Found：请求的资源在服务器上不存在或未找到

5xx

- 500 Internal Server Error：笼统通用错误码，服务器出错
- 501 Not Implemented：客户端请求的功能暂时还不能支持
- 502 Bad Gateway：通常是服务器作为网关或代理时返回的错误码，表示服务器自身正常工作，访问后端服务器发生了错误
- 503 Service Unavailable：服务器当前很忙，暂时无法响应客户端



**HTTP常见字段有哪些？**

Host：客户端发送请求时指定服务器域名。即便同一台服务器上也能访问不同网站。

Content-Length：服务器返回数据的长度。该字段作为HTTP body边界指明后续跟的数据长度，解决TCP粘包问题。另外HTTP协议通过回车符+换行符作为HTTP header边界也是解决粘包问题的一种方法。

<img src="C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\connection字段.png" alt="img" style="zoom:50%;" />

Connection：用于客户端要求服务器使用**HTTP长连接**机制，以便请求复用，直到任一端提出断开连接。HTTP/1.1默认都是长连接，老版本默认关闭，要开启需指定`Connection: Keep-Alive`。HTTP/1.1要关闭长连接则`Connection: close`。（HTTP的Keep-Alive由应用层实现，称HTTP长连接，TCP的Keepalive由TCP层实现，称TCP保活机制）

Content-Type：用于服务器回应，告诉客户端本次数据的格式。客户端请求时使用`Accept`字段声明接受的格式。

<img src="C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\content-type字段.png" alt="img" style="zoom:50%;" />

Content-Encoding：说明数据的压缩方式。

<img src="C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\content-encoding字段.png" alt="img" style="zoom:50%;" />





**2. GET与POST**

**GET和POST有什么区别？**

GET语义是从服务器获取指定资源，请求的资源位置一般写在URL中，URL规定只能支持ASCII字符，且URL长度有限。

POST语义是根据请求负荷（报文body）对指定的资源做出处理，处理方式根据资源类型而定，数据写在报文body，可以是任意格式只要双方协商好，且大小不限制。

**GET和POST方法都是安全和幂等的吗？**

**概念**

- 安全：HTTP协议里安全指请求方法不会破坏服务器上资源

- 幂等：多次执行相同操作，结果都相同

根据RFC规范定义的语义来看

- GET方法安全、幂等、可缓存：只读，每次结果相同，数据安全。可对GET请求的数据缓存。
- POST方法不安全、不幂等、（大部分实现）不可缓存：新增或提交数据会修改服务器资源。

不过如果开发者不依据RFC规范来做，GET也可以实现新增或删除数据，POST方法也可以实现查询数据。

GET和POST数据只是在URL或body里，用HTTP都是明文，要避免数据窃取要用HTTPS。

理论上任何请求都可以带body，URL查询参数也不是GET独有。



**3. HTTP缓存技术**

**HTTP缓存有哪些实现方式？**

对于重复性的HTTP请求，可以把请求-响应的数据缓存本地。

HTTP缓存有两种实现方法：**强制缓存**和**协商缓存**

**什么是强制缓存？**

浏览器判断缓存没过期则直接使用本地缓存，由浏览器决定是否使用缓存。

响应头的状态码后表示`(from disk cache)`就是使用了强制缓存。

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\强制缓存.png)

强缓存利用两个HTTP响应头部字段实现，`Cache-Control`（相对时间）、`Expires`（绝对时间），前者优先级高。

**什么是协商缓存？**

如出现响应码304时，就是告诉浏览器使用本地缓存，这种由服务器告知客户端的就是**协商缓存**。

<img src="C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\缓存etag.png" alt="img" style="zoom:50%;" />

协商缓存是与服务器协商后，通过协商结果来判断是否使用本地缓存。

协商缓存可以基于两种头部实现

- 第一种：响应头部`Last-Modified`（响应资源最后修改时间）和请求头部`If-Modified-Since`（资源过期时，若响应头部有`Last-Modified`则将值复制过来发给服务器做判断是否从该时间后有更新）
- 第二种：响应头部`Etag`（响应资源唯一标识）和请求头部`If-None-Match`（资源过期时，若响应头部有`Etag`则附上发给服务器判断是否改变）

第一种基于时间，第二种基于唯一标识，后者更准确判断是否修改，当同时存在`Etag`和`Last-Modified`时，`Etag`优先级高。

- 即时文件内容没有修改，最后修改时间还是可能改变
- 有些文件在秒级内修改，`If-Modified-Since`检查粒度只到秒级，检查不出来
- 有些服务器不能精确获得文件最后修改时间

协商缓存需要配合强制缓存的`Cache-Control`字段来使用。

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\http缓存.png)



**4. HTTP特性**

**HTTP/1.1的优点有哪些？**

- 简单

  HTTP基本报文格式header+body，头部信息也是key-value简单文本形式，易于理解。

- 灵活和易于扩展

  HTTP协议的各类请求方法、URI/URL、状态码、头字段等每个组成要求都没有固定死，允许开发人员自定义和扩充。

  且工作在应用层，下层可随意改变，如

  - HTTPS在HTTP与TCP层间增加SSL/TLS安全传输层
  - HTTP/1.1和HTTP/2.0传输协议使用TCP协议，HTTP/3.0使用UDP协议

- 应用广泛和跨平台

**HTTP/1.1的缺点有哪些？**

- 无状态双刃剑：无需额外资源记录状态信息，但完成关联性操作很麻烦。解决方案如**Cookie**技术，在请求和响应报文中写入Cookie信息来控制客户端状态。

  <img src="C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\cookie技术.png" alt="Cookie 技术" style="zoom:50%;" />

- 明文传输双刃剑：明文方便阅读，但信息泄露

- 不安全：**明文通信**内容窃听，无**身份验证**会有伪装攻击，**报文完整性无法证明**会遭篡改

**HTTP/1.1的性能如何？**

HTTP基于TCP/IP，使用请求-应答通信模式，性能关键在于

- 长连接：减少TCP连接的重复建立和断开带来额外开销

- 管道网络传输：减少整体响应时间，一个TCP连接里可连续发起多个请求，无需等待第一个请求的响应回来。不过服务器必须按照接收请求的顺序发送对这些管道化请求的响应，存在**队头堵塞**问题（前面请求处理较久，导致后面请求处理被阻塞）。HTTP/1.1管道解决了请求的队头阻塞，但没解决响应的队头阻塞。默认不开启。

  

**5. HTTP与HTTPS**

**HTTP与HTTPS有哪些区别？**

- HTTP超文本传输协议，明文传输，有安全风险。HTTPS在HTTP与TCP层间加入SSL/TLS安全协议，加密传输。
- HTTP只需要TCP三次握手，HTTPS在三次握手后还要进行SSL/TLS握手。
- HTTP默认端口80，HTTPS默认端口443。
- HTTPS需要向CA申请数字证书来保证服务器身份可信。

**HTTPS解决了HTTP的哪些问题？**

- 窃听风险 - 信息加密：混合加密实现信息机密性，HTTPS在通信建立前使用非对称加密交换会话秘钥，后续通信使用对称加密的会话秘钥加密明文。

- 篡改风险 - 校验机制：摘要算法实现完整性，另外为保证**内容+哈希值**不被篡改，使用非对称密钥生成数字签名和验证签名。

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\摘要算法.png)

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\数字签名.png)

  公钥加密+私钥解密，可以保证内容传输的安全；私钥加密+公钥解密，可以保证消息不被冒充。

- 冒充风险 - 身份证书：数字证书包含服务器公钥，由CA机构签发，保证服务器公钥身份。

  ![数子证书工作流程](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\数字证书工作流程.png)

**HTTPS是如何建立连接的？期间交互了什么？**

SSL/TLS协议基本流程

- 客户端向服务器索要并验证服务器公钥
- 双方协商生成会话秘钥
- 双方采用会话秘钥加密通信

SSL/TLS建立是前两步，即TLS握手阶段，涉及四次通信。采用不同密钥交换算法，握手流程也不同，有RSA和ECDHE两种常用密钥交换算法。以下以RSA密钥交换算法的为例

![HTTPS 连接建立过程](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\HTTPS工作流程.png)

三个随机数的作用：

- 防止重放攻击和伪造攻击：客户端和服务器随机数每次连接都不同
- 避免服务器控制密钥生成：双方参与提供随机数
- 增强密钥强度和不可预测性
- 适应不同的密钥交换算法

第一个客户端随机数应该是增强密钥安全性，RSA中`Master Secret = PRF(Pre-Master Secret, "master secret", Client Random + Server Random)`,没有客户端随机数，会使结果相对可预测。

基于RSA的HTTPS存在**前向安全**问题（服务器私钥泄露会导致过去被第三方截获的所有TLS通信密文被破解），为此大多数网站是使用ECDHE密钥协商算法。

客户端数字证书校验流程：

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\证书的校验.png)

证书验证过程存在证书信任链问题，需要逐层验证。根证书是自签证书，默认信任根证书。多层级将根证书隔离，保证根证书绝对安全性。

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\baidu证书.png)

**HTTPS的应用数据是如何保证完整性的？**

TLS分握手协议和记录协议两层

- 握手协议负责协商加密算法和生成对称密钥
- 记录协议负责包含应用程序数据并验证其完整性和来源，对HTTP数据进行了压缩、加密和认证

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\记录协议.png)

消息认证码（MAC值，由哈希生成），防篡改、防重放（加入了片段的编码）。

记录协议完成后将报文数据交给TCP协议传输。

**HTTPS一定安全可靠吗？**

存在中间人攻击问题，但这个前提是客户端接受了中间人服务器的证书，要么是自己点击接受，要么是计算机中毒恶意导入了中间人的根证书，并非HTTPS本身的安全问题。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/http/https%E4%B8%AD%E9%97%B4%E4%BA%BA.drawio.png)

抓包工具作为中间人截取HTTPS数据的原理是在客户端系统受信任的根证书列表中导入抓包工具生成的证书，使浏览器信任自己签发的证书。

要避免中间人抓取数据，可以通过HTTPS双向认证解决。

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\双向认证.png)



**6. HTTP/1.1、HTTP/2、HTTP/3演变**

**HTTP/1.1相比HTTP/1.0提高了什么性能？**

改进：

- 长连接
- 管道化传输

HTTP/1.1瓶颈：

- 请求/响应头部未经压缩就发送
- 发送冗长首部，相同首部重复浪费
- 服务器按请求顺序响应，队头阻塞
- 没有请求优先级控制
- 请求只能从客户端开始，服务器被动响应

**HTTP/2做了什么优化？**

![HTT/1 ~ HTTP/2](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\HTTP2.png)

- 头部压缩：HPACK算法在客户端和服务器同时维护一张头信息表，所有字段都存入表，生成一个索引号，相同字段就只发索引号，消除重复部分。

- 二进制格式：不像HTTP/1.1纯文本，头信息和数据体都是二进制，称帧，头信息帧、数据帧。

  ![HTTP/1 与 HTTP/2 ](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\二进制帧.png)

- 并发传输：引入Stream，多个Stream复用一条TCP连接。Stream包含一个或多个Message，Message对应HTTP/1的请求或响应，由HTTP头部和包体构成。Message里包含一条或多个Frame，Frame是HTTP/2最小单位，以二进制压缩格式存放HTTP/1的内容。不同的HTTP请求用独一无二的Stream ID区分，不同Stream帧可乱序发送。

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\stream.png)

- 服务器推送：服务器可主动向客户端发送消息。客户端和服务器双方都可建立Stream，客户端建立的Stream为奇数号，服务器的为偶数号。

  主动推送场景如：

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\push.png)

HTTP/2的队头阻塞问题，由于TCP层必须保证收到的字节数据完整连续，若有数据晚到或丢失重传，后续数据要在内核里等待。

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\http2阻塞.jpeg)

![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\tcp队头阻塞.gif)

**HTTP/3做了哪些优化？**

HTTP/2队头阻塞是因为TCP，HTTP/3将TCP改成UDP。

![HTTP/1 ~ HTTP/3](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\HTTP3.png)



UDP不管顺序也不管丢包，没有HTTP/2队头阻塞问题。但是UDP不可靠传输，需要基于UDP的QUIC协议来实现类似TCP的可靠传输。

- 无队头阻塞：QUIC也用类似Stream，不过当某个Stream发生丢包时只阻塞这个流，其他流不受影响。

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\quic无阻塞.jpeg)

- 更快的连接建立：QUIC内部包含TLS，自己的帧会携带TLS里的“记录”，再加上QUIC使用TLS/1.3，因此仅需1个RTT就可以同时完成建立连接和密钥协商。（HTTP/1和HTTP/2中TCP和TLS分层，先TCP握手后TLS握手）

  ![TCP HTTPS（TLS/1.3） 和 QUIC HTTPS ](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\HTTP3交互次数.png)

  ![img](C:\Users\Suyb\Desktop\springWorkAndStudy\Internship\计算机网络\img\HTTP篇\HTTP3更快连接.png)

- 连接迁移：HTTP/3中QUIC协议握手确认了双方的**连接ID**（客户端和服务器各自选择一组ID标记自己），在移动设备网络变化导致IP地址变化后（TCP连接采用了IP和端口四元组），仍然能够无缝复用原连接，消除重连成本，没有卡顿达到连接迁移功能。

QUIC是一个在UDP之上的伪TCP+TLS+HTTP/2的多路复用协议。



SSL和TLS其实是同一个东西，只是标准化前后的两种称呼。



### 2.2 HTTP/1.1如何优化？

![img](.\img\HTTP篇\优化http1.1提纲.png)



**1. 如何避免发送HTTP请求？**

缓存技术：

- 强制缓存：Cache-Control、Expired
- 协商缓存：Last-Modified、Etag，缓存可继续使用则响应304

**2. 如何减少HTTP请求次数？**

**减少重定向请求次数**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/http1.1%E4%BC%98%E5%8C%96/%E5%AE%A2%E6%88%B7%E7%AB%AF%E9%87%8D%E5%AE%9A%E5%90%91.png)

中间代理服务器负责重定向工作，减少轮次

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/http1.1%E4%BC%98%E5%8C%96/%E4%BB%A3%E7%90%86%E6%9C%8D%E5%8A%A1%E5%99%A8%E9%87%8D%E5%AE%9A%E5%90%91.png)

代理服务器事先知晓重定向规则，则进一步减少

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/http1.1%E4%BC%98%E5%8C%96/%E4%BB%A3%E7%90%86%E6%9C%8D%E5%8A%A1%E5%99%A8%E9%87%8D%E5%AE%9A%E5%90%912.png)

**合并请求**

HTTP/1.1默认不开启管道模式，为防止单个请求导致队头阻塞，通常同时发起许多TCP连接请求。

类似css小图标等，多个请求合并。

![图来源于：墨染枫林的CSDN](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/http1.1%E4%BC%98%E5%8C%96/css%E7%B2%BE%E7%81%B5.png)

像一些图片二进制数据用base64编码，可以将信息嵌入URL随HTML文件发送。

**延迟发送请求**

比如网页流量时按需获取即可，不用一次性把使用资源获取。

**3. 如何减少HTTP响应的数据大小？**

**无损压缩**

适合文本文件、程序源代码、可执行文件

用霍夫曼编码。常用如gzip等压缩算法，`Accept-Encoding: gzip, deflate, br`，`Content-Encoding: gzip`

Google的Brotli算法更牛，即`br`

**有损压缩**

适合图片、视频、音频

Google的WebP格式比Png格式图片压缩算法效果好。

**4. 总结**



### 2.3 HTTPS RSA握手解析

![img](.\img\HTTP篇\https RSA握手解析提纲.png)



**1. TLS握手过程**

HTTP风险：窃听（明文）、伪装（无身份验证）、篡改（无数据校验）

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https/tls%E6%8F%A1%E6%89%8B.png)

每一个框是一个记录（record），记录是TLS收发数据的基本单位，多个记录可组成一个TCP包发送。

HTTPS应用层协议，完成TCP连接后，还有进行TLS握手，才能进行安全通信。不同密钥交换算法，TLS握手过程有些区别。

**2. RSA握手过程**

**TLS第一次握手**

Client Hello：随机数、TLS版本、支持的密钥套件列表

**TLS第二次握手**

Server Hello：随机数、确认TLS版本、选择的密钥套件（RSA）

`Cipher Suite: TLS_RSA_WITH_AES_128_GCM_SHA256`

格式【密钥交换算法+签名算法+对称加密算法+摘要算法】

- TLS和WITH间只有一个RSA，所以密钥交换算法和签名算法都用RSA
- 握手后通信用AES对称算法，密钥长度128，使用GCM分组模式
- 摘要算法SHA256

Server Certificate：数字证书

Server Hello Done

**客户端验证证书**

数字证书内容：公钥、持有者信息、CA信息、CA对证书的签名和使用算法、有效期、额外信息。

CA对证书的信息的哈希值签名，生成数字签名附在数字证书上，验证时使用CA公钥对数字签名解得一个哈希值，与证书信息的哈希值对比即可确认是否证书有效。

多层CA将根证书隔离的更好，保证安全

**TLS第三次握手**

Client Key Exchange：经服务器公钥加密的pre-master随机数

Change Cipher Spec：通知服务端后续使用加密通信

Encrypted Handshake Message（Finished）：将先前发送的数据生成摘要，并加密发送给服务器做验证，验证是否可用以及是否有被篡改过。

Change Cipher Spec以及之前都是明文通信。

**TLS第四次握手**

Change Cipher Spec

Encrypted Handshake Message

**RSA算法的缺陷**

不支持前向保密，服务器私钥泄露，之前所有通信密文都被破解。



### 2.4 HTTPS ECDHE握手解析

![img](.\img\HTTP篇\ECDHE握手协议提纲.png)

比RSA支持前向保密

**1. 离散对数**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https/%E7%A6%BB%E6%95%A3%E5%AF%B9%E6%95%B0.png)

底数a和模数p是离散对数的公共参数，b是真数，i是对数，知道对数易推对数，但知道真数很难推对数。

但模数p是很大的质数时，知道a和b，现有的计算机也几乎无法算出对数。

**2. DH算法**

非对称加密算法，用于交换密钥，核心思想就是离散对数

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https/dh%E7%AE%97%E6%B3%95.png)

**3. DHE算法**

DH算法分

- static DH算法：弃用，一方私钥固定，通过大量密钥协商，可能暴力推导出服务器私钥，不具备前向安全性。
- DHE算法：每次协商，双方私钥都是随机生成的、临时的。E即ephemeral（临时的）

**4. ECDHE算法**

ECDHE是在DHE算法基础上利用ECC椭圆曲线特性，减少计算公钥的计算量。

**5. ECDHE握手过程**

ECDHE比RSA握手过程省一个消息往返时间，在第四次握手前客户端就可以发送加密的HTTP数据。即`TLS False Start`，和`TCP Fast Open`相像。

**TLS第一次握手**

Client Hello：随机数、TLS版本号、密码套件列表

**TLS第二次握手**

Server Hello：随机数、TLS版本、选择的密钥套件

`TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384`

签名算法RSA用在验证

Certificate：证书

Server Key Exchange：x25519椭圆曲线、生成随机数算出的公钥（经RSA签名确保不被篡改）

Server Hello Done

**TLS第三次握手**

Client Key Exchange：客户端生成随机数计算出的公钥

计算会话密钥：使用客户端随机数、服务端随机数、ECDHE算出的共享密钥生成

Change Cipher Spec：后续使用对称算法加密通信

Encrypted Handshake Message：之前发送数据的摘要经对称密钥加密，给服务器做验证

**TLS第四次握手**

Change Cipher Spec

Encrypted Handshake Message



ECDHE比RSA在TLS握手第二轮中多了Server Key Exchange。

### 2.5 HTTPS如何优化？

![img](.\img\HTTP篇\优化https提纲.png)



**1. 分析性能损耗**

- TLS协议握手过程：网络时延2RTT，ECDHE密钥协商算法握手中生成椭圆曲线公私钥，客户端验证证书时访问CA获取CRL或OCSP，双方计算Pre-Master生成对称密钥。

  ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/tls%E6%80%A7%E8%83%BD%E6%8D%9F%E8%80%97.png)

- 握手后对称加密通信：AES、ChaCha20等主流对称加密算法，硬件级别优化。

**2. 硬件优化**

HTTPS协议是计算密集型，而非I/O密集型，要花在CPU上。

选择支持AES-NI的CPU，能在指令级别优化AES算法。

**3. 软件优化**

- 软件升级：对于成百上千服务器的升级也会花费大量人力物力。
- 协议优化

**4. 协议优化**

针对密钥交换过程进行优化

**密钥交换算法优化**

RSA握手需要2RTT，且不具备前向安全性，所以尽量选用ECDHE密钥交换算法。

ECDHE有`False Start`抢跑，在第3次握手后，第4次握手前可以发送加密数据，将握手减少到1RTT，且具备前向安全性。

尽量用x25519椭圆曲线，Nginx上`ssl_ecdh_curve x25519:secp384r1;`配置需要使用的椭圆曲线，要优先的放前。

对称加密算法上如果对安全性要求不高，可以使用AES_128_GCM，密钥长度短一些，快一些。

Nginx上用`ssl_ciphers`配置想要使用的密钥套件。优先的放前，`ssl_ciphers 'EECDH+ECDSA+AES128+SHA:RSA+AES128+SHA';`

**TLS升级**

TLS1.3只需要1RTT，将Hello和公钥交换合并。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/tls1.2and1.3.png)

另外密钥交换算法上，废除不支持前向安全的RSA和DH算法，只支持ECDHE。对称加密和签名算法也只支持目前最安全的几个密码套件。废除古老不安全的密码套件防止中间人降级攻击导致被破密。

**5. 证书优化**

**证书传输优化**

减小证书大小，节约带宽，减少客户端运算量。使用椭圆曲线（ECDSA）证书，相同安全强度下，ECC要比RSA的密钥长度短。

**证书验证优化**

证书链逐级验证，验证过程包括用CA公钥解密证书、签名算法验证证书完整性，还要确认证书是否被CA吊销，客户端需要HTTP访问CA获取CRL或OCSP来确认证书有效性，这会存在一系列网络通信开销，如DNS查询、建立连接、收发数据。

CRL（Certificate Revocation List）存在问题：

- 实时性较差：CA定期更新吊销的证书
- 下载速度会慢：随着列表越来越大，下载和遍历都会很慢

OCSP在线证书状态协议，向CA查询证书有效状态。

OCSP Stapling：由服务器向CA周期性查询证书状态，客户端TLS握手时附带给。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/opscp-stapling.png)

**6. 会话复用**

将先前TLS握手协商出来的对称密钥接着复用。

**Session ID**

握手后内存缓存会话密钥，以键值对存储，key为Session ID，value为会话密钥。

再次连接时客户端带上Session ID，只需要1RTT就能建立安全通信。不过会话密钥会定期失效。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/sessionid.png)

缺点

- 客户端增多，内存保存会话密钥，服务器内存压力大
- 多服务器负载均衡提供服务时，客户端不一定命中上次访问的服务器，重新握手

**Session Ticket**

解决Session ID的问题，服务器不再缓存会话密钥，将缓存工作交给客户端，类似HTTP的Cookie，即Session Ticket。

连接时，服务器加密会话密钥作为Ticket给客户端，由客户端存储。再次连接时客户端发送Ticket，服务器解密并检查有效期，通过即可恢复会话。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/ticket.png)

Session ID和Session Ticket不具备前向安全性，应对重放攻击也困难。避免重放攻击的方式需要给会话密钥设置合理过期时间。

**Pre-shared Key**

重连TLS1.3只需要0RTT，客户端将Ticket和HTTP请求一同发送服务器

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/0-RTT.png" alt="img" style="zoom:50%;" />

Pre-shared Key也有重放攻击的危险，要设置合理会话密钥过期时间，且只针对安全的HTTP请求如GET/HEAD使用会话重用。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/https%E4%BC%98%E5%8C%96/0-rtt-attack.png)

**7. 总结**



### 2.6 HTTP/2牛逼在哪？

![img](.\img\HTTP篇\HTTP2优势提纲.png)



**1. HTTP/1.1协议的性能问题**

HTTP/1.1高延迟问题

- 延迟难以下降
- 并发连接有限
- 队头阻塞问题
- HTTP头部巨大且重复
- 不支持服务器推送消息

存在协议难以优化的地方，如请求-响应模型、头部巨大且重复、并发连接耗时、服务器不能主动推送等。

**2. 兼容HTTP/1.1**

首先HTTP/2没在URI引入新协议名，还是`http://`和`https://`，协议升级用户是意识不到的。

其次只在应用层做改变，还是基于TCP，且保持功能上的兼容（HTTP语义（请求方法、状态码等等规则）不变，只改动语法）

**3. 头部压缩**

HTTP/1.1中Body可用头字段`Content-Encoding`指定压缩方式，但Header没有专门的优化手段。

Header存在问题：

- 含很多固定字段，有必要压缩
- 大量请求和响应存在重复字段，需要避免重复性
- 字段ASCII编码，虽易于人类观察，但效率低，有必要二进制编码

HTTP/2采用HPACK算法压缩头部，由静态字典、动态字典和Huffman编码组成。字典中用长度较小的索引号表示重复的字符串，再用Huffman编码压缩数据。

**静态表编码**

<img src="https://cdn.xiaolincoding.com//picgo/image-20240105142818571.png" alt="img" style="zoom:50%;" />

![image-20240105142857712](https://cdn.xiaolincoding.com//picgo/image-20240105142857712.png)

![image-20240105142949762](https://cdn.xiaolincoding.com//picgo/image-20240105142949762.png)

<img src="https://cdn.xiaolincoding.com//picgo/image-20240105142917193.png" alt="img" style="zoom:50%;" />

**动态表编码**

静态表只包含61种高频出现在头部的字符串，其他的字符串要自行构建动态表。

第一次发送时添加入动态表，下次发送就从动态表获取。

动态表生效前提：同一个连接、重复传输完全相同的HTTP头部。如果在一个连接上只发送1次，或重复传输时字段总略有变化则无法充分利用动态表。

动态表会不断增大，占用内存，所以提供`http2_max_requests`限制一个连接传输的请求数量，达到上限后关闭连接释放内存。

![image-20240105143006681](https://cdn.xiaolincoding.com//picgo/image-20240105143006681.png)

**4. 二进制帧**

响应报文分为两类帧，HEADERS和DATA，并用二进制编码。

HEADERS帧：

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/http/index.png)

- 最前的1表示静态表中存在KV
- Index是在静态表中的索引

二进制帧结构

![image-20240105143208962](https://cdn.xiaolincoding.com//picgo/image-20240105143208962.png)

![image-20240105143150947](https://cdn.xiaolincoding.com//picgo/image-20240105143150947.png)

标志位可携带简单的控制信息，最后流标识符标识帧属于哪个Stream。帧数据是经过HPACK算法压缩过的HTTP头部和包体。

**5. 并发传输**

多个Stream复用一条TCP连接达到并发效果。

![image-20240105143224839](https://cdn.xiaolincoding.com//picgo/image-20240105143224839.png)

Stream：一个TCP连接可以有多个Stream

Message：一个Stream里可以有多个Message，Message对应HTTP/1的请求或响应，由HTTP头部和包体构成

Frame：一个Message里可以有多个Frame，Frame是HTTP/2最小单元，以二进制压缩格式存放HTTP/1内容。

不同Stream的帧可以乱序发送，所以可以并发。但一个Stream内的帧必须有序。

双方都可以建立Stream，客户端建立的Stream必须为奇数号，服务器的为偶数号。

同一个连接的Stream ID不能复用，只能顺序递增，当耗尽时发控制帧GOAWAY关闭TCP连接。Nginx中使用`http2_max_concurrent_Streams`配置Stream上限，默认128。

多个Stream复用TCP连接，减少TCP的连接数目，省去大量TCP握手、慢启动（拥塞控制逐步试探带宽）、TLS握手的过程。

**6. 服务器主动推送资源**

Nginx中如下设置，使客户端访问/test.html时，主动推送/test.css

```
location /test.html {
	http2_push /test.css
}
```

客户端发起奇数号Stream，服务器主动推送使用偶数号。通过PUSH_PROMISE帧传输HTTP**头部**，并用帧中的Promised Stream ID字段告知客户端，接下来在哪个Stream中发送**包体**。两个Stream是并发的。

<img src="https://cdn.xiaolincoding.com//picgo/image-20240105143338707.png" alt="image-20240105143338707" style="zoom:50%;" />

**7. 总结**

...

HTTP/2还是存在“对头阻塞”，基于TCP，而TCP是字节流协议，必须保证字节数据完整连续，内核才会将缓冲区数据返回应用层。HTTP/3解决了这个问题。

### 2.7 HTTP/3强势来袭

![image-20240105144308801](.\img\HTTP篇\HTTP3强势来袭提纲.png)



**1. 美中不足的HTTP/2**

**队头阻塞**

**TCP与TLS的握手时延迟**

**网络迁移需要重新连接**

![img](https://cdn.xiaolincoding.com//picgo/image-20240201140954238.png)

**2. QUIC协议的特点**

UDP简单、不可靠，UDP包间无序，UDP不需要连接，无握手挥手。

HTTP/3将传输协议换为UDP，基于此在应用层实现QUIC协议，具有类似TCP的连接管理、拥塞窗口、流量控制的网络特性，确保数据可靠。

**无队头阻塞**

HTTP/3某一流数据丢失只会影响到该流。

![image-20240105144623562](https://cdn.xiaolincoding.com//picgo/image-20240105144623562.png)

**更快的连接建立**

HTTP/1和HTTP/2协议中，TCP和TLS是分层的，分别属于内核实现的传输层、OpenSSL库实现的表示层，难以合并在一起，只能分批次握手。

QUIC协议握手只要1RTT，握手确认双方连接ID，连接迁移也是基于连接ID实现。

QUIC内部包含TLS，帧中携带TLS记录，另外使用TLS1.3，只需要1RTT同时完成建立连接和密钥协商。再次连接时只用0RTT。

**连接迁移**

使用连接ID替换四元组来绑定连接，不会因为网络迁移受影响。

**3. HTTP/3协议**

HTTP/3同HTTP/2采用同样二进制帧结构，但HTTP/2中要定义Stream，HTTP/3不用，直接使用QUIC里的Stream。

![image-20240105144457456](https://cdn.xiaolincoding.com//picgo/image-20240105144457456.png)

HTTP/3分数据帧和控制帧，Headers帧（HTTP头部）和DATA帧（HTTP包体）属于数据帧。

HTTP/3头部压缩采用QPACK算法，静态表相比HTTP/2的扩大到91项。HTTP/2中动态表具有时序性，如果首次出现的请求丢包，后续请求就无法解出HPACK头部（双方的动态表不一致了）。

QPACK中使用两个特殊的单向流来同步双方的动态表。

- QPACK Encoder Stream：将一个字典kv传给对方
- QPACK Decoder Stream：响应对方，说明收到字典并更新入了本地动态表，后续使用该字典来编码。

**4. 总结**





### 2.8 既然有HTTP协议，为什么还要有RPC？

**1. 从TCP聊起**

TCP特点：面向连接、可靠、基于字节流

基于字节流，无分界，存在粘包问题。需要定义规则区分消息边界，因此有了HTTP和RPC等协议，消息由消息头和消息体组成，消息体中说明了包长度等信息。

**2. HTTP和RPC**

同属于应用层协议，HTTP超文本传输协议，RPC远程方法调用。

RPC调用远程服务器暴露的方法，屏蔽网络细节，像本地调用函数一样。如gRPC、thrift协议。底层不一定非要TCP，还可以是UDP、HTTP。

TCP在70年代的协议，RPC是80年代，HTTP到90年代才流行。问题应该是“为什么有RPC还要HTTP”。

早期联网软件像xx卫士、xx管家，作为客户端要与服务端建立连接，因为都是自家软件，就用自家造的RPC协议。但浏览器出现后，不仅需要连接自家公司服务器，还要访问其他公司网站服务器，所以需要一个**统一标准**，于是HTTP就出来了。

HTTP主要用于B/S（Browser/Server）架构，RPC更多用于C/S架构，不过两个架构在慢慢融合。RPC多用于公司内部集群或各个微服务间的通讯。

**3. HTTP和RPC有什么区别**

**服务发现**

建立连接需要知道IP和端口，找到IP端口的过程就是服务发现。

HTTP中用DNS服务解析域名得到IP地址，默认80端口。

RPC中则向一些专门保存服务名和IP信息的中间服务去获取。DNS也是中间服务的一种。

**底层连接形式**

HTTP/1.1中TCP会保存长连接（Keep Alive），后续请求响应都复用。

RPC类似HTTP，使用TCP长连接，但还会建一个连接池，请求量大时建立多条连接放在池内，要发数据就从池中取连接来发送，用完放回去下次再复用。

连接池有助于提升网络请求性能，Go语言中也给HTTP加个连接池。

**传输的内容**

传输信息无非消息头和消息体，对于字符串或数字都能编码或直接转为二进制，但结构体就需要进行序列化转二进制，并反序列化恢复。像HTTP/1.1中body传输结构体时用json序列化，存在很多冗余。

RPC定制化程度高，可以采用体积更小的Protobuf或其他序列化协议保存结构体数据，并且不需要考虑HTTP一些浏览器行为，性能更好一些。因此像公司内部微服务都抛弃HTTP采用RPC。

![HTTP 原理](https://cdn.xiaolincoding.com//mysql/other/f4cef7331cabcfe56d9d6434f7ef907f.png)

![RPC 原理](https://cdn.xiaolincoding.com//mysql/other/12244fb0b19b2e61755fcab799198f68.png)

上图HTTP特指HTTP/1.1，HTTP/2性能可能要比RPC好，甚至gRPC底层直接用HTTP/2。

为什么不直接用HTTP/2，是因为其2015才出来，很多公司内部已跑RPC很多年。

**4. 总结**




### 2.9 既然有HTTP协议，为什么还要有WebSocket？

**1. 使用HTTP不断轮询**

如扫码登录

问题

- 大量轮询请求，消耗带宽
- 扫码后等1-2s才下一次请求页面跳转，明显卡顿

**2. 长轮询**

请求超时时间设置很大，较长时间内等待服务器响应

![图片](https://cdn.xiaolincoding.com//mysql/other/1058a96ba35215c0f30accc3ff5bb824.png)

服务器推送技术，comet。上述的不断轮询和长轮询都是客户端主动请求。

**3. WebSocket是什么**

HTTP/1.1同一时间里只有一方可以主动发数据，是半双工。早期看网页文本场景没有考虑到双方大量互相发送数据。

于是WebSocket应运而生，也是基于TCP的应用层协议。

**怎么建立WebSocket连接**

浏览器上为兼容刷图文的HTTP和网页游戏的WebSocket，在TCP三次握手后使用HTTP协议先通信一次，如果是普通HTTP请求则继续使用HTTP协议，如果HTTP请求带了特殊header头，则建立WebSocket连接

```
Connection: Upgrade
Upgrade: WebSocket
Sec-WebSocket-Key: T2a6wZlAwhgQNqruZ2YUyg==\r\n
```

升级协议为WebSocket协议，随机生成的base64码Sec-WebSocket-Key发给服务器。服务器支持升级成WebSocket协议则使用公开算法将base64码转换，后放在HTTP响应的Sec-WebSocket-Accept里，同时带上101状态码（协议切换）。浏览器也会使用公开算法转换base64码和响应的对比，相同则验证通过。

```
HTTP/1.1 101 Switching Protocols\r\n
Sec-WebSocket-Accept: iBJKv/ALIW2DobfoA4dmr3JHBCY=\r\n
Upgrade: WebSocket\r\n
Connection: Upgrade\r\n
```

![图片](https://cdn.xiaolincoding.com//mysql/other/f4edd3018914fe6eb38fad6aa3fd2d65.png)

**WebSocket消息格式**

数据包在WebSocket叫帧

![图片](https://cdn.xiaolincoding.com//mysql/other/3a63a86e5d7e72a37b9828fc6e65c21f.png)

**WebSocket使用场景**

WebSocket继承了TCP协议的全双工能力，而且解决粘包

应用场景如网页/小程序游戏、网页聊天室、网页协同办公软件等。

**4. 总结**





## 三、TCP篇

### 3.1 TCP三次握手与四次挥手面试题

![img](.\img\TCP篇\TCP三次握手与四次挥手面试题提纲.png)

**1. TCP基本认识**

**TCP头格式有哪些？**

![TCP 头格式](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230534096.png)

序列号：建立连接时生成随机数作为初始值，通过SYN包传给接收端主机，每发送一次数据就累加一次该数据字节数的大小。**解决网络包乱序问题**。

确认应答号：下次期望收到的数据的序列号，发送端收到确认应答后可以认为序号前的数据都被正确接收。**解决丢包问题**。

控制位

- ACK：为1则确认应答的字段有效，出初次连接的SYN包外都应该为1
- RST：出现异常强制端开连接
- SYN：建立连接
- FIN：断开连接

**为什么需要TCP协议？TCP工作在哪一层？**

IP层不可靠，TCP是工作在传输层的可靠数据传输的服务，确保网络包无损坏、无间隔、非冗余和按序的。

**什么是TCP？**

TCP面向连接、可靠、基于字节流

**什么是TCP连接？**

用于保证可靠性和流量控制维护的某些状态信息，这些信息的组合，包括Socket、序列号和窗口大小称为连接。

- Socket：IP地址和端口号
- 序列号：解决乱序问题
- 窗口大小：流量控制

**如何唯一确定一个TCP连接？**

四元组：源地址、源端口、目的地址、目的端口

理论上服务器监听某个端口时，最大TCP连接数是客户端IP数*客户端端口数

实际上最大并发TCP连接数受以下因素影响

- 文件描述符限制
  - 系统级
  - 用户级
  - 进程级
- 内存限制

**UDP和TCP有什么区别呢？分别的应用场景？**

UDP无复杂控制机制，利用IP提供面向无连接的通信服务。

![UDP 头部格式](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230439961.png)

TCP和UDP区别：

- 连接
- 服务对象：TCP一对一的两点服务，UDP还支持一对多、多对多
- 可靠性
- 拥塞控制、流量控制
- 首部开销：TCP无选项字段就已经要20字节了，UDP只有固定8字节
- 传输方式：TCP流式传输无边界但保证顺序可靠，UDP包传输有边界，但会丢包乱序
- 分片不同：TCP数据大于MSS会在传输层分片，UDP大于MTU会在IP分片

应用场景

- TCP：FTP、HTTP/HTTPS
- UDP：包总量较少的通信如DNS、SNMP，视频、音频等多媒体通信，广播通信

UDP头部长度固定，TCP有可变长的选项字段，因此多“首部长度”字段。

TCP数据长度=IP总长度-IP首部长度-TCP首部长度，因此不用“包长度”字段。UDP也可以如此计算，但有一些原因导致添加了该字段，例如首部对齐4字节、以前网络层可能不是IP协议。

**TCP和UDP可以使用同一个端口吗？**

可以。TCP和UDP是内核中两个独立的软件模块，IP包头的“协议号”区分了TCP/UDP

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/port/tcp%E5%92%8Cudp%E6%A8%A1%E5%9D%97.jpeg)

**2. TCP连接建立**

**TCP三次握手过程是怎样的？**

![TCP 三次握手](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4/%E7%BD%91%E7%BB%9C/TCP%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B.drawio.png)

![第一个报文 —— SYN 报文](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230500953.png)

![第二个报文 —— SYN + ACK 报文](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230504118.png)

![第三个报文 —— ACK 报文](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230508297.png)

第三次握手可以携带数据。

**如何在Linux系统查看TCP状态？**

`netstat -napt`

![TCP 连接状态查看](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230520683.png)

**为什么是三次握手？不是两次、四次？**

- 三次握手才能阻止历史连接对当前连接的初始化（主要原因）

  ![三次握手避免历史连接](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230525514.png)

  两次握手则浪费资源，服务器建立完连接后又因为RST释放资源

  ![两次握手无法阻止历史连接](https://cdn.xiaolincoding.com//mysql/other/fe898053d2e93abac950b1637645943f.png)

- 三次握手才能同步双方的初始序列号

  序列号作用：

  - 去重
  - 按序接收
  - 标识已收包（ACK报文中序列号）

  ![四次握手与三次握手](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230639121.png)

  两次握手只能同步一方的初始序列号。

- 三次握手才能避免资源浪费

  ![两次握手会造成资源浪费](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230636571.png)

**为什么每次建立TCP连接时，初始化的序列号都要求不一样呢？**

- 防止历史报文被下一个相同四元组连接接收（主要）

  ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/isn%E7%9B%B8%E5%90%8C.png)

  ![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/isn%E4%B8%8D%E7%9B%B8%E5%90%8C.png)

- 防止黑客伪造相同序列号TCP报文被对方接收

**既然IP层会分片，为什么TCP层还需要MSS呢？**

![MTU 与 MSS](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230633447.png)

如果一个IP分片丢失（接收方组装不成完整TCP报文，不会响应ACK，发送方会超时重传），整个IP报文的分片都得重传。

**第一次握手丢失，会发生什么？**

重传SYN报文，序列号一样。超时时间写死在内核，重传次数由`tcp_syn_retries`内核参数控制，默认5，`/proc/sys/net/ipv4/tcp_syn_retries`配置。每次重传超时时间翻倍。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC1%E6%AC%A1%E6%8F%A1%E6%89%8B%E4%B8%A2%E5%A4%B1.png)

**第二次握手丢失，会发生什么？**

除了客户端会超时重传，服务端由于发起了SYN连接，也会超时重传，最大重传次数由`tcp_synack_retries`决定，默认5。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC2%E6%AC%A1%E6%8F%A1%E6%89%8B%E4%B8%A2%E5%A4%B1.png)

**第三次握手丢失，会发生什么？**

只有服务端超时重传SYN-ACK报文。ACK报文不会重传，丢失了得由对方重传对应报文。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B%E4%B8%A2%E5%A4%B1.drawio.png)

**什么是SYN攻击？如何避免SYN攻击？**

![正常流程](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230622886.png)

队列满了就会丢弃报文，SYN攻击会构造不同的IP和服务器连接但不响应ACK，导致SYN队列满，而后SYN报文都被丢弃。

半连接队列（SYN队列），全连接队列（Accept队列）

避免SYN攻击方法

- 调大netdev_max_backlog：网卡接收数据包速度大于内核处理速度时，使用一个队列保存数据包，该传输控制其队列大小。`net.core.netdev_max_backlog = 10000`

- 增大TCP半连接队列：同时增大`net.ipv4.tcp_max_syn_backlog`、listen()函数的`backlog`、`net.core.somaxconn`

- 开启net.ipv4.tcp_syncookies：SYN半连接队列满时绕过SYN队列发送cookie给客户端，之后验证客户端的ACK即可。为0关闭，为2无条件开启，为1则SYN队列放不下再启用。

  ![tcp_syncookies 应对 SYN 攻击](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230618804.png)

- 减少SYN+ACK重传次数：更快断开处于SYN_REVC状态的TCP连接

**3. TCP连接断开**

**TCP四次挥手过程是怎样的？**

![客户端主动关闭连接 —— TCP 四次挥手](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230614791.png)

CLOSED_WAIT状态中，服务器处理完数据才发FIN报文。

主动关闭连接的，才有TIME_WAIT状态。

**为什么挥手需要四次？**

关闭连接时客户端发送FIN，仅表示客户端不再发送数据但仍能接收。

服务端收到客户端FIN后回复ACK，但服务端可能还有数据要处理和发送，服务端不再发送数据时才发送FIN给客户端，表示同意关闭连接。因此服务器的FIN和ACK分开发送。

**第一次挥手丢失了，会发生什么？**

超时重传FIN报文，由`tcp_orphan_retries`控制重传次数，如果都没收到第二次挥手则直接进入CLOSE状态。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC%E4%B8%80%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%A2%E5%A4%B1.png)

**第二次挥手丢失了，会发生什么？**

客户端超时重传，服务端发送ACK报文是不会重传的。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC%E4%BA%8C%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%A2%E5%A4%B1.png)

ACK若是接收到，进入FIN_WAIT_2状态，但处于此状态超过`tcp_fin_timeout`指定的时间则直接进入CLOSE状态，默认60s。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/fin_wait_2.drawio.png)

若是主动关闭方用shutdown函数关闭连接只关闭发送方向，则尽管一直没收到对方的FIN报文，也将一直处于FIN_WAIT2状态。（`tcp_fin_timeout`无法控制shutdown关闭的连接）

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/fin_wait_2%E6%AD%BB%E7%AD%89.drawio.png)

**第三次挥手丢失了，会发生什么？**

服务端进入CLOSE_WAIT状态等待应用进程调用close函数关闭连接，并发送FIN报文。

没有收到ACK会超时重传FIN报文，也由`tcp_orphan_retries`参数控制，

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC%E4%B8%89%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%A2%E5%A4%B1.drawio.png)

**第四次挥手丢失了，会发生什么？**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/%E7%AC%AC%E5%9B%9B%E6%AC%A1%E6%8C%A5%E6%89%8B%E4%B8%A2%E5%A4%B1drawio.drawio.png)

**为什么TIME_WAIT等待时间是2MSL？**

MSL最大报文生命周期，TTL为IP头一个字段，指明IP数据报可经过的最大路由数。MSL单位是时间，TTL是经过路由跳数，MSL应该大于等于TTL消耗为0的时间。

TTL一般为64，Linux将MSL设为30s。

TIME_WAIT等待2MSL，至少允许报文丢失一次，ACK在第一个MSL内丢失，则重传的FIN会在第二个MSL内到达。

如果TIME_WAIT状态下收到重发的FIN，2MSL将重新计时，默认60s。定义在内核代码，`TCP_TIMEWAIT_LEN`

**为什么需要TIME_WAIT状态？**

主动发起关闭连接一方才有TIME_WAIT状态。

原因：

- 防止历史连接中数据被后续相同四元组连接错误接收

  序列化SEQ、初始序列号ISN，32位无符号数，到4G循环回0。初始序列号数值每4us加1，循环一次4.55h。

  ![TIME-WAIT 时间过短，收到旧连接的数据报文](https://cdn.xiaolincoding.com//mysql/other/6385cc99500b01ba2ef288c27523c1e7-20230309230608128.png)

  2MSL足以让两个方向上的数据包都被丢弃。

- 保证被动关闭连接一方正确关闭

  ![TIME-WAIT 时间过短，没有确保连接正常关闭](https://cdn.xiaolincoding.com//mysql/other/3a81c23ce57c27cf63fc2b77e34de0ab-20230309230604522.png)

**TIME_WAIT过多有什么危害？**

- 占用系统资源：服务端作为主动发起关闭连接方时，TIME_WAIT过多导致TCP连接过多，占用系统资源，如文件描述符、内存资源、CPU资源、线程资源等。
- 占用端口资源：客户端作为主动发起关闭连接方时，要和同一个目的IP和目的端口的服务端连接时，当TIME_WAIT状态过多，会占满端口。一般开启的端口为32768-61000。

**如何优化TIME_WAIT？**

- net.ipv4.tcp_tw_reuse和tcp_timestamps：tcp_tw_reuse只用于客户端，调用connect函数时内核随机找一个time_wait状态超过1s的连接给新连接复用。同时需要开启tcp_timestamps，在TCP头部选项字段中有8个字节表示时间戳，前4字节保存发送该数据包时间，后4字节保存最近一次接收对方发送到达数据的时间。有了时间戳就不存在2MSL问题了，重复数据因为时间戳过期而丢弃。
- net.ipv4.tcp_max_tw_buckets：默认18000，系统中处于TIME_WAIT状态连接超过该数，系统将后面TIME_WAIT连接状态重置。
- 程序中使用SO_LINGER：so_linger.l_onoff=1，so_linger.l_linger=0，则调用close后立即发送RST直接断开连接。

服务端要避免过多的TIME_WAIT状态的连接，就不要主动断开连接，由客户端去承受。

**服务器出现大量TIME_WAIT状态的原因有哪些？**

服务器主动断开很多连接

- HTTP没有使用长连接：大多数Web服务实现中无论哪一方禁用HTTP Keep-Alive，都由服务端主动关闭连接。

- HTTP长连接超时：keepalive_timeout默认60s，超时触发回调函数关闭连接。

  ![HTTP 长连接超时](https://cdn.xiaolincoding.com//mysql/other/7e995ecb2e42941342f97256707496c9.png)

- HTTP长连接请求数量达到上限：nginx的keepalive_requests参数定义一个HTTP长连接最大的请求数量，超过则主动关闭。

**服务器出现大量CLOSE_WAIT状态的原因有哪些？**

被动关闭方没有调用close函数从CLOSE_WAIT状态转入LAST_ACK状态。

普通TCP服务端流程：

- 创建服务端socket，bind绑定端口，listen监听端口
- 将服务端socket注册到epoll
- epoll_wait等待连接到来，连接到来时调用accept获取已连接的socket
- 将已连接socket注册到epoll
- epoll_wait等待事件发生
- 对方连接关闭时，我方调用close

原因一：socket没注册到epoll，不过几乎不可能，属于代码bug

原因二：新连接到来时没有调用accept获取连接socket，服务端代码卡在accept函数前

原因三：没有将已连接socket注册到epoll

原因四：代码漏洞没有调用close函数，或代码卡在close函数之前。

基本都是代码问题。

**如果已经建立了连接，但是客户端突然出现故障了怎么办？**

如果服务端一直不发送数据则将感知不到客户端宕机，为此TCP保活机制经过一定时间间隔发送探测报文，连续没有响应则认为连接死亡，返回错误给应用。不过默认情况要经过2个多小时才会发现死亡连接。并且保活机制需要设置`SO_KEEPALIVE`才开启。

![img](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230552557.png)

TCP保活开启下需要考虑几种情况

- 正常响应探测报文，则TCP保活时间重置
- 宕机后重启，但由于没有连接的有效信息，产生RST报文重置连接
- 宕机或其他原因导致报文不可达，连接断开

TCP保活时间较长，为此可以在应用层实现心跳机制，如一般web服务软件提供`keepalive_timeout`参数指定HTTP长连接超时时间。

**如果已经建立了连接，但是服务端的进程崩溃会发生什么？**

进程崩溃，内核会回收该进程所有TCP连接资源，发起FIN报文断开连接。

**4. Socket编程**

**针对TCP应该如何Socket编程？**

![基于 TCP 协议的客户端和服务端工作](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230545997.png)

其中监听的socket和用来传送数据的socket是两个socket，监听socket和已完成连接socket。

**listen时候参数backlog的意义？**

![ SYN 队列 与 Accpet 队列 ](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230542373.png)

实际上限是内核参数somaxconn，即accept队列长度=min(backlog, somaxconn)

**accept发生在三次握手的哪一步？**

客户端connect成功返回在第二次握手，服务端accept成功返回在第三次握手。

![socket 三次握手](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4/%E7%BD%91%E7%BB%9C/socket%E4%B8%89%E6%AC%A1%E6%8F%A1%E6%89%8B.drawio.png)

**客户端调用close了，连接是断开的流程是什么？**

EOF放在已排队等候的其他已接收的数据之后，EOF表明连接再无额外数据到达。

![客户端调用 close 过程](https://cdn.xiaolincoding.com//mysql/other/format,png-20230309230538308.png)

**没有accept，能建立TCP连接吗？**

可以，accept只是从全连接队列取一个已连接的socket。

![半连接队列与全连接队列](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8D%8A%E8%BF%9E%E6%8E%A5%E5%92%8C%E5%85%A8%E8%BF%9E%E6%8E%A5/3.jpg)

**没有listen，能建立TCP连接吗？**

可以，客户端进行TCP自连接，或两个客户端同时向对方发出请求建立连接。



### 3.2 TCP重传、滑动窗口、流量控制、拥塞控制

![img](.\img\TCP篇\TCP可靠性的保证机制提纲.jpg)



**1. 重传机制**

**超时重传**

![超时重传的两种情况](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/5.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

RTT（Round-Trip Time往返时间）

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/6.jpg?" alt="RTT" style="zoom:50%;" />

RTO（Retransmission Timeout超时重传时间），过短过长都不行

![超时时间较长与较短](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/7.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/8.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0" alt="RTO 应略大于 RTT" style="zoom:50%;" />

RFC6289建议

![RFC6289 建议的 RTO 计算 ](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/9.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

SRTT是计算平滑的RTT，DevRTT是计算平滑的RTT与最新RTT的差距

Linux下**α = 0.125，β = 0.25， μ = 1，∂ = 4**

超时重发后再次超时，超时间隔翻倍。

**快速重传**

Fast Retransmit，不以时间为驱动，而是以数据驱动重传。

![快速重传机制](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/10.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

虽然解决了超时时间等待的问题，但仍有一个问题，是重传Seq2报文，还是重传所有报文，这并不好确定，因此有了SACK方法。

**SACK方法**

Selective Acknowledgement，选择性确认。

![选择性确认](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/11.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

TCP头部选项字段有SACK，将已收到数据信息发送给发送方，因此根据SACK可只重传丢失数据。SACK需要双方都支持，`net.ipv4.tcp_sack`参数打开功能。

**Duplicate SACK**

D-SACK，使用SACK告诉发送方哪些数据被重复接收。

- ACK丢包

![ACK 丢包](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/12.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)



- 网络延时

![网络延时](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/13.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

D-SACK好处

- 可以让发送方知道发出的包丢了，还是接收方回应的ACK包丢了
- 可以知道是不是发送方的数据包被网络延迟了
- 可以知道网络是不是把发送方的数据包复制了

**2. 滑动窗口**

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/14.jpg?" alt="按数据包进行确认应答" style="zoom:50%;" />

为此引入窗口，窗口大小指无需等待确认应答也可以继续发送数据的最大值。

窗口实现是在操作系统开辟一个缓存空间，在等到确认应答前在缓冲区保留已发送数据，按期收到确认应答后从缓冲区清除。

假设窗口大小为3个TCP段，则发送方可连续发送3个TCP段。如果有确认应答报文丢失，下一个确认应答能够**累计确认/累计应答**

![用滑动窗口方式并行处理](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/15.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

TCP头Window字段即窗口大小，由接收方告诉发送方自己还有多少缓冲区可以接收数据，发送方按接收方处理能力来发送数据。

发送方的窗口如下

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/16.jpg?)

TCP滑动窗口使用三个指针来跟踪四个传输类别中每一类的字节

- SND.WND：发送窗口大小（由接收方指定）
- SND.UNA（Send Unacknowledged）：绝对指针，指向已发送但未收到确认的第一个字节的序列号
- SND.NXT：绝对指针，指向未发送但可发送范围的第一个字节的序列号
- 指向#4的第一个字节是相对指针，SND.UNA+SND.WND

可用窗口大小=SND.WND-(SND.NXT-SND.UNA)

![SND.WND、SND.UN、SND.NXT](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/19.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

接收方的窗口如下

![接收窗口](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/20.jpg)

三个接收部分用两个指针划分

- RCV.WND：接收窗口大小
- RCV.NXT：指向期望从发送方发送来的下一个数据字节的序列号
- 指向#4的第一个字节是相对指针，RCV.NXT+RCV.WND

接收窗口大小约等于发送窗口大小，因为接收方在处理完数据后，新的接收窗口大小需要通过TCP报文Windows字段告诉发送方，存在时延。

**3. 流量控制**

窗口固定大小的情况

![流量控制](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/21.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

**操作系统缓冲区与滑动窗口的关系**

实际上，窗口存放的字节数存放在操作系统内存缓冲区中，会被操作系统调整。

例子1：应用进程没及时读取缓冲区内容时会影响缓冲区，从而影响窗口大小。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/22.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

例子2：系统资源紧张时操作系统可能直接减少接收缓冲区大小，这是如果应用程序没法及时读取缓存数据，就会导致严重的数据包丢失问题。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/23.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

为防止该情况发生，TCP规定不允许同时减少缓存又收缩窗口，而是先收缩窗口，过段时间再减少缓存，从而避免丢包。

**窗口关闭**

如果窗口大小为0，就会阻止发送方给接收方传递数据，直到窗口变为非0，即窗口关闭。

窗口关闭潜在危险：窗口关闭之后，接收方处理完数据向发送方通告窗口非0的ACK报文，但报文丢失。

![窗口关闭潜在的危险](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/24.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

解决办法：收到零窗口通知后启动**持续计时器**，计时器超时就发送**窗口探测（Window probe）报文**，对方接收到后返回自己当前接收窗口大小。一般探测次数3次，每次约30-60s。有的实现如果3次过后接收窗口还是0就RST断开连接。

![窗口探测](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/25.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

**糊涂窗口综合症**

如果接收方腾出几个字节就告诉发送方现在有几个字节的窗口，而发送方义无反顾发送这几字节，就是糊涂窗口综合征。

![糊涂窗口综合症](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/26.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

需要

- 让接收方不通告小窗口给发送方

  窗口大小小于min(MSS,缓存空间/2)，则通告窗口为0

- 让发送方避免发送小数据

  Nagle算法，延时处理，满足一下条件之一才发送数据

  - 窗口大小>=MSS且数据大小>=MSS
  - 收到之前发送数据的ack回包

需要同时满足`不通告小窗口给发送方`+`发送方开启Nagle算法`，才能避免糊涂窗口综合征。因为如果只开Nagle，接收方继续通告小窗口，那发送方接收到ack回包则会进行发送。

Nagle默认打开，但对于小数据包交互场景如telnet或ssh，需要关闭Nagle算法（没有全局参数，需要代码setsockopt函数传入TCP_NODELAY选项来关闭）。

**4. 拥塞控制**

**流量控制**是避免发送方的数据填满接收方的缓存，而**拥塞控制**是避免发送方数据填满整个网络，导致更加严重的丢包和时延。

拥塞窗口cwnd是发送方维护的一个状态变量，根据网络拥塞程度动态变化。发送窗口swnd=min(cwnd, rwnd)，即拥塞窗口和接收窗口中的较小者。

拥塞窗口cwnd变化规则

- 网络中没出现拥塞，cwnd增大
- 出现拥塞，cwnd减小

当发送方没在规定时间内接收到ACK应答，即发生超时重传，就认为网络拥塞。

拥塞控制的四个算法：慢启动、拥塞避免算法、拥塞发生、快速恢复

**慢启动**

TCP刚建立连接后，有一个慢启动过程，一点点提升发包数量。

**当发送方每收到一个ACK，cwnd大小+1**（一个MSS大小的数据）。

![慢启动算法](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/27.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

存在慢启动门限ssthresh（slow start threshold）

- 当cwnd < ssthresh，使用慢启动算法
- 当cwnd >= ssthresh，使用拥塞避免算法

**拥塞避免算法**

一般ssthresh大小是65535字节

每收到一个ACK，cwnd增加1/cwnd，变成线性增长

![拥塞避免](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/28.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

线性增长到出现丢包时，触发重传机制，进入拥塞发生算法。

**拥塞发生**

重传机制有超时重传和快速重传

发生**超时重传**的拥塞发生算法：

- ssthresh设为cwnd/2
- cwnd重置为1（恢复为cwnd初始化值，`ss -nli`命令查看TCP连接的cwnd初始值，默认10个MSS）

![拥塞发送 —— 超时重传](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost2/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%AF%E9%9D%A0%E7%89%B9%E6%80%A7/29.jpg?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

拥塞后突然减少数据量，反应强烈，会造成网络卡顿。

发生**快速重传**的拥塞发生算法

- cwnd=cwnd/2
- ssthresh=cwnd
- 进入快速恢复算法

**快速恢复**

快速重传和快速恢复算法一般同时使用，快速恢复算法认为还能收到3个重复ACK说明网络没那么糟糕，因此变化会缓和些。

快速恢复算法：

- cwnd=ssthresh+3（3是三个ACK）（降低cwnd减缓拥塞，且确认了3个重复ACK）
- 重传丢失的数据包
- 如果再收到重复的ACK，cwnd+1（尽快将丢失数据包发给目标）
- 如果收到新数据的ACK，cwnd设为第一步的ssthresh的值。因为收到新数据ACK说明重复ACK的数据已收到，可以恢复到之前的拥塞避免状态了。

![快速重传和快速恢复](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost4@main/%E7%BD%91%E7%BB%9C/%E6%8B%A5%E5%A1%9E%E5%8F%91%E7%94%9F-%E5%BF%AB%E9%80%9F%E9%87%8D%E4%BC%A0.drawio.png?image_process=watermark,text_5YWs5LyX5Y-377ya5bCP5p6XY29kaW5n,type_ZnpsdHpoaw,x_10,y_10,g_se,size_20,color_0000CD,t_70,fill_0)

### 3.3 TCP实战抓包分析

![提纲](.\img\TCP篇\TCP实战抓包分析提纲.jpg)



**1. 显形“不可见”的网络包**

tcpdump工具：命令行格式使用，常用于Linux服务器抓取和分析网络包

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/4.jpg)

![tcpdump 常用选项类](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/7.jpg)

![tcpdump 常用过滤表达式类](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/8.jpg)

wireshark工具：可视化

**2. 解密TCP三次握手和四次挥手**

如果挥手中，没有数据要发送，并且开启了TCP延迟确认机制，就会将第二和第三次挥手合并，形成三次挥手的现象。

**3. TCP三次握手异常情况实战分析**

**实验场景**

**实验一：TCP第一次握手SYN丢包**

**实验二：TCP第二次握手SYN、ACK丢包**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/30.jpg)

服务器在不断收到SYN包后都不会重置重传定时器，因此随后都会传两个SYN-ACK，一个是针对收到的SYN包，一个是重传的。

网络包进出主机的顺序

- 进入：Wire->NIC->tcpdump->netfilter/iptables
- 出去：iptables->tcpdump->NIC->Wire

**实验三：TCP第三次握手ACK丢包**

TCP建立连接后数据包传输的最大超时重传次数由`tcp_retries2`指定，默认15次。

**4. TCP快速建立连接**

![常规 HTTP 请求](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/44.jpg)

![常规 HTTP 请求 与 Fast  Open HTTP 请求](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/45.jpg)

`net.ipv4.tcp_fastopen`打开Fast Open功能，为0关闭，为1作为客户端使用，为2作为服务端时使用，为3时无论客户端还是服务端都使用。

**5. TCP重复确认和快速重传**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/48.jpg)

支持SACK需要双方都支持，`net.ipv4.tcp_sack`打开。

**6. TCP流量控制**

![服务端繁忙状态下的窗口变化](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/50.jpg)

**零窗口通知与窗口探测**

![窗口大小在收缩](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/51.jpg)

![零窗口 与 窗口探测](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/52.jpg)

**发送窗口的分析**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/53.jpg)

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/54.jpg)

**7. TCP延迟确认与Nagle算法**

![禁用 Nagle 算法 与 启用 Nagle 算法](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/55.jpg)

Nagle算法一定会有一个小报文，也就是最开始的时候

Nagle算法思路是延时处理。

没有携带数据的ACK报文网络效率低，40字节的IP头和TCP头却没有携带数据。为此衍生出TCP延迟确认

TCP延迟确认策略：

- 当有响应数据要发送，ACK随响应数据一起立刻发送
- 没有响应数据发送则ACK延迟一段时间，以等待是否有响应数据可以一起发送
- 如果延迟等待发送ACK期间对方第二个数据报文到达，则立刻发送ACK

![TCP 延迟确认](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/57.jpg)

延迟等待时间在Linux内核中定义

```
#define TCP_DELACK_MAX ((unsigned)(HZ/5))
#define TCP_DELACK_MIN ((unsigned)(HZ/25))
```

HZ大小一般会是1000，延迟等待时间单位ms，则最大延迟确认时间200ms，最短40ms。

TCP延迟确认通过Socket设置`TCP_QUICKACK`选项关闭。

TCP延迟确认和Nagle算法混合使用会导致时延增加

![TCP 延迟确认 和 Nagle 算法混合使用](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-Wireshark/61.jpg)

要么发送方关闭Nagle算法，要么接收方关闭TCP延迟确认。

**读者问答**

重传次数超过tcp_retries1会指示IP层进行MTU探测、刷新路由等过程，不会断开TCP连接，超过tcp_retries2则会断开。两个参数都受一个timeout值限制，其值根据这两个参数算出，重传时间超过了timeout则不再重传，即便次数没到。



### 3.4 TCP半连接队列和全连接队列

![本文提纲](.\img\TCP篇\TCP半连接队列和全连接队列提纲.jpg)



**1. 什么是TCP半连接队列和全连接队列**

- 半连接队列：SYN队列
- 全连接队列：accept队列

![半连接队列与全连接队列](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8D%8A%E8%BF%9E%E6%8E%A5%E5%92%8C%E5%85%A8%E8%BF%9E%E6%8E%A5/3.jpg)

两个队列溢出时内核会直接丢包，或返回RST包。

**2. 实战 - TCP全连接队列溢出**

`ss`命令查看TCP全连接队列，获取到的`Recv-Q/Send-Q`在LISTEN状态和非LISTEN状态所表达含义不同

LISTEN状态下

- Recv-Q：当前全连接队列大小
- Send-Q：全连接最大队列长度

非LISTEN状态

- Recv-Q：已收到但未被应用进程读取的字节数
- Send-Q：已发送但未收到确认的字节数

wrk工具，HTTP压测工具

全连接队列满后可以控制变量**tcp_abort_on_overflow**决定服务器行为

- 0：丢弃客户端发来的ACK
- 1：发送RST给客户端

默认设为0更有利于应对突发流量，当服务器丢弃客户端ACK，但客户端已处于ESTABLISHED状态，因此由于服务器没回复ACK，客户端会重发。如果服务器因为短暂繁忙造成队列满，当全连接队列有空位了，重发的报文带有ACK会触发服务器成功建立连接。

只有当全连接队列会长期溢出时才设置为1尽快通知客户端。

全连接队列最大值=min(somaxconn, backlog)，/proc/sys/net/core/somaxconn默认128，backlog通过listen(int sockfd, int backlog)设置，Nginx默认511。

不断有TCP全连接队列溢出，调大**backlog**和**somaxconn**

**3. 实战 - TCP半连接队列溢出**

半连接队列长度没法用ss命令查看，需要通过查看处于SYN_RECV状态的TCP连接数目，即为TCP半连接队列长度。`netstat -natp | grep SYN_RECV | wc -l`

模拟TCP半连接队列溢出：SYN洪泛、SYN攻击、DDos攻击，向服务器一直发SYN包但不回第三次握手的ACK包。hping3工具模拟SYN攻击。

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8D%8A%E8%BF%9E%E6%8E%A5%E5%92%8C%E5%85%A8%E8%BF%9E%E6%8E%A5/23.jpg)

不同版本的Linux内核，理论的半连接最大值计算方法会不同。

syncookies参数

- 0：关闭该功能
- 1：SYN队列放不下再启动
- 2：无条件启用

防御SYN攻击方法

- 增大半连接队列：同时增大`tcp_max_syn_backlog`、`somaxconn`、`backlog`
- 开启tcp_syncookies功能：`/proc/sys/net/ipv4/tcp_syncookies`
- 减少SYN+ACK重传次数：`tcp_synack_retries`

### 3.5 如何优化TCP？

![本节提纲](.\img\TCP篇\优化TCP提纲.jpg)



**1. TCP三次握手的性能提升**

**客户端优化**

**服务端优化**

**如何绕过三次握手？**

![三次握手优化策略](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%82%E6%95%B0/24.jpg)



**2. TCP四次挥手的性能提升**

**主动方的优化**

**被动方的优化**

![四次挥手的优化策略](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%82%E6%95%B0/39.jpg)



**3. TCP传输数据的性能提升**

**滑动窗口是如何影响传输速度的？**

**如何确定最大传输速度？**

**怎么调整缓冲区大小?**

![数据传输的优化策略](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/TCP-%E5%8F%82%E6%95%B0/49.jpg)



### 3.6 如何理解是TCP面向字节流协议？

**1. 如何理解字节流？**

**2. 如何解决粘包？**

**固定长度的消息**

**特殊字符作为边界**

**自定义消息结构**



### 3.7 为什么TCP每次建立连接时，初始化序列号都要不一样呢？



### 3.8 SYN报文什么情况下会被丢弃？

**1. 坑爹的tcp_tw_recycle**

**2. accept列表满了**

**半连接队列满了**

**全连接队列满了**



### **3.9 已建立连接的TCP，收到SYN会发生什么？**

收到SYN报文（初始序列化随机，其实是乱序的）后，由于当前处于Established状态，会回复一个携带正确序列号和确认号的ACK报文，即Challenge ACK。

对方收到Challenge ACK后发现确认号不是期望的就会回复RST报文，随后双方断开连接。

**1. RFC文档解释**

**2. 源码分析**

**3. 如何关闭一个TCP连接？**

**killcx的工具**

**tcpkill的工具**

**4. 总结**



### 3.10 四次挥手中收到乱序的FIN包会如何处理？

乱序FIN包会放入乱序队列，后续接收新数据后再判断乱序队列中有无能用数据包。

**1. TCP源码分析**

**2. 怎么看TCP源码？**



### 3.11 在TIME_WAIT状态的TCP连接，收到SYN后会发生什么？

**1. 先说结论**

**收到合法SYN**

重用四元组，跳过2MSL进入SYN_RECV状态。

![图片](https://cdn.xiaolincoding.com//mysql/other/39d0d04adf72fe3d37623acff9ae2507.png)

**收到非法的SYN**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/network/tcp/tw%E6%94%B6%E5%88%B0%E4%B8%8D%E5%90%88%E6%B3%95.png)

**2. 源码分析**

**3. 在TIME_WAIT状态， 收到RST会断开连接吗？**

**4. 总结**



### 3.12 TCP连接，一端断电和进程崩溃有什么区别？

**1. 主机崩溃**

**2. 进程崩溃**

**3. 有数据传输的场景**

**客户端主机宕机，又迅速重启**

**客户端主机宕机，一直没有重启**

**4. 总结**



### 3.13 拔掉网线后，原本的TCP连接还存在吗？

**1. 拔掉网线后，有数据传输**

**2. 拔掉网线后，没有数据传输**



### 3.14 tcp_tw_reuse为什么默认是关闭的？

**1. 什么是TIME_WAIT状态？**

**2. 为什么要设计TIME_WAIT状态？**

**3. tcp_tw_reuse是什么？**

**4. 为什么tcp_tw_reuse默认是关闭的？**

**第一个问题**

**第二个问题**



### 3.14 HTTPS中TLS和TCP能同时握手吗？

**1. TCP Fast Open**

**2. TLSv1.3**



### 3.15 TCP Keepalive和HTTP Keep-Alive是一个东西吗？

**1. HTTP的Keep-Alive**

**2. TCP的Keepalive**



### 3.16 TCP协议有什么缺陷？

**1. 升级TCP的工作很困难**

**2. TCP建立连接的延迟**

**3. TCP存在队头阻塞问题**

**4. 网络迁移需要重新建立TCP连接**

**5. 结尾**



### 3.17 如何基于UDP协议实现可靠传输？

![img](.\img\TCP篇\QUIC提纲.png)



**1. QUIC是如何实现可靠传输的？**

**Packet Header**

**QUIC Frame Header**



**2. QUIC是如何解决TCP队头阻塞问题的？**

**什么是TCP队头阻塞问题？**

**HTTP/2的队头阻塞**

**没有队头阻塞的QUIC**



**3. QUIC是如何做流量控制的？**

**Stream级别的流量控制**

**Connection流量控制**



**4. QUIC对拥塞控制改进**

**5. QUIC更快的连接建立**

**6. QUIC是如何迁移连接的？**



### 3.18 TCP和UDP可以使用同一个端口吗？

**1. TCP和UDP可以同时绑定相同的端口吗？**

**2. 多个TCP服务进程可以绑定同一个端口吗？**

**3. 客户端的端口可以重复使用吗？**

**4. 总结**



### 3.19 服务端没有listen，客户端发起连接建立，会发生什么？

服务端回复RST

**1. 做个实验**

**2. 源码分析**

**3. 没有listen，能建立TCP连接吗？**

TCP自连接，双方同时向对方发起请求建立连接



### 3.20 没有accept，能建立TCP连接吗？

**1. 三次握手的细节分析**

**半连接队列、全连接队列是什么**

**为什么半连接队列要设计成哈希表**

**怎么观察两个队列的大小**

**全连接队列满了会怎么样？**

**半连接队列要是满了会怎么样？**

**没有listen，为什么还能建立连接**

**2. 总结**



### 3.21 用了TCP协议，数据一定不会丢吗？

**1. 数据包的发送流程**

**2. 建立连接时丢包**

**3. 流量控制丢包**

**4. 网卡丢包**

**RingBuffer过小导致丢包**

**网卡性能不足**

**5. 接收缓冲区丢包**

**6. 两端之间的网络丢包**

**7. 发生丢包了怎么办**

**8. 用了TCP协议就一定不会丢包吗**

**9. 这类丢包问题怎么解决？**

**10. 总结**



### 3.22 TCP四次挥手，可以变成三次吗？

**1. TCP四次挥手**

**为什么TCP挥手需要四次呢？**

**粗暴关闭vs优雅关闭**

**2. 什么情况会出现三次挥手**

**实验验证**

**3. 总结**

没有数据要发送，同时没有开启TCP_QUICKACK，将ACK延迟到和FIN一起发。



### 3.23 TCP序列号和确认号是如何变化的？

**1. 万能公式**

发送的TCP报文：

- 公式一：序列号 = 上一次发送的序列号 + len（数据长度）。没携带数据下，SYN报文或FIN报文的数据长度算1，ACK报文算0
- 公式二：确认号 = 上一次收到的报文中的序列号 + len（数据长度）。没携带数据下，SYN报文或FIN报文的数据长度算1，ACK报文算0

**2. 三次握手阶段的变化**

**3. 数据传输阶段的变化**

**4. 四次挥手阶段的变化**

![在这里插入图片描述](.\img\TCP篇\四次挥手序列号和确认号变化.png)



**5. 实际抓包图**





## 四、IP篇

### 4.1 IP基础知识全家桶

![IP 基础知识全家桶](.\img\IP篇\IP基础知识提纲.jpg)



**1. 前菜 - IP基本认识**

网络层主要作用：实现点对点通信

![IP 的作用](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/2.jpg)

IP在主机间通信，MAC在直连设备间通信

![IP 的作用与 MAC 的作用](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/3.jpg)

源IP和目标IP地址传输中不会变（没使用NAT网络），但MAC地址都在变。

**2. 主菜 - IP地址的基础知识**

**IP地址的分类**

![IP 地址分类](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/7.jpg)

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/8.jpg)

最大主机数=2^(主机号位数)-2。全0指定某个网络，全1指定某个网络下所有主机（用于广播）

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/10.jpg)

在本网络内广播叫**本地广播**，在不同网络间广播叫**直接广播**

![本地广播与直接广播](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/11.jpg)

D类和E类没有主机号，不作为主机IP，D类常用于多播

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/12.jpg)

![单播、广播、多播通信](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/13.jpg)

分类简单明了、选路简单

<img src="https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/14.jpg" alt="IP 分类判断" style="zoom:50%;" />

缺点：

- 同一网络下没有地址层次，缺少地址的灵活性
- 不能与现实网络很好匹配，有的网络内IP地址数目可能太多又或者太少

**无分类地址CIDR**

`a.b.c.d/x`，前x为网络号，后x位主机号

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/15.jpg)

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/16.jpg)

子网掩码除了可以划分网络号和主机号，还可以**划分子网**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/18.jpg)

**公有IP地址与私有IP地址**

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/22.jpg)

公有IP由ICANN管理，互联网名称与数字地址分配机构。IANA是ICANN的一个机构

![img](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/24.jpg)

中国由CNNIC分配

**IP地址与路由控制**

![IP 地址与路由控制](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/25.jpg)

每个主机或路由器上都有路由控制表，选址时将目标IP地址的网络地址和控制表的IP地址的网络地址进行匹配，使用最长匹配原则。

环回地址127.0.0.1，也叫localhost，数据包不流向网络。

**IP分片与重组**

每种数据链路MTU不同，以太网1500字节，IP数据包大于时会被分片。路由器不会重组分片，只能由目标主机进行。

![分片与重组](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/26.jpg)

某分片丢失会造成整个IP数据包作废，全部分片需要重传，因此TCP引入MSS在传输层分片。UDP尽量不要发大于MTU的数据。

**IPv6基本认识**

128位，比IPv4更多地址、更好安全性、扩展性。

优点：

- 即插即用：可自动配置，无需DHCP实现自动分配IP地址
- 提高了传输性能：头部固定40字节，去掉了包头校验和等
- 提升了安全性：有应对伪造IP地址的网络安全功能以及防止线路窃听功能

![IPv6 地址表示方法](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/27.jpg)

连续0可用::隔开，不过只能用一次

![Pv6 地址缺省表示方](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/28.jpg)

![IPv6地址结构](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/29.jpg)

![ IPv6 中的单播通信](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/30.jpg)



**IPv4首部与IPv6首部**

![IPv4 首部与 IPv6 首部的差异](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/31.jpg)

- 取消首部校验和
- 取消分片/重组相关字段
- 取消选项字段

**3. 点心 - IP协议相关技术**

**DNS**

DNS域名解析

- 根DNS服务器（服务器信息保存在互联网中所有DNS服务器中）
- 顶级域DNS服务器
- 权威DNS服务器

![DNS 树状结构](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/32.jpg)

浏览器首先查看缓存，再向操作系统缓存，接着是本机域名解析文件hosts，都没有就继续DNS服务器查询。

![域名解析的工作流程](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/33.jpg)

**ARP**

通过ARP请求与ARP响应两种包确定MAC地址

![ARP 广播](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/34.jpg)

MAC地址缓存有一定期限

RARP协议已知MAC地址得IP地址，通常如打印机服务器等小型嵌入式设备接入网络时使用。

![RARP](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/35.jpg)



**DHCP**

动态获取IP地址，UDP广播通信，DHCP客户端监听地址端口68，DHCP服务器监听地址端口67

![DHCP 工作流程](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/36.jpg)

DHCP中继代理，使得对不同网段的IP地址分配也可以由一个DHCP服务器统一进行管理。中继代理和DHCP服务器间以单播形式通信。

![ DHCP 中继代理](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/37.jpg)

**NAT**

无分类地址也不够消耗，因此提出**网络地址转换NAT**

![NAT](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/38.jpg)

网络地址与端口转换NAPT

![NAPT](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/39.jpg)

NAT路由器中实际上是将私有IP及其端口一起转换成同一个公有IP但是不同端口的情况。

转换表在NAT路由器上自动生成，例如在TCP发送SYN包时生成，在关闭连接时收到FIN包的确认应答从表中删除

缺点：

- 外部无法主动与NAT内部服务器建立连接，因为NAPT转换表没有记录
- 转换表生成与转换都会产生性能开销
- 通信中如果NAT路由器重启，使用TCP连接都将重置

解决方法：

- 改用IPv6：地址足够，无需NAT转换
- NAT穿透技术：客户端主动从NAT设备获取公有IP地址，然后自己建立端口映射条目，用该条目对外通信，就不需要NAT设备来转换了

**ICMP**

互联网控制报文协议

主要功能：确认IP包是否成功送达目标地址、报告发送过程中IP包被废弃的原因和改善网络设置等。

![ICMP 目标不可达消息](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/40.jpg)

报文类型

- 查询报文类型：用于诊断的查询消息
- 差错报文类型：通知出错原因的错误消息

![常见的 ICMP 类型](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/41.jpg)



**IGMP**

因特网组（播成员）管理协议，网络层协议，工作在主机和最后一条路由之间（下图蓝色部分）

不在一组的主机不能收到数据包

![组播模型](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/42.jpg)

IGMP报文采用IP封装，头部协议号2，TTL通常为1

三个版本IGMPv1、IGMPv2、IGMPv3，以下例IGMPv2

常规查询与响应工作机制：

![ IGMP 常规查询与响应工作机制](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/43.jpg)

离开组播组工作机制

情况一：离开后网段内仍有该组播组

![ IGMPv2 离开组播组工作机制 情况1](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/44.jpg)

情况二：离开后没有该组播组

![ IGMPv2 离开组播组工作机制 情况2](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/IP/45.jpg)

组播一般用UDP协议，发送UDP组播数据时，目的地址填组播地址

### 4.2 ping的工作原理

 **1. IP协议的助手 - ICMP协议**

ping基于ICMP协议，ICMP报文封装在IP包里，工作在网络层。ICMP协议号1

![ICMP 报文](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/5.jpg)

![常见的 ICMP 类型](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/6.jpg)

**2. 查询报文类型**

**回送消息**，类型0和8，用于判断数据包是否成功到达对端，ping命令基于此实现

![ICMP 回送消息](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/7.jpg)

![ICMP 回送请求和回送应答报文](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/8.jpg)

相比原生ICMP多了些字段

- 标识符：区分哪个应用程序发的ICMP包
- 序号：确认网络包是否有丢失
- 选项数据中ping还会存放发送请求的时间以计算往返时间。

**3. 差错报文类型**

- 目标不可达消息 —— 类型3

  ICMP包头的代码字段说明具体原因

  ![目标不可达类型的常见代码号](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/9.jpg)

- 原点抑制消息 —— 类型4：路由器向低速线路发送数据时可能拥堵，可以向IP包源地址发送该消息，从而增大IP包发送的间隔。

- 重定向消息 —— 类型5：路由器发现发送端主机使用了不是最优的路径，使用重定向消息通知最合适的路由信息和源数据。

- 超时消息 —— 类型11：IP包TTL字段减到0时IP包被丢弃，路由器会通告发送端主机。TTL目的在防止环路。

  ![ICMP 时间超过消息](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/11.jpg)

**4. ping - 查询报文类型的使用**

![主机 A ping 主机 B 期间发送的事情](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/17.png)

**5. traceroute - 差错报文类型的使用**

traceroute工具使用ICMP差错报文类型

- 作用一：设置特殊TTL来追踪途径的路由器，利用IP包TTL从1开始递增的同时发生UDP包，强制接收**ICMP超时消息**。不过有的公网地址是看不到途径的路由器的。UDP包中填入不可能的端口号33434作为UDP目标端口，下一个UDP端口号再加一，UDP包到达目的主机后会返回ICMP差错报文消息中的**目标不可达消息**，类型**端口不可达**。

- 作用二：设置不分片以确定路径的MTU。ICMP返回目标不可达消息，代码4（需要进行分片但设置了不分片）

  ![MTU 路径发现（UDP的情况下）](https://cdn.xiaolincoding.com/gh/xiaolincoder/ImageHost/%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%BD%91%E7%BB%9C/ping/18.jpg)

### 4.3 断网了，还能ping通127.0.0.1吗？

**1. 什么是127.0.0.1**

127开头属于回环地址，127.0.0.1是其中之一，但源码中将回环地址定死成127.0.0.1。IPv6的回环地址是::1

**2. 什么是ping**

ping是应用层命令，底层用的是网络层的ICMP协议

虽然ICMP协议和IP协议都属于网络层协议，但ICMP利用了IP协议传输消息

**3. TCP发数据和ping的区别**

![图片](https://cdn.xiaolincoding.com//mysql/other/eb0963a11439dff361dbe0e7a8876abd.png)

`socket(AF_INET, SOCK_STREAM, 0)`，AF_INET说明使用IPv4的host:port方式解析待会输入的网络地址。SOCK_STREAM面向字节流的TCP协议，SOCK_RAW原始套接字，SOCK_DGRAM面向无连接的UDP协议。

**4. 为什么断网了还能ping通127.0.0.1**

![图片](https://cdn.xiaolincoding.com//mysql/other/c1019a8be584b27c4fc8b8abda9d3cf1.png)

数据到网络层后，判断目标IP是回环地址的话就走本地网卡，本地网卡是个假网卡，本地网卡会将数据推送到input_pkt_queue链表中，该链表被所有网卡共享，上面挂着发给本机的各种消息。消息发送到链表后会触发软中断，再由ksoftirqd内核线程来处理软中断，拆包交给应用程序。

**5. ping回环地址和ping本机地址有什么区别**

没有区别，都走本地回环接口

**6. 127.0.0.1和localhost以及0.0.0.0有区别吗**

localhost不是IP而是域名，不过默认会被解析成127.0.0.1，可在/etc/hosts文件里修改。

ping 0.0.0.0会失败，因为0.0.0.0在IPv4中表示无效的目的地址。不过在listen时使用0.0.0.0则能够监听本机所有IPv4地址。connect时不能连接0.0.0.0，要指明IP。



## 五、学习方法

### 5.1 入门系列

**《图解HTTP》**

**《图解TCP/IP》**

**《网络是怎样连接的》**

[《计算机网络微课堂》](https://www.bilibili.com/video/BV1c4411d7jb?p=1)

### 5.2  深入学习系列

**《计算机网络 - 自顶向下方法》**

**《TCP/IP详解 卷一：协议》**

**《The TCP/IP GUIDE》**

[TCP协议的RFC文档](https://datatracker.ietf.org/doc/rfc1644/)

### 5.3 实战系列

**《Wireshark网络分析就这么简单》**

**《Wireshark网络分析的艺术》**

