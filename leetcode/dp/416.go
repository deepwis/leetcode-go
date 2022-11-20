package leetcode

// 416. Partition Equal Subset Sum
//
// Given a non-empty array nums containing only positive integers, find if the array can be 
// partitioned into two subsets such that the sum of elements in both subsets is equal.
//
// Runtime: 315 ms, faster than 37.62% of Go online submissions for Partition Equal Subset Sum.
// Memory Usage: 6.5 MB, less than 74.92% of Go online submissions for Partition Equal Subset Sum.
//
func canPartition(nums []int) bool {
    totalSum := 0
    for _, v := range nums {
        totalSum += v
    }
    if totalSum & 1 == 1 {
        return false
    }
    // if dp[halfSum] == true, we can find a group of numbers have sum halfSum,
    // because total sum is totalSum == halfSum * 2, result is true
    halfSum := totalSum >> 1
    dp := make(map[int]bool, halfSum+1)
    dp[0] = true
    sums := make([]int,halfSum+1)
    sums = append(sums, 0)
    for _, num := range nums {
        l := len(sums)
        for i := 0; i < l; i++ {
            if _, ok := dp[sums[i] + num]; !ok && sums[i]+num <= halfSum  {
                dp[sums[i] + num] = true
                sums = append(sums, sums[i] + num)
            }
        }
    }

    return dp[halfSum]
}

