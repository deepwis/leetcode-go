package leetcode

// 122. Best Time to Buy and Sell Stock II
//
// You are given an integer array prices where prices[i] is the price of a given stock on the ith day.
// On each day, you may decide to buy and/or sell the stock. You can only hold at most one share of the stock 
// at any time. However, you can buy it then immediately sell it on the same day.
// Find and return the maximum profit you can achieve.
//
// Runtime: 6 ms, faster than 80.27% of Go online submissions for Best Time to Buy and Sell Stock II.
// Memory Usage: 3.2 MB, less than 11.72% of Go online submissions for Best Time to Buy and Sell Stock II.
//
func maxProfit(prices []int) int {
    dp := make([]int, len(prices))
    for i := 1; i < len(prices); i++ {
        dp[i] = dp[i-1]
        if prices[i] > prices[i-1] {
            dp[i] += prices[i] - prices[i-1]
        }
    }
    return dp[len(prices)-1]
}
