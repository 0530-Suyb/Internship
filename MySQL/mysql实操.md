## 安装与配置

```shell
# linux上
sudo apt update
sudo apt install mysql-server

# 启动
sudo systemctl start mysql
sudo systemctl status mysql

# 默认密码在/etc/mysql/debian.cnf中记录，其中的password就是默认密码，先用其登录后再修改
mysql -u debian-sys-maint -p 
... # 输入password密码

# 修改root使用sql的验证密码
use mysql;
# 将root用户在localhost上的身份验证方式改为mysql_native_password，并且设置一个新的密码为123456
ALTER USER 'root'@'localhost' IDENTIFIED WITH mysql_native_password BY '123456'
flush privileges; # 刷新授权 
quit; # 退出sql

sudo service mysql restart # 重启
mysql -u root -p # 输入新密码
```

## 用户操作

```shell
mysql -u root -p
# 若连接本地MySQL服务，无需-h指定IP地址
# -p可以在命令行输入密码，但建议在交互对话里输入

use mysql;

# 创建用户
CREATE USER 'username'@'host' IDENTIFIED BY 'password';
# 例
CREATE USER 'suyb'@'localhost' IDENTIFIED BY '123456'; # 无密码的话就''
CREATE USER 'suyb'@'%' IDENTIFIED BY '123456'; # 允许用户suyb从任意地址%登录，远程登录权限

# 用户权限
GRANT SELECT ON databasename.tablename TO 'username'@'host'; # 授予用户SELECT操作权限，也可以是INSERT、UPDATE等等，*.*表示授予该用户所有数据库和表的相应操作权限
REVOKE privileges ON databasename.tablename FROM 'username'@'host'; # 撤销权限
ALTER USER 'username'@'host' IDENTIFIED BY 'password' PASSWORD EXPIRE NEVER; # 设置修改密码
DROP USER 'username'@'host'; # 删除用户

SELECT User, Host FROM user; # 查看用户
show grants for 'username'@'host'; # 查看用户权限

SELECT Host, User, plugin FROM mysql.user; # 查看加密规则
ALTER USER 'username'@'host' IDENTIFIED WITH mysql_native_password BY 'password'; # 修改加密方式


```

## 数据库

### 数据库的操作

```shell
CREATE DATABASE test; # 创建
SHOW DATABASES;	# 显示存在的所有数据库
USE test; # 使用
DROP DATABASE test; # 删除
```

### 数据类型

**数值类型**

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/03d008e6068a44dba3ab6467e5ba1335.png)

DECIMAL和NUMERIC都用于存储精确数据，M是整数位数和小数位数之和的最大值，D表示小数位数。

BIT类型数据，赋值时b'1'或b'0'来存储，如 `INSERT INTO users (user_id, username, is_active) VALUE(1, 'john', b'1');`

```shell
# 查询、显示、比较BIT类型数据时需要转换
# 查询 users 表中所有记录，并将 is_active 字段转换为整数类型
SELECT user_id, username, CAST(is_active AS UNSIGNED) AS is_active FROM users;
# 查询已激活的用户
SELECT * FROM users WHERE CAST(is_active AS UNSIGNED) = 1;
```

**字符串类型**

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/30eadfd1a0454c86980756d1ce63e01c.png)

VARCHAR(SIZE)最大程度SIZE，一个汉字也算一个字符；可变长度除了字符占用的空间，还需要额外1-2字节记录字符串长度，小于256字符则1字节记录，否则2字节。

CHAR固定长度，如CHAR(10)，无论存的多少都会占满10个字符，如存'abc'也会占10字符。

**日期类型**

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/d9b3f5ca063e4bc2a7bffe4dda28b5fb.png)

DATE日期、TIME时间、DATATIME日期时间、TIMESTAMP时间戳

```shell
mysql> CREATE TABLE tab (
    -> id INT,
    -> name VARCHAR(10),
    -> active BIT,
    -> high FLOAT(4,1),
    -> weight DECIMAL(4,1),
    -> motto TEXT,
    -> birthday DATE,
    -> sleepT TIME);
Query OK, 0 rows affected, 1 warning (0.01 sec)
```

### 数据表的操作

```shell
# 创建表
mysql> CREATE TABLE tab (
    -> id INT,
    -> name VARCHAR(10),
    -> active BIT,
    -> high FLOAT(4,1),
    -> weight DECIMAL(4,1),
    -> motto TEXT,
    -> birthday DATE,
    -> sleepT TIME);
Query OK, 0 rows affected, 1 warning (0.01 sec)

# 查看存在的表
SHOW TABLES;
# 具体表内容
DESC 表名;
# 删除表
DROP TABLE 表名;
```

### 数据表的增删改查

```shell
# 新增
INSERT INTO table (id, name, ...) VALUES(列, 列, ...), (列, 列, ...)...;

# 查询
SELECT * FROM 表名; # 查询列越多，传输数据量越大，操作吃大量硬盘IO和网络IO导致服务器无法正常响应其他用户请求。
SELECT 列 FROM 表名;
SELECT 表达式 FROM 表名;
SELECT 表达式 AS 别名 FROM 表名;
SELECT DISTINCT 列名, ... FROM 表名; # distinct去重
SELECT 列名, ... FROM 表名 ORDER BY 列名, ... [DESC]; # 加DESC将降序排序
SELECT 列名, ... FROM 表名 WHERE 条件; # 条件运算符有
```

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/ef08916c2ab5480f98a5f94617dedc98.png)

NULL不能用=判断，WHERE XX!=NULL和WHERE XX=NULL永远都是返回0行。

WHERE条件可以用表达式但是不能用别名。

NOT、AND、OR优先级依此递减，可以用()优先执行。

LIKE模糊匹配，%匹配任意个字符，_匹配一个字符。

IS NOT NULL与=、<=>区别。

```shell
# 分页查询
SELECT 列名, ... FROM 表名 LIMIT 行数n OFFSET 行偏移m; # 从m行开始查询n行
# 修改
UPDATE 表名 SET 列名=新值, ... [WHERE 条件/ORDER BY 表达式 LIMIT n]; # ORDER是排序取前n行操作
# 删除
DELETE FROM 表名 [WHERE 条件/ORDER BY 表达式 LIMIT n]; # 删除若干行，若没有指定任何条件将删除所有行

```

### 数据库约束

```shell
NOT NULL
UNIQUE
DEFAULT
PRIMARY KEY # NOT NULL和UNIQUE的结合，AUTO_INCREMENT关键字配合键会自增，多mysql集群自增失效
FOREIGN KEY # 外键，如foreign key (classes_id) references classes(id)，将从表的classes_id外键和主表classes的主键id关联

```

### 表的增查（进阶）

```shell
# 插入查询结果
insert into 表名 (列名, ...) select 列名, ... from 被引用的表名;

# 聚合查询
COUNT # count(*)包括null，count(列名)不包括null
SUM
AVG
MAX
MIN

# 分组查询
select 列名，聚合函数(列名), ... from 表名 group by 列名 [having 条件]; # having从分组结果进行条件过滤
# 分组前条件筛选用where，分组后用having

# 联合查询：多表查询，取笛卡尔积得到一张更大的表，新表的列数为两表列之和，新表的行数为两表行之积
# 内连接，并集
select 字段 from 表1 别名1 [inner] join 表2 别名2 on 连接条件 and 其他条件;
select 字段 from 表1 别名1,表2 别名2 where 连接条件 and 其他条件;
# 外连接，分左外连接和右外连接
-- 左外连接，表1完全显示
select 字段名 from 表名1 left join 表名2 on 连接条件;
-- 右外连接，表2完全显示
select 字段 from 表名1 right join 表名2 on 连接条件;

# 自连接：行关系转列关系使得可以行间比较
-- 先查询“计算机原理”和“Java”课程的id
select id,name from course where name='Java' or name='计算机原理';
 
-- 也可以使用join on 语句来进行自连接查询
SELECT s1.* FROM score s1 JOIN score s2 ON s1.student_id = s2.student_id AND s1.score < s2.score AND s1.course_id = 1 AND s2.course_id = 3;

SELECT stu.*, s1.score Java, s2.score 计算机原理 FROM score s1 JOIN score s2 ON s1.student_id = s2.student_id JOIN student stu ON s1.student_id = stu.id JOIN course c1 ON s1.course_id = c1.id JOIN course c2 ON s2.course_id = c2.id AND s1.score < s2.score AND c1.NAME = 'Java' AND c2.NAME = '计算机原理';

# 子查询（嵌套查询）
select * from student where classes_id=(select classes_id from student where name='不想毕业');

# 合并查询
select * from course where id<3 union select * from course where name='英文';
-- 或者使用or来实现
select * from course where id<3 or name='英文';

```

### 索引基础

索引是一种特殊文件，包含对数据表里所有记录的引用指针。可针对表中一列或多列创建索引，并指定索引的类型，各类索引各有数据结构实现。

特点：加快查询、额外存储、新增删除表行时索引要更新（额开开销）

使用场景：数据量大、经常条件查询、插入和修改频率低、有额外可存储的空间

普通索引、主键约束、唯一约束、外键约束

```shell
# 普通索引
create index 索引名 on 表名(列名, ...)
# 唯一索引
create unique index 索引名 on 表名(列名, ...)
# 主键索引，建表时定义，PRIMARY KEY [AUTO_INCREMENT]
# 全文索引
create fulltext index 索引名 on 表名(列名, ...)

# 建表时创建索引
create table 表名 (
	c1 INT,
	c2 VARCHAR(10),
	...
	INDEX 索引名 (c1, c2, ...),
	UNIQUE INDEX 唯一索引名 (c3),
	PRIMARY KEY (c4)
)


show index from 表名;
create index 索引名 on 表名(列名, ...);
drop index 索引名 on 表名;

```

索引主要用B+树结构，数据都在叶子节点，且可顺序访问，叶子节点指向相邻的叶子节点。

![在这里插入图片描述](https://i-blog.csdnimg.cn/direct/be99c21e4fb041cf8845454db05bb15e.png)

hash可快速定位但没有顺序、二叉树不能自平衡、红黑树树高随数据量增加而增加。

B+树范围查询先找到初始点，后沿链表遍历即可；查询时间稳定；

### 事务基础

使用事务原因：

- 数据不一致问题：事务原子性
- 并发操作问题：数据冲突（赃读、不可重复读、幻读等）
- 系统故障问题：数据丢失
- 复杂操作的可靠性：业务逻辑包含多步骤，要作为一个整体去执行，全部成功或失败

事务的使用

```shell
# 开启事务
start transaction;
# 执行多条SQL语句
# 回滚或提交
rollback/commit;
```

事务特性ACID

- A原子性
- C一致性：事务前后，数据库完整性不被破坏
- I隔离性：防止多个事务对某些数据并发操作导致数据不一致
- D持久性：事务处理后对数据的修改是永久的，系统故障也不会丢失

事务并发问题

- 赃读：一个事务读到另一个事务未提交的数据，但未提交的数据最终被回滚，则读到的数据是无效的
- 不可重复读：一个事务内多次读同一个数据，由于其他事务的修改或删除导致读到的结果不一样。锁或MVCC多版本并发控制来防止
- 幻读：一个事务多次执行相同查询操作，但由于其他事务插入或删除符合查询条件的数据，导致前后读取到的结果集不一致。数据行数增减，而非单行数据变化。可以完全串行化，或使用范围锁，则查询条件涉及的数据范围不允许其他事务插入或删除

事务隔离级别：

- 读未提交：存在赃读、不可重复度、幻读
- 读已提交：存在不可重复度、幻读
- 可重复读：存在幻读
- 串行化：最高隔离级别

### JDBC

## 数据库导出导入

```shell
# 导出：mysqldump工具操作
# 导出整个数据库
mysqldump -u [用户名] -p [数据库名] > [导出文件名].sql

# 导出特定表
mysqldump -u [用户名] -p [数据库名] [表名1] [表名2] ... > [导出文件名].sql

# 导出数据库结构(不含数据)
mysqldump -u [用户名] -p -d [数据库名] > [导出文件名].sql

# 导入
# mysql导入数据库
mysql -u [用户名] -p [数据库名] < [导入文件名].sql

# 或者在mysql客户端使用source命令导入
mysql -u [用户名] -p
USE [数据库名];
SOURCE [导入文件的完整路径].sql;
```
