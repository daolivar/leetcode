// Using Sorting
class Solution {
public:
    bool containsDuplicate(vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        for (int i = 0; i < nums.size()-1; i++) {
            if (nums[i] == nums[i+1]) {
                return true;
            }
        }
        return false;
    }
};

// Using Hashmap
class Solution2 {
public:
    bool containsDuplicate(vector<int>& nums) {
        std::unordered_map<int, bool> seen;
        for (auto i : nums) {
           if (seen[i]) {
               return true;
           } else {
               seen[i] = true;
           }
        }
        return false;
    }
};
