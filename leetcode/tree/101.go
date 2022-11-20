package leetcode

// 101. Symmetric Tree
//
// Given the root of a binary tree, check whether it is a mirror of itself (i.e., symmetric around its center).
//
// Runtime: 4 ms, faster than 61.49% of Go online submissions for Symmetric Tree.
// Memory Usage: 2.9 MB, less than 84.46% of Go online submissions for Symmetric Tree.
//
func isSymmetric(root *TreeNode) bool {
    if root == nil {
        return true
    }
    if root.Left == nil && root.Right == nil {
        return true
    }
    if root.Left == nil || root.Right == nil {
        return false
    }
    return compareTwo(root.Left, root.Right)
}

func compareTwo(n1, n2 *TreeNode) bool {
    if n1.Val != n2.Val {
        return false
    }

    if n1.Left == nil && n1.Right == nil && n2.Left == nil && n2.Right == nil {
        return true
    }

    if n1.Left == nil {
        if n2.Right != nil {
            return false
        }
    } else if n2.Right == nil {
        return false
    }
    if n1.Right == nil {
        if n2.Left != nil {
            return false
        }
    } else if n2.Left == nil {
        return false
    }

    if n1.Left != nil && n2.Right != nil && !compareTwo(n1.Left, n2.Right) {
        return false
    }
    if n1.Right != nil && n2.Left != nil && !compareTwo(n1.Right, n2.Left) {
        return false
    }
    return true
}

