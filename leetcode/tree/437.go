package leetcode

// 437. Path Sum III
//
// Given the root of a binary tree and an integer targetSum, return the number of paths 
// where the sum of the values along the path equals targetSum.
// The path does not need to start or end at the root or a leaf, but it must go downwards 
// (i.e., traveling only from parent nodes to child nodes).
//
// Runtime: 35 ms, faster than 27.68% of Go online submissions for Path Sum III.
// Memory Usage: 10.4 MB, less than 7.59% of Go online submissions for Path Sum III.
//
func pathSumIII(root *TreeNode, targetSum int) int {
    pathSums := []int{}
    return recPathSum(root, targetSum, pathSums)
}

func recPathSum(node *TreeNode, targetSum int, pathSums []int) int {
    if node == nil {
        return 0
    }
    cnt := 0
    for i := 0; i < len(pathSums); i++ {
        (pathSums)[i] += node.Val
        if (pathSums)[i] == targetSum {
            cnt++
        }
    }
    pathSums = append(pathSums, node.Val)
    if node.Val == targetSum {
        cnt++
    }
    rightSums := make([]int, len(pathSums))
    copy(rightSums, pathSums)
    cnt += recPathSum(node.Left, targetSum, pathSums)
    cnt += recPathSum(node.Right, targetSum, rightSums)
    return cnt
}

