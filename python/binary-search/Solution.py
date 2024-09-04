class Solution:
    def search(self, nums: List[int], target: int) -> int:
        """
        Performs a binary search to find the index of the target value in a sorted list.

        Binary search is an efficient algorithm for finding an item from a sorted list of items.
        It works by repeatedly dividing the search interval in half.

        Parameters:
            nums (List[int]): A list of integers sorted in ascending order.
            target (int): The integer value to search for within the list.

        Returns:
            int: The index of the target in the list if found, otherwise -1.
        """
        left = 0
        right = len(nums) - 1

        while left <= right:
            middle = left + (right - left) // 2  # Prevents overflow in some languages

            if nums[middle] == target:
                return middle
            elif nums[middle] < target:
                left = middle + 1
            else:
                right = middle - 1

        return -1  # Target not found
