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