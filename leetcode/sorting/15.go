package leetcode

import "sort"

// 15. 3Sum
//
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] 
// such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
//
// Runtime: 28 ms, faster than 98.04% of Go online submissions for 3Sum.
// Memory Usage: 7.2 MB, less than 84.91% of Go online submissions for 3Sum.
//
func threeSum(nums []int) [][]int {
    var ret [][]int
    n := len(nums)
    sort.Ints(nums)

    for i := 0; i < n - 2 && nums[i] <= 0; i++ {
        if i > 0 && nums[i] == nums[i-1] {
            continue
        }
        j, k := i + 1, n - 1
        target := -1 * nums[i]
        for j < k {
            sum2 := nums[j] + nums[k]
            if sum2 > target {
                k--
            } else if sum2 < target {
                j++
            } else {
                if !(j > i + 1 && k < n - 1 && nums[j] == nums[j-1] && nums[k] == nums[k+1]) {
                    ret = append(ret, []int{nums[i], nums[j], nums[k]})
                }
                j++
                k--
            }
        }
    }

    return ret
}
