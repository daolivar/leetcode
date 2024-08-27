class MinStack:
    """
    A stack that supports push, pop, top, and retrieving the minimum element in constant time.

    Attributes:
        stack (List[int]): A stack to store the elements.
        min_stack (List[int]): A stack to keep track of the minimum elements.
    """
    
    def __init__(self):
        """
        Initializes an empty MinStack.
        """
        self.stack = []       # Main stack to store the elements
        self.min_stack = []   # Stack to keep track of the minimum elements

    def push(self, val: int) -> None:
        """
        Pushes an element onto the stack and updates the minimum stack accordingly.

        Parameters:
            val (int): The value to be pushed onto the stack.
        """
        self.stack.append(val)
        # Update min_stack to keep track of the minimum value
        if not self.min_stack or val <= self.min_stack[-1]:
            self.min_stack.append(val)

    def pop(self) -> None:
        """
        Removes the top element from the stack and updates the minimum stack if necessary.
        """
        top = self.stack.pop()
        # If the popped element is the current minimum, update min_stack
        if top == self.min_stack[-1]:
            self.min_stack.pop()

    def top(self) -> int:
        """
        Retrieves the top element of the stack without removing it.

        Returns:
            int: The top element of the stack.
        """
        return self.stack[-1]

    def getMin(self) -> int:
        """
        Retrieves the minimum element from the stack without removing it.

        Returns:
            int: The minimum element in the stack.
        """
        return self.min_stack[-1]
