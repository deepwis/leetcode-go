package leetcode

// 94. Binary Tree Inorder Traversal
//
// Given the root of a binary tree, return the inorder traversal of its nodes' values.
//
// Runtime: 2 ms, faster than 43.97% of Go online submissions for Binary Tree Inorder Traversal.
// Memory Usage: 1.9 MB, less than 89.17% of Go online submissions for Binary Tree Inorder Traversal.
//
func inorderTraversal(root *TreeNode) []int {
    ret := []int{}
    if root == nil {
        return ret
    }
    inorder(root, &ret)
    return ret
}

func inorder(root *TreeNode, result *[]int) {
    if root.Left != nil {
        inorder(root.Left, result)
    }
    *result = append(*result, root.Val)
    if root.Right != nil {
        inorder(root.Right, result)
    }
}

