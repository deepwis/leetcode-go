package leetcode

// 96. Unique Binary Search Trees
//
// Given an integer n, return the number of structurally unique BST's (binary search trees) 
// which has exactly n nodes of unique values from 1 to n.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Binary Search Trees.
// Memory Usage: 2 MB, less than 33.33% of Go online submissions for Unique Binary Search Trees.
//
func numTrees(n int) int {
    nTrees := make([]int, n+1)
    nTrees[0] = 1 // n = 0
    nTrees[1] = 1 // n = 1
    for i := 2; i < n+1; i++ {
        num := 0
        for p := 0; p < i; p++ {
            num += nTrees[p] * nTrees[i-p-1]
        }
        nTrees[i] = num
    }
    return nTrees[n]
}
