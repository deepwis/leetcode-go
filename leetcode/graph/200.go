package leetcode

// 200. Number of Islands
//
// Given an m x n 2D binary grid grid which represents a map of '1's (land) and '0's (water), 
// return the number of islands.
// An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. 
// You may assume all four edges of the grid are all surrounded by water.
//
// Runtime: 2 ms, faster than 93.67% of Go online submissions for Number of Islands.
// Memory Usage: 3.9 MB, less than 80.96% of Go online submissions for Number of Islands.
//
func numIslands(grid [][]byte) int {
    var num int
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] == '1' {
                num++
                gridDFS(grid, i, j)
            }
        }
    }
    return num
}

func gridDFS(grid [][]byte, r, c int) {
    if r < 0 || c < 0 || r >= len(grid) || c >= len(grid[r]) {
        return
    }
    v := grid[r][c]
    if  v != '1' {
        return
    }

    grid[r][c] = '2'
    gridDFS(grid, r - 1, c)
    gridDFS(grid, r + 1, c)
    gridDFS(grid, r, c - 1)
    gridDFS(grid, r, c + 1)
}
