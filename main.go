package main

import "fmt"

// RemoveDuplicates compacts a sorted slice in place so each value appears
// once, returning the new length. The write pointer only advances for values
// that differ from the last kept one.
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	w := 1
	for r := 1; r < len(nums); r++ {
		if nums[r] != nums[w-1] {
			nums[w] = nums[r]
			w++
		}
	}
	return w
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	n := RemoveDuplicates(nums)
	fmt.Println(n, nums[:n])
}
