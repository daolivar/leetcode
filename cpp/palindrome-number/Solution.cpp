class Solution {
public:
    bool isPalindrome(int x) {
        long copyX = x, reverse = 0;
        while (copyX > 0) {
            reverse = reverse * 10 + copyX % 10;
            copyX /= 10;
        }
        return reverse == x ? 1 : 0;
    }
};
