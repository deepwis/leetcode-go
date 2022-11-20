package leetcode

// 1337. The K Weakest Rows in a Matrix
//
// You are given an m x n binary matrix mat of 1's (representing soldiers) and 0's (representing civilians). 
// The soldiers are positioned in front of the civilians. That is, all the 1's will appear to the left of 
// all the 0's in each row.
// A row i is weaker than a row j if one of the following is true:
// * The number of soldiers in row i is less than the number of soldiers in row j.
// * Both rows have the same number of soldiers and i < j.
// Return the indices of the k weakest rows in the matrix ordered from weakest to strongest.
//
// Runtime: 21 ms, faster than 63.44% of Go online submissions for The K Weakest Rows in a Matrix.
// Memory Usage: 4.9 MB, less than 80.65% of Go online submissions for The K Weakest Rows in a Matrix.
//
func kWeakestRows(mat [][]int, k int) []int {
    m, n, idx, foundRows, ret := len(mat), len(mat[0]), 0, make(map[int]bool, k), make([]int, k)
    for i := 0; i < n; i++ {
        for j := 0; j < m; j++ {
            if _, ok := foundRows[j]; ok {
                continue
            }
            if mat[j][i] == 0 {
                ret[idx] = j
                foundRows[j] = true
                idx++
                if idx == k {
                    return ret
                }
            }
        }
    }
    if idx < k {
        for j := 0; j < m; j++ {
            if _, ok := foundRows[j]; ok {
                continue
            }
            ret[idx] = j
            foundRows[j] = true
            idx++
            if idx == k {
                return ret
            }
        }
    }
    return ret
}
