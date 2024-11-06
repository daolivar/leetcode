class Solution(object):
    def isAnagram(self, s, t):
        """
        Check if two strings s and t are anagrams of each other.

        :type s: str  # s is the first string
        :type t: str  # t is the second string
        :rtype: bool  # The function returns a boolean value
        """

        # If the strings don't have the same length, they can't be anagrams.
        if len(s) != len(t):
            return False

        # Dictionary to keep track of character counts for both strings.
        char_count = {}

        # Iterate through both strings simultaneously.
        for i in range(len(s)):
            # Increment the count for the character from string s.
            char_count[s[i]] = char_count.get(s[i], 0) + 1
            # Decrement the count for the character from string t.
            char_count[t[i]] = char_count.get(t[i], 0) - 1

        # Check if all values in char_count are zero, meaning both strings
        # have the same frequency of each character.
        for val in char_count.values():
            # If any character has a non-zero count, return False.
            if val != 0:
                return False

        # If all counts are zero, the strings are anagrams, so return True.
        return True
