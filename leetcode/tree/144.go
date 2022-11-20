package leetcode

// 144. Binary Tree Preorder Traversal
//
// Given the root of a binary tree, return the preorder traversal of its nodes' values.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Preorder Traversal.
// Memory Usage: 2 MB, less than 43.41% of Go online submissions for Binary Tree Preorder Traversal.
//
func preorderTraversal(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    traversal := []int{}
    preorder(root, &traversal)
    return traversal
}

func preorder(root *TreeNode, traversal *[]int) {
    if root == nil {
        return
    }
    *traversal = append(*traversal, root.Val)
    preorder(root.Left, traversal)
    preorder(root.Right, traversal)
}
