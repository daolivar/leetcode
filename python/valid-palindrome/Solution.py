class Solution:
    def isPalindrome(self, s: str) -> bool:
        """
        Determines if the given string is a palindrome, considering only alphanumeric characters
        and ignoring cases.

        A palindrome is a word, phrase, number, or other sequence of characters that reads the 
        same forward and backward (ignoring spaces, punctuation, and case differences).

        Parameters:
            s (str): The input string to check for being a palindrome.

        Returns:
            bool: True if the string is a palindrome, False otherwise.
        """
        left = 0
        right = len(s) - 1

        while start < end:
            # Move start forward if not alphanumeric
            while start < end and not s[start].isalnum():
                start += 1

            # Move end backward if not alphanumeric
            while start < end and not s[end].isalnum():
                end -= 1

            # Compare characters
            if s[start].lower() != s[end].lower():
                return False

            # Move pointers closer towards the middle
            start += 1
            end -= 1
            
        return True
