package leetcode

// 113. Path Sum II
//
// Given the root of a binary tree and an integer targetSum, return all root-to-leaf paths 
// where the sum of the node values in the path equals targetSum. Each path should be returned 
// as a list of the node values, not node references.
// A root-to-leaf path is a path starting from the root and ending at any leaf node. 
// A leaf is a node with no children.
//
// Runtime: 9 ms, faster than 47.64% of Go online submissions for Path Sum II.
// Memory Usage: 4.4 MB, less than 89.27% of Go online submissions for Path Sum II.
//
func pathSum(root *TreeNode, targetSum int) [][]int {
    ret := [][]int{}
    path := []int{}
    checkSum(root, &targetSum, 0, &path, &ret)
    return ret
}

func checkSum(root *TreeNode, target *int, sum int, path *[]int, ret *[][]int) {
    if root == nil {
        return
    }
    sum += root.Val
    if root.Left == nil && root.Right == nil {
        if sum == *target {
            p := make([]int, len(*path)+1)
            copy(p, *path)
            p[len(*path)] = root.Val
            *ret = append(*ret, p)
        }
        return
    }
    *path = append(*path, root.Val)
    checkSum(root.Left, target, sum, path, ret)
    checkSum(root.Right, target, sum, path, ret)
    *path = (*path)[:len(*path)-1]
}
