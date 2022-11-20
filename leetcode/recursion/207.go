package leetcode

// 207. Course Schedule
//
// There are a total of numCourses courses you have to take, labeled from 0 to numCourses - 1. 
// You are given an array prerequisites where prerequisites[i] = [ai, bi] indicates that you must take course bi 
// first if you want to take course ai.
// For example, the pair [0, 1], indicates that to take course 0 you have to first take course 1.
// Return true if you can finish all courses. Otherwise, return false.
//
// Runtime: 114 ms, faster than 5.15% of Go online submissions for Course Schedule.
// Memory Usage: 6.6 MB, less than 51.83% of Go online submissions for Course Schedule.
//
func canFinish(numCourses int, prerequisites [][]int) bool {
    if len(prerequisites) == 0 {
        return true
    }
    edges := make(map[int][]int, numCourses)
    for _, v := range prerequisites {
        if v[0] == v[1] {
            return false
        }
        if _, ok := edges[v[1]]; ok {
            edges[v[1]] = append(edges[v[1]], v[0])
        } else {
            edges[v[1]] = []int{v[0]}
        }
        if _, ok := edges[v[0]]; !ok {
            edges[v[0]] = []int{}
        }
        for _, src := range edges[v[0]] {
            if hasConnection(&edges, src, v[0]) {
                return false
            }
        }
    }
    return len(edges) <= numCourses
}

func hasConnection(edges *map[int][]int, src, tgt int) bool {
    if ovs, ok := (*edges)[src]; ok {
        for _, ov := range ovs {
            if ov == tgt {
                return true
            }
            if hasConnection(edges, ov, tgt) {
                return true
            }
        }
    }
    return false
}

