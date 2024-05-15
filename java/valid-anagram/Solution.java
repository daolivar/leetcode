/**
 * Solution class for LeetCode problem 242. Valid Anagram
 */
public class Solution {

    /**
     * Checks whether two strings are anagrams of each other.
     *
     * An anagram is a word or phrase formed by rearranging the letters of a different word or phrase,
     * typically using all the original letters exactly once.
     *
     * @param s the first string
     * @param t the second string
     * @return true if the strings are anagrams, false otherwise
     */
    public boolean isAnagram(String s, String t) {
        // Check if the lengths of the two strings are different
        if (s.length() != t.length()) {
            return false;
        }

        // Create a map to store the frequency of characters in the first string
        Map<Character, Integer> frequencyMap = new HashMap<>();

        // Iterate through each character in the first string and update its frequency in the map
        for (char c : s.toCharArray()) {
            frequencyMap.put(c, frequencyMap.getOrDefault(c, 0) + 1);
        }

        // Iterate through each character in the second string
        for (char c : t.toCharArray()) {
            // Update the frequency of the character in the map
            frequencyMap.put(c, frequencyMap.getOrDefault(c, 0) - 1);

            // If the frequency becomes zero, remove the character from the map
            if (frequencyMap.get(c) == 0) {
                frequencyMap.remove(c);
            }
        }

        // If the map is empty, the strings are anagrams; otherwise, they are not
        return frequencyMap.isEmpty();
    }
}
