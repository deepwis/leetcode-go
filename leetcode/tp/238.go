package leetcode

// 238. Product of Array Except Self
//
// Given an integer array nums, return an array answer such that answer[i] is equal to the product of 
// all the elements of nums except nums[i].
// The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.
// You must write an algorithm that runs in O(n) time and without using the division operation.
//
// Runtime: 39 ms, faster than 77.70% of Go online submissions for Product of Array Except Self.
// Memory Usage: 9.1 MB, less than 6.14% of Go online submissions for Product of Array Except Self.
//
func productExceptSelf(nums []int) []int {
    n := len(nums)
    product := make([]int, n)
    fracs := make([][]int, n)
    fracs[0] = make([]int, 2)
    fracs[0][0] = 1
    fracs[n-1] = make([]int, 2)
    fracs[n-1][1] = 1
    left, right := 1, n-2
    for left < n && right >= 0 {
        if left <= right {
            fracs[left] = make([]int, 2)
        }
        if left < right {
            fracs[right] = make([]int, 2)
        }
        fracs[left][0] = fracs[left-1][0] * nums[left-1] // nums[0] * nums[1] * ... * nums[left-1]
        fracs[right][1] = fracs[right+1][1] * nums[right+1] // nums[right+1] * nums[right+2] * ... * nums[n-1]
        left++
        right--
    }
    for i, f := range fracs {
        product[i] = f[0] * f[1]
    }
    return product
}
