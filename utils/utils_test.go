package utils

import (
    "testing"
    "reflect"
)

func TestUtils(t *testing.T) {
    t.Run("Test Sort Int Slice", func(t *testing.T) {
        input := []int{3,1,5,2}
        want := []int{1,2,3,5}
        SortSlice(input)
        got := input
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
    })

    t.Run("Test Sort Float Slice", func(t *testing.T) {
        input := []float64{3.1,1.2,-2.3,0.1}
        want := []float64{-2.3,0.1,1.2,3.1}
        SortSlice(input)
        got := input
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
    })

    t.Run("Test Sort String Slice", func(t *testing.T) {
        input := []string{"ab","aaa","c","ca"}
        want := []string{"aaa","ab","c","ca"}
        SortSlice(input)
        got := input
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
    })

    t.Run("Test Sort Slice of Int Slice", func(t *testing.T) {
        input := [][]int{
            []int{3,1},
            []int{5,2},
            []int{3,5},
        }
        want := [][]int{
            []int{1,3},
            []int{2,5},
            []int{3,5},
        }
        SortSlice(input)
        got := input
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
    })


    t.Run("Test Sort Slice of Int Slice", func(t *testing.T) {
        input := [][]string{
            []string{"tan", "nat"},
            []string{"eat","ate","tea"},
            []string{"bat"},
        }
        want := [][]string{
            []string{"ate","eat","tea"},
            []string{"bat"},
            []string{"nat","tan"},
        }
        SortSlice(input)
        got := input
        if !reflect.DeepEqual(got, want) {
            t.Errorf("got: %v (%T), want: %v (%T)", got, got, want, want)
        }
    })
}
