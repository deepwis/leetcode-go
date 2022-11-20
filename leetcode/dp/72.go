package leetcode

// 72. Edit Distance
//
// Given two strings word1 and word2, return the minimum number of operations required to convert word1 to word2.
// You have the following three operations permitted on a word:
// * Insert a character
// * Delete a character
// * Replace a character
//
// Runtime: 5 ms, faster than 60.69% of Go online submissions for Edit Distance.
// Memory Usage: 5.6 MB, less than 42.77% of Go online submissions for Edit Distance.
//
func minDistance(word1, word2 string) int {
    if len(word1) == 0 {
        return len(word2)
    }
    if len(word2) == 0 {
        return len(word1)
    }
    dist := make([][]int, len(word1)+1)
    dist[0] = make([]int, len(word2)+1)
    for i := 0; i < len(word2)+1; i++ {
        dist[0][i] = i // dist[0][i] represents distance convert "" to word2[:i]
    }
    for i := 1; i < len(word1)+1; i++ {
        // the i-th character in word1, word1[i-1]
        dist[i] = make([]int, len(word2)+1)
        dist[i][0] = dist[i-1][0] + 1 // dist[i][0] represents distance convert word1[:i] to ""
        for j := 1; j < len(word2)+1; j++ {
            // the j-th character in word2, word1[j-1]
            d1 := dist[i][j-1] + 1 // case1: convert word1[:i] to word2[:j-1], and then add word2[j]
            d2 := dist[i-1][j] + 1 // case2: convert word1[:i-1] to word2[:j], and then delete word1[i]
            d3 := dist[i-1][j-1] // case3: convert word1[:i-1] to word2[:j-1], and then compare word1[i] and word2[j]
            if word1[i-1] != word2[j-1] {
                d3 += 1 // need replace word1[i] to word2[j]
            }
            dist[i][j] = d1
            if d2 < dist[i][j] {
                dist[i][j] = d2
            }
            if d3 < dist[i][j] {
                dist[i][j] = d3
            }
        }
    }
    return dist[len(word1)][len(word2)]
}
