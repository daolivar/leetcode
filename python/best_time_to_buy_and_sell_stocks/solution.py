class Solution(object):
    def maxProfit(self, prices):
        """
        Calculate the maximum profit that can be achieved by buying and selling stock,
        given a list of daily prices.

        :type prices: List[int]  # A list of integers where each element represents the stock price on a given day
        :rtype: int  # The function returns an integer representing the maximum profit
        """

        # Initialize the minimum price to the price on the first day
        minPrice = prices[0]
        # Initialize the maximum profit to zero (no transaction has occurred yet)
        maxProfit = 0

        # Iterate through each price in the list
        for currentPrice in prices:
            # If the current price is lower than the minimum price seen so far, update minPrice
            if currentPrice < minPrice:
                minPrice = currentPrice
            else:
                # If selling at the current price yields a higher profit than maxProfit, update maxProfit
                if currentPrice - minPrice > maxProfit:
                    maxProfit = currentPrice - minPrice

        # Return the maximum profit found
        return maxProfit
