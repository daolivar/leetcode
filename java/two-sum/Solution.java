import java.util.*;

class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map < Integer, Integer > valToIndex = new HashMap < Integer, Integer > ();
        for (int i = 0; i < nums.length; i++) {
            if (valToIndex.containsKey(target - nums[i])) {
                return new int[] {
                    valToIndex.get(target - nums[i]), i
                };
            } else {
                valToIndex.put(nums[i], i);
            }
        }
        return new int[] {};
    }
}
