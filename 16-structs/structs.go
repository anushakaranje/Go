package main

import (
	"fmt"
	"time"
)

// order struct
// type order struct {
// 	id        string
// 	amount    float32
// 	status    string
// 	createdAt time.Time // nanosecond precision
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}

// 	fmt.Println("Order struct", myOrder)
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}

// 	myOrder.createdAt = time.Now()

// 	fmt.Println(myOrder.status)

// 	fmt.Println("Order struct", myOrder)

// 	myOrder2 := order{
// 		id:        "2",
// 		amount:    100,
// 		status:    "delivered",
// 		createdAt: time.Now(),
// 	}

// 	myOrder.status = "paid"
// 	fmt.Println("Order struct", myOrder2)
// 	fmt.Println("Order struct", myOrder)

// receiver type
// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}

// 	myOrder.changeStatus("confirmed")
// 	fmt.Println(myOrder)

// func (o *order) changeStatus(status string) {
// 	o.status = status
// }

// func (o order) getAmount() float32 {
// 	return o.amount
// }

// func main() {
// 	myOrder := order{
// 		id:     "1",
// 		amount: 50.00,
// 		status: "received",
// 	}

// 	myOrder.changeStatus("confirmed")
// 	fmt.Println(myOrder.getAmount())

// func main() {
// 	language := struct {
// 		name   string
// 		isGood bool
// 	}{"golang", true}

// 	fmt.Println(language)
// }

type customer struct {
	name  string
	phone string
}

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // nanosecond precision
	customer
}

func main() {
	newCustomer := customer{
		name:  "john",
		phone: "1234567890",
	}
	newOrder := order{
		id:       "1",
		amount:   30,
		status:   "received",
		customer: newCustomer,
	}
	newOrder.customer.name = "robin"
	fmt.Println(newOrder)
}
