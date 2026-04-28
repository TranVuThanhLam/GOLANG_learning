#include<iostream>
#include<sstream>
#include<vector>
#include<string>
#include<algorithm>

using namespace std;

// 1. input: vì không có dấu cách nên cin là dc
// 2. biến các dấu | thành space
// 3. dùng stringstream để tách ra và dùng vector để  lưu và in ra

int main() {
    string s;
    cin>>s;
    cout<<s;

    for (int i = 0; i < s.length();i++) {
        if (s[i] == '|') s[i] = ' ';
    }

    stringstream ss(s);
    string tmp;
    vector<string> res;

    while (ss >> tmp) {
        res.push_back(tmp);
    }

    for (int i = 0; i < res.size(); i++) {
        cout << res[i] << "\n";
    }
}
