package leetcode

import "bytes"

// 402. Remove K Digits
//
// Given string num representing a non-negative integer num, and an integer k, 
// return the smallest possible integer after removing k digits from num.
//
// Runtime: 4 ms, faster than 85.53% of Go online submissions for Remove K Digits.
// Memory Usage: 4.1 MB, less than 90.79% of Go online submissions for Remove K Digits.
//
func removeKdigits(num string, k int) string {
    // remove first digit witch bigger than next one
    n := len(num)
    if k == n {
        return "0"
    }
    stack := make([]byte, 0, n)
    stack = append(stack, num[0])
    for i, j := 1, k; i < n; i++ {
        for j > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
            j--
            stack = stack[:len(stack)-1]
        }
        stack = append(stack, num[i])
    }

    buf := new(bytes.Buffer)
    buf.Grow(n-k)
    buf.Write(stack[:n-k])
    ans := buf.String()
    i := 0
    for i < len(ans) && ans[i] == '0' {
        i++
    }
    if i == len(ans) {
        return "0"
    }
    return ans[i:]
}
