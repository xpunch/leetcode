#
# @lc app=leetcode id=24 lang=python3
#
# [24] Swap Nodes in Pairs
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def swapPairs(self, head: ListNode) -> ListNode:
        root = head
        if head and head.next:
            root = head.next
        cur, pre, pre_pre = head, None, None
        while cur:
            tmp = cur.next
            if pre:
                pre.next = cur.next
                cur.next = pre
                if pre_pre:
                    pre_pre.next = cur
                pre_pre = pre
                pre = None
            else:
                pre = cur
            cur = tmp
        return root
        
# @lc code=end

