#
# @lc app=leetcode id=30 lang=python3
#
# [30] Substring with Concatenation of All Words
#

# @lc code=start


class Solution:

    # segments matched sub word cache
    seg_cache = []
    # last segment matched, work with seg cache
    seg_index = []
    # first segment position
    seg_cur = []
    word_index, word_count, word_len = 0, 0, 0

    def findSubstring(self, s: str, words: List[str]) -> List[int]:
        str_len = len(s)
        if str_len == 0:
            return []
        self.word_count = len(words)
        if self.word_count == 0:
            return []
        self.word_len = len(words[0])
        if self.word_len == 0:
            return []
        self.seg_cache = [{} for i in range(self.word_len)]
        self.seg_index = [0 for i in range(self.word_len)]
        self.seg_cur = [0 for i in range(self.word_len)]
        substr_len = self.word_len * self.word_count
        if str_len < substr_len:
            return []
        result = []
        for i in range(str_len-substr_len+1):
            if self.canBeStartingIndex(s[i:], words):
                result.append(i)
        return result

    def canBeStartingIndex(self, s: str, words: List[str]) -> bool:
        cache = self.seg_cache[self.word_index]
        index = self.seg_index[self.word_index]  # 3
        cur = self.seg_cur[self.word_index]
        match = True
        for i in range(max(index, cur), self.word_count+cur):
            seg = s[(i-cur)*self.word_len:(i-cur+1)*self.word_len]
            found = False
            for j in range(self.word_count):
                if words[j] == seg:
                    if cache.__contains__(j) and cache[j] >= cur:
                        continue
                    else:
                        found = True
                        cache[j] = i
                        self.seg_index[self.word_index] = i+1
                        break
            if not found:
                match = False
                break
        # remove first segment related cache
        for k in list(cache.keys()):
            if cache[k] <= cur:
                del cache[k]
        self.seg_cur[self.word_index] += 1
        self.word_index = (self.word_index+1) % self.word_len
        return match
# @lc code=end
