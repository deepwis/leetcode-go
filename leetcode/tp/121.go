package leetcode

// 121. Best Time to Buy and Sell Stock
//
// You are given an array prices where prices[i] is the price of a given stock on the ith day.
// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day 
// in the future to sell that stock.
// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.
//
// Runtime: 274 ms, faster than 39.40% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 9 MB, less than 38.81% of Go online submissions for Best Time to Buy and Sell Stock.
//
func maxProfit(prices []int) int {
    minP, maxP := prices[0], 0
    for i := 1; i < len(prices); i++ {
        if prices[i] - minP > maxP {
            maxP = prices[i] - minP
        }
        if prices[i] < minP {
            minP = prices[i]
        }
    }
    return maxP
}

// Runtime: 237 ms, faster than 61.98% of Go online submissions for Best Time to Buy and Sell Stock.
// Memory Usage: 8.8 MB, less than 66.22% of Go online submissions for Best Time to Buy and Sell Stock.
//
func maxProfit2(prices []int) int {
    curMin, curMax, maxP := prices[0], prices[0], 0
    for i := 1; i < len(prices); i++ {
        if prices[i] > curMax {
            curMax = prices[i]
            if curMax - curMin > maxP {
                maxP = curMax - curMin
            }
        }
        if prices[i] < curMin {
            curMin, curMax = prices[i], prices[i]
        }
    }
    return maxP
}
