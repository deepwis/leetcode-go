package leetcode

/*
// 1143. Longest Common Subsequence
//
// Given two strings text1 and text2, return the length of their longest common subsequence. 
// If there is no common subsequence, return 0.
// A subsequence of a string is a new string generated from the original string with some characters 
// (can be none) deleted without changing the relative order of the remaining characters.
// For example, "ace" is a subsequence of "abcde".
// A common subsequence of two strings is a subsequence that is common to both strings.
//
// Runtime: 7 ms, faster than 70.03% of Go online submissions for Longest Common Subsequence.
// Memory Usage: 11 MB, less than 36.50% of Go online submissions for Longest Common Subsequence.
//
func longestCommonSubsequence(text1, text2 string) int {
    base, commonChars := make(map[byte]bool), make(map[byte][]int)
    for i := 0; i < len(text2); i++ {
        base[text2[i]] = true
    }
    for i := 0; i < len(text1); i++ {
        if _, ok := commonChars[text1[i]]; ok {
            commonChars[text1[i]] = append(commonChars[text1[i]], i)
            continue
        }
        if _, ok := base[text1[i]]; ok {
            commonChars[text1[i]] = []int{i}
        }
    }

    fmt.Printf("common sequence %s vs %s\n", text1, text2)
    sequence := []int{}
    for i := 0; i < len(text2); i++ {
        c := text2[i]
        if _, ok := commonChars[c]; !ok {
            continue
        }
        pos := commonChars[c]
        if len(sequence) == 0 {
            sequence = append(sequence, pos[0])
            continue
        }
        if pos[len(pos)-1] > sequence[len(sequence)-1] {
            for j := 0; j < len(pos); j++ {
                if pos[j] > sequence[len(sequence)-1] {
                    sequence = append(sequence, pos[j])
                    break
                }
            }
        } else {
            for j := 0; j < len(pos); j++ {
                for k := 0; k <= len(sequence)-1; k++ {
                    if sequence[k] == pos[j] {
                        break
                    }
                    if sequence[k] > pos[j] {
                        sequence[k] = pos[j]
                        break
                    }
                }
            }
        }
    }
    return len(sequence)
}
*/
