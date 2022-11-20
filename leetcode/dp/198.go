package leetcode

// 198. House Robber
//
// You are a professional robber planning to rob houses along a street. 
// Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them 
// is that adjacent houses have security systems connected and it will automatically contact the police 
// if two adjacent houses were broken into on the same night.
// Given an integer array nums representing the amount of money of each house, 
// return the maximum amount of money you can rob tonight without alerting the police.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for House Robber.
// Memory Usage: 2.1 MB, less than 42.51% of Go online submissions for House Robber.
//
func rob(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    n := len(nums)
    if n == 1 {
        return nums[0]
    }
    maxAmount := make([]int, n + 1)
    maxAmount[0] = 0
    maxAmount[1] = nums[0]
    for i := 2; i <= n; i++ {
        a1, a2 := maxAmount[i-1], maxAmount[i-2]
        a2 += nums[i-1]
        if a1 > a2 {
            maxAmount[i] = a1
        } else {
            maxAmount[i] = a2
        }
    }
    return maxAmount[n]
}

