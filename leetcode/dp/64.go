package leetcode

// 64. Minimum Path Sum
//
// Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right, 
// which minimizes the sum of all numbers along its path.
// Note: You can only move either down or right at any point in time.
//
// Runtime: 15 ms, faster than 59.29% of Go online submissions for Minimum Path Sum.
// Memory Usage: 3.9 MB, less than 91.74% of Go online submissions for Minimum Path Sum.
//
func minPathSum(grid [][]int) int {
    m := len(grid)
    if m == 0 {
        return 0
    }
    n := len(grid[0])
    for i := 0; i < m; i++ {
        if i == 0 {
            for j := 1; j < n; j++ {
                grid[i][j] += grid[i][j-1]
            }
        } else {
            grid[i][0] += grid[i-1][0]
            for j := 1; j < n; j++ {
                if grid[i-1][j] < grid[i][j-1] {
                    grid[i][j] += grid[i-1][j]
                } else {
                    grid[i][j] += grid[i][j-1]
                }
            }
        }
    }
    return grid[m-1][n-1]
}
