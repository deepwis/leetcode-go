package leetcode

// 763. Partition Labels
//
// You are given a string s. We want to partition the string into as many parts as possible 
// so that each letter appears in at most one part.
// Note that the partition is done so that after concatenating all the parts in order, 
// the resultant string should be s.
// Return a list of integers representing the size of these parts.
//
// Runtime: 5 ms, faster than 39.57% of Go online submissions for Partition Labels.
// Memory Usage: 2.2 MB, less than 66.19% of Go online submissions for Partition Labels.
//
func partitionLabels(s string) []int {
    charPos := make([][]int, 26)
    for i := 0; i < len(s); i++ {
        idx := int(s[i] - 'a')
        if charPos[idx] == nil {
            charPos[idx] = make([]int, 1, 2)
            charPos[idx][0] = i
        } else if len(charPos[idx]) == 1 {
            charPos[idx] = append(charPos[idx], i)
        } else {
            charPos[idx][1] = i
        }
    }

    if len(s) == 1{
        return []int{len(s)}
    }
    partitions := []int{}
    // low := 0
    high := 0
    i := 0
    for i < len(s) && high <= len(s) {
        idx := int(s[i] - 'a')
        if len(charPos[idx]) == 2 {
            pos := charPos[idx]
            if pos[0] == i {
                if pos[0] > high {
                    singles := make([]int, pos[0]-high)
                    for j := 0; j < len(singles); j++ {
                        singles[j] = 1
                    }
                    partitions = append(partitions, singles...)
                    high = pos[0]
                }
                if pos[0] == high {
                    partitions = append(partitions, pos[1]+1-pos[0])
                    //low = pos[0]
                    high = pos[1]+1
                } else if pos[1]+1 > high{
                    partitions[len(partitions)-1] = partitions[len(partitions)-1] + (pos[1] + 1 - high)
                    high = pos[1]+1
                }
            }
        }
        i++
    }
    if i == len(s) && high < len(s) {
        singles := make([]int, i-high)
        for j := 0; j < len(singles); j++ {
            singles[j] = 1
        }
        partitions = append(partitions, singles...)
    }
    return partitions
}

