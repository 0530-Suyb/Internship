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