class Solution(object):
    def search(self, nums, target):
        """
        Perform binary search on a sorted array to find the target's index.

        :type nums: List[int]  # nums is a sorted list of integers
        :type target: int  # target is the integer value we are searching for
        :rtype: int  # The function returns the index of the target, or -1 if the target is not found
        """

        # Initialize the low and high pointers for binary search
        low = 0
        high = len(nums) - 1

        # Continue searching while the low pointer is less than or equal to the high pointer
        while low <= high:
            # Calculate the middle index (prevents overflow in languages where large integer sums can cause overflow)
            middle = low + (high - low) // 2  

            # If the target is found at the middle index, return the index
            if nums[middle] == target:
                return middle
            # If the target is greater than the middle element, search the right half
            elif nums[middle] < target:
                low = middle + 1
            # If the target is smaller than the middle element, search the left half
            else:
                high = middle - 1

        # If the target is not found, return -1
        return -1
