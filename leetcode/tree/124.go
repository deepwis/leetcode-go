package leetcode

// 124. Binary Tree Maximum Path Sum
//
// A path in a binary tree is a sequence of nodes where each pair of adjacent nodes 
// in the sequence has an edge connecting them. A node can only appear in the sequence at most once. 
// Note that the path does not need to pass through the root.
// The path sum of a path is the sum of the node's values in the path.
// Given the root of a binary tree, return the maximum path sum of any non-empty path.
//
// Runtime: 301 ms, faster than 5.30% of Go online submissions for Binary Tree Maximum Path Sum.
// Memory Usage: 7.4 MB, less than 82.07% of Go online submissions for Binary Tree Maximum Path Sum.
//
func maxPathSum(root *TreeNode) int {
    maxSum, lmax, rmax, lbranch, rbranch := root.Val, root.Val-1, root.Val-1, 0, 0
    if root.Left != nil {
        lmax = maxPathSum(root.Left)
        lbranch = maxBranchSum(root.Left)
        if lbranch > maxSum {
            maxSum = lbranch
        }
        if lbranch + root.Val > maxSum {
            maxSum = lbranch + root.Val
        }
    }
    if root.Right != nil {
        rmax = maxPathSum(root.Right)
        rbranch = maxBranchSum(root.Right)
        if rbranch > maxSum {
            maxSum = rbranch
        }
        if rbranch + root.Val > maxSum {
            maxSum = rbranch + root.Val
        }
    }
    if lbranch + root.Val + rbranch > maxSum {
        maxSum = lbranch + root.Val + rbranch
    }

    if maxSum < lmax {
        maxSum = lmax
    }

    if maxSum < rmax {
        maxSum = rmax
    }

    return maxSum
}

func maxBranchSum(root *TreeNode) int {
    lmax, rmax, bmax := root.Val - 1, root.Val - 1, root.Val
    if root.Left != nil {
        lmax = maxBranchSum(root.Left) + root.Val
    }
    if root.Right != nil {
        rmax = maxBranchSum(root.Right) + root.Val
    }
    if lmax > bmax {
        bmax = lmax
    }
    if bmax < rmax {
        return rmax
    }
    return bmax
}

