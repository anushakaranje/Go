package main

import (
	"fmt"
)

// slice -> dynamic
// most used construct in golang
// + useful methods

func main() {
	// uninitialized slice is nil
	// var num []int
	// fmt.Println(num == nil)

	// var num []int
	// fmt.Println(len(num))

	// var num = make([]int, 2)
	// fmt.Println(num)

	// var num = make([]int, 2)
	// fmt.Println(cap(num))

	// var num = make([]int, 2, 5)
	// num = append(num, 1)
	// num = append(num, 2)
	// num = append(num, 3)
	// num = append(num, 5)
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// num := []int{}
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// num := []int{}
	// num = append(num, 1)
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// var num = make([]int, 2)
	// num[0] = 3
	// num[1] = 5
	// fmt.Println(num)
	// fmt.Println(cap(num))
	// fmt.Println(len(num))

	// var num = make([]int, 0, 5)
	// var num1 = make([]int, len(num))
	// num = append(num, 2)
	// fmt.Println(num, num1)

	// copy function
	// var num = make([]int, 0, 5)
	// num = append(num, 2)
	// var num1 = make([]int, len(num))
	// copy(num1, num)
	// fmt.Println(num, num1)

	// slice operator
	// var num = []int{1, 2, 3, 4, 5}
	// fmt.Println(num[0:1])
	// fmt.Println(num[:2])
	// fmt.Println(num[1:])

	// slice
	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 3}
	// fmt.Println(slices.Equal(num1, num2))

	// var num1 = []int{1, 2, 3}
	// var num2 = []int{1, 2, 4}
	// fmt.Println(slices.Equal(num1, num2))

	var num = [][]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(num)
}
