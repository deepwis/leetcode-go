package leetcode

// 103. Binary Tree Zigzag Level Order Traversal
//
// Given the root of a binary tree, return the zigzag level order traversal of its nodes' values. 
// (i.e., from left to right, then right to left for the next level and alternate between).
//
// Runtime: 3 ms, faster than 44.27% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// Memory Usage: 2.6 MB, less than 64.89% of Go online submissions for Binary Tree Zigzag Level Order Traversal.
// 
func zigzagLevelOrder(root *TreeNode)[][]int {
    if root == nil {
        return [][]int{}
    }
    ret := [][]int{}
    q := []*TreeNode{root}
    level := 1
    ret = append(ret, []int{})
    for lp, rp, next := 0, 0, 1; lp < next; lp++ {
        if level & 1 == 1 {
            ret[level-1] = append(ret[level-1], q[lp].Val)
        } else {
            ret[level-1] = append(ret[level-1], q[rp].Val)
        }

        if q[lp].Left != nil {
            q = append(q, q[lp].Left)
        }
        if q[lp].Right != nil {
            q = append(q, q[lp].Right)
        }
        if lp == next -1 && next < len(q) {
            ret = append(ret, make([]int, 0, len(q) - next))
            level++
            next = len(q)
            rp = next
        }
        rp--
    }
    return ret
}
