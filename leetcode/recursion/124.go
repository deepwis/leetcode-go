package leetcode

// 124. Binary Tree Maximum Path Sum
//
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes 
// in the sequence has an edge connecting them. A node can only appear in the sequence at most once. 
// Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
// Runtime: 31 ms, faster than 56.84% of Go online submissions for Binary Tree Maximum Path Sum.
// Memory Usage: 7.4 MB, less than 94.37% of Go online submissions for Binary Tree Maximum Path Sum.
//
func maxPathSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    p := -1001
    recurse(root, &p)
    return p
}

func recurse(root *TreeNode, p *int) int {
    if root == nil {
        return -1001
    }
    lmax := recurse(root.Left, p)
    rmax := recurse(root.Right, p)
    cmax := root.Val
    if lmax > 0 {
        cmax += lmax
    }
    if rmax > 0 {
        cmax += rmax
    }
    if cmax > *p {
        *p = cmax
    }

    if lmax < 0 && rmax < 0 {
        return root.Val
    }
    if lmax > rmax {
        return lmax + root.Val
    }
    return rmax + root.Val
}
