package leetcode

// 110. Balanced Binary Tree
//
// Given a binary tree, determine if it is height-balanced.
//
// Runtime: 9 ms, faster than 73.93% of Go online submissions for Balanced Binary Tree.
// Memory Usage: 5.9 MB, less than 10.52% of Go online submissions for Balanced Binary Tree.
//
func isBalanced(root *TreeNode) bool {
    h := 0
    return isHeightBalanced(root, &h)
}
func isHeightBalanced(root *TreeNode, h *int) bool {
    if root == nil {
        return true
    }
    lh, rh := 0, 0
    lb := isHeightBalanced(root.Left, &lh)
    rb := isHeightBalanced(root.Right, &rh)
    if !lb || !rb {
        return false
    }
    sub := 0
    if lh > rh {
        *h = lh + 1
        sub = lh - rh
    } else {
        *h = rh + 1
        sub = rh - lh
    }
    if sub <= 1 {
        return true
    }
    return false
}
