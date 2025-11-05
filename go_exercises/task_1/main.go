package main

import "fmt"

func sum(nums []int) int {
	ans := 0
	for _, val := range nums {
		ans += val
	}
	return ans
}

func main() {
	var nums = []int{1,2,3,4,5}
	fmt.Print("{1, 2, 3, 4, 5}: ")
	fmt.Println(sum(nums))
	nums = []int{}
	fmt.Print("{}: ")
	fmt.Println(sum(nums))
}
