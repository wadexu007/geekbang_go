/*
Given a non-empty array of integers nums, every element appears twice except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

*/

package main

import "log"

func singleNumber3(nums []int) int {
	ret := 0
	for _, num := range nums {
		ret ^= num // ret = ret ^ num 异或操作
	}
	return ret
}

func main() {
	nums := []int{2, 1, 1, 4, 4}
	log.Println(singleNumber3(nums))
}
