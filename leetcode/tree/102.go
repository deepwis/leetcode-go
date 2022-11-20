package leetcode

// 102. Binary Tree Level Order Traversal
//
// Given the root of a binary tree, return the level order traversal of its nodes' values. 
// (i.e., from left to right, level by level).
//
// Runtime: 6 ms, faster than 26.25% of Go online submissions for Binary Tree Level Order Traversal.
// Memory Usage: 3.1 MB, less than 23.49% of Go online submissions for Binary Tree Level Order Traversal.
//
func levelOrder(root *TreeNode) [][]int {
    var traversal [][]int
    if root == nil {
        return traversal
    }
    levelOrderRecursive(root, &traversal, 0)
    return traversal
}

func levelOrderRecursive(n *TreeNode, traversal *[][]int, depth int) {
    if n == nil {
        return
    }
    if len(*traversal) > depth {
        (*traversal)[depth] = append((*traversal)[depth], n.Val)
    } else {
        *traversal = append(*traversal, []int{n.Val})
    }

    levelOrderRecursive(n.Left, traversal, depth + 1)
    levelOrderRecursive(n.Right, traversal, depth + 1)
}

// Runtime: 3 ms, faster than 67.41% of Go online submissions for Binary Tree Level Order Traversal.
// Memory Usage: 2.8 MB, less than 53.22% of Go online submissions for Binary Tree Level Order Traversal.
//
func levelOrder2(root *TreeNode) [][]int {
    var traversal [][]int
    if root == nil {
        return traversal
    }
    q := []*TreeNode{root}
    low, high, depth, next := 0, 1, -1, 0
    for low < len(q) {
        node := q[low]
        high = len(q)

        if low == next { // first node of the depth
            traversal = append(traversal, []int{node.Val})
            depth++
        } else {
            traversal[depth] = append(traversal[depth], node.Val)
        }

        if node.Left != nil {
            q = append(q, node.Left)
        }
        if node.Right != nil {
            q = append(q, node.Right)
        }
        if low >= next && high < len(q) { // node in new depth has pushed into q
            next = high
        }
        low++
    }
    return traversal
}

