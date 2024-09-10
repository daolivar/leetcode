class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        """
        Checks if the input list contains any duplicate elements.

        Parameters:
            nums (List[int]): A list of integers to check for duplicates.

        Returns:
            bool: True if there is at least one duplicate element in the list, False otherwise.
        
        Explanation:
            A set named `seen` tracks elements as they are encountered.
            During iteration, if an element is found in the `seen` set, a duplicate is detected, and True is returned.
            If no duplicates are found by the end of the iteration, False is returned.
        """

        # Initialize a set to track seen values.
        seen = set()

        # Iterate through the list of numbers.
        for num in nums:
            # Return True if the current number is already in the set.
            if num in seen:
                return True
            # Otherwise, add the number to the set.
            seen.add(num)
        
        # Return False if no duplicates are found.
        return False
