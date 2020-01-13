#
# @lc app=leetcode id=1298 lang=python3
#
# [1298] Maximum Candies You Can Get from Boxes
#

# @lc code=start


class Solution:

    allKeys = {}
    boxesCollected = []

    def maxCandies(self, status: List[int], candies: List[int], keys: List[List[int]], containedBoxes: List[List[int]], initialBoxes: List[int]) -> int:
        if len(initialBoxes) == 0:
            return 0
        if len(status) == 0:
            return 0
        result = 0
        self.boxesCollected = [False for i in range(len(status))]
        for b in initialBoxes:
            result += self.getBoxCandies(status, candies, keys, containedBoxes, b)
        return result

    def getBoxCandies(self, status: List[int], candies: List[int], keys: List[List[int]], containedBoxes: List[List[int]], initialBox: int) -> int:
        if self.boxesCollected[initialBox]:
            return 0
        count = 0
        if status[initialBox] == 1 or self.allKeys.__contains__(initialBox):
            count += candies[initialBox]
            self.boxesCollected[initialBox] = True
        for k in keys[initialBox]:
            if not self.allKeys.__contains__(k):
                self.allKeys[k] = initialBox
                if not self.boxesCollected[k] and status[k] == 0:
                    count += self.getBoxCandies(status, candies, keys, containedBoxes, k)
        for b in containedBoxes[initialBox]:
            if not self.boxesCollected[b]:
                count += self.getBoxCandies(status, candies, keys, containedBoxes, b)
        return count
# @lc code=end
