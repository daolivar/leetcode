class Solution(object):
    def twoSum(self, nums, target):
        """
        Given an array of integers 'nums' and an integer 'target',
        return the indices of the two numbers such that they add up to the target.

        :type nums: List[int]  # nums is a list of integers
        :type target: int  # target is the integer value we need two numbers to sum up to
        :rtype: List[int]  # The function returns a list of two indices
        """

        # Dictionary to store the indices of previously seen numbers
        indices = {}

        # Loop through the list of numbers with both index and value
        for idx, num in enumerate(nums):
            # Calculate the number that, when added to the current number, gives the target
            compliment = target - num

            # If the compliment is already in the dictionary, it means we have found
            # the two numbers that sum up to the target, so return their indices
            if compliment in indices:
                return [indices[compliment], idx]

            # Store the current number's index in the dictionary
            # for future reference if its compliment is found later
            indices[num] = idx

        # No need to return anything here because the problem guarantees that
        # there is exactly one solution. The function will always find and return
        # the correct indices before reaching this point.
