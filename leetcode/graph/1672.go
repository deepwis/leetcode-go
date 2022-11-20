package leetcode

// 1672. Richest Customer Wealth
//
// You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i^th customer 
// has in the j^th bank. Return the wealth that the richest customer has.
// A customer's wealth is the amount of money they have in all their bank accounts. 
// The richest customer is the customer that has the maximum wealth.
//
// Runtime: 2 ms, faster than 95.32% of Go online submissions for Richest Customer Wealth.
// Memory Usage: 3 MB, less than 100.00% of Go online submissions for Richest Customer Wealth.
//
func maximumWealth(accounts [][]int) int {
    var maximum int
    m, n, total := len(accounts), len(accounts[0]), 0
    for i := 0; i < m; i++ {
        total = 0
        for j := 0; j < n; j++ {
            total += accounts[i][j]
        }
        if total > maximum {
            maximum = total
        }
    }
    return maximum
}

