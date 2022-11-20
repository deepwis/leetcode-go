package leetcode

// 104. Maximum Depth of Binary Tree
//
// Given the root of a binary tree, return its maximum depth.
// A binary tree's maximum depth is the number of nodes along the longest path from the root node down 
// to the farthest leaf node.
//
// Runtime: 12 ms, faster than 17.00% of Go online submissions for Maximum Depth of Binary Tree.
// Memory Usage: 4.3 MB, less than 13.27% of Go online submissions for Maximum Depth of Binary Tree.
//
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    return findDepth(root, 1)
}

func findDepth(root *TreeNode, d int) int {
    max, ld, rd := d, 0, 0
    if root.Left != nil {
        ld = findDepth(root.Left, d+1)
    }
    if root.Right != nil {
        rd = findDepth(root.Right, d+1)
    }
    if ld > max {
        max = ld
    }
    if rd > max {
        max = rd
    }
    return max
}

// Runtime: 3 ms, faster than 91.65% of Go online submissions for Maximum Depth of Binary Tree.
// Memory Usage: 4.4 MB, less than 13.27% of Go online submissions for Maximum Depth of Binary Tree.
//
func maxDepth2(root *TreeNode) int {
    if root == nil {
        return 0
    }
    q, next := []*TreeNode{root}, 0
    md := 0
    for i := next; i < len(q); i++ {
        n := len(q)
        if q[i].Left != nil {
            q = append(q, q[i].Left)
        }
        if q[i].Right != nil {
            q = append(q, q[i].Right)
        }
        if i == next {
            md++
            next = n
        }
    }
    return md
}

