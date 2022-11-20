package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 1. Two Sum
	t.Run("Test twoSum()", func(t *testing.T) {
        inputs1 := [][]int{
            []int{2,7,11,15},
            []int{3,2,4},
            []int{3,3},
        }
        inputs2 := []int{9,6,6}
		wants := [][]int{
            []int{0,1},
            []int{1,2},
            []int{0,1},
        }

        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := twoSum(in1, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 13. Roman to Integer
	t.Run("Test romanToInt()", func(t *testing.T) {
        inputs := []string{
            "III",
            "LVIII",
            "MCMXCIV",
        }
		wants := []int{3, 58, 1994}

        for i, s := range inputs {
            want := wants[i]
            got := romanToInt(s)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 13. Roman to Integer
	t.Run("Test romanToInt()", func(t *testing.T) {
        inputs := []string{
            "III",
            "LVIII",
            "MCMXCIV",
        }
		wants := []int{3, 58, 1994}

        for i, s := range inputs {
            want := wants[i]
            got := romanToInt2(s)
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 139. Word Break
	t.Run("Test wordBreak()", func(t *testing.T) {
        inputs1 := []string{
            "leetcode",
            "applepenapple",
            "catsandog",
            "dogs",
            "aebbbbs",
            "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
        }
        inputs2 := [][]string{
            []string{"leet","code"},
            []string{"apple","pen"},
            []string{"cats","dog","sand","and","cat"},
            []string{"dog","s","gs"},
            []string{"a","aeb","ebbbb","s","eb"},
            []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"},
        }
		wants := []bool{true, true, false, true, true, false}
        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := wordBreak(in1, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 169. Majority Element
	t.Run("Test majorityElement()", func(t *testing.T) {
        inputs := [][]int{
            []int{3,2,3},
            []int{2,2,1,1,1,2,2},
            []int{6,5,5},
            []int{1},
        }
		wants := []int{3, 2, 5, 1}
        for i, in := range inputs { want := wants[i]
            got := majorityElement(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 17. Letter Combinations of a Phone Number
	t.Run("Test letterCombinations()", func(t *testing.T) {
        inputs := []string{"23", "2", "234"}
		wants := [][]string{
            []string{"ad","ae","af","bd","be","bf","cd","ce","cf"},
            []string{"a","b","c"},
            []string{"adg","adh","adi","aeg","aeh","aei","afg","afh","afi","bdg","bdh","bdi","beg","beh","bei","bfg","bfh","bfi","cdg","cdh","cdi","ceg","ceh","cei","cfg","cfh","cfi"},
        }
        for i, in := range inputs {
            want := wants[i]
            got := letterCombinations(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 287. Find the Duplicate Number
	t.Run("Test findDuplicate()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,3,4,2,2},
            []int{3,1,3,4,2},
            []int{1,1},
        }
		wants := []int{2,3,1}
        for i, in := range inputs {
            want := wants[i]
            got := findDuplicate(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 347. Top K Frequent Elements
	t.Run("Test topKFrequent()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,1,1,2,2,3},
            []int{3,1,3,4,2},
            []int{1},
        }
        inputs2 := []int{2, 1, 1}
		wants := [][]int{
            []int{1,2},
            []int{3},
            []int{1},
        }
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := topKFrequent(in, in2)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 383. Ransom Note
	t.Run("Test canConstruct()", func(t *testing.T) {
        inputs := [][2]string{
            [2]string{"a", "b"},
            [2]string{"aa", "ab"},
            [2]string{"aa", "aab"},
        }
		wants := []bool{false, false, true}

        for i, in := range inputs {
            want := wants[i]
            got := canConstruct(in[0], in[1])
            if want != got {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 41. First Missing Positive
	t.Run("Test firstMissingPositive()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,0},
            []int{3,4,-1,1},
            []int{7,8,9,11,12},
            []int{0,3,8},
            []int{},
            []int{2,1,2},
        }
		wants := []int{3,2,1,1,1,3}

        for i, in := range inputs {
            want := wants[i]
            got := firstMissingPositive(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 49. Group Anagrams
	t.Run("Test groupAnagrams()", func(t *testing.T) {
        inputs := [][]string{
            []string{"eat","tea","tan","ate","nat","bat"},
            []string{""},
            []string{"a"},
            []string{
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
                "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
            },
        }
		wants := [][][]string{
            [][]string{
                []string{"bat"},
                []string{"nat","tan"},
                []string{"ate","eat","tea"},
            },
            [][]string{
                []string{""},
            },
            [][]string{
                []string{"a"},
            },
            [][]string{
                []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab"},
                []string{"aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"},
            },
        }
        for i, in := range inputs {
            want := wants[i]
            utils.SortSlice(want)
            got := groupAnagrams(in)
            utils.SortSlice(got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 560. Subarray Sum Equals K
	t.Run("Test subarraySum()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,1,1},
            []int{1,2,3},
        }
        inputs2 := []int{2, 3}
		wants := []int{2, 2}

        for i, in := range inputs {
            want := wants[i]
            in2 := inputs2[i]
            got := subarraySum(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 73. Set Matrix Zeroes
	t.Run("Test setZeroes()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,1,1},
                []int{1,0,1},
                []int{1,1,1},
            },
            [][]int{
                []int{0,1,2,0},
                []int{3,4,5,2},
                []int{1,3,1,5},
            },
        }
		wants := [][][]int{
            [][]int{
                []int{1,0,1},
                []int{0,0,0},
                []int{1,0,1},
            },
            [][]int{
                []int{0,0,0,0},
                []int{0,4,5,0},
                []int{0,3,1,0},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            setZeroes(in)
            got := in
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 76. Minimum Window Substring
	t.Run("Test minWindow()", func(t *testing.T) {
        inputs := [][]string{
            []string{"ADOBECODEBANC", "ABC"},
            []string{"a", "a"},
            []string{"a", "aa"},
            []string{"a", "b"},
            []string{"bdab", "ab"},
            []string{"cabwefgewcwaefgcf", "cae"},
        }
		wants := []string{"BANC", "a", "", "", "ab", "cwae"}
        for i, in := range inputs {
            want := wants[i]
            got := minWindow(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 138. Copy List with Random Pointer
    // 141. Linked List Cycle
    // 142. Linked List Cycle II
    // 160. Intersection of Two Linked Lists
}
