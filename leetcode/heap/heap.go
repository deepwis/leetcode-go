package leetcode

type Heap []int

func (h *Heap) Push(val int) {
    *h = append(*h, val)
    i := len(*h) - 1
    for i > 0 {
        p := (i-1) / 2
        if (*h)[p] > (*h)[i] {
            (*h)[p], (*h)[i] = (*h)[i], (*h)[p]
            i = p
        } else {
            break
        }
    }
}

func (h *Heap) Pop() int {
    n := len(*h) - 1
    if n < 0 {
        return -1
    }
    val := (*h)[0]
    (*h)[0], (*h)[n] = (*h)[n], (*h)[0]
    *h = (*h)[:n]
    p := 0
    for p < n-1 {
        left := p * 2 + 1
        if left >= n || left < 0 { // left < 0 after int overflow
            return val
        }
        i := left
        right := left + 1
        if right < n && (*h)[left] > (*h)[right] {
            i = right
        }
        if (*h)[i] < (*h)[p] {
            (*h)[i], (*h)[p] = (*h)[p], (*h)[i]
        }
    }
    return val
}

/*
func heap_push(nums *[]int, values *[]int, key int) {
    n := len(*nums)
    *nums = append(*nums, key)
    i := n
    for i > 0 {
        p := (i-1) / 2
        if (*values)[(*nums)[p]] > (*values)[(*nums)[i]] {
            (*nums)[p], (*nums)[i] = (*nums)[i], (*nums)[p]
            i = p
        } else {
            break
        }
    }
}

func heap_pop(nums *[]int, values *[]int) {
    n := len(*nums)
    (*nums)[0], (*nums)[n-1] = (*nums)[n-1], (*nums)[0]
    *nums = (*nums)[:n-1]
    i := 0
    for i < n-1 {
        p := i * 2 + 1
        if p+1 < n-1 && (*values)[(*nums)[p+1]] < (*values)[(*nums)[p]] {
            p = p+1
        }
        if (*values)[(*nums)[i]] > (*values)[(*nums)[p]] {
            (*nums)[p], (*nums)[i] = (*nums)[i], (*nums)[p]
            i = p
        } else {
            break
        }
    }
}
*/
