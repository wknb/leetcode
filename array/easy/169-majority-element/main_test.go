// main_test.go

package main

import (
	"testing"
)

func TestTwoSum(t *testing.T) {
	// 编写测试用例
	tests := []struct {
		nums     []int
		target   int
		expected []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
	}

	for _, test := range tests {
		result := twoSum(test.nums, test.target)
		if !equal(result, test.expected) {
			t.Errorf("For nums=%v, target=%d, expected=%v, got=%v", test.nums, test.target, test.expected, result)
		}
	}
}

// 辅助函数，用于比较两个切片是否相等
func equal(slice1, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}
