package leetcode

import (
	"testing"
)

func TestSolution(t *testing.T) {
    // 20. Valid Parentheses
	t.Run("Test isValid()", func(t *testing.T) {
        inputs := []string{"()", "()[]{}", "(]", "][", "()[[[", "[([{}"}
		wants := []bool{true, true, false, false, false, false}
        for i, in := range inputs {
            want := wants[i]
            got := isValid(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 32. Longest Valid Parentheses
	t.Run("Test longestValidParentheses()", func(t *testing.T) {
        inputs := []string{
            "(",
            "()",
            ")(",
            "(()",
            ")()())",
            "",
            "()(()",
        }
		wants := []int{0,2,0,2,4,0,2}

        for i, in := range inputs {
            want := wants[i]
            got := longestValidParentheses(in)
            if got != want {
                t.Errorf("input: %v (%T), got: %v (%T), want: %v (%T)", in, in, got, got, want, want)
            }
        }
	})

    // 224. Basic Calculator
	t.Run("Test caculate()", func(t *testing.T) {
        inputs := []string{"1+1", "2-1 + 2 ", "(1+(4+5+2)-3)+(6+8)"}
		wants := []int{2, 3, 23}

        for i, in := range inputs {
            want := wants[i]
            got := calculate(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 227. Basic Calculator II
	t.Run("Test caculateII()", func(t *testing.T) {
        inputs := []string{"3+2*2", "3/2 ", "3+5 / 2 "}
		wants := []int{7, 1, 5}

        for i, in := range inputs {
            want := wants[i]
            got := calculateII(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 394. Decode String
	t.Run("Test decodeString()", func(t *testing.T) {
        inputs := []string{"3[a]2[bc]", "3[a2[c]]", "2[abc]3[cd]ef"}
		wants := []string{"aaabcbc", "accaccacc", "abcabccdcdcdef"}

        for i, in := range inputs {
            want := wants[i]
            got := decodeString(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 402. Remove K Digits
	t.Run("Test removeKdigits()", func(t *testing.T) {
        inputs := []string{"1432219", "10200", "10", "112","1234567890"}
        inputs2 := []int{3, 1, 2, 1, 9}
		wants := []string{"1219", "200", "0", "11", "0"}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := removeKdigits(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
