package leetcode

// 54. Spiral Matrix
//
// Given an m x n matrix, return all elements of the matrix in spiral order.
//
// Runtime: 2 ms, faster than 42.16% of Go online submissions for Spiral Matrix.
// Memory Usage: 2 MB, less than 9.95% of Go online submissions for Spiral Matrix.
//
func spiralOrder(matrix [][]int) []int {
    m := len(matrix)
    if m ==0 {
        return []int{}
    }

    n := len(matrix[0])
    ret := make([]int, m * n)
    k := 0
    for depth := 0; depth < m/2 && depth < n/2; depth++ {
        for j := depth; j < n-depth-1; j++ {
            ret[k] = matrix[depth][j]
            k++
        }
        for i := depth; i < m-depth-1; i++ {
            ret[k] = matrix[i][n-depth-1]
            k++
        }
        for j := n-depth-1; j > depth; j-- {
            ret[k] = matrix[m-depth-1][j]
            k++
        }
        for i := m-depth-1; i > depth; i-- {
            ret[k] = matrix[i][depth]
            k++
        }
    }

    if m < n && m % 2 == 1 {
        i := m / 2
        for j := i; j < n-i; j++ {
            ret[k] = matrix[i][j]
            k++
        }
        return ret
    }
    if m > n && n % 2 == 1 {
        j := n / 2
        for i := j; i < m-j; i++ {
            ret[k] = matrix[i][j]
            k++
        }
        return ret
    }
    if m == n && m % 2 == 1 {
        i := m / 2
        ret[k] = matrix[i][i]
    }
    return ret
}

// Runtime: 2 ms, faster than 42.01% of Go online submissions for Spiral Matrix.
// Memory Usage: 2 MB, less than 74.49% of Go online submissions for Spiral Matrix.
// 
func spiralOrder2(matrix [][]int) []int {
    m := len(matrix)
    if m ==0 {
        return []int{}
    }

    n := len(matrix[0])
    ret := make([]int, m * n)
    k := 0
    lr, hr, lc, hc, i, j := 0, m, 0, n, 0, 0
    for lr < hr && lc < hc  {
        for j = lc; j < hc; j++ {
            ret[k] = matrix[lr][j]
            k++
        }
        lr++
        if lr == hr {
            break
        }

        for i = lr; i < hr; i++ {
            ret[k] = matrix[i][hc-1]
            k++
        }
        hc--
        if lc == hc {
            break
        }

        for j = hc-1; j >= lc; j-- {
            ret[k] = matrix[hr-1][j]
            k++
        }
        hr--
        if lr == hr {
            break
        }

        for i = hr-1; i >=lr; i-- {
            ret[k] = matrix[i][lc]
            k++
        }
        lc++
    }
    return ret
}

