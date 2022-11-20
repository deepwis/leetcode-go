package leetcode

// 22. Generate Parentheses
//
// Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.
//
// Runtime: 2 ms, faster than 68.65% of Go online submissions for Generate Parentheses.
// Memory Usage: 2.7 MB, less than 72.79% of Go online submissions for Generate Parentheses.
//
func generateParentheses(n int) []string {
    all_combinations := make([][]string, n+1)
    all_combinations[0] = []string{""}
    all_combinations[1] = []string{"()"}
    for i := 2; i < n+1; i++ {
        ilist := []string{}
        for p := 0; p < i; p++ {
            q := i - p - 1
            for _, u := range all_combinations[p] {
                for _, v := range all_combinations[q] {
                    ilist = append(ilist, "(" + u + ")" + v)
                }
            }
        }
        all_combinations[i] = ilist
    }
    return all_combinations[n]
}


