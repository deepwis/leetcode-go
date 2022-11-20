package leetcode

// 1342. Number of Steps to Reduce a Number to Zero
//
// Given an integer num, return the number of steps to reduce it to zero.
// In one step, if the current number is even, you have to divide it by 2, 
// otherwise, you have to subtract 1 from it.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Number of Steps to Reduce a Number to Zero.
// Memory Usage: 1.9 MB, less than 13.84% of Go online submissions for Number of Steps to Reduce a Number to Zero.
//
func numberOfSteps(num int) int {
    if num == 0 {
        return 0
    }
    n, c, offset := num, 0, 8
    for num >> offset != 0 {
        c = offset
        offset += 8
    }
    offset = offset - 7
    for num >> offset != 0 {
        c++
        offset++
    }
    for n & (n - 1) != 0 {
        n = n & (n - 1)
        c++
    }
    return c + 1
}
