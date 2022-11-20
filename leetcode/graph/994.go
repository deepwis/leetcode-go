package leetcode

// 994. Rotting Oranges
//
// You are given an m x n grid where each cell can have one of three values:
// * 0 representing an empty cell,
// * 1 representing a fresh orange, or
// * 2 representing a rotten orange.
// Every minute, any fresh orange that is 4-directionally adjacent to a rotten orange becomes rotten.
// Return the minimum number of minutes that must elapse until no cell has a fresh orange. 
// If this is impossible, return -1.
//
// Runtime: 6 ms, faster than 60.00% of Go online submissions for Rotting Oranges.
// Memory Usage: 3.1 MB, less than 24.35% of Go online submissions for Rotting Oranges.
//
func orangesRotting(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    newRottens  := make([][]int, 0, m * n)
    freshCnt := 0
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 2 {
                newRottens = append(newRottens, []int{i,j})
            }
            if grid[i][j] == 1 {
                freshCnt += 1
            }
        }
    }
    if freshCnt == 0 {
        return 0
    }
    if len(newRottens) == 0 {
        return -1
    }
    minutes := 0
    for len(newRottens) > 0 {
        minutes++
        nextRottens := [][]int{}
        for _, r := range newRottens {
            doRotten(&grid, r[0], r[1], &nextRottens)
        }
        cnt := len(nextRottens)
        if cnt > 0 {
            freshCnt -= cnt
            newRottens = nextRottens
        } else {
            newRottens = [][]int{}
        }
        if freshCnt == 0 {
            break
        }
    }
    if freshCnt > 0 {
        return -1
    }
    return minutes
}

func doRotten(grid *[][]int, i, j int, rottens *[][]int) {
    if i > 0 && (*grid)[i-1][j] == 1 {
        (*grid)[i-1][j] = 2
        *rottens = append(*rottens, []int{i-1, j})
    }
    if i < len(*grid)-1 && (*grid)[i+1][j] == 1 {
        (*grid)[i+1][j] = 2
        *rottens = append(*rottens, []int{i+1, j})
    }
    if j > 0 && (*grid)[i][j-1] == 1 {
        (*grid)[i][j-1] = 2
        *rottens = append(*rottens, []int{i, j-1})
    }
    if j < len((*grid)[0])-1 && (*grid)[i][j+1] == 1 {
        (*grid)[i][j+1] = 2
        *rottens = append(*rottens, []int{i, j+1})
    }
}
