package leetcode

// 224. Basic Calculator
//
// Given a string s representing a valid expression, implement a basic calculator to evaluate it, 
// and return the result of the evaluation.
// Note: You are not allowed to use any built-in function which evaluates strings as mathematical expressions, 
// such as eval().
//
// Runtime: 4 ms, faster than 81.40% of Go online submissions for Basic Calculator.
// Memory Usage: 2.9 MB, less than 100.00% of Go online submissions for Basic Calculator.
//
func calculate(s string) int {
    stack := []int{}
    total, sign := 0, 1
    for i := 0; i < len(s); i++ {
        if s[i] >= '0' && s[i] <= '9' {
            num := 0
            for i < len(s) && s[i] >= '0' && s[i] <= '9' {
                num = num * 10 + int(s[i] - '0')
                i++
            }
            total += num * sign
            sign = 1
            i--
            continue
        }
        switch s[i] {
        case '+':
            sign = 1
        case '-':
            sign = -1
        case '(':
            stack = append(stack, total)
            stack = append(stack, sign)
            total, sign = 0, 1
        case ')':
            num := total
            sign, total = stack[len(stack)-1], stack[len(stack)-2]
            total = total + sign * num
            stack = stack[:len(stack)-2]
        }
    }
    return total
}
