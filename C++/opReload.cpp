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
    return p1;
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

bool operator==(const Point& p1, const Point& p2) {
    return p1.x == p2.x && p1.y == p2.y;
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
    p3 += p1 += Point{1, 1}; // 连续赋值
    cout << p1; // (2, 3)
    cout << p3; // (6, 9)
    ++p3;
    cout << p3; // (7, 10)
    cout << p3++; // (7, 10)
    cout << p3; // (8, 11)

    bool res = p1 == p2;
    cout << res << endl; // 0

    return 0;

    // c++编译
    // g++ main.cpp -o main

    // 运算符重载的英文
    // operator overloading
}