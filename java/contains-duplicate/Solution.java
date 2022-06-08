import java.util.*;

// Using Sorting
class Solution {
    public boolean containsDuplicate(int[] nums) {
        Arrays.sort(nums);
        for (int i = 0; i < nums.length - 1; i++) {
            if (nums[i] == nums[i + 1]) {
                return true;
            }
        }
        return false;
    }
}

// Using Hashmap
class Solution2 {
    public boolean containsDuplicate(int[] nums) {
        Map < Integer, Boolean > mp = new HashMap < Integer, Boolean > ();
        for (int i: nums) {
            if (mp.containsKey(i)) {
                return true;
            } else {
                mp.put(i, true);
            }
        }
        return false;
    }
}
