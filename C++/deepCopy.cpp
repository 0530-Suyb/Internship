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