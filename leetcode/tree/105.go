package leetcode

// 105. Construct Binary Tree from Preorder and Inorder Traversal
//
// Given two integer arrays preorder and inorder where preorder is the preorder traversal of a binary tree and inorder is the inorder traversal of the same tree, construct and return the binary tree.
//
// Runtime: 7 ms, faster than 69.10% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
// Memory Usage: 4.2 MB, less than 74.62% of Go online submissions for Construct Binary Tree from Preorder and Inorder Traversal.
//
func buildTree(preorder, inorder []int) *TreeNode {
    if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
        return &TreeNode{}
    }
    tree, ok := buildNode(preorder, inorder)
    if !ok {
        return &TreeNode{}
    }
    return tree
}
func buildNode(preorder, inorder []int) (*TreeNode, bool) {
    root := &TreeNode{
        Val: preorder[0],
        Left: nil,
        Right: nil,
    }

    i := 0
    for ; i < len(inorder); i++ {
        if inorder[i] == preorder[0] {
            if i > 0 {
                lnode, ok := buildNode(preorder[1:i+1], inorder[0:i])
                if ok {
                    root.Left = lnode
                } else {
                    continue
                }
            }
            if i < len(inorder) - 1 {
                rnode, ok := buildNode(preorder[i+1:], inorder[i+1:])
                if ok {
                    root.Right = rnode
                } else {
                    root.Left = nil
                    continue
                }
            }
            break
        }
    }
    if i >= len(inorder) {
        return root, false
    }
    return root, true
}
