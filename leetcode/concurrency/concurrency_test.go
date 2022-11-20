package leetcode

import (
	"testing"
)

func TestSolution(t *testing.T) {
    // 165. Compare Version Numbers
	t.Run("Test compareVersion()", func(t *testing.T) {
        inputs := [][]string{
            []string{"1.01", "1.001"},
            []string{"1.0", "1.0.0"},
            []string{"0.1", "1.1"},
            []string{"1.1", "1.0.5.8"},
            []string{"1.0.1", "1"},
        }
		wants := []int{0, 0, -1, 1, 1}
        for i, in := range inputs {
            want := wants[i]
            got := compareVersion(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
