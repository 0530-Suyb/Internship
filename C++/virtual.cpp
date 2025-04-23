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