package leetcode

// 498. Diagonal Traverse
//
// Given an m x n mat mat, return an array of all the elements of the array in a diagonal order.
//
// Runtime: 48 ms, faster than 52.17% of Go online submissions for Diagonal Traverse.
// Memory Usage: 7.2 MB, less than 53.91% of Go online submissions for Diagonal Traverse.
//
func findDiagonalOrder(mat [][]int) []int {
    m, n := len(mat), len(mat[0])
    ans := make([]int, m * n)
    i, j, k := 0, 0, 0
    for {
        if i >= 0 && i < m && j >= 0 && j < n {
            ans[k] = mat[i][j]
            k++
            if k == m * n {
                break
            }
        }
        if (i+j) & 1 == 0 {
            if j == n || i < 0 {
                i++
            } else {
                i--
                j++
            }
        } else {
            if i == m || j < 0 {
                j++
            } else {
                i++
                j--
            }
        }
    }
    return ans
}
