/**
 * Solution class provides methods to count the number of jewels in a set of
 * stones.
 */
class Solution {
    /**
     * Counts the number of jewels in the given set of stones.
     * 
     * Time Complexity: O(m + n) where m is the length of the string jewels and n
     * is the length of the string stones.
     * Converting jewels to a set takes O(m) time, and iterating through stones
     * takes O(n) time.
     * 
     * Space Complexity: O(m) where m is the length of the string jewels. The space
     * used is proportional to the number
     * of unique characters in the jewels string.
     * 
     * @param jewels The string representing types of jewels.
     * @param stones The string representing a set of stones.
     * @return The count of jewels present in the set of stones.
     */
    public int numJewelsInStones(String jewels, String stones) {
        int jewelCount = 0;

        // Convert jewels to a set for efficient lookup
        Set<Character> jewelSet = toSet(jewels);

        // Iterate through each stone and check if it's a jewel
        for (char c : stones.toCharArray()) {
            if (jewelSet.contains(c)) {
                // Increment the count if the stone is a jewel
                jewelCount++;
            }
        }

        // Return the total count of jewels in the set of stones
        return jewelCount;
    }

    /**
     * Converts a string to a set of characters.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input string.
     * Iterates through each character in the string.
     * 
     * Space Complexity: O(m), where 'm' is the number of unique characters in the
     * string.
     * The space used is proportional to the number of unique characters.
     * 
     * @param s The input string.
     * @return A set containing unique characters from the input string.
     */
    public static Set<Character> toSet(String s) {
        Set<Character> set = new HashSet<>();

        // Iterate through each character in the string and add it to the set
        for (char c : s.toCharArray()) {
            set.add(c);
        }

        // Return the set containing unique characters from the string
        return set;
    }
}
