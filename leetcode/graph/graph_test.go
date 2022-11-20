package leetcode

import (
	"testing"
)

func TestSolution(t *testing.T) {
    // 207. Course Schedule
	t.Run("Test canFinish()", func(t *testing.T) {
        inputs1 := []int{2, 2, 3, 3, 5}
        inputs2 := [][][]int{
            [][]int{
                []int{1,0},
            },
            [][]int{
                []int{1,0},
                []int{0,1},
            },
            [][]int{
                []int{0,1},
                []int{0,2},
                []int{1,2},
            },
            [][]int{
                []int{1,0},
                []int{1,2},
                []int{0,1},
            },
            [][]int{
                []int{1,4},
                []int{2,4},
                []int{3,1},
                []int{3,2},
            },
        }
		wants := []bool{true, false, true, false, true}
        for i, in1 := range inputs1 {
            in2 := inputs2[i]
            want := wants[i]
            got := canFinish(in1, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 1672. Richest Customer Wealth
	t.Run("Test maximumWealth()", func(t *testing.T) {
        input := [][]int{
            []int{1,2,3},
            []int{3,2,1},
        }
		want := 6
        got := maximumWealth(input)
        if got != want {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }

        input = [][]int{
            []int{1,5},
            []int{7,3},
            []int{3,5},
        }
		want = 10
        got = maximumWealth(input)
        if got != want {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }

        input = [][]int{
            []int{2,8,7},
            []int{7,1,3},
            []int{1,9,5},
        }
		want = 17
        got = maximumWealth(input)
        if got != want {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
	})

    // 200. Number of Islands
	t.Run("Test numIslands()", func(t *testing.T) {
        inputs := [][][]byte{
            [][]byte{
                []byte{'1','1','1','1','0'},
                []byte{'1','1','0','1','0'},
                []byte{'1','1','0','0','0'},
                []byte{'0','0','0','0','0'},
            },
            [][]byte{
                []byte{'1','1','0','0','0'},
                []byte{'1','1','0','0','0'},
                []byte{'0','0','1','0','0'},
                []byte{'0','0','0','1','1'},
            },
        }
		wants := []int{1, 3}
        for i, in := range inputs {
            want := wants[i]
            got := numIslands(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 79. Word Search
	t.Run("Test exist()", func(t *testing.T) {
        inputs := [][][]byte{
            [][]byte{
                []byte{'A','B','C','E'},
                []byte{'S','F','C','S'},
                []byte{'A','D','E','E'},
            },
            [][]byte{
                []byte{'A','B','C','E'},
                []byte{'S','F','C','S'},
                []byte{'A','D','E','E'},
            },
            [][]byte{
                []byte{'A','B','C','E'},
                []byte{'S','F','C','S'},
                []byte{'A','D','E','E'},
            },
            [][]byte{
                []byte{'a','b'},
                []byte{'c','d'},
            },
            [][]byte{
                []byte{'A','B','C','E'},
                []byte{'S','F','E','S'},
                []byte{'A','D','E','E'},
            },
        }
        inputs2 := []string{"ABCCED", "SEE", "ABCB","cdba", "ABCESEEEFS"}
		wants := []bool{true, true, false, true, true}
        for i, in := range inputs {
            in2 := inputs2[i]
            want := wants[i]
            got := exist(in, in2)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

    // 994. Rotting Oranges
	t.Run("Test orangesRotting()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{2,1,1},
                []int{1,1,0},
                []int{0,1,1},
            },
            [][]int{
                []int{2,1,1},
                []int{0,1,1},
                []int{1,0,1},
            },
            [][]int{
                []int{1},
                []int{2},
            },
        }
		wants := []int{4, -1, 1}

        for i, in := range inputs {
            want := wants[i]
            got := orangesRotting(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
    })

    // 221. Maximal Square
	t.Run("maximalSquare()", func(t *testing.T) {
		inputs := [][][]byte{
            [][]byte{
                []byte{'1','0','1','0','0'},
                []byte{'1','0','1','1','1'},
                []byte{'1','1','1','1','1'},
                []byte{'1','0','0','1','0'},
            },
            [][]byte{
                []byte{'0', '1'},
                []byte{'1', '0'},
            },
            [][]byte{
                []byte{'0'},
            },
        }
        wants := []int{4, 1, 0}

        for i, in := range inputs {
            want := wants[i]
            got := maximalSquare(in)
            if got != want {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})
    // 695. Max Area of Island
	t.Run("Test maxAreaOfIsland()", func(t *testing.T) {
        inputs := [][][]int{
            [][]int{
                []int{0,0,1,0,0,0,0,1,0,0,0,0,0},
                []int{0,0,0,0,0,0,0,1,1,1,0,0,0},
                []int{0,1,1,0,1,0,0,0,0,0,0,0,0},
                []int{0,1,0,0,1,1,0,0,1,0,1,0,0},
                []int{0,1,0,0,1,1,0,0,1,1,1,0,0},
                []int{0,0,0,0,0,0,0,0,0,0,1,0,0},
                []int{0,0,0,0,0,0,0,1,1,1,0,0,0},
                []int{0,0,0,0,0,0,0,1,1,0,0,0,0},
            },
            [][]int{
                []int{0,0,0,0,0,0},
            },
        }
		wants := []int{6, 0}
        for i, in := range inputs {
            want := wants[i]
            got := maxAreaOfIsland(in)
            if got != want {
                t.Errorf("got: %v, want: %v", got, want)
            }
        }
	})

}
