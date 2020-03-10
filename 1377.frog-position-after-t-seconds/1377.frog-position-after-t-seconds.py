#
# @lc app=leetcode id=1377 lang=python3
#
# [1377] Frog Position After T Seconds
#
from typing import List
# @lc code=start


class Solution:
    def frogPosition(self, n: int, edges: List[List[int]], t: int, target: int) -> float:
        if target > n or target < 1:
            return 0
        tree = {}
        for edge in edges:
            fromi, toi = edge[0], edge[1]
            if toi != 1:
                if tree.__contains__(fromi):
                    tree[fromi].append(toi)
                else:
                    tree[fromi] = [toi]
            if fromi != 1:
                if tree.__contains__(toi):
                    tree[toi].append(fromi)
                else:
                    tree[toi] = [fromi]
        visited = []
        return self.probability(tree, visited, t, 1, 1, target)

    def probability(self, tree: {}, visited: List[int], t: int, p: float, v: int, target: int) -> float:
        if t == 0:
            if v == target:
                return p
            else:
                return 0
        if v == target:
            hasUnvisitedChildren = False
            if tree.__contains__(v):
                for c in tree[v]:
                    if not visited.__contains__(c):
                        hasUnvisitedChildren = True
                        break
            if hasUnvisitedChildren:
                return 0
            else:
                return p
        else:
            if tree.__contains__(v):
                visited.append(v)
                unvisitedChildren = []
                for c in tree[v]:
                    if not visited.__contains__(c):
                        unvisitedChildren.append(c)
                if len(unvisitedChildren)==0:
                    return 0
                p2 = p * 1/len(unvisitedChildren)
                for c in unvisitedChildren:
                        result = self.probability(
                            tree, visited, t-1,  p2, c, target)
                        if result > 0:
                            return result
                return 0
            else:
                return 0

# @lc code=end
