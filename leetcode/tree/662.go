package leetcode

// 662. Maximum Width of Binary Tree
//
// Given the root of a binary tree, return the maximum width of the given tree.
// The maximum width of a tree is the maximum width among all levels.
// The width of one level is defined as the length between the end-nodes 
// (the leftmost and rightmost non-null nodes), where the null nodes between the end-nodes 
// that would be present in a complete binary tree extending down to that level 
// are also counted into the length calculation.
// It is guaranteed that the answer will in the range of a 32-bit signed integer.
//
// Runtime: 3 ms, faster than 91.49% of Go online submissions for Maximum Width of Binary Tree.
// Memory Usage: 4.6 MB, less than 42.55% of Go online submissions for Maximum Width of Binary Tree.
//
func widthOfBinaryTree(root *TreeNode) int {
    mw := 0
    if root == nil {
        return mw
    }
    indices := []int{0}
    q := []*TreeNode{root}
    for len(q) > 0 {
        n := len(q)
        for i := 0; i < n; i++ {
            n := q[i]
            p := indices[i]
            if n.Left != nil {
                q = append(q, n.Left)
                indices = append(indices, p*2+1)
            }
            if n.Right != nil {
                q = append(q, n.Right)
                indices = append(indices, p*2+2)
            }
        }
        w := indices[n-1] - indices[0] + 1
        if w > mw {
            mw = w
        }
        q = q[n:]
        indices = indices[n:]
    }
    return mw
}

