#include <iostream>
#include <vector>
using namespace std;

int main() {
    vector<double> v1;
    vector<double> v2(10);
    vector<double> v3(10, 3.14);
    vector<double> v4(v3);
    vector<double> v5={1.1, 2.2, 3.3, 4.4, 5.5};
    vector<double> v6(v5.begin(), v5.end());

    cout << "v1 size: " << v1.size() << endl;
    v1.resize(5);
    cout << "v1 size: " << v1.size() << endl;
    v1.clear();
    cout << "v1 size: " << v1.size() << endl;
    cout << "v2[2]: " << v2[2] << endl;
    cout << "v5.at(2): " << v5.at(2) << endl; // at()会检查下标是否越界 
    cout << "v5.front(): " << v5.front() << endl;
    cout << "v5.back(): " << v5.back() << endl;
    v5.push_back(6.6);
    cout << "v5.back(): " << v5.back() << endl;
    v5.pop_back();
    cout << "v5.back(): " << v5.back() << endl;

    vector<double>::iterator it;
    for (it = v5.begin(); it != v5.end(); it++) {
        cout << *it << " ";
    }
    cout << endl;

    for (double num : v5) {    // c++11新特性，遍历容器
        cout << num << " ";
    }

    return 0;
}