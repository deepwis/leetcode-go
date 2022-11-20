package leetcode

// 155. Min Stack
//
// Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.
// Implement the MinStack class:
// * MinStack() initializes the stack object.
// * void push(int val) pushes the element val onto the stack.
// * void pop() removes the element on the top of the stack.
// * int top() gets the top element of the stack.
// * int getMin() retrieves the minimum element in the stack.
// You must implement a solution with O(1) time complexity for each function.
//
// Runtime: 33 ms, faster than 51.83% of Go online submissions for Min Stack.
// Memory Usage: 9 MB, less than 29.25% of Go online submissions for Min Stack.
//
type MinStack struct {
    data []int
    minPos int
}

//func Constructor() MinStack {
func Constructor2() MinStack {
    return MinStack{[]int{}, -1}
}

func (this *MinStack) Push(val int) {
    this.data = append(this.data, val)
    if this.minPos == -1 || this.data[this.minPos] > val {
        this.minPos = len(this.data) - 1
    }
}

func (this *MinStack) Pop() {
    n := len(this.data)
    this.data = this.data[:n-1]
    if n == 1 {
        this.minPos = -1
    } else if this.minPos == n - 1 {
        this.minPos = 0
        for i := 1; i < n-1; i++ {
            if this.data[i] < this.data[this.minPos] {
                this.minPos = i
            }
        }
    }
}

func (this *MinStack) Top() int {
    n := len(this.data)
    if n == 0 {
        return -1
    }
    return this.data[n-1]
}

func (this *MinStack) GetMin() int {
    if this.minPos == -1 {
        return -1
    }
    return this.data[this.minPos]
}
