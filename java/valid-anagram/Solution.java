import java.util.Arrays;

class Solution {
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }
        char array_s[] = s.toCharArray();
        char array_t[] = t.toCharArray();
        Arrays.sort(array_s);
        Arrays.sort(array_t);
        for (int i = 0; i < array_t.length; i++) {
            if (array_s[i] != array_t[i]) {
                return false;
            }
        }
        return true;
    }
}
