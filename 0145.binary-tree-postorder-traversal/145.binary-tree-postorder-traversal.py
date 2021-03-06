#
# @lc app=leetcode id=145 lang=python3
#
# [145] Binary Tree Postorder Traversal
#

# @lc code=start
# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, x):
#         self.val = x
#         self.left = None
#         self.right = None


class Solution:
    def postorderTraversal(self, root: TreeNode) -> List[int]:
        if root == None:
            return []
        result = []
        if root.left:
            result = result+self.postorderTraversal(root.left)
        if root.right:
            result = result+self.postorderTraversal(root.right)
        if root.val:
            result.append(root.val)
        return result
# @lc code=end
