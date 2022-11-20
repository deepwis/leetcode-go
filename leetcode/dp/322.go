package leetcode

import "sort"

// 322. Coin Change
//
// You are given an integer array coins representing coins of different denominations and an integer amount 
// representing a total amount of money.
// Return the fewest number of coins that you need to make up that amount. 
// If that amount of money cannot be made up by any combination of the coins, return -1.
// You may assume that you have an infinite number of each kind of coin.
//
// Runtime: 478 ms, faster than 10.37% of Go online submissions for Coin Change.
// Memory Usage: 7.1 MB, less than 48.69% of Go online submissions for Coin Change.
//
func coinChange(coins []int, amount int) int {
    if amount == 0 {
        return 0
    }
    minChanges := make(map[int]int, amount+1)
    sort.Ints(coins)
    if amount < coins[0] {
        return -1
    }
    for i := 0; i < len(coins); i++ {
        if coins[i] > amount {
            coins = coins[:i]
            break
        }
    }
    minChanges[0] = 0
    for i := 0; i < len(coins); i++ {
        for j := coins[i]; j <= coins[len(coins)-1]; j++ {
            minC := j+1
            if v, ok := minChanges[j]; ok {
                minC = v
            }
            if v, ok := minChanges[j - coins[i]]; ok && v+1 < minC {
                minChanges[j] = v+1
            }
        }
    }
    for i := 0; i < len(coins); i++ {
        for j := coins[len(coins)-1]+1; j <= amount; j++ {
            minC := j+1
            if v, ok := minChanges[j]; ok {
                minC = v
            }
            if v, ok := minChanges[j - coins[i]]; ok && v+1 < minC {
                minChanges[j] = v+1
            }
        }
    }
    if _, ok := minChanges[amount]; !ok {
        return -1
    }
    return minChanges[amount]
}

