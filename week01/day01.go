package main

import "fmt"

/*
	https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		numMap[num] = i
	}

	var result []int
	for i, num := range nums {
		value := target - num

		if index, ok := numMap[value]; ok && index != i {
			result = append(result, i, index)
			return result
		}
	}

	return result
}

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", twoSum([]int{3, 2, 4}, 6))
	fmt.Printf("%v\n", twoSum([]int{3, 3}, 6))
}
