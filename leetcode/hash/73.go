package leetcode

// 73. Set Matrix Zeroes
//
// Given an m x n integer matrix matrix, if an element is 0, set its entire row and column to 0's.
// You must do it in place.
//
// Runtime: 17 ms, faster than 76.88% of Go online submissions for Set Matrix Zeroes.
// Memory Usage: 6.3 MB, less than 97.58% of Go online submissions for Set Matrix Zeroes.
//
func setZeroes(matrix [][]int) {
    m := len(matrix)
    if m == 0 {
        return
    }
    zeroRows, zeroCols := make(map[int]bool), make(map[int]bool)

    n := len(matrix[0])
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                zeroRows[i] = true
                if _, ok := zeroCols[j]; !ok {
                    zeroCols[j] = true
                }
            }
            if len(zeroRows) == m || len(zeroCols) == n {
                break
            }
        }
    }
    for i := 0; i < m; i++ {
        if _, ok := zeroRows[i]; ok {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
            continue
        }
        for j := 0; j < n; j++ {
            if _, ok := zeroCols[j]; ok {
                matrix[i][j] = 0
            }
        }
    }
}

