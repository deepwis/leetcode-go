package structure

type ListNode struct {
    Val int
    Next *ListNode
    Prev *ListNode
}

func NewListNode(val int) *ListNode {
    return &ListNode{val, nil, nil}
}

