package leetcode

// 17. Letter Combinations of a Phone Number
//
// Given a string containing digits from 2-9 inclusive, return all possible letter combinations 
// that the number could represent. Return the answer in any order.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Letter Combinations of a Phone Number.
// Memory Usage: 2.1 MB, less than 17.09% of Go online submissions for Letter Combinations of a Phone Number.
//
func letterCombinations(digits string) []string {
    var ret []string
    if len(digits) == 0 {
        return ret
    }
    letterMap := make(map[rune][]rune)
    letterMap['2'] = []rune{'a','b','c'}
    letterMap['3'] = []rune{'d','e','f'}
    letterMap['4'] = []rune{'g','h','i'}
    letterMap['5'] = []rune{'j','k','l'}
    letterMap['6'] = []rune{'m','n','o'}
    letterMap['7'] = []rune{'p','q','r','s'}
    letterMap['8'] = []rune{'t','u','v'}
    letterMap['9'] = []rune{'w','x','y','z'}
    var q [][]rune
    for _, c := range(letterMap[rune(digits[0])]) {
        q = append(q, []rune{c})
    }
    digits = string(digits[1:])
    qSize := len(q)
    for i, c := range digits {
        if i == 3 {
            break
        }
        qSize = len(q)
        for j := 0; j < qSize; j++ {
            for k := 0; k < len(letterMap[c]); k++ {
                if i == len(digits) - 1 {
                    ret = append(ret, string(append(q[j], letterMap[c][k])))

                } else {
                    q = append(q, append(q[j], letterMap[c][k]))
                }
            }
        }
        q = q[qSize:]
    }
    for _, v := range q {
        ret = append(ret, string(v))
    }
    return ret
}
