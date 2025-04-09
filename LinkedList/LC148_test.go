package LinkedList

import (
	"fmt"
	"testing"
)

// Helper function to create a linked list from a slice
func arrayToList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for _, num := range nums {
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return dummy.Next
}

// Helper function to convert a linked list to a slice
func listToArray(head *ListNode) []int {
	var result []int
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func TestSortListIII(t *testing.T) {
	// 测试用例
	tests := []struct {
		input    []int
		expected []int
	}{
		{input: []int{-1, 5, 3, 4, 0}, expected: []int{-1, 0, 3, 4, 5}},
		{input: []int{4, 2, 1, 3}, expected: []int{1, 2, 3, 4}},
		{input: []int{}, expected: []int{}},
		{input: []int{1}, expected: []int{1}},
	}

	for i, test := range tests {
		head := arrayToList(test.input)
		sortedHead := sortListIII(head)
		output := listToArray(sortedHead)

		if fmt.Sprintf("%v", output) != fmt.Sprintf("%v", test.expected) {
			fmt.Printf("Test %d failed: input=%v, expected=%v, got=%v\n", i+1, test.input, test.expected, output)
		} else {
			fmt.Printf("Test %d passed: input=%v, output=%v\n", i+1, test.input, output)
		}
	}
}
