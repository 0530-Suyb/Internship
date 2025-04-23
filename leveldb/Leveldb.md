参考

[【深度知识】LevelDB从入门到原理详解](https://cloud.tencent.com/developer/article/1602204)

[一文彻底搞懂leveldb架构](https://blog.csdn.net/songguangfan/article/details/124828824)

### 简介

leveldb写性能较佳，底层基于LSM(Log Structured Merge)-Tree实现。非关系型数据库，k-v存储，单机。适用于查询较少，写多的场景。

LSM核心思想通过牺牲部分读性能来换取最大的写性能，具体来说是先将数据写入内存，减少随机写磁盘的次数，但内存数据量达到一定阈值后再将数据真正刷新到磁盘。通常顺序写60MB/s，随机写45MB/s。

（达到阈值刷新磁盘，防止日志中写操作过多，当异常发生时恢复时间过长）



![在这里插入图片描述](https://i-blog.csdnimg.cn/blog_migrate/b64890490de82113e751adab94955a59.jpeg#pic_center)

leveldb由以下部件组成：

- memtable：数据首先写入到内存的memtable，达到阈值（4MB）后转换成immutable memtable，同时创建新memtable供用户读写。底层采用跳表，增删改查复杂度都在O(logN)
- immutable memtable：只读，被创建后leveldb的后台压缩进程会利用其内容创建sstable并持久化到磁盘
- sstable：除一些元数据文件，leveldb数据主要都用sstable存储。多个memtable数据持久化到磁盘，对应的sstable间可能存在交集，从而导致读操作需要遍历所有sstable。对此leveldb定期整合campaction，将sstable文件在逻辑上分若干层，内存数据直接dump出来的是0层，其他每一层的数据都由上一层进行compaction得到。每一层包含多个SST文件，文件内数据有序。sstable本身不可修改。
- manifest：记录了各层SST文件的元数据（文件大小、最大最小key等）。每次compaction完成都会新建一个version，表示在旧version上变化了哪些内容（如增删的sstable文件、日志文件编号、当前compaction下表等），这些变化信息都会编码成一条记录写入manifest文件中。leveldb快照需要维护多版本，所以可能存在多个manifest文件。
- current：记录当前manifest文件名，每次leveldb启动都会新建一个manifest
- log：leveldb写操作将数据写入内存前，会将写操作写入日志文件，以防异常或宕机而丢失内存数据等（日志的写入是顺序写入，效率高）



### 特点

- k、v都是任意长度的字节数组
- entry（一条k-v记录）默认按key字典序存储
- 基本操作接口：Put、Delete、Get、Batch
- 支持批量操作原子进行
- 可创建数据快照snapshot并从中查找
- 可向前向后迭代遍历数据（迭代器）
- 自动使用snappy压缩数据
- 可移植性
- 单线程写、多线程读
- leveldb是单进程服务，性能高，每秒写40W，随机读10W，压缩后读会更好些



### 基础部件

#### 跳表

#### SSTable

- Footer
- IndexBlock
- MetaIndexBlock
- DataBlock
- FilterBlock

#### 日志

#### 版本控制



### 操作接口





### 知识遗漏点

- 不支持写并发
- 