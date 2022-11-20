package leetcode

// 74. Search a 2D Matrix
//
// Write an efficient algorithm that searches for a value target in an m x n integer matrix matrix. 
// This matrix has the following properties:
// Integers in each row are sorted from left to right.
// The first integer of each row is greater than the last integer of the previous row.
//
// Runtime: 4 ms, faster than 69.72% of Go online submissions for Search a 2D Matrix.
// Memory Usage: 2.6 MB, less than 100.00% of Go online submissions for Search a 2D Matrix.
//
// Two pairs of Double Pointer
func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    if m == 0 {
        return false
    }
    n := len(matrix[0])
    lr, hr, lc, hc := 0, m, 0, n
    for lr < hr && lc < hc {
        if matrix[lr][0] > target || matrix[hr-1][hc-1] < target {
            return false
        }
        mr := (hr + lr - 1) / 2
        if matrix[mr][0] > target {
            hr = mr
        } else if matrix[mr][hc-1] < target {
            lr = mr+1
        } else {
            if matrix[mr][lc] > target || matrix[mr][hc-1] < target {
                return false
            }
            mc := (hc + lc -1) / 2
            if matrix[mr][mc] > target {
                hc = mc
            } else if matrix[mr][mc] < target {
                lc = mc+1
            } else {
                return true
            }
        }
    }
    return false
}

