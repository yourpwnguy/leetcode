package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)

	for i, v := range nums {
		missing := target - v
		if j, ok := seen[missing]; ok {
			return []int{i, j}
		}
		seen[v] = i
	}
	return nil
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
}
