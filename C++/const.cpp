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