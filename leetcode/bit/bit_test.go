package leetcode

import (
	"testing"
)

func TestSolution(t *testing.T) {
    // 8. String to Integer (atoi)
	t.Run("Test myAtoi()", func(t *testing.T) {
        inputs := []string{"42", "     -42", "  +444abc"}
		wants := []int{42, -42, 444}

        for i, in := range inputs {
            want := wants[i]
            got := myAtoi(in)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 43. Multiply Strings
	t.Run("Test multiply()", func(t *testing.T) {
        inputs := [][]string{
            []string{"2", "3"},
            []string{"123", "456"},
            []string{"9623", "0"},
        }
		wants := []string{"6", "56088", "0"}

        for i, in := range inputs {
            want := wants[i]
            got := multiply(in[0], in[1])
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 69. Sqrt(x) 
	t.Run("Test mySqrt()", func(t *testing.T) {
        inputs := []int{1, 4, 8, 10000}
		wants := []int{1, 2, 2, 100}

        for i, in := range inputs {
            want := wants[i]
            got := mySqrt(in)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 1342. Number of Steps to Reduce a Number to Zero
	t.Run("Test numberOfSteps()", func(t *testing.T) {
        inputs := []int{14, 8, 9, 123, 83962}
		wants := []int{6, 4, 5, 12, 27}

        for i, in := range inputs {
            want := wants[i]
            got := numberOfSteps(in)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 136. Single Number
	t.Run("Test singleNumber()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,2,1},
            []int{4,1,2,1,2},
            []int{1},
        }
		wants := []int{1,4,1}

        for i, in := range inputs {
            want := wants[i]
            got := singleNumber(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
            }
        }
    })

}
