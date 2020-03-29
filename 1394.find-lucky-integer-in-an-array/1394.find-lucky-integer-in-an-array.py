#
# @lc app=leetcode id=1394 lang=python3
#
# [1394] Find Lucky Integer in an Array
#

# @lc code=start
class Solution:
    def findLucky(self, arr: List[int]) -> int:
        times = {}
        for n in arr:
            if times.__contains__(n):
                times[n] += 1
            else:
                times[n] = 1
        result = -1
        for n in times:
            if n == times[n]:
                if n > result:
                    result = n
        return result
# @lc code=end

