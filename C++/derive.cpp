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