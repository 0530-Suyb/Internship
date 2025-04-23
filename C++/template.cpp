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

