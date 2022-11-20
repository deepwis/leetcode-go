package leetcode

// 139. Word Break
//
// Given a string s and a dictionary of strings wordDict, return true if s can be segmented into 
// a space-separated sequence of one or more dictionary words.
// Note that the same word in the dictionary may be reused multiple times in the segmentation.
//
// Runtime: 4 ms, faster than 39.92% of Go online submissions for Word Break.
// Memory Usage: 2.2 MB, less than 64.97% of Go online submissions for Word Break.
//
func wordBreak(s string, wordDict []string) bool {
    wordSet := make(map[string]bool, len(wordDict))
    minLen, maxLen := len(s), 0
    for _, word := range wordDict {
        wordSet[word] = true
        if len(word) > maxLen {
            maxLen = len(word)
        }
        if len(word) < minLen {
            minLen = len(word)
        }
    }
    if len(s) < minLen {
        return false
    }
    validPos := []int{0}
    for low, high := 0, minLen; high - low <= maxLen && high <= len(s); high++ {
        for i := len(validPos) - 1; i >= 0; i-- {
            if high - validPos[i] > maxLen {
                break
            }
            if high - validPos[i] < minLen {
                continue
            }
            if _, ok := wordSet[s[validPos[i] : high]]; ok {
                if high == len(s) {
                    return true
                }
                validPos = append(validPos, high)
                low = high
                break
            }
        }
    }
    return false
}
