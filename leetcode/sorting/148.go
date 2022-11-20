package leetcode

import "sort"

// 148. Sort List
//
// Given the head of a linked list, return the list after sorting it in ascending order.
//
// Runtime: 78 ms, faster than 77.78% of Go online submissions for Sort List.
// Memory Usage: 9.8 MB, less than 12.65% of Go online submissions for Sort List.
//
func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    nodeList := []*ListNode{}
    cur := head
    for cur != nil {
        nodeList = append(nodeList, cur)
        cur = cur.Next
    }
    sortable := ListArray(nodeList)
    sort.Sort(sortable)
    return sortable[0]
}

type ListArray []*ListNode
func (this ListArray) Len() int { return len(this) }
func (this ListArray) Less(i, j int) bool { return this[i].Val < this[j].Val }
func (this ListArray) Swap(i, j int) {
    this[i], this[j] = this[j], this[i]
    if i == len(this) - 1 {
        this[i].Next = nil
    } else {
        this[i].Next = this[i+1]
    }
    if j == len(this) - 1 {
        this[j].Next = nil
    } else {
        this[j].Next = this[j+1]
    }
    if i != 0 {
        this[i-1].Next = this[i]
    }
    if j != 0 {
        this[j-1].Next = this[j]
    }
}
