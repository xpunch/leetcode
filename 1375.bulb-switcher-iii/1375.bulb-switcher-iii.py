#
# @lc app=leetcode id=1375 lang=python3
#
# [1375] Bulb Switcher III
#
from typing import List
# @lc code=start


class Solution:
    def numTimesAllBlue(self, light: List[int]) -> int:
        result, lastBlue, maxLight, cache = 0, 0, 0, {}
        for i in range(len(light)):
            l = light[i]
            cache[l] = True
            if l > maxLight:
                maxLight = l
            if lastBlue == l-1:
                lastBlue = l
            else:
                continue
            allBlue = True
            for j in range(lastBlue+1, maxLight+1):
                if cache.__contains__(j) and cache[j]:
                    lastBlue = j
                else:
                    allBlue = False
                    break
            if allBlue:
                result += 1
        return result

# @lc code=end
