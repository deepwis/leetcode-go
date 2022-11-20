package leetcode

import (
	"reflect"
	"testing"
    "github.com/deepwis/leetcode-go/utils"
)

func TestSolution(t *testing.T) {
    // 51. N-Queens
	t.Run("Test solveNQueens()", func(t *testing.T) {
        inputs := []int{4,5}
		wants := [][][]string{
            [][]string{
                []string{".Q..","...Q","Q...","..Q."},
                []string{"..Q.","Q...","...Q",".Q.."},
            },
            [][]string{
                []string{"Q....","..Q..","....Q",".Q...","...Q."},
                []string{"Q....","...Q.",".Q...","....Q","..Q.."},
                []string{".Q...","...Q.","Q....","..Q..","....Q"},
                []string{".Q...","....Q","..Q..","Q....","...Q."},
                []string{"..Q..","Q....","...Q.",".Q...","....Q"},
                []string{"..Q..","....Q",".Q...","...Q.","Q...."},
                []string{"...Q.","Q....","..Q..","....Q",".Q..."},
                []string{"...Q.",".Q...","....Q","..Q..","Q...."},
                []string{"....Q",".Q...","...Q.","Q....","..Q.."},
                []string{"....Q","..Q..","Q....","...Q.",".Q..."},
            },
        }

        for i, in := range inputs {
            want := wants[i]
            got := solveNQueens(in)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})

    // 93. Restore IP Addresses
	t.Run("restoreIpAddresses()", func(t *testing.T) {
		inputs := []string{"25525511135","0000", "101023"}
        wants := [][]string{
            []string{"255.255.11.135", "255.255.111.35"},
            []string{"0.0.0.0"},
            []string{"10.10.2.3","10.1.0.23", "1.0.10.23", "1.0.102.3", "101.0.2.3"},
        }

        for i, in := range inputs {
            want := wants[i]
            utils.SortSlice(want)
            got := restoreIpAddresses(in)
            utils.SortSlice(got)
            if !reflect.DeepEqual(got, want) {
                t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
            }
        }
	})
}
