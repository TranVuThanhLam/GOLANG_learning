package main

import "fmt"

type stackInt struct {
	elements []int
}

func (s *stackInt) Push(r int) {
	s.elements = append(s.elements, r)
}

func (s *stackInt) Pop() {
	if s.isEmpty() {
		return
	}
	s.elements = s.elements[:(len(s.elements) - 1)]
}

func (s *stackInt) Peak() int {
	var r int
	if s.isEmpty() {
		return r
	}

	return s.elements[len(s.elements)-1]
}

func (s *stackInt) isEmpty() bool {
	if len(s.elements) == 0 {
		return true
	} else {
		return false
	}
}

func removeAnElementInArrayByIndex(intArray []int, index int) []int {
	if index >= len(intArray)-2 {
		return intArray
	}

	return append(intArray[:index-1], intArray[index:]...)
}

func removeDuplicates(nums []int) []int {
	var si stackInt
	for index, value := range nums {
		if value == si.Peak() {
			nums = removeAnElementInArrayByIndex(nums, index)
		} else {
			si.Push(value)
		}
	}
	return si.elements
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4, 4, 4, 6, 6, 7}
	fmt.Println(nums)
	nums = removeAnElementInArrayByIndex(nums, 1)
	// nums = removeDuplicates(nums)
	fmt.Println(nums)
}
