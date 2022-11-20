package leetcode

// 131. Palindrome Partitioning
//
// Given a string s, partition s such that every substring of the partition is a palindrome. 
// Return all possible palindrome partitioning of s.
//
// Runtime: 540 ms, faster than 32.49% of Go online submissions for Palindrome Partitioning.
// Memory Usage: 22.4 MB, less than 36.29% of Go online submissions for Palindrome Partitioning.
//
func partition(s string) [][]string {
    partitions := make([][][]string, len(s)+1)
    partitions[0] = [][]string{}
    partitions[1] = [][]string{[]string{string(s[0])}}
    for i := 1; i < len(s); i++ {
        partitions[i+1] = make([][]string, len(partitions[i]))
        copy(partitions[i+1], partitions[i])
        cs := string(s[i])
        plen := len(partitions[i])
        for j := 0; j < plen; j++ {
            partitions[i+1][j] = append(partitions[i+1][j], cs)
        }
        for left := 0; left < i; left++ {
            if isPalindromeString(s[left:i+1]) {
                if left == 0 {
                    partitions[i+1] = append(partitions[i+1], []string{s[left:i+1]})
                } else {
                    for k := 0; k < len(partitions[left]); k++ {
                        llen := len(partitions[left][k])
                        newPart := make([]string, llen+1)
                        copy(newPart[:llen], partitions[left][k])
                        newPart[llen] = s[left:i+1]
                        partitions[i+1] = append(partitions[i+1], newPart)
                    }
                }
            }
        }
    }
    return partitions[len(s)]
}

func isPalindromeString(s string) bool {
    for i := 0; i < len(s) / 2; i++ {
        if s[i] != s[len(s)-1-i] {
            return false
        }
    }
    return true
}

