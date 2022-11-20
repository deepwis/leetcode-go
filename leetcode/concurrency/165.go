package leetcode

// 165. Compare Version Numbers
//
// Given two version numbers, version1 and version2, compare them.
// Version numbers consist of one or more revisions joined by a dot '.'. 
// Each revision consists of digits and may contain leading zeros. Every revision contains at least one character. 
// Revisions are 0-indexed from left to right, with the leftmost revision being revision 0, 
// the next revision being revision 1, and so on. For example 2.5.33 and 0.1 are valid version numbers.
// To compare version numbers, compare their revisions in left-to-right order. 
// Revisions are compared using their integer value ignoring any leading zeros. This means 
// that revisions 1 and 001 are considered equal. If a version number does not specify a revision at an index, 
// then treat the revision as 0. For example, version 1.0 is less than version 1.1 because their revision 0s 
// are the same, but their revision 1s are 0 and 1 respectively, and 0 < 1.
// Return the following:
// * If version1 < version2, return -1.
// * If version1 > version2, return 1.
// Otherwise, return 0.
//
// Runtime: 2 ms, faster than 31.58% of Go online submissions for Compare Version Numbers.
// Memory Usage: 2 MB, less than 80.70% of Go online submissions for Compare Version Numbers.
//
func compareVersion(version1, version2 string) int {
    ch1, ch2 := make(chan int), make(chan int)
    go parseRevision(version1, ch1)
    go parseRevision(version2, ch2)
    for {
        v1, ok1 := <-ch1
        v2, ok2 := <-ch2
        if !ok1 && !ok2 {
            break
        }
        if v1 < v2 {
            return -1
        }
        if v1 > v2 {
            return 1
        }
    }
    return 0
}

func parseRevision(s string, ch chan int) {
    offset := 0
    for offset < len(s) {
        v := 0
        i := offset
        for i < len(s) && s[i] != '.' {
            v = v * 10 + int(s[i] - '0')
            i++
        }
        offset = i+1
        ch <- v
    }
    close(ch)
}
