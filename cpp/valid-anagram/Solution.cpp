#include <string>
#include <unordered_map>
using namespace std;

class Solution {
public:
    bool isAnagram(string s, string t) {
        if (s.length() != t.length()) {
            return false;
        }
        unordered_map<int, int> mp;
        for (int i = 0; i < s.length(); i++) {
            mp[s[i]]++;
            mp[t[i]]--;
        }
        for (auto k : mp) {
            if (k.second != 0) {
                return false;
            }
        }
        return true;
    }
};
