package leetcode

// 152. Maximum Product Subarray
//
// Given an integer array nums, find a contiguous non-empty subarray within the array 
// that has the largest product, and return the product.
// The test cases are generated so that the answer will fit in a 32-bit integer.
// A subarray is a contiguous subsequence of the array.
//
// Runtime: 13 ms, faster than 36.62% of Go online submissions for Maximum Product Subarray.
// Memory Usage: 3.4 MB, less than 79.11% of Go online submissions for Maximum Product Subarray.
//
func maxProduct(nums []int) int {
    if len(nums) == 1 {
        return nums[0]
    }
    maxP, lProduct, rProduct := nums[0]-1, 1, 1 // lProduct: Product from first no-zero value to last no-zero value
    negP := -1 // negP: first negative values index from array left or last zero value
    for i := 0; i < len(nums); i++ {
        lProduct *= nums[i]
        if negP > -1 { // already has at least one negative value
            rProduct *= nums[i]
            if rProduct > maxP {
                maxP = rProduct
            }
        } else if nums[i] < 0 { // nums[i] is first negative value
            negP = i
        }

        if lProduct > maxP {
            maxP = lProduct
        }
        if nums[i] == 0 { // found zero value, reset
            if maxP < 0 {
                maxP = 0
            }
            negP = -1
            lProduct, rProduct = 1, 1
        }
    }
    return maxP
}
