package leetcode

// 199. Binary Tree Right Side View
//
// Given the root of a binary tree, imagine yourself standing on the right side of it, 
// return the values of the nodes you can see ordered from top to bottom.
//
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Binary Tree Right Side View.
// Memory Usage: 2.2 MB, less than 99.45% of Go online submissions for Binary Tree Right Side View.
//
func rightSideView(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }
    q := []*TreeNode{root}
    view, next := []int{}, 0
    for i := 0; i < len(q); i++ {
        p := len(q)
        if q[i].Right != nil {
            q = append(q, q[i].Right)
        }
        if q[i].Left != nil {
            q = append(q, q[i].Left)
        }
        if i == next {
            view = append(view, q[i].Val)
        }
        if i >= next && len(q) > p{
            next = p
        }
    }
    return view
}

