package leetcode

// 51. N-Queens
//
// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens 
// attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle. You may return the answer in any order.
// Each solution contains a distinct board configuration of the n-queens' placement, 
// where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
// Runtime: 12 ms, faster than 39.61% of Go online submissions for N-Queens.
// Memory Usage: 3.5 MB, less than 38.16% of Go online submissions for N-Queens.
//
func solveNQueens(n int) [][]string {
    if n == 1 {
        return [][]string{[]string{"Q"}}
    } else if n < 4 {
        return [][]string{}
    }
    ret := [][]string{}
    queens := []int{}
    path := make([]byte, n)
    for i := 0; i < n; i++ {
        path[i] = byte('.')
    }
    next := 0
    for i := 0; i < n; {
        invalidCols := make(map[int]bool)
        for r, c := range queens {
            invalidCols[c] = true
            sub := i - r
            if sub == 0 {
                continue
            }
            if c - sub >= 0 {
                invalidCols[c-sub] = true
            }
            if c + sub < n {
                invalidCols[c+sub] = true
            }
        }

        found := -1
        for j := next; j < n; j++ {
            if _, ok := invalidCols[j]; ok {
                continue
            }
            found = j
            break
        }
        if found == -1 {
            if i == 0 {
                break
            }
            next = queens[i-1]+1
            i--
            queens = queens[:i]
        } else {
            queens = append(queens, found)
            if len(queens) == n {
                next = queens[0] + 1
                solution := make([]string, n)
                for r, c := range queens {
                    path[c] = byte('Q')
                    solution[r] = string(path)
                    path[c] = byte('.')
                }
                ret = append(ret, solution)
                next = queens[i-1]+1
                i--
                queens = queens[:i]
            } else {
                i++
                next = 0
            }
        }
    }
    return ret
}

