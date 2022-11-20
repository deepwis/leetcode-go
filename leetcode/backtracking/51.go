package leetcode

// 51. N-Queens
//
// The n-queens puzzle is the problem of placing n queens on an n x n chessboard such that no two queens 
// attack each other.
// Given an integer n, return all distinct solutions to the n-queens puzzle. 
// You may return the answer in any order.
// Each solution contains a distinct board configuration of the n-queens' placement, 
// where 'Q' and '.' both indicate a queen and an empty space, respectively.
//
// Runtime: 7 ms, faster than 76.33% of Go online submissions for N-Queens.
// Memory Usage: 3.1 MB, less than 94.20% of Go online submissions for N-Queens.
//
func solveNQueens(n int) [][]string {
    if n == 1 {
        return [][]string{[]string{"Q"}}
    } else if n < 4 {
        return [][]string{}
    }
    ret := [][]string{}
    solution := make([][]byte, n)
    path := make([]byte, n)
    for i := 0; i < n; i++ {
        path[i] = byte('.')
    }
    for i := 0; i < n; i++ {
        solution[i] = make([]byte, n)
        copy(solution[i], path)
    }
    solve(&ret, &solution, 0, n)
    return ret
}

func solve(ret *[][]string, solution *[][]byte, i, n int) {
    if i == n {
        s := make([]string, n)
        for i, v := range *solution {
            s[i] = string(v)
        }
        *ret = append(*ret, s)
        return
    }
    for j := 0; j < n; j++ {
        if checkQueen(solution, i, j) {
            (*solution)[i][j] = 'Q'
            solve(ret, solution, i+1, n)
            (*solution)[i][j] = '.'
        }
    }
}

func checkQueen(solution *[][]byte, i, j int) bool {
    for k := 0; k < i; k++ {
        if (*solution)[k][j] == byte('Q') {
            return false
        }
        c1, c2 := j-i+k, j+i-k
        if c1 >= 0 && (*solution)[k][c1] == byte('Q') {
            return false
        }
        if c2 < len(*solution) && (*solution)[k][c2] == byte('Q') {
            return false
        }
    }
    return true
}
