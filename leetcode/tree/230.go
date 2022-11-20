package leetcode

// 230. Kth Smallest Element in BST
//
// Given the root of a binary search tree, and an integer k, return the kth smallest value (1-indexed) of 
// all the values of the nodes in the tree.
//
// Runtime: 4 ms, faster than 98.74% of Go online submissions for Kth Smallest Element in a BST.
// Memory Usage: 6.3 MB, less than 97.85% of Go online submissions for Kth Smallest Element in a BST.
//
func kthSmallest(root *TreeNode, k int) int {
    var kth **TreeNode = &root // pointer of pointer, we needs to store and edit the node address found in traversal.
    findKthNode(root, k, kth)
    return (*kth).Val
}
func findKthNode(root *TreeNode, k int, kth **TreeNode) int {
    c := 0
    if root.Left != nil {
        c += findKthNode(root.Left, k, kth)
    }
    if c == k {
        return c
    }
    c++
    if c == k {
        *kth = root
        return c
    }
    if root.Right != nil {
        c += findKthNode(root.Right, k-c, kth)
    }
    return c
}

