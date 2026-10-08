package main

import "fmt"

// func sum(num ...int) int {
// 	total := 0

// 	for _, nums := range num
// 	total = total + nums
// }
// func main() {
// 	result := sum(3, 4, 5)
// 	fmt.Println(result)
// }

func sum(num ...int) int {
	total := 0

	for _, nums := range num {
		total = total + nums
	}
	return total
}
func main() {
	num := []int{3, 4, 5, 6}
	result := sum(num...)
	fmt.Println(result)
}
