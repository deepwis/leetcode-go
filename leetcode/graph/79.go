package leetcode

// 79. Word Search
//
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells 
// are horizontally or vertically neighboring. The same letter cell may not be used more than once.
//
// Runtime: 293 ms, faster than 30.51% of Go online submissions for Word Search.
// Memory Usage: 2.2 MB, less than 26.86% of Go online submissions for Word Search.
//
func exist(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }
    m, n := len(board), len(board[0])
    if m * n < len(word) {
        return false
    }
    chars := make(map[byte]bool, len(word))
    visited := make(map[int]bool, m*n)
    for i := 0; i < len(word); i++ {
        chars[word[i]] = true
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !chars[board[i][j]] {
                visited[i*n+j] = true
            }
            if board[i][j] == word[0] {
                visited := make(map[int]bool)
                if bfs(&board, word, i, j, &visited) {
                    return true
                }
            }
        }
    }
    return false
}

func bfs(board *[][]byte, word string, i, j int, visited *map[int]bool) bool {
    if (*board)[i][j] != word[0] {
        return false
    }
    if len(word) == 1 {
        return true
    }
    m, n := len(*board), len((*board)[0])
    (*visited)[i*n+j] = true
    if i > 0 && !(*visited)[(i-1)*n+j] && bfs(board, word[1:], i-1, j, visited) {
        return true
    }
    if i < m-1 && !(*visited)[(i+1)*n+j] && bfs(board, word[1:], i+1, j, visited) {
        return true
    }
    if j > 0 && !(*visited)[i*n+j-1] && bfs(board, word[1:], i, j-1, visited) {
        return true
    }
    if j < n-1 && !(*visited)[i*n+j+1] && bfs(board, word[1:], i, j+1, visited) {
        return true
    }
    (*visited)[i*n+j] = false
    return false
}

// Runtime: 930 ms, faster than 16.12% of Go online submissions for Word Search.
// Memory Usage: 2.1 MB, less than 31.89% of Go online submissions for Word Search.
//
func exist2(board [][]byte, word string) bool {
    if len(board) == 0 {
        return false
    }
    m, n := len(board), len(board[0])
    if m * n < len(word) {
        return false
    }
    chars := make(map[byte]bool, len(word))
    visited := make(map[int]bool, m*n)
    for i := 0; i < len(word); i++ {
        chars[word[i]] = true
    }
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if !chars[board[i][j]] {
                visited[i*n+j] = true
            }
            if board[i][j] == word[0] {
                if bfs2(&board, word, i, j, &visited, &chars) {
                    return true
                }
            }
        }
    }
    return false
}

func bfs2(board *[][]byte, word string, i, j int, visited *map[int]bool, chars *map[byte]bool) bool {
    if (*board)[i][j] != word[0] {
        return false
    }
    m, n := len(*board), len((*board)[0])
    (*visited)[i*n+j] = true
    if len(word) == 1 {
        return true
    }
    adjacents := [][]int{[]int{i-1,j}, []int{i+1,j}, []int{i,j-1}, []int{i,j+1}}
    for _, pos := range adjacents {
        if pos[0] < 0 || pos[1] < 0 || pos[0] > m-1 || pos[1] > n-1 || (*visited)[pos[0]*n+pos[1]] {
            continue
        }
        if !(*chars)[(*board)[pos[0]][pos[1]]] {
            (*visited)[pos[0]*n+pos[1]] = true
        } else if bfs2(board, word[1:], pos[0], pos[1], visited, chars) {
            return true
        }
    }
    (*visited)[i*n+j] = false
    return false
}
