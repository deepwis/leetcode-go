package leetcode

// 62. Unique Paths
//
// There is a robot on an m x n grid. The robot is initially located at the top-left corner (i.e., grid[0][0]). 
// The robot tries to move to the bottom-right corner (i.e., grid[m - 1][n - 1]). 
// The robot can only move either down or right at any point in time.
// Given the two integers m and n, return the number of possible unique paths that the robot 
// can take to reach the bottom-right corner.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Unique Paths.
// Memory Usage: 2.1 MB, less than 58.64% of Go online submissions for Unique Paths.
//
func uniquePaths(m, n int) int {
    paths := make([][]int, m)
    paths[0] = make([]int, n)
    for i := range paths[0] {
        paths[0][i] = 1
    }
    for i := 1; i < m; i++ {
        paths[i] = make([]int, n)
        paths[i][0] = 1
        for j := 1; j < n; j++ {
            paths[i][j] = paths[i-1][j] + paths[i][j-1]
        }
    }
    return paths[m-1][n-1]
}
