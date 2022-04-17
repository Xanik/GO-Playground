package main

import "fmt"

func main() {
	checker := newNums()

	checker.print()
}

type numbers []int

func newNums() numbers {
	nums := []int{}
	for i := 1; i < 11; i++ {
		nums = append(nums, i)
	}
	return nums
}

func (n numbers) print() {
	for _, v := range n {
		if v%2 == 0 {
			fmt.Println(v, ": even")
		} else {
			fmt.Println(v, ": odd")
		}
	}
}
