#
# @lc app=leetcode id=57 lang=python3
#
# [57] Insert Interval
#
from typing import List
# @lc code=start


class Solution:
    def insert(self, intervals: List[List[int]], newInterval: List[int]) -> List[List[int]]:
        n = len(intervals)
        if n == 0:
            return [newInterval]
        left = intervals[0][0]
        if left > newInterval[1]:
            intervals.insert(0, newInterval)
            return intervals
        if left == newInterval[1]:
            intervals[0][0] = newInterval[0]
            return intervals
        right = intervals[-1][1]
        if right < newInterval[0]:
            intervals.append(newInterval)
            return intervals
        if right == newInterval[0]:
            intervals[-1][1] = newInterval[1]
            return intervals
        result = []
        for i in range(n):
            interval = intervals[i]
            li, ri = interval[0], interval[1]
            if ri < newInterval[0]:
                result.append(interval)
            elif li > newInterval[1]:
                result.append(newInterval)
                result.extend(intervals[i:])
                break
            else:
                newInterval[0] = min(li, newInterval[0])
                newInterval[1] = max(ri, newInterval[1])
                if i == n-1:
                    result.append(newInterval)
        return result
# @lc code=end
