package leetcode

// 112. Path Sum
//
// Given the root of a binary tree and an integer targetSum, return true 
// if the tree has a root-to-leaf path such that adding up all the values along the path equals targetSum.
//
// Runtime: 8 ms, faster than 68.29% of Go online submissions for Path Sum.
// Memory Usage: 4.6 MB, less than 92.45% of Go online submissions for Path Sum.
//
func hasPathSum(root *TreeNode, targetSum int) bool {
    if root == nil {
        return false
    }
    if root.Left == nil && root.Right == nil {
        if root.Val == targetSum {
            return true
        }
        return false
    }

    if hasPathSum(root.Left, targetSum - root.Val) {
        return true
    }
    if hasPathSum(root.Right, targetSum - root.Val) {
        return true
    }
    return false
}
