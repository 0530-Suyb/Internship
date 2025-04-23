[toc]

速通C++

[C++入门/2小时从C到C++快速入门（2018，C++教程）](https://www.bilibili.com/video/BV1kW411Y76d/?share_source=copy_web&vd_source=0382d972655d9a40d7bd297e9b630baf)



### 文件命名

```c++
// .cpp/.hpp 或直接就不用后缀
#include <math.hpp>
#include <cmath>
#include <iostream>	// 拥有std::cout、std::cin等标准输入输出
```



### 命名空间

```c++
using namespace std; // 使用命名空间从而多个文件内可以有同名变量
using std::cout; // 文件中调用std::cout时直接用cout就行
std::cout << "hello" << std::endl; // 或者直接带上命名空间std，endl插入换行符并刷新输出缓冲区（将缓冲区内容都输出）
```



### 引用

```c++
int a = 1;
int &b = a; // 引用是变量的别名，声明时必须初始化
cout << b; // 1
a = 2;
cout << b; // 2

// 与指针很像但引用不可修改指向新地址，引用本身不占用额外内存空间，而指针需要额外的4/8字节
// 引用更简洁，不需解引用
// 引用地址和引用变量的地址相同，指针的地址和指针指向的变量地址不同

void swap(int &a, int &b) {
    int t = a;
    a = b;
    b = t;
} 

// 引用不适合动态内存管理，因为不能更改引用对象，无法灵活管理内存分配和释放
int *ptr = new int(10);
delete ptr;
```



### 内联函数

内联函数用inline关键字声明，编译器会将内联函数在调用处将其展开（内联展开），以避免函数调用开销。内联函数应该是不包含循环的简单函数。

```c++
#include <iostream>
#include <cmath>
using namespace std;

inline double distance(double a, double b) {
    return sqrt(a*a + b*b)
}

int main() {
    double k=6, m=9;
    cout << distance(k, m) << endl;
    // 等同于
    cout << sqrt(k*k + m*m) << endl;
    return 0;
}
```



### 异常捕获try-catch

允许程序运行时检测到错误或异常情况时，采取相应措施，而不是让程序崩溃。

```c++
#include <iostream>
#include <stdexcept>

double divide(double numerator, double denominator) {
    if (denominator == 0) {
        // 抛出一个 std::runtime_error 类型的异常
        throw std::runtime_error("除数不能为零");
    }
    return numerator / denominator;
}

int main() {
    double num1 = 10.0;
    double num2 = 0.0;

    try {
        double result = divide(num1, num2);
        std::cout << "结果: " << result << std::endl;
    }
    catch (const std::runtime_error& e) {
        // 捕获并处理 std::runtime_error 类型的异常
        std::cerr << "捕获到异常: " << e.what() << std::endl;
    }
    catch (...) {
        // 捕获所有类型异常
    }
    // catch块顺序影响处理先后，只会进入一个catch

    return 0;
}
```



### 默认形参

函数形参可带默认值，不过必须一律靠右。

```c++
int add(int a, int b = 0) {
    return a+b;
}
```



### 函数重载

允许函数同名，只要形参不同（数目、类型）。

调用时根据实参和形参的匹配来选择。

返回类型不同没用。



### 运算符重载

运算符重载是一种特殊的函数，允许用户定义运算法如何作用于自定义类型的对象，可提高代码可读性和可维护性。

```c++
返回类型 operator 运算符 (参数列表) {
	// 函数体
}
```

规则：

- 不能创造新运算符，只能重载已存在的，如+、-、*、/、==、!=
- 运算符优先级和结合性不变。右结合性（从右往左：单目、赋值、三目）
- 运算符重载函数的参数至少一个是自定义类型
- 部分运算符不能重载：.（成员访问）、.*（成员指针访问）、::（作用域解析）、?:（条件运算法）

```c++
#include <iostream>
using namespace std;

struct Point {
    int x;
    int y;
};

// 二元运算符
// 重载+运算符
Point operator+(const Point& p1, const Point& p2) {
    Point p;
    p.x = p1.x + p2.x;
    p.y = p1.y + p2.y;
    return p;
}
// 重载+=运算符，复合赋值运算符
Point& operator+=(Point& p1, const Point& p2) {
    p1.x += p2.x;
    p1.y += p2.y;
    return p1;	// 返回自引用，可以连续使用运算符
}
// 重载==运算符
bool operator==(const Point& p1, const Point& p2) {
    return p1.x == p2.x && p1.y == p2.y;
}

// 一元运算符
// 前置++
Point& operator++(Point& p) {
    p.x++;
    p.y++;
    return p;
}
// 后置++
// 为什么要加int，因为编译器需要区分前置++和后置++
Point operator++(Point& p, int) {
    Point old = p;
    p.x++;
    p.y++;
    return old;
}

ostream& operator<<(ostream& os, const Point& p) {
    os << "(" << p.x << ", " << p.y << ")" << endl;
    return os;
}

int main() {
    Point p1 = {1, 2};
    Point p2 = {3, 4};
    Point p3 = p1 + p2;
    cout << p3; // (4, 6)
    p3 += p1;
    cout << p3; // (5, 8)
    ++p3;
    cout << p3; // (6, 9)
    cout << p3++; // (6, 9)
    cout << p3; // (7, 10)

    bool res = p1 == p2;
    cout << res << endl; // 0
    
    return 0;
}
```



### 模板函数

使用模板让编译器自动生成一个针对该数据类型的具体函数

模板函数是泛型编程的主要工具，提高代码复用性。

```c++
template <typename T>
返回类型 函数名(参数列表) {
	// 函数体
}
// typename关键字也可以替换成class关键字
```

```c++
#include <iostream>
using namespace std;

// 模板函数分两部分，第一部分是声明，第二部分是定义，必须放在一起
template <class T>
T add(T a, T b) {
    return a + b;
}

template <class T>  
T sub(T a, T b) {
    return a - b;
}

// 参数不同类型
template <class T1, class T2>
T1 add(T1 a, T2 b) {
    return a + b;
}

template <class T1, class T2>
T1 sub(T1 a, T2 b) {
    return a - b;
}

int main() {
    cout << add(1, 2) << endl; // 3
    cout << add(1.1, 2.2) << endl; // 3.3
    cout << add(1, 2.2) << endl; // 3
    cout << sub(1, 2) << endl; // -1
    cout << sub(1.1, 2.2) << endl; // -1.1
    cout << sub(1, 2.2) << endl; // -1
    return 0;
}
```



### 动态内存分配

关键字new和delete比C语言的malloc/calloc/realloc和free更好，可以对类对象调用初始化构造函数或销毁析构函数。以上都基于堆内存空间。

```c
int *ptr = (int*)malloc(5*sizeof(int));
free(ptr);
int *ptr2 = (int*)calloc(5, sizeof(int)); // 分配并初始化为0
int *ptr3 = (int*)realloc(ptr2, 7*sizeof(int)); //可扩大缩小内存 
free(ptr3);
```

new和delete

```c++
#include <iostream>
#include <cstring>
using namespace std;

int main() {
    double *dp = new double;
    *dp = 7.77;
    cout << "dp: " << dp << endl;
    cout << "*dp: " << *dp << endl;
    delete dp; // 只释放dp指向的内存，不会释放dp本身
    cout << "dp: " << dp << endl;
    cout << "*dp: " << *dp << endl;

    int n = 5;
    dp = new double[n];
    for (int i = 0; i < n; i++) {
        dp[i] = i + 0.1;
    }
    for (int i = 0; i < n; i++) {
        cout <<*(dp+i) << " ";
    }
    cout << endl;
    delete[] dp; // 释放dp指向的内存数组

    char *s = new char[100];
    strcpy(s, "hello");
    cout << s << endl;
    delete[] s;

    return 0;
}
```



### 类

类Class：默认成员为private（访问控制的权限），类将成员变量和成员函数封装，是面向对象编程的核心概念，涉及成员访问控制、构造函数、析构函数、继承、多态等内容

结构Struct：默认成员为public

```c++
class 类名 {
public: // 访问控制修饰符 public/private，类中默认private
    // 成员变量，一般在类中成员变量不公开，而是通过成员函数去操作
    // 成员函数
}
```

类和结构体的一些使用和对比

```c++
#include <iostream>
#include <cstring>
using namespace std;

class P {
    // 抽象类，定义纯虚函数，后续派生类要重写纯虚函数
public: // 注意配合上public，否则外部不能调用
    virtual int getAge() = 0;
    virtual char *getName() = 0;
};

class Person:P {
    // 默认声明都为private
    // 类的成员变量一般要声明为private，成员函数一般声明为public
    // public的公开成员也称“对外接口”
private:
    int age;
    char *name;
public:
    // 默认构造函数，无参，无返回值。一旦自定义了构造函数，系统就不会再提供默认构造函数
    Person(int a, char *name = "no name") { // 默认参数
        age = a;
        int len = strlen(name); // 计算字符串长度, 不包括'\0'
        this->name = new char[len + 1]; // 为name分配内存
        strcpy(this->name, name); // 复制字符串, 包括'\0'
        // this指针指向当前对象, 存在同名变量时，用this指针区分
        // (*this).name = name; // 也可以这样写
    }
    
    // 成员函数
    int getAge() override { return age; } 
    char *getName() override {return name; }

    void show() {
        cout << "age: " << age << ", name: " << name << endl;
    }
    void doWork();

    Person& old() {	// 自引用，从而可连续调用
        cout << "olding..." << endl;
        age++;
        return *this;
    }

    Person& operator+=(int i) {
        this->age += i;
        return *this;
    }

    // 析构函数，无参，无返回值，没有重载
    ~Person() {
        delete[] name; // 释放name指向的内存
    }
};  // 类定义结束后要加分号

// 成员函数的定义, 类外定义成员函数时，要加类名和作用域解析符
void Person::doWork() {
    cout << "working..." << endl;
}

// 通过结构体来实现类似类的功能
struct Person2:P {
    // 默认声明都为public
    // 可以设置private进行访问控制
    Person2(int a, char *name = "no name") {
        age = a;
        int len = strlen(name);
        this->name = new char[len + 1];
        strcpy(this->name, name);
    }

    int getAge() override { return age; }
    char *getName() override { return name; } // 不安全，将name的指针返回，不过此次暂且如此

    void show() {
        cout << "age: " << age << ", name: " << name << endl;
    }
    void doWork();
    Person2& old() {
        cout << "olding..." << endl;
        age++;
        return *this;
    }
    
    Person2& operator+=(int i) {
        age+=i;
        return *this;
    }

    ~Person2() {
        delete[] name;
    }
private:
    int age;
    char *name;
};

void Person2::doWork() {
    cout << "working..." << endl;
}

// ostream& operator<<(ostream& o, Person& p) {
//     o << "age: " << p.getAge() << ", name: " << p.getName() << endl;
// }

// ostream& operator<<(ostream& o, Person2& p) {
//     o << "age: " << p.getAge() << ", name: " << p.getName() << endl;
// }

// 通过抽象类P提供的统一接口，按实际传入的派生类调用不同派生类的具体实现，从而减少重复代码量
ostream& operator<<(ostream& o, P& p) {
    o << "age: " << p.getAge() << ", name: " << p.getName() << endl;
}

int main() {
    Person p1(18, "Tom");
    p1.show();
    p1.doWork();
    p1.old().show();
    p1+=1;
    p1.show();

    Person2 p2(20, "Jerry");
    p2.show();
    p2.doWork();
    p2.old().show();
    cout << (p2+=1);

    return 0;
}
```



！！！

类如果没有显示定义**拷贝构造函数**和**赋值运算符重载函数**，编译器默认生成的都是进行浅拷贝（仅仅复制对象数据成员的值，对于指针所指向的内容是不会复制的）。

因此多个对象可能会共享同一块内存，当一个对象释放内存后，其他对象指针就变成悬空指针了，会引发未定义行为。

为了防止上述问题，可以实现**深拷贝**，需要显示定义并分配内存。

```c++
#include <iostream>
#include <cstring>
using namespace std;

class Person {
    char *name;
public:
    Person(char *name) {
        cout << "constructor" << endl;
        int len = strlen(name);
        this->name = new char[len + 1];
        strcpy(this->name, name);
    }
    // 拷贝构造函数：深拷贝 
    Person(Person& p) {
        cout << "copy constructor" << endl;
        int len = strlen(p.name);
        this->name = new char[len + 1];
        strcpy(this->name, p.name);
    }
    // 赋值运算符重载：深拷贝
    Person& operator=(Person& p) {
        cout << "operator overload" << endl;
        int len = strlen(p.name);
        this->name = new char[len + 1];
        strcpy(this->name, p.name);
    }
    ~Person() {
        cout << "destructor" << endl;
        delete[] name;
    }
    void show() {
        cout << "name: " << name << endl;
    }
};

void invoke() {
    Person p1("Tom");
    p1.show();
    Person p2(p1);
    p2.show();
    Person p3 = p1; // 用的是拷贝构造函数
    p3.show();
    Person p4("su");
    p4 = p1;    // 赋值运算符重载
    p4.show();
}

int main() {
    try {
        invoke();
        cout << "exec success" << endl;
    } catch(...) {
        // 捕获不到，因为重复释放内存触发的错误是操作系统或底层的问题，C++异常机制是处理不到的。
        cerr << "err" << endl;
    }
    return 0;
}
```



### 类模板

```c++
#include <iostream>
#include <cstring>
using namespace std;

template <class T>
class Array {	// 声明定义实例时，如Array<int> arr 
    int size;
    T *data;
public:
    Array(int n) : size(n) {
        data = new T[n];
    }
    T& operator[] (int i) {
        if ( i < 0 || i >= size) {
            cerr << "out of range" << endl;
            throw "index out of range";
        } else {
            return data[i];
        }
    }
    ~Array() {
        delete[] data;
    }
};

int main() {
    Array<int> intArr(5);
    cout << intArr[0] << endl;
    intArr[0] = 1;
    cout << intArr[0] << endl;

    Array<string> strArr(5);
    cout << strArr[0] << endl;
    strArr[0] = "suyb";
    cout << strArr[0] << endl;
}
```



### 类型别名 

```c++
typedef int INT;
int main(){
    INT i = 3;
    return 0;
}

// c++11
using INT = int;
template <class T>
using Vec = std::vector<T>;
int main(){
    INT i = 3;
    Vec<int> vec = {1,2,3};
    return 0;
}
```



### string - vector

string类

```c++
#include <iostream>
#include <cstring>
#include <string>
using namespace std;

typedef string str;

int main() {
    system("chcp 65001");   // windows上设置字符集为UTF-8，否则中文乱码
    str s = "hello";
    cout << s << endl;
    str s2(s);
    cout << s2 << endl;
    str s3 = s + s2;
    cout << s3 << endl;
    str s4("你好");
    cout << s4 << endl;
    str s5("hello, world", 5);
    cout << s5 << endl;
    cout << s5.size() << endl;
    cout << s5.length() << endl;

    str s6 = u8"我好";  // C++11后使用u8指明UTF-8编码
    cout << s6 << endl;

    str s7(s, 1, 3);    // ell
    cout << s7 << endl;

    str s8(10, 'a');    // aaaaaaaaaa
    cout << s8 << endl;

    str s9(s.begin(), s.end()); // hello 
    cout << s9 << endl;

    str s10 = " world!";
    s.append(s10);
    cout << s << endl;

    s2 += s10;
    cout << s2 << endl;

    s9.insert(5, s10); // hello world!
    cout << s9 << endl;

    s9.erase(5, 6); // hello!
    cout << s9 << endl;

    for (int i = 0; i<s9.size(); i++) {
        cout << s9[i] << " ";
    }
    cout << endl;

    string::iterator it; // 迭代器，可读写，const_iterator只读
    for (it = s9.begin(); it != s9.end(); it++) {
        cout << *it << " ";
    }

    return 0;
}
```



vector类模板

```c++
#include <iostream>
#include <vector>
using namespace std;

int main() {
    vector<double> v1;
    vector<double> v2(10);
    vector<double> v3(10, 3.14);
    vector<double> v4(v3);
    vector<double> v5={1.1, 2.2, 3.3, 4.4, 5.5};
    vector<double> v6(v5.begin(), v5.end());

    cout << "v1 size: " << v1.size() << endl;
    v1.resize(5);
    cout << "v1 size: " << v1.size() << endl;
    v1.clear();
    cout << "v1 size: " << v1.size() << endl;
    cout << "v2[2]: " << v2[2] << endl;
    cout << "v5.at(2): " << v5.at(2) << endl; // at()会检查下标是否越界 
    cout << "v5.front(): " << v5.front() << endl;
    cout << "v5.back(): " << v5.back() << endl;
    v5.push_back(6.6);
    cout << "v5.back(): " << v5.back() << endl;
    v5.pop_back();
    cout << "v5.back(): " << v5.back() << endl;

    vector<double>::iterator it;
    for (it = v5.begin(); it != v5.end(); it++) {
        cout << *it << " ";
    }
    cout << endl;

    for (double num : v5) {    // c++11新特性，遍历容器
        cout << num << " ";
    }

    return 0;
}
```



### 派生类

inheritance继承（derivation派生）：一个派生类（derived class）从一个或多个父类（parent class）/基类（base class）继承，即继承父类的属性和行为，但也有自己的特有属性和行为。

```c++
#include <iostream>
using namespace std;

class Employee {
    string name;
    int age;
public:
    Employee(string name, int age) : name(name), age(age) {}   // 初始化成员列表name(name), age(age)
    void display() {
        cout << "name: " << name << ", age: " << age << endl;
    }
};

class Engineer : public Employee {  // 以public方式继承Employee类，public继承方式，基类的public成员在派生类中仍然是public成员，默认则是private继承方式
    int level;
public:
    // 构造函数初始化列表，先调用基类的构造函数，再调用派生类的构造函数
    // 不允许派生类直接访问基类的私有成员！例如：
    // Engineer(string name, int age, int level) : name(name), age(age), level(level) {}
    Engineer(string name, int age, int level) : Employee(name, age), level(level) {}
    void display() {
        Employee::display();    // 调用基类的display函数
        cout << "level: " << level << endl;
    }
};

int main() {
    Employee e1("Tom", 20);
    e1.display();
    Engineer e2("Jerry", 22, 3);
    e2.display();
    Engineer *p = &e2;
    p->display();
    return 0;
}
```



### 虚函数和多态

```c++
#include <iostream>
using namespace std;

class Employee {
    string name;
    int age;
public:
    Employee(string name, int age) : name(name), age(age) {}
    virtual void display() {  // 虚函数，用virtual关键字修饰，表示该函数是虚函数，可以被派生类重写
        cout << "name: " << name << ", age: " << age << endl;
    }
};

class Engineer : public Employee {
    int level;
public:
    Engineer(string name, int age, int level) : Employee(name, age), level(level) {}
    void display() override {   // override关键字，表示该函数是重写基类的虚函数。可以不显示使用override关键字，但是建议使用，因为可以增加代码可读性，并且帮助编译器检查是否真的重写了基类的虚函数
        Employee::display();
        cout << "level: " << level << endl;
    }
};

int main() {
    Employee e1("Tom", 20);
    Engineer e2("Jerry", 22, 3);
    Employee *p = &e1;
    p->display();
    p = &e2;   // 派生类对象赋值给基类指针, 派生类指针可以自动转换为基类指针
    p->display();   // 如果基类中的成员函数是虚函数，那么调用的是派生类的成员函数，否则调用的是基类的成员函数
    
    // 多态性，不同对象调用同一个函数，产生不同的结果
    return 0;
}
```



### 多重继承

```c++
class Person {};
class Student : public Person {};
class Father: public Person {};
class Me: public Student, public Father {};
```

可能存在菱形继承问题，通过虚继承解决

```c++
class Person {
    int age;
public:
    Person(int age) : age(age) {}
};
class Student : virtual public Person {}; // 虚继承，解决菱形继承问题
class Father: virtual public Person {}; // 虚继承，解决菱形继承问题
class Me: public Student, public Father {}; // 只存在一份Person对象
```



### 纯虚函数和抽象类

```c++
// 纯虚函数
class Person {
public:
    // 抽象类通常用来定义接口，不能实例化对象，只能被继承
    virtual void display() = 0;  // 纯虚函数，没有函数体，派生类必须重写该函数
};
class Man : public Person {
public:
    void display() override {
        cout << "man" << endl;
    }
};
class Woman : public Person {
public:
    void display() override {
        cout << "woman" << endl;
    }
};
```



### 条件编译

```c++
#pragma once // 用于头文件只编译一次，避免重复定义

#define _STRING_

#ifdef _STRING_
#include <iostream>
#include <cstring>
using namespace std;
#elif defined _VECTOR_
#include <iostream>
#include <vector>
using namespace std;
#elif defined _VIRTUAL_
#include <iostream>
using namespace std;
#else 
#include <iostream>
using namespace std;
#endif


#ifndef _STRING_
int main() {
    cout << "Hello, world!" << endl;
    return 0;
}
#endif

// const int a = 5; // 没用，因为宏定义在预处理阶段，而const在编译阶段
#define a 5
#if a > 5
int b = 10;
#elif a < 5
int b = 0;
#else
int b = 5;
#endif

int main() {
    cout << b << endl;
    return 0;
}
```



### const

```c++
#include <iostream>
using namespace std;

int main() {
    // 1. const修饰变量
    const int num = 5;  // const修饰的变量是只读变量，不能被修改
    // num = 6; // error: assignment of read-only variable 'num'

    // 2. const与指针
    // 2.1 常量指针
    int a = 10;
    const int *p1 = &a; // 指向常量的指针，指针指向的值不能被修改
    // *p1 = 20; // error: assignment of read-only location '* p1'
    int b = 20;
    p1 = &b;    // 指针本身可以被修改
    
    // 2.2 指针常量
    int c = 30;
    int *const p2 = &c; // 常量指针，指针本身不能被修改
    *p2 = 40;   // 指针指向的值可以被修改

    // 2.3 常量指针常量
    int d = 50;
    const int *const p3 = &d;   // 指向常量的常量指针，指针和指针指向的值都不能被修改
    // *p3 = 60; // error: assignment of read-only location '* p3'
    // p3 = &a; // error: assignment of read-only variable 'p3'

    // 3. const修饰函数参数

    // 4. const修饰函数返回值

    // 5. const修饰成员函数
    // const修饰成员函数，表示该函数不会修改成员变量，这些函数称为常量成员函数
    class Test {
        int num;
    public:
        Test(int num) : num(num) {}
        int getNum() const {    // 常量成员函数
            // num = 10; // error: assignment of member 'Test::num' in read-only object
            return num;
        }
    };

    return 0;
}
```



### 友类

友类允许一个类访问另一个类的私有(private)和保护(protected)成员。

```c++
class ClassA {
    // 声明 ClassB 为 ClassA 的友类
    friend class ClassB;
private:
    int privateData;
protected:
    int protectedData;
public:
    ClassA(int p, int pr) : privateData(p), protectedData(pr) {}
};

class ClassB {
public:
    void accessClassAData(ClassA& obj) {
        // 由于 ClassB 是 ClassA 的友类，可以访问其私有和保护成员
        std::cout << "ClassA 的私有数据: " << obj.privateData << std::endl;
        std::cout << "ClassA 的保护数据: " << obj.protectedData << std::endl;
    }
};
```

1. 友类的特点

- **单向性**：友类关系是单向的。如果 `ClassA` 声明 `ClassB` 为其友类，并不意味着 `ClassB` 也会自动将 `ClassA` 视为友类。也就是说，`ClassB` 可以访问 `ClassA` 的私有和保护成员，但 `ClassA` 不能自动访问 `ClassB` 的私有和保护成员。
- **非继承性**：友类关系不能被继承。如果 `ClassA` 是 `ClassB` 的友类，`ClassC` 是 `ClassB` 的派生类，`ClassA` 不会自动成为 `ClassC` 的友类。
- **打破封装性**：友类打破了类的封装性原则，因为它允许外部类访问本类的私有和保护成员。所以应该谨慎使用友类，只有在确实需要的情况下才使用，以避免破坏类的信息隐藏和封装特性。

2. 适用场景

- **数据共享**：当两个类之间需要紧密合作，一个类需要频繁访问另一个类的私有或保护成员时，可以使用友类。例如，一个类负责管理数据，另一个类负责对这些数据进行特定的处理。
- **运算符重载**：在实现某些运算符重载时，可能需要访问另一个类的私有成员，这时可以将重载运算符的函数所在的类声明为友类。



### 标准模板库STL

1. **容器**：顺序容器（`vector`、`list`、`deque`）、关联容器（`map`、`set`、`unordered_map`、`unordered_set`）
2. **迭代器**：容器的遍历工具，如 `iterator`、`const_iterator`
3. **算法**：通用算法（排序、查找、复制等）
4. **适配器**：如 `stack`、`queue`、`priority_queue`



### 智能指针

```
std::unique_ptr`、`std::shared_ptr`、`std::weak_ptr
```



### 类型转换

1. 隐式转换
2. 显示转换：静态转换（`static_cast`）、动态转换（`dynamic_cast`）、常量转换（`const_cast`）、重新解释转换（`reinterpret_cast`）
3. C风格转换：`double d = 1.1; int i = (int)d;`



### Lambda表达式

匿名函数对象



### 线程库

多线程编程



### auto/decltype关键字

c++11引入的类型推导机制，用于自动识别变量的类型，都在**编译阶段**推导。

```c++
auto n = 10; // n的类型为int
auto f = 3.14; // f的类型为double
auto p = &n; // p的类型为int*
auto url = "https://example.com"; // url的类型为const char*

// decltype获取表达式的类型
int a = 10;
int b = 20;
decltype(a + b) c = a + b; // c的类型为a + b的结果类型，通常是int
```

