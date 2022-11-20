package leetcode

// 695. Max Area of Island
//
// You are given an m x n binary matrix grid. An island is a group of 1's (representing land) 
// connected 4-directionally (horizontal or vertical.) You may assume all four edges of the grid 
// are surrounded by water.
// The area of an island is the number of cells with a value 1 in the island.
// Return the maximum area of an island in grid. If there is no island, return 0.
//
// Runtime: 10 ms, faster than 94.97% of Go online submissions for Max Area of Island.
// Memory Usage: 4.9 MB, less than 97.04% of Go online submissions for Max Area of Island.
//
func maxAreaOfIsland(grid [][]int) int {
    ma := 0
    m, n := len(grid), len(grid[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                a := dfsArea(&grid, i, j)
                if a > ma {
                    ma = a
                }
            }
        }
    }
    return ma
}

func dfsArea(grid *[][]int, i, j int) int {
    if i < 0 || i == len(*grid) || j < 0 || j == len((*grid)[0]) || (*grid)[i][j] != 1 {
        return 0
    }
    a := 1
    (*grid)[i][j] = 2
    a += dfsArea(grid, i-1, j)
    a += dfsArea(grid, i+1, j)
    a += dfsArea(grid, i, j-1)
    a += dfsArea(grid, i, j+1)
    return a
}
