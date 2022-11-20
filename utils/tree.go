package utils

const NULL = -99999

func GenTree(nums []int) (root *TreeNode) {
    if len(nums) == 0 {
        return
    }
    nodes := []*TreeNode{&TreeNode{nums[0], nil, nil}}
    cur := 0
    i := 1
    for i < len(nums) {
        if nums[i] != NULL {
            left := &TreeNode{nums[i], nil, nil}
            nodes = append(nodes, left)
            nodes[cur].Left = left
        }
        if i+1 < len(nums) && nums[i+1] != NULL {
            right := &TreeNode{nums[i+1], nil, nil}
            nodes = append(nodes, right)
            nodes[cur].Right = right
        }
        i += 2
        cur += 1
    }
    root = nodes[0]
    return
}

func SearchNode(root *TreeNode, val int) *TreeNode {
    if root == nil {
        return nil
    }
    if root.Val == val {
        return root
    }
    left := SearchNode(root.Left, val)
    if left != nil {
        return left
    }
    return SearchNode(root.Right, val)
}

func TreeToArray(root *TreeNode, ret *[]int) {
    if root == nil {
        return
    }
    traversal := []*TreeNode{root}
    for i := 0; i < len(traversal); i++ {
        if traversal[i] == nil {
            *ret = append(*ret, NULL)
        } else {
            *ret = append(*ret, traversal[i].Val)
            traversal = append(traversal, traversal[i].Left, traversal[i].Right)
        }
    }
    last := len(*ret)
    for last > 0 {
        if (*ret)[last-1] == NULL {
            last--
        } else {
            break
        }
    }
    *ret = (*ret)[:last]
}

