#
# @lc app=leetcode id=554 lang=python3
#
# [554] Brick Wall
#

# @lc code=start
import collections


class Solution:
    def leastBricks(self, wall: List[List[int]]) -> int:
        edges = collections.defaultdict(int)
        count = 0
        for line in wall:
            tmp = 0
            for brick in line[:-1]:
                tmp += brick
                edges[tmp] += 1
                if edges[tmp] > count:
                    count = edges[tmp]
        return len(wall) - count
# @lc code=end
