package main

import "fmt"

// func add(a int, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(3, 5)
// 	fmt.Println(result)
// }

// func add(a, b int) int {
// 	return a + b
// }

// func main() {
// 	result := add(3, 5)
// 	fmt.Println(result)
// }

// func getLanguages() (string, string, string) {
// 	return "golang", "javascript", "c"
// }

// func main() {
// 	fmt.Println(getLanguages())
// }

// func getLanguages() (string, string, bool) {
// 	return "golang", "javascript", true
// }

// func main() {
// 	fmt.Println(getLanguages())
// }

func processIt(fn func(a int) int) func(a int) int {
	return func(a int) int {
		return 4
	}
}

func main() {
	f1 := func(a int) int {
		return a * 2
	}
	fn := processIt(f1)
	result := fn(6)
	fmt.Println(result)
}
