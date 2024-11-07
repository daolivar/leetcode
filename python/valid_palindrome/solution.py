class Solution(object):
    def isPalindrome(self, s):
        """
        Given a string s, determine if it is a valid palindrome,
        considering only alphanumeric characters and ignoring case.

        :type s: str  # s is the input string
        :rtype: bool  # The function returns a boolean value (True if palindrome, False otherwise)
        """

        # Initialize two pointers: one at the start and one at the end of the string
        left = 0
        right = len(s) - 1

        # Continue comparing characters while the left pointer is less than the right
        while left < right:
            # Move the left pointer to the right as long as it's pointing to a non-alphanumeric character
            while left < right and not s[left].isalnum():
                left += 1

            # Move the right pointer to the left as long as it's pointing to a non-alphanumeric character
            while left < right and not s[right].isalnum():
                right -= 1

            # If the characters at both pointers are not equal (case insensitive), return False
            if s[left].lower() != s[right].lower():
                return False

            # Move the pointers towards each other after the comparison
            left += 1
            right -= 1

        # If no mismatch was found, the string is a palindrome
        return True
