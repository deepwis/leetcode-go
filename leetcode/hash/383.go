package leetcode

// 383. Ransom Note
//
// Given two strings ransomNote and magazine, return true if ransomNote can be constructed 
// by using the letters from magazine and false otherwise.
// Each letter in magazine can only be used once in ransomNote.
//
// Runtime: 37 ms, faster than 20.93% of Go online submissions for Ransom Note.
// Memory Usage: 3.9 MB, less than 82.91% of Go online submissions for Ransom Note.
//
func canConstruct(ransomNote, magazine string) bool {
   charCount := make(map[rune]int)
    for _, c := range magazine {
       charCount[c] = charCount[c] + 1
    }

    for _, c := range ransomNote {
        left := charCount[c] - 1
        if left < 0 {
            return false
        }
       charCount[c] = left
    }
    return true
}
