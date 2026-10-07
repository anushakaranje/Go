package main

import (
	"fmt"
)

func main() {
	// simple switch
	// 'break' isn't required in switch statement
	// i := 3

	// switch i {
	// case 1:
	// 	fmt.Println("One")
	// case 2:
	// 	fmt.Println("Two")
	// case 3:
	// 	fmt.Println("Three")
	// default:
	// 	fmt.Println("Other")
	// }

	// multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("It's a weekend.")
	// default:
	// 	fmt.Println("It's a workday.")
	// }

	// type switch
	// whoAmI := func(i interface{}) {
	// 	switch t := i.(type) {
	// 	case int:
	// 		fmt.Println("It's an integer.")
	// 	case string:
	// 		fmt.Println("It's a string.")
	// 	case bool:
	// 		fmt.Println("It's a boolean.")
	// 	default:
	// 		fmt.Println("Other", t)
	// 	}
	// }
	// whoAmI(7)

	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("It's an integer.")
		case string:
			fmt.Println("It's a string.")
		case bool:
			fmt.Println("It's a boolean.")
		default:
			fmt.Println("Other")
		}
	}
	whoAmI("anusha")
}
