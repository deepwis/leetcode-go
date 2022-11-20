package leetcode

// 226. Invert Binary Tree
//
// Given the root of a binary tree, invert the tree, and return its root.
//
// Runtime: 2 ms, faster than 48.55% of Go online submissions for Invert Binary Tree.
// Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Invert Binary Tree.
//
func invertTree(root *TreeNode) *TreeNode {
    if root == nil {
        return nil
    }
    left := invertTree(root.Left)
    right := invertTree(root.Right)
    root.Left = right
    root.Right = left
    return root
}

