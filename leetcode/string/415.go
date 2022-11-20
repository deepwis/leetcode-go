package leetcode

import "strings"

// 415. Add Strings
//
// Given two non-negative integers, num1 and num2 represented as string, return the sum of num1 and num2 
// as a string.
// You must solve the problem without using any built-in library for handling large integers (such as BigInteger). 
// You must also not convert the inputs to integers directly.
//
// Runtime: 4 ms, faster than 71.28% of Go online submissions for Add Strings.
// Memory Usage: 2.5 MB, less than 93.62% of Go online submissions for Add Strings.
//
func addStrings(num1, num2 string) string {
    sum, next := 0, 0
    digits := "0123456789"
    b := []byte{}
    for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0; {
        n1, n2 := 0, 0
        if i >= 0 {
            n1 = int(num1[i] - '0')
            i--
        }
        if j >= 0 {
            n2 = int(num2[j] - '0')
            j--
        }
        sum = n1 + n2 + next
        sum, next = sum % 10, sum / 10
        b = append(b, digits[sum])
    }
    if next == 1 {
        b = append(b, '1')
    }
    for i, j := 0, len(b) - 1; i < j; {
        b[i], b[j] = b[j], b[i]
        i++
        j--
    }

    sb := new(strings.Builder)
    sb.Grow(len(b))
    sb.Write(b)

    return sb.String()
}
