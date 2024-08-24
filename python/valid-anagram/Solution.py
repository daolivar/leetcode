class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        """
        Determines if two strings, s and t, are anagrams of each other.

        An anagram is a word formed by rearranging the letters of another, 
        using all the original letters exactly once.

        Parameters:
            s (str): The first string.
            t (str): The second string.

        Returns:
            bool: True if s and t are anagrams, False otherwise.
        """

        # If the lengths of the strings are not the same, they can't be anagrams.
        if len(s) != len(t):
            return False

        # Initialize a frequency map using defaultdict to count character occurrences.
        freq_map = defaultdict(int)

        # Iterate over each character in both strings.
        for i in range(len(s)):
            freq_map[s[i]] += 1
            freq_map[t[i]] -= 1
            
            # If the count for the character in s drops to 0, remove it from the map.
            if freq_map[s[i]] == 0:
                del freq_map[s[i]]
            # If the count for the character in t drops to 0, remove it from the map.
            if freq_map[t[i]] == 0:
                del freq_map[t[i]]

        # If freq_map is empty, it means all counts balanced out, and the strings are anagrams.
        return not freq_map
