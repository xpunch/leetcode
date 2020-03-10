#
# @lc app=leetcode id=56 lang=python3
#
# [56] Merge Intervals
#
from typing import List
# @lc code=start


class Solution:
    def merge(self, intervals: List[List[int]]) -> List[List[int]]:
        n = len(intervals)
        if n <= 1:
            return intervals
        intervals.sort(key=lambda i: i[0])
        result = [intervals[0]]
        for i in range(1, n):
            last = result[-1]
            interval = intervals[i]
            if interval[0] > last[1]:
                result.append(interval)
            else:
                last[1] = max(last[1], interval[1])
        return result
# @lc code=end
