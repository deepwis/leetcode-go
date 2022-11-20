package leetcode

import "strings"

// 151. Reverse Words in a String
//
// Given an input string s, reverse the order of the words.
// A word is defined as a sequence of non-space characters. The words in s will be separated by at least one space.
// Return a string of the words in reverse order concatenated by a single space.
// Note that s may contain leading or trailing spaces or multiple spaces between two words. 
// The returned string should only have a single space separating the words. Do not include any extra spaces.
//
// Runtime: 2 ms, faster than 70.99% of Go online submissions for Reverse Words in a String.
// Memory Usage: 3.1 MB, less than 65.27% of Go online submissions for Reverse Words in a String.
//
func reverseWords(s string) string {
    words := []string{}
    for offset, i := 0, 0; i < len(s); i++ {
        if s[i] == ' ' {
            offset = i+1
            continue
        }
        if i == len(s) - 1 || s[i+1] == ' ' {
            words = append(words, s[offset:i+1])
        }
    }
    switch len(words) {
    case 0:
        return ""
    case 1:
        return words[0]
    }
    n := len(words) - 1 // number of space charactor
    for _, w := range words {
        n += len(w)
    }
    sb := new(strings.Builder)
    sb.Grow(n)
    sb.WriteString(words[len(words)-1])
    for i := len(words)-2; i>= 0; i-- {
        sb.WriteString(" ")
        sb.WriteString(words[i])
    }
    return sb.String()
}
