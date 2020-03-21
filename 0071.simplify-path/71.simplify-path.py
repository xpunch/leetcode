#
# @lc app=leetcode id=71 lang=python3
#
# [71] Simplify Path
#

# @lc code=start


class Solution:
    def simplifyPath(self, path: str) -> str:
        folders, tmp, result = [], '', '/'
        for i in range(len(path)):
            c = path[i]
            if c == '/':
                if len(tmp) == 0:
                    continue
                if tmp == '..':
                    if len(folders) > 0:
                        folders.pop()
                elif tmp != '.':
                    folders.append(tmp)
                tmp = ''
            else:
                tmp += c
                if i == len(path)-1:
                    if tmp == '..':
                        if len(folders) > 0:
                            folders.pop()
                    elif tmp != '.':
                        folders.append(tmp)
        for i in range(len(folders)):
            result += folders[i]
            if i != len(folders)-1:
                result += '/'
        return result
# @lc code=end
