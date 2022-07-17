/*
Write a function solution that, given an array A of N integers, returns the largest integer K > 0 such that both values K and −K (the opposite number) exist in array A.
If there is no such integer, the function should return 0.

Examples:

1. Given A = [3, 2, −2, 5, −3], the function should return 3 (both 3 and −3 exist in array A).
2. Given A = [1, 1, 2, −1, 2, −1], the function should return 1 (both 1 and −1 exist in array A).
3. Given A = [1, 2, 3, −4], the function should return 0 (there is no such K for which both values K and −K exist in array A).

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000,000..1,000,000,000].

*/

package main

import (
	"log"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func solution1(A []int) int {
	K := 0
	myArray := []int{}

	for i := 0; i < len(A); i++ {
		myArray = append(myArray, A[i])
		if contains(myArray, -1*A[i]) {
			K = max(K, abs(A[i]))
		}
	}
	return K
}

func solution2(A []int) int {
	ret := 0
	tmp_map := map[int]bool{}
	for _, num := range A {
		if _, exist := tmp_map[-num]; exist {
			if abs(num) > ret {
				ret = abs(num)
			}
			continue
		}
		tmp_map[num] = true
	}
	return ret
}

func main() {

	A := []int{3, 2, -2, 5, -3}
	B := []int{1, 1, 2, -1, 2, -1}
	C := []int{1, 2, 3, -4}

	log.Println(solution1(A))
	log.Println(solution1(B))
	log.Println(solution1(C))

	log.Println(solution2(A))
	log.Println(solution2(B))
	log.Println(solution2(C))

	//go mod init main
	//go mod tidy
	// Run and Debug
}
