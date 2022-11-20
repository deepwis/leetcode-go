package leetcode

import (
	"testing"
)

// 155. Min Stack
func TestMinStack(t *testing.T) {
    var s MinStack
	t.Run("Test Constructor()", func(t *testing.T) {
        s = Constructor2()
        want := -1
        got := s.Top()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Push()", func(t *testing.T) {
        s.Push(-2)
        s.Push(0)
        s.Push(-3)
        want := -3
        got := s.Top()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test GetMin()", func(t *testing.T) {
        want := -3
        got := s.GetMin()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Pop()", func(t *testing.T) {
        s.Pop()
        want := 0
        got := s.Top()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test Top()", func(t *testing.T) {
        want := 0
        got := s.Top()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })

	t.Run("Test GetMin()", func(t *testing.T) {
        want := -2
        got := s.GetMin()
        if got != want {
            t.Errorf("got: %v, want: %v", got, want)
        }
    })
}
