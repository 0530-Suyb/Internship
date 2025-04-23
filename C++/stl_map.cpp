#include <iostream>
#include <map>
#include <set>
using namespace std;

int main() {
    set<int> s1={1,2,4,5,3,0,1};    // 不可重复
    for (int num : s1) {
        cout << num << " ";
    }
    cout << endl;

    multiset<int> s2={1,2,4,5,3,0,1};    // 可重复
    for (int num : s2) {
        cout << num << " ";
    }
    cout << endl;

    // 不可重复
    map<int, string> m1={{1, "one"}, {0, "nil"}, {2, "two"}, {3, "three"}, {0, "zero"}};
    for (const pair<int, string> &p :m1) {
        cout << p.first << ": " << p.second << endl;
    }

    // 可重复
    multimap<int, string> m2={{1, "one"}, {0, "nil"}, {2, "two"}, {3, "three"}, {0, "zero"}};
    for (const auto &p :m2) { // auto自动推导类型
        cout << p.first << ": " << p.second << endl;
    }
}