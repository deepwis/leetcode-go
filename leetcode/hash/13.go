package leetcode

// 13. Roman to Integer
//
// Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.
// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000
// For example, 2 is written as II in Roman numeral, just two ones added together. 
// 12 is written as XII, which is simply X + II. The number 27 is written as XXVII, which is XX + V + II.
// Roman numerals are usually written largest to smallest from left to right. 
// However, the numeral for four is not IIII. Instead, the number four is written as IV. 
// Because the one is before the five we subtract it making four. The same principle applies to the number nine, 
// which is written as IX. There are six instances where subtraction is used:
// * I can be placed before V (5) and X (10) to make 4 and 9. 
// * X can be placed before L (50) and C (100) to make 40 and 90. 
// * C can be placed before D (500) and M (1000) to make 400 and 900.
// Given a roman numeral, convert it to an integer.
//
// Runtime: 32 ms, faster than 6.18% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 28.45% of Go online submissions for Roman to Integer.
//
func romanToInt(s string) int {
    simbol2Value := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    substractPairs := map[byte]byte{
        'V': 'I',
        'X': 'I',
        'L': 'X',
        'C': 'X',
        'D': 'C',
        'M': 'C',
    }

    if len(s) == 0 {
        return 0
    }
    if len(s) == 1 {
        return simbol2Value[s[0]]
    }

    var ret int
    for i :=0; i < len(s); {
        cur := s[i]
        add := simbol2Value[cur]
        if i + 1 == len(s) {
            ret += add
            return ret
        }
        next := s[i+1]
        left, ok := substractPairs[next]
        if ok && left == cur {
            add = simbol2Value[next] - simbol2Value[cur]
            i += 2
        } else {
            i++
        }
        ret += add
    }
    return ret
}

// Runtime: 26 ms, faster than 12.54% of Go online submissions for Roman to Integer.
// Memory Usage: 2.9 MB, less than 84.08% of Go online submissions for Roman to Integer.
//
func romanToInt2(s string) int {
    simbol2Value := map[byte]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    substractPairs := map[byte]byte{
        'V': 'I',
        'X': 'I',
        'L': 'X',
        'C': 'X',
        'D': 'C',
        'M': 'C',
    }

    if len(s) == 0 {
        return 0
    }
    if len(s) == 1 {
        return simbol2Value[s[0]]
    }

    var ret int
    for i := 0; i < len(s); i++ {
        ret += simbol2Value[s[i]]
        if prev, ok := substractPairs[s[i]]; i > 0 && ok && s[i-1] == prev {
            ret -= (2 * simbol2Value[prev])
        }
    }
    return ret
}
