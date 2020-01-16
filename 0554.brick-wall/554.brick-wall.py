#
# @lc app=leetcode id=554 lang=python3
#
# [554] Brick Wall
#

# @lc code=start
import copy


class Solution:
    def leastBricks(self, wall: List[List[int]]) -> int:
        height = len(wall)
        edgeCache = {}
        count = 0
        for i in range(height):
            tmp = 0
            for j in range(len(wall[i])-1):
                tmp += wall[i][j]
                edgeCache[tmp] = edgeCache.setdefault(tmp, 0)+1
                if edgeCache[tmp] > count:
                    count = edgeCache[tmp]
        return height - count
# @lc code=end
