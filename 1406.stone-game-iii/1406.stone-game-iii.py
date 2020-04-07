#
# @lc app=leetcode id=1406 lang=python3
#
# [1406] Stone Game III
#
from typing import List
# @lc code=start


class Solution:
    def stoneGameIII(self, stoneValue: List[int]) -> str:
        n = len(stoneValue)
        alice, bob = [0 for i in range(n)], [0 for i in range(n)]
        for i in range(n-1, -1, -1):
            last = n-i
            if last == 1:
                s1 = stoneValue[i]
                alice[i] = s1
                bob[i] = -s1
            elif last == 2:
                s1 = stoneValue[i]
                s2 = s1+stoneValue[i+1]
                alice[i] = max(s1+bob[i+1], s2)
                bob[i] = min(-s1+alice[i+1], -s2)
            elif last == 3:
                s1 = stoneValue[i]
                s2 = s1+stoneValue[i+1]
                s3 = s2+stoneValue[i+2]
                alice[i] = max(s1+bob[i+1], max(s2+bob[i+2], s3))
                bob[i] = min(-s1+alice[i+1], min(-s2+alice[i+2], -s3))
            else:
                s1 = stoneValue[i]
                s2 = s1 + stoneValue[i+1]
                s3 = s2+stoneValue[i+2]
                alice[i] = max(s1+bob[i+1], max(s2+bob[i+2], s3+bob[i+3]))
                bob[i] = min(-s1+alice[i+1], min(-s2+alice[i+2], -s3+alice[i+3]))
        if alice[0] > 0:
            return "Alice"
        elif alice[0] < 0:
            return "Bob"
        else:
            return "Tie"

# @lc code=end
