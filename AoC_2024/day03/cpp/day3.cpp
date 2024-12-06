#include <iostream>
#include <fstream>
#include <list>
#include <sstream>
#include <algorithm>
#include <vector>
using namespace std;

int main() {
    cout << "Hello world!\n";

    string s;
    ifstream f ("test.txt");
    int len = 0;
    list<string> listStrings;

    while (getline (f, s)) {
        listStrings.push_back(s);
        len++;
    }
    f.close();

    for (string s : listStrings) {
        cout << s << '\n';
    }

    cout << "len: " << len << '\n';
    cout << "s: " << s << '\n';


    return 0;
}