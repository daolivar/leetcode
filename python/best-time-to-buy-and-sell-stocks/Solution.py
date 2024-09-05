class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        """
        Finds the maximum profit that can be achieved from buying and selling a stock.

        The function uses a sliding window approach to calculate the maximum possible profit 
        from one buy and one sell operation. The 'window' slides over the list of prices, 
        tracking the lowest price seen so far (buying opportunity) and calculating the profit 
        based on the current price (selling opportunity).

        Parameters:
            prices (List[int]): A list of integers representing the stock prices on different days.

        Returns:
            int: The maximum profit that can be achieved. If no profit is possible, returns 0.

        Explanation:
            - The algorithm maintains two variables:
                - `minPrice`: The lowest price encountered so far (starting with infinity).
                - `maxProfit`: The maximum profit found so far (starting with 0).
            - As we iterate through the prices, the window of comparison expands with each price:
                - If the current price is lower than `minPrice`, we shift the "buying" window 
                  to the current price.
                - If the current price is higher, we calculate the profit using the current 
                  price as a "selling" point and compare it to the maximum profit.
            - The sliding window ensures that for each selling price, the algorithm only considers 
              prices that occurred before it as buying opportunities.
        """
        minPrice = float('inf')  # Initialize to infinity, representing the left side of the window
        maxProfit = 0  # Initialize max profit to 0, as the starting point for the right side of the window

        for price in prices:
            # Move the "buy" window to a lower price if found
            if price < minPrice:
                minPrice = price
            else:
                # Calculate the potential profit for the current "sell" window
                profit = price - minPrice
                # Move the "sell" window to the right if a higher profit is found
                if profit > maxProfit:
                    maxProfit = profit

        return maxProfit  # Return the maximum profit found within the window
