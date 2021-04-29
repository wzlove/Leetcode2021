package main

/*
	https://leetcode-cn.com/problems/two-sum/
*/
func twoSum(nums []int, target int) []int {
	tempMap := make(map[int]int, 0)
	for index, value := range nums {
		sub := target - value
		if result, ok := tempMap[sub]; ok {
			return []int{result, index}
		}
		tempMap[value] = index
	}
	return []int{}
}
