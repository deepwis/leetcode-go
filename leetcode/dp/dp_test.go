package leetcode

import (
	"reflect"
	"testing"
)

func TestSolution(t *testing.T) {
    // 1143. Longest Common Subsequence
	t.Run("Test longestCommonSubsequence()", func(t *testing.T) {
        inputs := [][]string{
            []string{"abcba", "abcbcba"},
            []string{"oxcpqrsvwf", "shmtulqrypy"},
            []string{"pqrsqp", "sqrp"},
            []string{"pqrs", "sqrp"},
            []string{"bsbininm", "jmjkbkjkv"},
        }
		wants := []int{5,2,3,2,1}

        for i, in := range inputs {
            want := wants[i]
            got := longestCommonSubsequence(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 122. Best Time to Buy and Sell Stock II
	t.Run("Test maxProfit()", func(t *testing.T) {
        inputs := [][]int{
            []int{7,1,5,3,6,4},
            []int{1,2,3,4,5},
            []int{7,6,4,3,1},
        }
		wants := []int{7, 4, 0}
        for i, in := range inputs {
            want := wants[i]
            got := maxProfit(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 128. Longest Consecutive Sequence
	t.Run("Test longestConsecutive()", func(t *testing.T) {
        inputs := [][]int{
            []int{100,4,200,1,3,2},
            []int{0,3,7,2,5,8,4,6,0,1},
            []int{1},
            []int{},
            []int{3,8,-2,0,-1,1,0,2},
        }
		wants := []int{4, 9, 1, 0, 6}
        for i, in := range inputs {
            want := wants[i]
            got := longestConsecutive(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 131. Palindrome Partitioning
	t.Run("Test partition()", func(t *testing.T) {
        inputs := []string{"a", "ab", "aab", "fff", "abbab"}
		wants := [][][]string{
            [][]string{
                []string{"a"},
            },
            [][]string{
                []string{"a", "b"},
            },
            [][]string{
                []string{"a","a","b"},
                []string{"aa","b"},
            },
            [][]string{
                []string{"f","f","f"},
                []string{"ff","f"},
                []string{"fff"},
                []string{"f","ff"},
            },
            [][]string{
                []string{"a","b","b","a","b"},
                []string{"a","bb","a","b"},
                []string{"abba","b"},
                []string{"a","b","bab"},
            },
        }
        for i, in := range inputs {
            want := wants[i]
            got := partition(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 152. Maximum Product Subarray
	t.Run("Test maxProduct()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,3,-2,4},
            []int{-2,0,-1},
            []int{1,1,5},
            []int{-1,-3,-2},
            []int{-3,-1,-1},
        }
		wants := []int{6, 0, 5, 6, 3}
        for i, in := range inputs {
            want := wants[i]
            got := maxProduct(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 198. House Robber
	t.Run("Test rob()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3,1},
            []int{2,7,9,3,1},
            []int{1},
        }
		wants := []int{4, 12, 1}
        for i, in := range inputs {
            want := wants[i]
            got := rob(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 22. Generate Parentheses
	t.Run("Test generateParentheses()", func(t *testing.T) {
        inputs := []int{1,2,3}
		wants := [][]string{
            []string{"()"},
            []string{"()()","(())"},
            []string{"()()()","()(())","(())()","(()())","((()))"},
        }

        for i, in := range inputs {
            want := wants[i]
            got := generateParentheses(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 300. Longest Increasing Subsequence 
	t.Run("Test lengthOfLIS()", func(t *testing.T) {
        inputs := [][]int{
            []int{10,9,2,5,3,7,101,18},
            []int{0,1,0,3,2,3},
            []int{7,7,7,7,7,7,7},
            []int{1,3,6,7,9,4,10,5,6},
        }
		wants := []int{4, 4, 1, 6}

        for i, in := range inputs {
            want := wants[i]
            got := lengthOfLIS(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 322. Coin Change
	t.Run("Test coinChange()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,5},
            []int{2},
            []int{1},
            []int{1},
            []int{1,2147483647},
            []int{442,295,365,485},
        }
        inputs2 := []int{11, 3, 0, 2, 2, 8091}
		wants := []int{3, -1, 0, 2, 2, 20}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := coinChange(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 416. Partition Equal Subset Sum
	t.Run("Test canPartition()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,5,11,5},
            []int{1,2,3,5},
            []int{1,2,5},
        }
		wants := []bool{true, false, false}

        for i, in := range inputs {
            want := wants[i]
            got := canPartition(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 45. Jump Game II
	t.Run("Test jump()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,3,1,1,4},
            []int{3,2,0,1,4},
            []int{0},
        }
		wants := []int{2,2,0}

        for i, in := range inputs {
            want := wants[i]
            got := jump(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
            }
        }
    })

    // 46. Permutations
	t.Run("Test permute()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3},
            []int{0,1},
            []int{1},
        }
		wants := [][][]int{
            [][]int{
                []int{3,2,1},
                []int{2,3,1},
                []int{2,1,3},
                []int{3,1,2},
                []int{1,3,2},
                []int{1,2,3},
            },
            [][]int{
                []int{1,0},
                []int{0,1},
            },
            [][]int{
                []int{1},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := permute(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 55. Jump Game
	t.Run("Test canJump()", func(t *testing.T) {
        inputs := [][]int{
            []int{2,3,1,1,4},
            []int{3,2,1,0,4},
        }
		wants := []bool{true,false}

        for i, in := range inputs {
            want := wants[i]
            got := canJump(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 62. Unique Paths
	t.Run("Test uniquePaths()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,1},
            []int{1,2},
            []int{2,2},
            []int{2,3},
            []int{3,2},
            []int{3,7},
        }
		wants := []int{1,1,2,3,3,28}

        for i, in := range inputs {
            want := wants[i]
            got := uniquePaths(in[0], in[1])
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 64. Minimum Path Sum
	t.Run("Test minPathSum()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,3,1},
                []int{1,5,1},
                []int{4,2,1},
            },
            [][]int{
                []int{1,2,3},
                []int{4,5,6},
            },
            [][]int{
                []int{1},
            },
            [][]int{
                []int{1,2},
            },
            [][]int{
                []int{1},
                []int{3},
            },
        }
		wants := []int{7,12,1,3,4}

        for i, in := range inputs {
            want := wants[i]
            got := minPathSum(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 647. Palindromic Substrings
	t.Run("Test countSubstrings()", func(t *testing.T) {
        inputs := []string{
            "abc",
            "aaa",
        }
		wants := []int{3, 6}

        for i, in := range inputs {
            want := wants[i]
            got := countSubstrings(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 70. Climbing Stairs
	t.Run("Test climbStairs()", func(t *testing.T) {
        inputs := []int{1,2,3,4}
		wants := []int{1,2,3,5}

        for i, in := range inputs {
            want := wants[i]
            got := climbStairs(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 72. Edit Distance
	t.Run("Test minDistance()", func(t *testing.T) {
        inputs := [][]string{
            []string{"Horse","ros"},
            []string{"intention","execution"},
        }
		wants := []int{3, 5}
        for i, in := range inputs {
            want := wants[i]
            got := minDistance(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 739. Daily Temperatures
	t.Run("Test dailyTemperatures()", func(t *testing.T) {
        inputs := [][]int{
            []int{73,74,75,71,69,72,76,73},
            []int{30,40,50,60},
            []int{30,60,90},
        }
		wants := [][]int{
            []int{1,1,4,2,1,1,0,0},
            []int{1,1,1,0},
            []int{1,1,0},
        }

        for i, in := range inputs {
            want := wants[i]
            got := dailyTemperatures(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 78. Subsets
	t.Run("Test subsets()", func(t *testing.T) {
        inputs := [][]int{
            []int{1,2,3},
            []int{0},
            []int{9,0,3,5,7},
        }
		wants := [][][]int{
            [][]int{
                []int{},
                []int{1},
                []int{2},
                []int{1,2},
                []int{3},
                []int{1,3},
                []int{2,3},
                []int{1,2,3},
            },
            [][]int{
                []int{},
                []int{0},
            },
            [][]int{
                []int{},
                []int{9},
                []int{0},
                []int{9,0},
                []int{3},
                []int{9,3},
                []int{0,3},
                []int{9,0,3},
                []int{5},
                []int{9,5},
                []int{0,5},
                []int{9,0,5},
                []int{3,5},
                []int{9,3,5},
                []int{0,3,5},
                []int{9,0,3,5},
                []int{7},
                []int{9,7},
                []int{0,7},
                []int{9,0,7},
                []int{3,7},
                []int{9,3,7},
                []int{0,3,7},
                []int{9,0,3,7},
                []int{5,7},
                []int{9,5,7},
                []int{0,5,7},
                []int{9,0,5,7},
                []int{3,5,7},
                []int{9,3,5,7},
                []int{0,3,5,7},
                []int{9,0,3,5,7},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := subsets2(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T), input: %v (%T)", got, got, want, want, in, in)
            }
        }
    })

    // 96. Unique Binary Search Trees
	t.Run("Test numTrees()", func(t *testing.T) {
        inputs := []int{1,2,3}
		wants := []int{1,2,5}

        for i, in := range inputs {
            want := wants[i]
            got := numTrees(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
    })

    // 718. Maximum Length of Repeated Subarrray
	t.Run("Test findLength()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{1,2,3,2,1},
                []int{3,2,1,4,7},
            },
            [][]int{
                []int{0,0,0,0,0},
                []int{0,0,0,0,0},
            },
        }
		wants := []int{3,5}
        for i, in := range inputs {
            want := wants[i]
            got := findLength(in[0], in[1])
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
