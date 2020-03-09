#
# @lc app=leetcode id=1376 lang=python3
#
# [1376] Time Needed to Inform All Employees
#
from typing import List
# @lc code=start


class Solution:
    def numOfMinutes(self, n: int, headID: int, manager: List[int], informTime: List[int]) -> int:
        employees = {}
        for i in range(n):
            m = manager[i]
            if m == -1:
                continue
            if employees.__contains__(m):
                employees[m].append(i)
            else:
                employees[m] = [i]
        return self.getMinutes(headID, employees, informTime)

    def getMinutes(self, e: int, employees: {}, informTime: List[int]) -> int:
        result = informTime[e]
        if employees.__contains__(e):
            maxChildrenTime = 0
            for c in employees[e]:
                childTime = self.getMinutes(c, employees, informTime)
                if childTime > maxChildrenTime:
                    maxChildrenTime = childTime
            result += maxChildrenTime
        return result
