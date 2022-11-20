package leetcode

// 207. Course Schedule
//
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi 
// first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
//
// Runtime: 9 ms, faster than 96.70% of Go online submissions for Course Schedule.
// Memory Usage: 6 MB, less than 92.67% of Go online submissions for Course Schedule.
//
func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0 {
        return true
    }
    adj := make([][]int, numCourses)
    inDegrees := make([]int, numCourses)
    for _, dep := range prerequisites {
        adj[dep[1]] = append(adj[dep[1]], dep[0])
        inDegrees[dep[0]]++
    }

    q := []int{}
    for v, d := range inDegrees {
        if d == 0 {
            q = append(q, v)
        }
    }
    if len(q) == 0 {
        return false
    }
    for len(q) > 0 {
        v := q[0]
        q = q[1:]
        for _, out := range adj[v] {
            inDegrees[out]--
            if inDegrees[out] == 0 {
                q = append(q, out)
            }
        }
    }
    for _, d := range inDegrees {
        if d > 0 {
            return false
        }
    }
    return true
}


