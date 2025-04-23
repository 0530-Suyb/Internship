参考【[深入理解Linux中网络I/O复用并发模型 —— 刘丹冰Aceld](https://www.bilibili.com/video/BV1jK4y1N7ST/?spm_id_from=333.1387.homepage.video_card.click&vd_source=b9cd24f9035eb62c95c54f64275cce57)】



### 引入

阻塞等待，不占用CPU时间片

非阻塞忙轮询，占用CPU、系统资源

但阻塞等待不能很好处理多IO请求场景，同一个阻塞一次只能处理一个流的阻塞监听。

为此引入IO多路复用，就能够阻塞等待不浪费资源，也能同时监听多个IO请求状态。



### IO模型分类

- 同步IO模型：
  - 阻塞IO（BIO）
  - 非阻塞IO（NIO）
  - 多路IO复用
  - 信号驱动IO
- 异步IO模型（AIO）



阻塞/非阻塞（进程等待状态）、同步/异步（消息通讯机制）



### 大量IO请求方案

方法一：阻塞等待+多进程/多线程，需要开辟线程浪费资源

方法二：非阻塞+忙轮询，CPU大部分时间都在判断，利用率不高

方法三：多路IO复用——select（与平台无关），监听IO数量有限1024，将监听的所有流传入select函数，有数据后还要对所有流遍历找到具体哪些流可读写

方法四：多路IO复用——epoll（Linux），可监听大量IO（系统可打开的文件数目），epoll_wait函数返回可处理的流，无需遍历



### epoll API

1. 创建epoll

   ```c
   // size为内核监听的数目，返回一个epoll句柄（文件描述符）
   int epoll_create(int size);
   ```

   在内核创建一颗红黑树的根节点

2. 控制epoll

   ![image-20250420222702901](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\epoll_ctr.png)

   ![image-20250420222911936](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\epoll_ctr案例.png)

3. 监听epoll

   ![image-20250420223818810](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\epoll_wait.png)

   ![image-20250420224024109](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\epoll_wait案例.png)



编程架构

![image-20250420224403331](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\epoll编程架构.png)



### epoll触发模式

1. 水平触发

   ![image-20250420224835609](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\水平触发.png)

2. 边缘触发

   ![image-20250420224927758](C:\Users\hp-pc\Desktop\实习备战记\操作系统\Linux网络IO复用并发模型\img\边缘触发.png)

水平触发，用户监听epoll事件，内核有事件就拷贝给用户态，如果用户没处理完，剩余的事件在下次epoll_wait时还会再次返回。如果存在大量未处理的事件，那每次都拷贝就消耗大。默认情况下是水平触发模式。

边缘触发，只会通知一次，如果用户不处理，就不会再通知了，可能存在未处理事件丢失。通过在epoll_ctl传入的epoll_event中设置EPOLLET来使用。