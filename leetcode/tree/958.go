package leetcode

// 958. Check Completeness of a Binary Tree
//
// Given the root of a binary tree, determine if it is a complete binary tree.
// In a complete binary tree, every level, except possibly the last, is completely filled, and all nodes 
// in the last level are as far left as possible. It can have between 1 and 2h nodes inclusive at the last level h.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Check Completeness of a Binary Tree.
// Memory Usage: 3.1 MB, less than 56.41% of Go online submissions for Check Completeness of a Binary Tree.
//
func isCompleteTree(root *TreeNode) bool {
    q := []*TreeNode{root}
    maxLen := 1
    shouldNoChild := false

    for len(q) != 0 {
        n := len(q)
        for i := 0; i < n; i++ {
            left := q[i].Left
            right := q[i].Right
            if left != nil {
                if shouldNoChild || n < maxLen {
                    return false
                }
                q = append(q, left)
            } else {
                shouldNoChild = true
            }
            if right != nil {
                if shouldNoChild || n < maxLen {
                    return false
                }
                q = append(q, right)
            } else {
                shouldNoChild = true
            }
        }
        q = q[n:]
        maxLen *= 2
    }
    return true
}
