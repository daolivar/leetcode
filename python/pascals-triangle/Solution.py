class Solution:
    def generate(self, numRows: int) -> List[List[int]]:
        """
        Generate the first numRows of Pascal's triangle.

        Parameters:
        numRows (int): The number of rows of Pascal's triangle to generate.

        Returns:
        List[List[int]]: A list of lists representing the first numRows of Pascal's triangle.
        """
        triangle = []  # Initialize an empty list to hold the rows of Pascal's triangle

        for rowNum in range(numRows):  # Iterate over each row number from 0 to numRows-1
            row = [1] * (rowNum + 1)  # Create a row with all elements initialized to 1

            for i in range(1, rowNum):  # Update the row elements starting from the second element
                row[i] = triangle[rowNum-1][i-1] + triangle[rowNum-1][i]  # Sum the two elements above the current position

            triangle.append(row)  # Append the current row to the triangle

        return triangle  # Return the completed Pascal's triangle
