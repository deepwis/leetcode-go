package leetcode

import "math/rand"

// 470. Implement Rand10() using Rand7()
//
// Given the API rand7() that generates a uniform random integer in the range [1, 7], 
// write a function rand10() that generates a uniform random integer in the range [1, 10]. 
// You can only call the API rand7(), and you shouldn't call any other API. 
// Please do not use a language's built-in random API.
// Each test case will have one internal argument n, the number of times that your implemented function 
// rand10() will be called while testing. Note that this is not an argument passed to rand10().
//
// Runtime: 15 ms, faster than 100.00% of Go online submissions for Implement Rand10() Using Rand7().
// Memory Usage: 5.6 MB, less than 100.00% of Go online submissions for Implement Rand10() Using Rand7().
//
func rand10() int {
    r := 100
    for r >= 40 {
        r = (rand7() -1) * 7 + rand7() - 1
    }
    return r % 10 + 1
}

func rand7() int {
    return rand.Intn(7) + 1
}

