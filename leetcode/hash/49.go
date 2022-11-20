package leetcode

// 49. Group Anagrams
//
// Given an array of strings strs, group the anagrams together. You can return the answer in any order.
// An Anagram is a word or phrase formed by rearranging the letters of a different word or phrase, 
// typically using all the original letters exactly once.
//
// Runtime: 17 ms, faster than 98.89% of Go online submissions for Group Anagrams.
// Memory Usage: 7.4 MB, less than 94.57% of Go online submissions for Group Anagrams.
//
var primes = [26]int{2,3,5,7,11,13,17,19,23,29,31,37,41,43,47,53,59,61,67,71,73,79,83,89,97,101}

func groupAnagrams(strs []string) [][]string {
    if len(strs) == 0 {
        return [][]string{}
    }
    if len(strs) == 1 {
        return [][]string{strs}
    }
    groups := make(map[float64][]string)
    for _, s := range strs {
        key := 1.0
        for _, c := range s {
            key *= float64(primes[c - 'a'])
        }
        if _, ok := groups[key]; ok {
            groups[key] = append(groups[key], s)
        } else {
            groups[key] = []string{s}
        }
    }
    ret := make([][]string, len(groups))
    i := 0
    for _, v := range groups {
        ret[i] = v
        i++
    }
    return ret
}
