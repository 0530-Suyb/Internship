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