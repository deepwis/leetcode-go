package leetcode

import (
	"testing"
    "reflect"
    "github.com/deepwis/leetcode-go/utils"
)

// 155. Min Stack
func TestMinStack(t *testing.T) {
    var c Codec

	t.Run("Test Constructor()", func(t *testing.T) {
        c = Constructor()
    })

	t.Run("Test serialize()", func(t *testing.T) {
        inputs := []*TreeNode{
            utils.GenTree([]int{3,9,20,utils.NULL,utils.NULL,15,7}),
            utils.GenTree([]int{1}),
        }
        wants := []string{"3,9,20,,,15,7", "1"}
        for i, in := range inputs {
            want := wants[i]
            got := c.Serialize(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

	t.Run("Test Deserialize()", func(t *testing.T) {
        inputs := []string{"3,9,20,,,15,7", "1"}
        wants := [][]int{
            []int{3,9,20,utils.NULL,utils.NULL,15,7},
            []int{1},
        }
        for i, in := range inputs {
            want := wants[i]
            got := []int{}
            utils.TreeToArray(c.Deserialize(in), &got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })
}
