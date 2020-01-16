#
# @lc app=leetcode id=309 lang=python3
#
# [309] Best Time to Buy and Sell Stock with Cooldown
#

# @lc code=start


class Solution:

    def maxProfit(self, prices: List[int]) -> int:
        days = len(prices)
        if days <= 1:
            return 0
        buy = [0 for i in range(days)]
        sell = [0 for i in range(days)]
        buy[0] = -prices[0]
        buy[1] = max(-prices[0], -prices[1])
        sell[1] = max(prices[1]-prices[0], 0)
        for i in range(2, days):
            sell[i] = max(buy[i-1]+prices[i], sell[i-1])
            buy[i] = max(sell[i-2]-prices[i], buy[i-1])
        return max(buy[-1], sell[-1])
        # @lc code=end
