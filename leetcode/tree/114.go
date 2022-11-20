package leetcode

// 114. Flatten Binary Tree to Linked List
//
// Given the root of a binary tree, flatten the tree into a "linked list":
// The "linked list" should use the same TreeNode class where the right child pointer points to 
// the next node in the list and the left child pointer is always null.
// The "linked list" should be in the same order as a pre-order traversal of the binary tree.
//
// Runtime: 8 ms, faster than 16.43% of Go online submissions for Flatten Binary Tree to Linked List.
// Memory Usage: 2.9 MB, less than 33.57% of Go online submissions for Flatten Binary Tree to Linked List.
//
func flatten(root *TreeNode) {
    if root == nil {
        return
    }
    flattenPreorder(root)
}
func flattenPreorder(root *TreeNode) *TreeNode {
    tail, right := root, root.Right
    if root.Left != nil {
        tail = flattenPreorder(root.Left)
        root.Right = root.Left
        root.Left = nil
    }
    if right != nil {
        tail.Right = right
        tail = flattenPreorder(right)
    }
    return tail
}

