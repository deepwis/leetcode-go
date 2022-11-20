package leetcode

// 221. Maximal Square
//
// Given an m x n binary matrix filled with 0's and 1's, find the largest square 
// containing only 1's and return its area.
//
// Runtime: 10 ms, faster than 79.78% of Go online submissions for Maximal Square.
// Memory Usage: 6.7 MB, less than 89.89% of Go online submissions for Maximal Square.
//
func maximalSquare(matrix [][]byte) int {
    ma := 0
    m, n := len(matrix), len(matrix[0])
    nums := make([][]int, m)
    for i := 0; i < m; i++ {
        nums[i] = make([]int, n)
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == '1' && nums[i][j] == 0 {
                dfsNum(&matrix, &nums, i, j, &ma)
            }
        }
    }
    return ma * ma
}
func dfsNum(matrix *[][]byte, nums *[][]int, i, j int, ma *int) int {
    if i == len(*matrix) || j == len((*matrix)[0]) {
        return 0
    }
    if (*matrix)[i][j] == '0' {
        return 0
    }
    if (*nums)[i][j] > 0 {
        return (*nums)[i][j]
    }

    (*nums)[i][j] = 1
    nr := dfsNum(matrix, nums, i, j+1, ma)
    nb := dfsNum(matrix, nums, i+1, j, ma)
    nrb := dfsNum(matrix, nums, i+1, j+1, ma)
    n := 0
    if nr < nb {
        n = nr
    } else {
        n = nb
    }
    if nrb < n {
        n = nrb
    }
    (*nums)[i][j] = n+1
    if *ma < (*nums)[i][j] {
        *ma = (*nums)[i][j]
    }
    return (*nums)[i][j]
}

