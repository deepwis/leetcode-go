package leetcode

// 20. Valid Parentheses
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', 
// determine if the input string is valid.
//
// Runtime: 3 ms, faster than 38.04% of Go online submissions for Valid Parentheses.
// Memory Usage: 2.1 MB, less than 45.97% of Go online submissions for Valid Parentheses.
//
func isValid(s string) bool {
    pairs := make(map[rune]rune)
    pairs['}'] = '{'
    pairs[']'] = '['
    pairs[')'] = '('
    var stack []rune
    l, c := 0, 0
    for _, char := range s {
        v, ok := pairs[char]
        if !ok {
            if l >= len(s) / 2 {
                return false
            }
            if l < c {
                stack[l] = char
            } else {
                stack = append(stack, char)
                c++
            }
            l++
        } else if l == 0 || stack[l-1] != v {
            return false
        } else {
            stack[l-1] = 0
            l--
        }
    }
    if l > 0 {
        return false
    }
    return true
}
