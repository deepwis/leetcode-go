package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 3. Longest Substring Without Repeating Characters
	t.Run("Test lengthOfLongestSubstring()", func(t *testing.T) {
        inputs := []string{"abcabcbb", "bbbbb", "pwwkew", " ", "au", "aabcd", "dvdf", "abba"}
		wants := []int{3, 1, 3, 1, 2, 4, 3, 2}
        for i, in := range inputs {
            want := wants[i]
            got := lengthOfLongestSubstring(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 14. Longest Common Prefix
	t.Run("Test longestCommonPrefix()", func(t *testing.T) {
        inputs := [][]string{
            []string{"flower", "flow", "flight"},
            []string{"dog", "racecar", "car"},
        }
		wants := []string{"fl", ""}
        for i, in := range inputs {
            want := wants[i]
            got := longestCommonPrefix(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 412. Fizz Buzz
	t.Run("Test fizzBuzz()", func(t *testing.T) {
		inputs := []int{3, 5, 15}
        wants := [][]string{
            []string{"1","2","Fizz"},
            []string{"1","2","Fizz","4","Buzz"},
            []string{"1","2","Fizz","4","Buzz","Fizz","7","8","Fizz","Buzz","11","Fizz","13","14","FizzBuzz"},
        }

        for i, in := range inputs {
            want := wants[i]
            got := fizzBuzz(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 5. Longest Palindromic Substring
	t.Run("Test longestPalindrome()", func(t *testing.T) {
        inputs := []string{"babad", "cbbd", "b", "abcdeedcba", "abcba"}
		wants := []string{"bab", "bb", "b", "abcdeedcba", "abcba"}
        for i, in := range inputs {
            want := wants[i]
            got := longestPalindrome(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 151. Reverse Words in a String
	t.Run("Test reverseWords()", func(t *testing.T) {
        inputs := []string{"the sky is blue", "  hello world  ", "a good    example"}
		wants := []string{"blue is sky the", "world hello", "example good a"}
        for i, in := range inputs {
            want := wants[i]
            got := reverseWords(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 415. Add Strings
	t.Run("Test addStrings()", func(t *testing.T) {
        inputs := [][]string{
            []string{"11", "123"},
            []string{"456", "77"},
            []string{"0", "0"},
        }
		wants := []string{"134", "533", "0"}
        for i, in := range inputs {
            want := wants[i]
            got := addStrings(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 468. Validate IP Address
	t.Run("Test validateIPAddress()", func(t *testing.T) {
        inputs := []string{
            "172.16.254.1",
            "2001:0db8:85a3:0:0:8A2E:0370:7334",
            "256.256.256.256",
            "2001:0db8:85a3:0:0:8A2Ea:0370:7334",
            "20EE:FGb8:85a3:0:0:8A2E:0370:7334",
        }
		wants := []string{"IPv4", "IPv6", "Neither", "Neither", "Neither"}
        for i, in := range inputs {
            want := wants[i]
            got := validateIPAddress(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
