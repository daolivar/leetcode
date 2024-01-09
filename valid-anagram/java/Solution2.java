/**
 * Solution to determine whether two strings are anagrams. An anagram is a word
 * or phrase formed by rearranging the letters of another.
 */
class Solution2 {
    /**
     * Checks if two strings are anagrams using an array to track character counts.
     *
     * Time Complexity: O(n), where 'n' is the length of the strings. In the worst
     * case, the algorithm iterates through both strings once.
     * 
     * Space Complexity: O(1), this is because the size of the array (int[26]) is
     * constant and does not depend on the size of the input strings.
     * 
     * @param s The first string to be compared.
     * @param t The second string to be compared.
     * @return {@code true} if the strings are anagrams, {@code false} otherwise.
     */
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        // Array to store character counts (assuming lowercase English letters)
        int[] array = new int[26];

        // Increment counts for characters in the first string and decrement counts for
        // characters in the second string
        for (int i = 0; i < s.length(); i++) {
            array[s.charAt(i) - 'a']++;
            array[t.charAt(i) - 'a']--;
        }

        // Check if any character counts are non-zero, indicating the strings are not
        // anagrams
        for (int count : array) {
            if (count != 0) {
                return false;
            }
        }

        // If all character counts are zero, the strings are anagrams
        return true;
    }
}
