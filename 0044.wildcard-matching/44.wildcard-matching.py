#
# @lc app=leetcode id=44 lang=python3
#
# [44] Wildcard Matching
#

# @lc code=start


class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        sn = len(s)
        if len(p) == 0:
            return sn == 0
        if sn == 0:
            return len(p) == 0 or p == '*'
        patterns, minLength = self.split(p)
        pn = len(patterns)
        if sn < minLength:
            return False
        si, pi, cache, rollback = 0, 0, [], False
        for pattern in patterns:
            if pattern == '*':
                cache.append(0)
            else:
                cache.append(len(pattern))
        while pi < pn:
            if rollback:
                if patterns[pi] == '*':
                    minLastSegmentsLength = 0
                    for i in range(pi+1, pn):
                        minLastSegmentsLength += cache[i]
                    maxSegLength = sn - si - minLastSegmentsLength+1
                    if cache[pi] < maxSegLength:
                        cache[pi] += 1
                        si += cache[pi]
                        pi += 1
                        rollback = False
                    else:
                        if pi == 0:
                            return False
                        pi -= 1
                        si -= cache[pi]
                        cache[pi] = 0
                else:
                    if pi == 0:
                        return False
                    pi -= 1
                    si -= cache[pi]
            else:
                if patterns[pi] == '*':
                    if pi == pn-1:
                        return True
                    else:
                        cache[pi] = 0
                        pi += 1
                else:
                    if self.isWildcardMatch(s[si:si+cache[pi]], patterns[pi]):
                        if pi == pn-1:
                            if si + cache[pi] == sn:
                                return True
                            else:
                                if pi == 0:
                                    return False
                                pi -= 1
                                si -= cache[pi]
                                rollback = True
                        else:
                            si += cache[pi]
                            pi += 1
                    else:
                        if pi == 0:
                            return False
                        pi -= 1
                        si -= cache[pi]
                        rollback = True

    def isWildcardMatch(self, s: str, p: str) -> bool:
        print(s, p)
        sn, pn = len(s), len(p)
        if sn != pn:
            return False
        for i in range(sn):
            if p[i] != '?' and s[i] != p[i]:
                return False
        return True

    def split(self, p: str) -> ([], int):
        segments, minLength, n, head = [], 0, len(p), 0
        for i in range(n):
            if p[i] == '*':
                if i == 0 or p[i-1] != '*':
                    if head != i:
                        segments.append(p[head:i])
                        minLength += i-head
                    segments.append('*')
                head = i+1
            elif i == n-1:
                segments.append(p[head:i+1])
        return segments, minLength

# @lc code=end
