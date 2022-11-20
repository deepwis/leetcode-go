package leetcode

// 378. Kth Smallest Element in a Sorted Matrix
//
// Given an n x n matrix where each of the rows and columns is sorted in ascending order, return the kth 
// smallest element in the matrix.
// Note that it is the kth smallest element in the sorted order, not the kth distinct element.
// You must find a solution with a memory complexity better than O(n2).
//
// Runtime: 42 ms, faster than 77.42% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
// Memory Usage: 7.6 MB, less than 45.97% of Go online submissions for Kth Smallest Element in a Sorted Matrix.
//
func kthSmallestII(matrix [][]int, k int) int {
    n := len(matrix)
    if k == 1 {
        return matrix[0][0]
    }
    if k == n*n {
        return matrix[n-1][n-1]
    }
    min, max := matrix[0][0], matrix[n-1][n-1]
    found := 0
    for min < max {
        // mid := (min + max) / 2
        mid := min + (max-min) / 2
        found = foundLessThan(&matrix, mid)
        if found < k {
            min = mid+1
        } else {
            max = mid
        }
    }
    return min
}

func foundLessThan(matrix *[][]int, target int) int {
    count := 0
    i, j := 0, len(*matrix) - 1
    for i < len(*matrix) && j >= 0 {
        if (*matrix)[i][j] <= target {
            count += (j+1)
            i++
        } else {
            j--
        }
    }
    return count
}
