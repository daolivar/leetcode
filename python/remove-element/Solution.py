class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        """
        Remove all instances of `val` in-place and return the new length.

        Parameters:
        nums (List[int]): The list of integers from which to remove `val`.
        val (int): The value to be removed from the list.

        Returns:
        int: The new length of the list after removing `val`.
        """
        j = 0  # Pointer to place the next valid element

        for i in range(len(nums)):
            if nums[i] != val:  # Check if the current element is not equal to val
                nums[j] = nums[i]  # Place the valid element at position j
                j += 1  # Increment the pointer j

        return j  # Return the length of the array after removing val
