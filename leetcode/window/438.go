package leetcode

// 438. Find All Anagrams in a String
//
// Given two strings s and p, return an array of all the start indices of p's anagrams in s. 
// You may return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
// typically using all the original letters exactly once.
//
// Runtime: 7 ms, faster than 90.63% of Go online submissions for Find All Anagrams in a String.
// Memory Usage: 5.2 MB, less than 56.25% of Go online submissions for Find All Anagrams in a String.
//
func findAnagrams(s, p string) []int {
    w := len(p)
    if len(s) < w {
        return []int{}
    }
    ret := []int{}
    needCnt := make(map[int]int, 26)
    windowCnt := make(map[int]int, 26)
    for i := 0; i < len(p); i++ {
        idx := int(p[i] - 'a')
        needCnt[idx] += 1
    }
    needs := w
    extra := 0
    for i := 0; i < w; i++ {
        idx := int(s[i] - 'a')
        if _, ok := needCnt[idx]; ok {
            windowCnt[idx] += 1
            if windowCnt[idx] <= needCnt[idx] {
                needs -= 1
            } else {
                extra += 1
            }
        } else {
            extra += 1
        }
    }
    for i := 0; i+w <= len(s); i++ {
        if needs == 0 && extra == 0 {
            ret = append(ret, i)
        }
        if i+w == len(s) {
            break
        }
        left := int(s[i] - 'a')
        if _, ok := needCnt[left]; ok {
            windowCnt[left] -= 1
            if windowCnt[left] < needCnt[left] {
                needs += 1
            } else {
                extra -= 1
            }
        } else {
            extra -= 1
        }
        right := int(s[i+w] - 'a')
        if _, ok := needCnt[right]; ok {
            windowCnt[right] += 1
            if windowCnt[right] <= needCnt[right] {
                needs -= 1
            } else {
                extra += 1
            }
        } else {
            extra += 1
        }
    }

    return ret
}
