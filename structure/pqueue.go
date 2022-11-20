package structure

// Priority Queue

type PQueue []int

func (q PQueue) Len() int { return len(q) }
func (q PQueue) Less(i, j int) bool { return q[i] > q[j] }
func (q PQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }
func (q *PQueue) Push(x interface{}) {
    *q = append(*q, x.(int))
}

func (q *PQueue) Pop() interface{} {
    old := *q
    n := len(old)
    x := old[n-1]
    *q = old[: n-1]
    return x
}

