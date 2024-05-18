/**
 * Solution class for LeetCode problem 125. Valid Palindrome
 */
public class Solution {

    /**
     * Checks whether the given string is a palindrome, considering only alphanumeric characters and ignoring cases.
     * 
     * This method uses the two-pointer technique to compare characters from the start and end of the string
     * moving towards the center. Non-alphanumeric characters are ignored.
     * 
     * @param s the input string
     * @return true if the string is a palindrome, false otherwise
     */
    public boolean isPalindrome(String s) {
        int i = 0; // Initialize the left pointer
        int j = s.length() - 1; // Initialize the right pointer

        while (i < j) { // Continue until the pointers meet in the middle
            // Move the left pointer right past any non-alphanumeric characters
            while (i < j && !Character.isLetterOrDigit(s.charAt(i))) {
                i++;
            }
            // Move the right pointer left past any non-alphanumeric characters
            while (i < j && !Character.isLetterOrDigit(s.charAt(j))) {
                j--;
            }
            // Compare the characters, ignoring case
            if (Character.toLowerCase(s.charAt(i)) != Character.toLowerCase(s.charAt(j))) {
                return false; // If they don't match, it's not a palindrome
            }
            i++; // Move the left pointer right
            j--; // Move the right pointer left
        }

        return true; // If all characters matched, it's a palindrome
    }
}
