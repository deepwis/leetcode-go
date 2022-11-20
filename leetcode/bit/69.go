package leetcode

// 69. Sqrt(x)
//
// Given a non-negative integer x, return the square root of x rounded down to the nearest integer. 
// The returned integer should be non-negative as well.
// You must not use any built-in exponent function or operator.
// For example, do not use pow(x, 0.5) in c++ or x ** 0.5 in python.
//
// Runtime: 16 ms, faster than 30.32% of Go online submissions for Sqrt(x).
// Memory Usage: 2.1 MB, less than 23.07% of Go online submissions for Sqrt(x).
//
func mySqrt(x int) int {
    if x == 0 {
        return 0
    }
    i := 1
    for i*i + 2*i < x { // Use gradient: d(x^2) = 2x
        i++
    }
    return i
}
