// Your First C++ Program

#include <iostream>
#include <fstream>
#include <list>
#include <sstream>
#include <algorithm>
#include <vector>
using namespace std;

void sortArray(int arr[], int size) {
    std::sort(arr, arr + size); // Sorts the array in ascending order
}

int main() {

    // create text string which is used to output the text file
    string s;

    // open textfile
    ifstream f ("input.txt");

    int len = 0;

    // declare a list of strings
    list<string> listStrings;

    // loop through file
    while (getline (f, s)) {
        // push each new line from file to the back of the list
        listStrings.push_back(s);
        len++;
    }

    // close the file
    f.close();

    string arrStrings[len];
    int i = 0;
    for (string s : listStrings) {
        // cout << s;
        // cout << '\n';
        arrStrings[i] = s;
        i++;
    }

    int numListA[len];
    int numListB[len];

    for (int i = 0; i < len; i++) {
        std::stringstream ss(arrStrings[i]);
        int num1, num2;
        ss >> num1 >> num2;
        numListA[i] = num1;
        numListB[i] = num2;
    }

    sortArray(numListA, len);
    sortArray(numListB, len);

    int sum = 0;
    for (int i = 0; i < len; i++) {
        sum += numListB[i] - numListA[i];
    }
    cout << "Part 1: " << sum << '\n';

    int sum2 = 0;

    for (int i = 0; i < len; i++) {
        int occurences = 0;
        for (int j = 0; j < len; j++) {
            if (numListA[i] == numListB[j]) {
                occurences++;
            }
        }
        sum2 += (numListA[i] * occurences);
        occurences = 0;
    }

    cout << "Part 2: " << sum2 << '\n';

    return 0;
}
