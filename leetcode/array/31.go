package leetcode

// 31. Next permutation
//
// A permutation of an array of integers is an arrangement of its members into a sequence or linear order.
// For example, for arr = [1,2,3], the following are all the permutations of arr: 
// [1,2,3], [1,3,2], [2, 1, 3], [2, 3, 1], [3,1,2], [3,2,1].
// The next permutation of an array of integers is the next lexicographically greater permutation of its integer. 
// More formally, if all the permutations of the array are sorted in one container according to 
// their lexicographical order, then the next permutation of that array is the permutation 
// that follows it in the sorted container. If such arrangement is not possible, the array 
// must be rearranged as the lowest possible order (i.e., sorted in ascending order).
// For example, the next permutation of arr = [1,2,3] is [1,3,2].
// Similarly, the next permutation of arr = [2,3,1] is [3,1,2].
// While the next permutation of arr = [3,2,1] is [1,2,3] because [3,2,1] does not have a lexicographical 
// larger rearrangement.
// Given an array of integers nums, find the next permutation of nums.
// The replacement must be in place and use only constant extra memory.
//
// Runtime: 9 ms, faster than 15.94% of Go online submissions for Next Permutation.
// Memory Usage: 2.5 MB, less than 8.70% of Go online submissions for Next Permutation.
//
func nextPermutation(nums []int) {
    n := len(nums)
    if n == 1 {
        return
    }
    pos := n - 1
    // find the position need to reverse
    for pos > 0 {
        if nums[pos] > nums[pos-1] {
            break
        }
        pos--
    }
    // reverse elements
    for i := 0; 2*i < n-1-pos; i++ {
        nums[pos+i], nums[n-1-i] = nums[n-1-i], nums[pos+i]
    }
    // next permutation
    if pos > 0 {
        for i:= pos; i < n; i++ {
            if nums[i] > nums[pos-1] {
                nums[i], nums[pos-1] = nums[pos-1], nums[i]
                break
            }
        }
    }
}
