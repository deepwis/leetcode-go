package leetcode

// 543. Diameter of Binary Tree
//
// Given the root of a binary tree, return the length of the diameter of the tree.
// The diameter of a binary tree is the length of the longest path between any two nodes in a tree. 
// This path may or may not pass through the root.
// The length of a path between two nodes is represented by the number of edges between them.
//
// Runtime: 10 ms, faster than 42.77% of Go online submissions for Diameter of Binary Tree.
// Memory Usage: 4.4 MB, less than 76.63% of Go online submissions for Diameter of Binary Tree.
//
func diameterOfBinaryTree(root *TreeNode) int {
    d := 0 // *d is number of edge, not number of node. So *d = num(node)-1
    recDiameter(root, &d)
    return d
}

func recDiameter(node *TreeNode, d *int) int {
    if node == nil {
        return 0
    }
    ld := recDiameter(node.Left, d)
    rd := recDiameter(node.Right, d)
    if (ld+rd) > *d {
        *d = ld+rd
    }
    if rd > ld {
        return rd+1
    }
    return ld+1
}

