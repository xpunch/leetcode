#
# @lc app=leetcode id=25 lang=python3
#
# [25] Reverse Nodes in k-Group
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def reverseKGroup(self, head: ListNode, k: int) -> ListNode:
        if k <= 1:
            return head
        root, pre, left = head, None, head
        while left and left.next:
            i, right = 1, left.next
            while right.next and i < k-1:
                right = right.next
                i+=1
            if i == k-1:
                tmp = right.next
                self.reverseList(left, right)
                if root == head:
                    root = right
                if pre:
                    pre.next = right
                pre = left
                left = tmp
            else:
                left = None
        return root
        

    def reverseList(self, head: ListNode, tail: ListNode) -> ListNode:
        if head==None or head.next == None or tail == None:
            return head
        cur = head
        while cur and cur != tail:
            tmp = head.next.next
            head.next.next = cur
            cur = head.next
            head.next = tmp
        return tail

# @lc code=end

