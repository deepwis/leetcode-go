package leetcode

// 8. String to Integer (atoi)
//
// Implement the myAtoi(string s) function, which converts a string to a 32-bit signed integer 
// (similar to C/C++'s atoi function).
// The algorithm for myAtoi(string s) is as follows:
// * Read in and ignore any leading whitespace.
// * Check if the next character (if not already at the end of the string) is '-' or '+'. 
//   Read this character in if it is either. This determines if the final result is negative or positive 
//   respectively. Assume the result is positive if neither is present.
// * Read in next the characters until the next non-digit character or the end of the input is reached. 
//   The rest of the string is ignored.
// * Convert these digits into an integer (i.e. "123" -> 123, "0032" -> 32). 
//   If no digits were read, then the integer is 0. Change the sign as necessary (from step 2).
// * If the integer is out of the 32-bit signed integer range [-231, 231 - 1], then clamp the integer so that 
//   it remains in the range. Specifically, integers less than -231 should be clamped to -231, 
//   and integers greater than 231 - 1 should be clamped to 231 - 1.
// * Return the integer as the final result.
// Note:
// * Only the space character ' ' is considered a whitespace character.
// * Do not ignore any characters other than the leading whitespace or the rest of the string after the digits.
//
// Runtime: 3 ms, faster than 71.06% of Go online submissions for String to Integer (atoi).
// Memory Usage: 2.2 MB, less than 86.06% of Go online submissions for String to Integer (atoi).
//
func myAtoi(s string) int {
    if len(s) == 0 {
        return 0
    }
    i := 0
    for i < len(s) {
        if s[i] != ' ' {
            break
        }
        i++
    }
    if i == len(s) {
        return 0
    }
    isNeg := false
    if s[i] == '-' {
        isNeg = true
        i++
    } else if s[i] == '+' {
        i++
    }
    if i == len(s) || s[i] - '0' < 0 || s[i] - '0' > 9 { 
        return 0
    }
    num := 0
    max := 1 << 31 - 1
    for i < len(s) {
        d := s[i] - '0'
        if d < 0 || d > 9 {
            break
        }
        num = num * 10 + int(d)
        i++
        if !isNeg && num > max {
            num = max
            break
        }
        if isNeg && num > max+1 {
            num = max+1
            break
        }
    }
    if isNeg {
        num = 0 - num
    }
    return num
}
