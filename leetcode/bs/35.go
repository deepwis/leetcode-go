package leetcode

// 35. Search Insert Position
//
// Given a sorted array of distinct integers and a target value, return the index 
// if the target is found. If not, return the index where it would be if it were inserted in order.
// You must write an algorithm with O(log n) runtime complexity.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Search Insert Position.
// Memory Usage: 3 MB, less than 9.01% of Go online submissions for Search Insert Position.
//
func searchInsert(nums []int, target int) int {
    if len(nums) == 0 || nums[0] >= target {
        return 0
    }
    low, high := 0, len(nums)
    for low < high {
        mid := (low + high) / 2
        if nums[mid] == target {
            return mid
        }
        if nums[mid] < target {
            low = mid + 1
        } else {
            high = mid
        }
    }
    return high
}

