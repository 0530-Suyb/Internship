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

    Person& old() {
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

    int getAge () override { return age; }
    char *getName() override  { return name; } // 不安全，将name的指针返回，不过此次暂且如此

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