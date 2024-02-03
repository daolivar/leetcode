/**
 * Solution to check if a string is a palindrome, ignoring spaces and
 * punctuation.
 */
class Solution {
    /**
     * Checks if the given string is a palindrome, ignoring spaces and punctuation.
     *
     * Time Complexity: O(n) where n is the length of the input string s. In the
     * worst case, the algorithm iterates through the entire string once.
     * 
     * Space Complexity: O(1) the algorithm uses a constant amount of additional
     * space regardless of the size of the input string.
     * 
     * @param s The input string to be checked.
     * @return {@code true} if the string is a palindrome, {@code false} otherwise.
     */
    public boolean isPalindrome(String s) {
        // Check if the string is empty, in which case it is considered a palindrome
        if (s.isEmpty()) {
            return true;
        }

        // Initialize pointers for the start and end of the string
        int left = 0;
        int right = s.length() - 1;

        // Iterate through the string characters
        while (left <= right) {
            char leftChar = s.charAt(left);
            char rightChar = s.charAt(right);

            // Skip non-alphanumeric characters from the left
            if (!Character.isLetterOrDigit(leftChar)) {
                left++;
            } else if (!Character.isLetterOrDigit(rightChar)) { // Skip non-alphanumeric characters from the right
                right--;
            } else {
                // Compare alphanumeric characters, ignoring case
                if (Character.toLowerCase(leftChar) != Character.toLowerCase(rightChar)) {
                    return false;
                }
                left++;
                right--;
            }
        }

        // If the loop completes, the string is a palindrome
        return true;
    }
}
