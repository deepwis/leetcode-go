package utils

func GenSingly(nums []int) (head *ListNode) {
    if len(nums) == 0 {
        return
    }
    head = NewListNode(nums[0])
    cur := head
    for i := 1; i < len(nums); i++ {
        next := NewListNode(nums[i])
        cur.Next = next
        cur = next
    }
    return
}

func ToArray(head *ListNode) (ret []int) {
    if head == nil {
        return
    }
    cur := head
    for cur != nil {
        ret = append(ret, cur.Val)
        cur = cur.Next
    }
    return
}
