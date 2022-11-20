package leetcode

import (
    "strconv"
)

// 412. Fizz Buzz
//
// Given an integer n, return a string array answer (1-indexed) where:
// * answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
// * answer[i] == "Fizz" if i is divisible by 3.
// * answer[i] == "Buzz" if i is divisible by 5.
// * answer[i] == i (as a string) if none of the above conditions are true.
//
// Runtime: 12 ms, faster than 23.25% of Go online submissions for Fizz Buzz.
// Memory Usage: 3.6 MB, less than 65.38% of Go online submissions for Fizz Buzz.
//
func fizzBuzz(n int) []string {
    answer := make([]string, n, n)
    for i:= 1; i <= n; i++ {
        if i % 3 == 0 {
            if i % 5 == 0 {
                answer[i-1] = "FizzBuzz"
            } else {
                answer[i-1] = "Fizz"
            }
        } else if i % 5 == 0 {
            answer[i-1] = "Buzz"
        } else {
            answer[i-1] = strconv.Itoa(i)
        }
    }
    return answer
}

