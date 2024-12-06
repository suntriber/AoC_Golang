#include <iostream>
#include <fstream>
#include <list>
#include <sstream>
#include <algorithm>
#include <vector>
using namespace std;

bool isSafe(int nums[], int size){
    int v = nums[0];

    for (int i = 1; i < size; i++) {
        if (v < nums[i]) {
            if ((nums[i] - v) > 3) {
                return false;
            }
            v = nums[i];
        } else if (v == nums[i]) {
            return false;
        }
    }

    int vv = nums[0];
    for (int i = 1; i < size; i++) {
        if (vv > nums[i]) {
            if ((vv - nums[i]) > 3) {
                return false;
            }
            vv = nums[i];
        } else if (vv == nums[i]) {
            return false;
        }
    }

    int vvv = nums[0];
    bool increasing = false;
    bool decreasing = false;
    for (int i = 1; i < size; i++) {
        if (vvv == nums[i]) {
            return false;
        } else if (vvv < nums[i]) {
            decreasing = true;
            if (increasing) {
                return false;
            }
        } else if (vvv > nums[i]) {
            increasing = true;
            if (decreasing) {
                return false;
            }
        }
        vvv = nums[i];
    }
    return true;
}

int main() {

    string s;
    ifstream f ("input.txt");
    int len = 0;
    list<string> listStrings;

    while (getline (f, s)) {
        listStrings.push_back(s);
        len++;
    }
    f.close();

    string arrStrings[len];
    int i = 0;
    for (string s : listStrings) {
        arrStrings[i] = s;
        i++;
    }

    int part1 = 0;
    for (int i = 0; i < len; i++) {
        std::stringstream ss(arrStrings[i]);
        std::vector<int> numbers;
        int n;
        while (ss >> n) {
            numbers.push_back(n);
        }
        int nums[numbers.size()];
        int j = 0;
        for (int m : numbers) {
            nums[j] = m;
            j++;
        }
        if (isSafe(nums, numbers.size())) {
            part1++;
        }
    }

    cout << "Part 1: " << part1 << '\n';

    int part2 = 0;
    for (int i = 0; i < len; i++) {
        std::stringstream ss(arrStrings[i]);
        std::vector<int> numbers;
        int n;
        while(ss >> n) {
            numbers.push_back(n);
        }
        int nums[numbers.size()];
        int j = 0;
        for (int m : numbers) {
            nums[j] = m;
            j++;
        }
        if (isSafe(nums, numbers.size())) {
            part2++;
            continue;
        } else {
            for (int j = 0; j < numbers.size(); j++) {
                std::vector<int> tmpNums;
                for (int k = 0; k < numbers.size(); k++) {
                    if (j != k) {
                        tmpNums.push_back(nums[k]);
                    }
                }
                int n2[tmpNums.size()];
                int l = 0;
                for (int h : tmpNums) {
                    n2[l] = h;
                    l++;
                }
                if (isSafe(n2, tmpNums.size())) {
                    part2++;
                    break;
                }
            }
        }
    }
    cout << "Part 2: " << part2 << '\n';

    return 0;
}