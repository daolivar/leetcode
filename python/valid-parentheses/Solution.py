class Solution:
    def isValid(self, s: str) -> bool:
        """
        Checks if the input string `s` contains valid parentheses.

        Valid parentheses are balanced and properly nested. The function verifies
        if each opening bracket has a corresponding and correctly ordered closing
        bracket.

        Parameters:
            s (str): The string containing the brackets to be checked.

        Returns:
            bool: True if the string contains valid parentheses, False otherwise.
        """
        stack = []

        for c in s:
            if c in '([{':
                # Push the opening bracket onto the stack
                stack.append(c)
            else:
                # Check if stack is empty before accessing the top element
                if not stack:
                    return False
                top = stack.pop()  # Pop the top element for comparison
                if (c == ')' and top != '(') or (c == ']' and top != '[') or (c == '}' and top != '{'):
                    return False

        # At the end, the stack should be empty if all brackets matched correctly
        return len(stack) == 0
