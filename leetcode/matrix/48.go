package leetcode

// 48. Rotate Image
//
// You are given an n x n 2D matrix representing an image, rotate the image by 90 degrees (clockwise).
// You have to rotate the image in-place, which means you have to modify the input 2D matrix directly. 
// DO NOT allocate another 2D matrix and do the rotation.
//
// Runtime: 4 ms, faster than 32.05% of Go online submissions for Rotate Image.
// Memory Usage: 2.3 MB, less than 67.11% of Go online submissions for Rotate Image.
//
func rotate(matrix [][]int) {
    n := len(matrix)
    if n <= 1 {
        return
    }
    for i := 0; i < n / 2; i++ {
        for j := i; j < n-1-i; j++ {
            matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j], matrix[n-1-j][i] = matrix[n-1-j][i], matrix[i][j], matrix[j][n-1-i], matrix[n-1-i][n-1-j]
        }
    }
}

