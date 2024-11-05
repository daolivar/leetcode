class Solution(object):
    def containsDuplicate(self, nums):
        """
        Check if the list contains any duplicates.

        :type nums: List[int]  # nums is a list of integers
        :rtype: bool  # The function returns a boolean value
        """

        # Initialize an empty set to keep track of seen numbers.
        s = set()

        # Loop through each number in the list.
        for num in nums:
            # If the current number is already in the set, it's a duplicate, so return True.
            if num in s:
                return True
            # If it's not a duplicate, add the number to the set.
            s.add(num)

        # If no duplicates were found after checking all elements, return False.
        return False
