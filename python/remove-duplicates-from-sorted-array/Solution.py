class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        """
        Remove duplicates from a sorted array in-place and return the new length.

        Parameters:
        nums (List[int]): The sorted list of integers from which to remove duplicates.

        Returns:
        int: The new length of the list after removing duplicates.
        """
        if not nums:
            return 0  # If the list is empty, return 0

        j = 0  # Pointer to place the next unique element
        for i in range(1, len(nums)):  # Start iterating from the second element
            if nums[i] != nums[j]:  # If the current element is different from the last unique element
                j += 1  # Move the pointer for the unique element forward
                nums[j] = nums[i]  # Place the current element at the unique position

        return j + 1  # Return the length of the unique elements list
