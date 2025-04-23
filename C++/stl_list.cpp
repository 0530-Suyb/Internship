#include <iostream>
#include <list>
using namespace std;

int main() {
    list<int> l1 = {1,2,3};
    l1.push_back(4);
    l1.push_front(0);
    l1.insert(++l1.begin(), 10);
    list<int>::iterator it;
    for (it = l1.begin(); it != l1.end(); it++) {
        cout << *it << " ";
    }
    cout << endl;

    for (int num : l1) { 
        cout << num << " ";
    }
}