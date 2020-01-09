#
# @lc app=leetcode id=23 lang=python3
#
# [23] Merge k Sorted Lists
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def mergeKLists(self, lists: List[ListNode]) -> ListNode:
        n = len(lists)
        if n == 0:
            return None
        if n == 1:
            return lists[0]
        pivot = int(n//2)
        left = self.mergeKLists(lists[:pivot])
        right = self.mergeKLists(lists[pivot:])
        head, tail = None, None
        while left or right:
            tmp = None
            if left == None:
                tmp = right
                right = None
            elif right == None:
                tmp = left
                left = None
            else:
                if left.val <= right.val:
                    tmp = left
                    left = left.next
                else:
                    tmp = right
                    right = right.next
            if head == None:
                head = tail = tmp
            else:
                tail.next = tmp
                tail = tmp
        return head
# @lc code=end
