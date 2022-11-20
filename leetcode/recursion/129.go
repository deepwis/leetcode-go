package leetcode

// 129. Sum Root to Leaf numbers
//
// You are given the root of a binary tree containing digits from 0 to 9 only.
// Each root-to-leaf path in the tree represents a number.
// For example, the root-to-leaf path 1 -> 2 -> 3 represents the number 123.
// Return the total sum of all root-to-leaf numbers. Test cases are generated 
// so that the answer will fit in a 32-bit integer.
// A leaf node is a node with no children.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Sum Root to Leaf Numbers.
// Memory Usage: 2.2 MB, less than 77.70% of Go online submissions for Sum Root to Leaf Numbers.
//
func sumNumbers(root *TreeNode) int {
    f := 0
    return sumRecursive(root, f)
}

func sumRecursive(root *TreeNode, f int) int {
    if root == nil {
        return 0
    }
    f = f * 10 + root.Val
    if root.Left == nil && root.Right == nil {
        return f
    }
    ls := sumRecursive(root.Left, f)
    rs := sumRecursive(root.Right, f)
    return ls + rs
}
