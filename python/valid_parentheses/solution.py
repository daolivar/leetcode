class Solution(object):
    def isValid(self, s):
        """
        Given a string s containing only the characters '(', ')', '{', '}', '[' and ']',
        determine if the input string is valid.

        A valid string must satisfy the following conditions:
        1. Open brackets must be closed by the same type of brackets.
        2. Open brackets must be closed in the correct order.

        :type s: str  # s is the input string containing brackets
        :rtype: bool  # The function returns a boolean value (True if valid, False otherwise)
        """

        # Initialize an empty stack to store opening brackets.
        stack = []

        # Iterate over each character in the string.
        for char in s:
            # If the character is an opening bracket, push it onto the stack.
            if char in '([{':
                stack.append(char)
            else:
                # If the stack is empty (no opening bracket to match), or the current closing 
                # bracket does not match the last opening bracket in the stack, return False.
                if not stack or (char == ')' and stack[-1] != '(') or (char == ']' and stack[-1] != '[') or (char == '}' and stack[-1] != '{'):
                    return False

                # Pop the last opening bracket from the stack as it has been matched.
                stack.pop()

        # If the stack is empty at the end, all brackets were matched correctly, so return True.
        # If the stack is not empty, it means there are unmatched opening brackets, return False.
        return not stack
