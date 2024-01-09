/**
 * Solution to determine whether two strings are anagrams.
 * An anagram is a word or phrase formed by rearranging the letters of another.
 */
class Solution {
    /**
     * Checks if two strings are anagrams.
     * 
     * Uses a map to store the frequency of characters.
     * 
     * Time Complexity: O(n), where 'n' is the length of the input strings. In the
     * worst case, the algorithm iterates through both strings once.
     * 
     * Space Complexity: O(c), where 'c' is the number of unique characters
     * across both strings. In the worst case, all unique characters from both
     * strings are stored in the map. The maximum space required is proportional to
     * the number of unique characters, and it does not depend on the length of the
     * strings.
     * 
     * @param s The first string to be compared.
     * @param t The second string to be compared.
     * @return {@code true} if the strings are anagrams, {@code false} otherwise.
     */
    public boolean isAnagram(String s, String t) {
        if (s.length() != t.length()) {
            return false;
        }

        // Hash map to store the count of each character
        Map<Character, Integer> map = new HashMap<>();

        // Increment counts for characters in the first string
        for (int i = 0; i < s.length(); i++) {
            map.put(s.charAt(i), map.getOrDefault(s.charAt(i), 0) + 1);
        }

        // Decrement counts for characters in the second string
        for (int i = 0; i < t.length(); i++) {
            map.put(t.charAt(i), map.getOrDefault(t.charAt(i), 0) - 1);

            // If count becomes zero, remove the entry from the map
            if (map.get(t.charAt(i)) == 0) {
                map.remove(t.charAt(i));
            }
        }

        return map.isEmpty();
    }
}
