package leetcode

// 34. Find First and Last Position of Element fo Sorted Array
//
// Given an array of integers nums sorted in non-decreasing order, 
// find the starting and ending position of a given target value.
// If target is not found in the array, return [-1, -1].
// You must write an algorithm with O(log n) runtime complexity.
//
// Runtime: 5 ms, faster than 94.67% of Go online submissions for Find First and Last Position of Element in Sorted Array.
// Memory Usage: 3.9 MB, less than 99.89% of Go online submissions for Find First and Last Position of Element in Sorted Array.
//
func searchRange(nums []int, target int) []int {
    ret := []int{-1,-1}
    if len(nums) == 0 {
        return ret
    }
    low, high := 0, len(nums) - 1
    lp, rp := high, low
    for low <= high {
        if nums[low] == target && ret[0] == -1 {
            ret[0] = low
        }
        if nums[high] == target && ret[1] == -1 {
            ret[1] = high
        }
        if (ret[0] != -1 && ret[1] != -1) || low == high {
            return ret
        }

        if low < lp {
            lm := (lp + low) / 2
            if target > nums[lm] {
                low = lm + 1
            } else {
                lp = lm
            }
        }

        if rp < high {
            rm := (rp + high + 1) / 2
            if target < nums[rm] {
                high = rm - 1
            } else {
                rp = rm
            }
        }
    }
    return ret
}

