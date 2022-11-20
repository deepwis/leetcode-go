package leetcode

import "strconv"

// 227. Basic Calculator II
//
// Given a string s which represents an expression, evaluate this expression and return its value. 
// The integer division should truncate toward zero.
// You may assume that the given expression is always valid. All intermediate results will be 
// in the range of [-2^31, 2^31 - 1].
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, 
// such as eval().
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Basic Calculator II.
// Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Basic Calculator II.
//
func calculateII(s string) int {
    exp := make([]int, 0, 2)
    op := ADD
    for i := 0; i < len(s); {
        if s[i] == ' ' {
            i++
            continue
        }
        offset, token := parseToken(s[i:])
        if token == DIGITS {
            d, _ := strconv.Atoi(s[i:i+offset])
            switch op {
            case ADD:
                if len(exp) == 2 {
                    exp[0] = exp[0] + exp[1]
                    exp = exp[:1]
                }
                exp = append(exp, d)
            case MINUS:
                if len(exp) == 2 {
                    exp[0] = exp[0] + exp[1]
                    exp = exp[:1]
                }
                exp = append(exp, 0-d)
            case MULTIPLY:
                l := len(exp)
                exp[l-1] = exp[l-1] * d
            case DIVISE:
                l := len(exp)
                exp[l-1] = exp[l-1] / d
            }
        } else {
            op = token
        }
        i += offset
    }
    if len(exp) == 2 {
        return exp[0] + exp[1]
    }
    return exp[0]
}

const (
    _ = iota

    DIGITS
    ADD
    MINUS
    MULTIPLY
    DIVISE
)
func parseToken(s string) (int, int) {
    switch s[0] {
    case '+':
        return 1, ADD
    case '-':
        return 1, MINUS
    case '*':
        return 1, MULTIPLY
    case '/':
        return 1, DIVISE
    }
    i := 1
    for i < len(s) && s[i] >= '0' && s[i] <= '9' {
        i++
    }
    return i, DIGITS
}
