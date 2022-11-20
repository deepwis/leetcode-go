package leetcode

// 32. Longest Valid Parentheses
//
// Given a string containing just the characters '(' and ')', find the length of 
// the longest valid (well-formed) parentheses substring.
//
// Runtime: 8 ms, faster than 26.36% of Go online submissions for Longest Valid Parentheses.
// Memory Usage: 3.5 MB, less than 54.09% of Go online submissions for Longest Valid Parentheses.
//
func longestValidParentheses(s string) int {
    if len(s) <=1 {
        return 0
    }
    longest := 0
    l := 0
    unPaired := []int{}
    for i, c := range s {
        if c == ')' {
            if len(unPaired) == 0 {
                if i - l > longest {
                    longest = i - l
                }
                l = i+1
                continue
            }
            unPaired = unPaired[: len(unPaired)-1]
        } else {
            unPaired = append(unPaired, i)
        }
    }
    for _, v := range unPaired {
        if v - l > longest {
            longest = v - l
        }
        l = v+1
    }
    if len(s) - l > longest {
        longest = len(s) - l
    }
    return longest
}
