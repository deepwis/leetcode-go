package leetcode

// 240. Search a 2D Matrix II
//
// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. 
// This matrix has the following properties:
// Integers in each row are sorted in ascending from left to right.
// Integers in each column are sorted in ascending from top to bottom.
//
// Runtime: 56 ms, faster than 19.93% of Go online submissions for Search a 2D Matrix II.
// Memory Usage: 7.4 MB, less than 34.27% of Go online submissions for Search a 2D Matrix II.
//
// Runtime: 17 ms, faster than 97.98% of Go online submissions for Search a 2D Matrix II.
// Memory Usage: 7 MB, less than 79.76% of Go online submissions for Search a 2D Matrix II.
//
func searchMatrixII(matrix [][]int, target int) bool {
    m, n := len(matrix), len(matrix[0])
    i, j := 0, n-1 // from right top

    for i < m && j >= 0 {
        if matrix[i][j] > target {
            j--
        } else if matrix[i][j] < target {
            i++
        } else {
            return true
        }
    }
    return false
}
