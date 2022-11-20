package leetcode

// 189. Rotate Array
//
// Given an array, rotate the array to the right by k steps, where k is non-negative.
//
// Runtime: 64 ms, faster than 55.97% of Go online submissions for Rotate Array.
// Memory Usage: 8.8 MB, less than 34.95% of Go online submissions for Rotate Array.
//
func rotate(nums []int, k int) {
    k = k % len(nums)
    if k == 0 {
        return
    }
    low, high, cur := 0, len(nums), len(nums) - k
    for low < cur && cur < high {
        lsize ,rsize := cur - low, high - cur
        if rsize < lsize {
            for i := cur; i < high; i++ {
                nums[i], nums[i-rsize] = nums[i-rsize], nums[i]
            }
            cur -= rsize
            high -= rsize
        } else {
            for i := low; i < cur; i++ {
                nums[i], nums[i+lsize] = nums[i+lsize], nums[i]
            }
            cur += lsize
            low += lsize
        }
    }
}

