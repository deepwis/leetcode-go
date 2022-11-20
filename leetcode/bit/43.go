package leetcode

import "bytes"

// 43. Multiply Strings
//
// Given two non-negative integers num1 and num2 represented as strings, return the product of num1 and num2, 
// also represented as a string.
// Note: You must not use any built-in BigInteger library or convert the inputs to integer directly.
//
// Runtime: 4 ms, faster than 70.45% of Go online submissions for Multiply Strings.
// Memory Usage: 2.3 MB, less than 76.89% of Go online submissions for Multiply Strings.
//
func multiply(num1, num2 string) string {
    digits := "0123456789"
    l1, l2 := len(num1), len(num2)
    mul := make([]int, l1 + l2) //maximum digits in multiply
    b := make([]byte, l1+l2)
    for i := l1-1; i >= 0; i-- {
        n1 := int(num1[i] - '0')
        for j := l2-1 ; j >= 0; j-- {
            n2 := int(num2[j] - '0')
            pos := i+j+1
            mul[pos] += n1 * n2
        }
    }
    for i := len(mul)-1; i >= 0; i-- {
        n1, n2 := mul[i] % 10, mul[i] / 10
        b[i] = digits[n1]
        if i > 0 {
            mul[i-1] += n2
        }
    }
    offset := 0
    for offset < len(b)-1 {
        if b[offset] != '0' {
            break
        }
        offset++
    }
    b = b[offset:]
    buf := new(bytes.Buffer)
    buf.Grow(len(b))
    buf.Write(b)
    return buf.String()
}
