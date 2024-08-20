class Solution:
    def majorityElement(self, nums: List[int]) -> int:
        """
        Find the majority element in a list of integers.

        The majority element is the element that appears more than half of the time in the list.
        
        Parameters:
            nums (List[int]): A list of integers where we want to find the majority element.

        Returns:
            int: The majority element if one exists, otherwise -1.
        
        Note:
            This function assumes that there is always a majority element in the input list,
            but the return value -1 is used as a fallback in case the input list is empty or no
            element satisfies the majority condition (although by problem constraints, a majority 
            element is expected).
        """
        counts = defaultdict(int)
        majorElem = len(nums) // 2

        for num in nums:
            counts[num] += 1

            if counts[num] > majorElem:
                return num

        return -1
