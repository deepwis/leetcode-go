package utils

import (
    "sort"
    "fmt"
)

func InArray(x any, arr any) bool {
    switch v := x.(type) {
    case int:
        if a, ok := arr.([]int); !ok {
            return false
        } else {
            for _, vv := range a {
                if v == vv {
                    return true
                }
            }
            return false
        }
    case string:
        if a, ok := arr.([]string); !ok {
            return false
        } else {
            for _, vv := range a {
                if v == vv {
                    return true
                }
            }
            return false
        }
    case float64:
        if a, ok := arr.([]float64); !ok {
            return false
        } else {
            for _, vv := range a {
                if v == vv {
                    return true
                }
            }
            return false
        }
    }
    return false
}

func SortSlice(x any) {
    switch v := x.(type) {
    case []int:
        sort.Ints(v)
    case []string:
        sort.Strings(v)
    case []float64:
        sort.Float64s(v)
    case [][]int:
        for i := 0; i < len(v); i++ {
            sort.Ints(v[i])
        }
        sort.Slice(v, func(i, j int) bool {
            for k := 0; k < len(v[i]); k++ {
                if k == len(v[j]) || v[i][k] > v[j][k] {
                    return false
                } else if v[i][k] < v[j][k] {
                    return true
                }
            }
            return true
        })
    case [][]string:
        for i := 0; i < len(v); i++ {
            sort.Strings(v[i])
        }
        sort.Slice(v, func(i, j int) bool {
            for k := 0; k < len(v[i]); k++ {
                if k == len(v) || v[i][k] > v[j][k] {
                    return false
                } else if v[i][k] < v[j][k] {
                    return true
                }
            }
            return true
        })
    case [][]float64:
        for i := 0; i < len(v); i++ {
            sort.Float64s(v[i])
        }
        sort.Slice(v, func(i, j int) bool {
            for k := 0; k < len(v[i]); k++ {
                if k == len(v[j]) || v[i][k] > v[j][k] {
                    return false
                } else if v[i][k] < v[j][k] {
                    return true
                }
            }
            return true
        })
    default:
        fmt.Printf("Unknown type x: %v, (%T)\n", x, x)
    }
}

