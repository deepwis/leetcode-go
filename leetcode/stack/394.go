package leetcode

import (
    "strconv"
    "bytes"
    "strings"
)

// 394. Decode String
//
// Given an encoded string, return its decoded string.
// The encoding rule is: k[encoded_string], where the encoded_string inside the square brackets is being 
// repeated exactly k times. Note that k is guaranteed to be a positive integer.
// You may assume that the input string is always valid; there are no extra white spaces, square brackets 
// are well-formed, etc. Furthermore, you may assume that the original data does not contain any digits and that 
// digits are only for those repeat numbers, k. For example, there will not be input like 3a or 2[4].
// The test cases are generated so that the length of the output will never exceed 105.
//
// Runtime: 3 ms, faster than 26.65% of Go online submissions for Decode String.
// Memory Usage: 2 MB, less than 91.85% of Go online submissions for Decode String.
//
const (
    _ = iota
    DIGIT
    LETTER

    LEFT_BRACKET
    RIGHT_BRACKET

    UNKNOWN
)
func decodeString(s string) string {
    parsed := []string{}
    buf := new(bytes.Buffer)
    for low := 0; low < len(s); {
        off, pattern := parseTerm(s[low:])
        if pattern != RIGHT_BRACKET {
            parsed = append(parsed, s[low:low+off])
        }
        if pattern == RIGHT_BRACKET {
            plen := len(parsed)
            letters := []string{}
            i := plen-1
            repeatSize := 0
            for i >= 0 {
                if parsed[i] == "[" {
                    break
                }
                letters = append(letters, parsed[i])
                repeatSize += len(parsed[i])
                i--
            }
            times, _ := strconv.Atoi(parsed[i-1])
            buf.Grow(repeatSize)
            for j := len(letters)-1; j >=0; j-- {
                buf.WriteString(letters[j])
            }
            repeatStr := buf.String()
            buf.Reset()

            buf.Grow(repeatSize*times)
            for i := 0; i < times; i++ {
                buf.WriteString(repeatStr)
            }
            parsed[i-1] = buf.String()
            parsed = parsed[:i]
            buf.Reset()
        }
        low += off
    }
    if len(parsed) > 1 {
        return strings.Join(parsed, "")
    }
    return parsed[0]
}

func parseChar(c byte) int {
    if c >= '0' && c <= '9' {
        return DIGIT
    }
    if c >= 'a' && c <= 'z' {
        return LETTER
    }
    if c == '[' {
        return LEFT_BRACKET
    }
    if c == ']' {
        return RIGHT_BRACKET
    }
    return UNKNOWN
}

func parseTerm(s string) (off, pattern int) {
    pattern = parseChar(s[0])
    off = 1
    if pattern == LEFT_BRACKET || pattern == RIGHT_BRACKET {
        return
    }
    i := 1
    for i < len(s) {
        if parseChar(s[i]) != pattern {
            break
        }
        i++
    }
    off = i
    return
}
