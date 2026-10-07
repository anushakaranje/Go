package main

import (
	"fmt"
	"maps"
)

func main() {
	// maps-> hashtable, object, dictionary

	// creating a map
	// m := make(map[string]string)

	// setting an element
	// m["name"] = "golang"

	// getting an element
	// fmt.Println(m["name"])

	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["area"] = "backend"
	// fmt.Println(m["name"], m["area"])

	// if a key value doesn't exist in a map, then it returns zero

	// m := make(map[string]string)
	// m["name"] = "golang"
	// m["area"] = "backend"
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["age"])

	// m := make(map[string]int)
	// m["age"] = 30
	// fmt.Println(m["phone"])

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// fmt.Println(len(m))

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// delete(m, "price")
	// fmt.Println(m)

	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50
	// clear(m)
	// fmt.Println(m)

	// alternate way to create a map
	// m := map[string]int{"price": 40, "phones": 3}
	// fmt.Println(m)

	// m := map[string]int{"price": 40, "phones": 3}
	// _, ok := m["price"]
	// if ok {
	// 	fmt.Println("all ok")
	// } else {
	// 	fmt.Println("not ok")
	// }

	// 	m1 := map[string]int{"price": 40, "phones": 3}
	// 	m2 := map[string]int{"price": 40, "phones": 3}
	// 	fmt.Println(maps.Equal(m1, m2))

	m1 := map[string]int{"price": 40, "phones": 3}
	m2 := map[string]int{"price": 40, "phones": 7}
	fmt.Println(maps.Equal(m1, m2))
}
