#include <iostream>
#include <cstring>
using namespace std;

template <class T>
class Array {
    int size;
    T *data;
public:
    Array(int n) : size(n) {
        data = new T[n];
    }
    T& operator[] (int i) {
        if ( i < 0 || i >= size) {
            cerr << "out of range" << endl;
            throw "index out of range";
        } else {
            return data[i];
        }
    }
    ~Array() {
        delete[] data;
    }
};

int main() {
    Array<int> intArr(5);
    cout << intArr[0] << endl;
    intArr[0] = 1;
    cout << intArr[0] << endl;

    Array<string> strArr(5);
    cout << strArr[0] << endl;
    strArr[0] = "suyb";
    cout << strArr[0] << endl;
}