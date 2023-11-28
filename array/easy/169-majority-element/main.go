package main

import "fmt"

func twoSum(nums []int, target int) []int {
	// key为值，value为下标
	m := make(map[int]int, len(nums))

	for i := range nums {
		dev := target - nums[i]
		fmt.Println(dev, m)
		_, exist := m[dev]
		if exist {
			return []int{m[dev], i}
		}
		m[nums[i]] = i
	}

	return []int{}
}
