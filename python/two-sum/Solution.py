class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        """
        Find two numbers in the list that add up to the target.

        Parameters:
        nums (List[int]): List of integers.
        target (int): Target sum.

        Returns:
        List[int]: Indices of the two numbers that add up to the target.
        """
        diffs = {}  # Dictionary to store numbers and their indices

        for i, num in enumerate(nums):
            complement = target - num  # Calculate the complement of the current number
            if complement in diffs:  # Check if the complement is already in the dictionary
                return [diffs[complement], i]  # Return the indices of the complement and the current number

            diffs[num] = i  # Store the current number and its index in the dictionary

        return []  # Return an empty list if no solution is found
