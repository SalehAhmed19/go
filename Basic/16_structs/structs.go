package main

import (
	"fmt"
	"time"
)

// order struct
type Order struct {
	id        int
	amount    float32
	status    string
	createdAt time.Time // time.Time is a struct from time inbuild package - nanosecond precision
}

// go does not have constructors - we can create methods to work with structs - newFuncname is convention
func newOrder(id int, amount float32, status string) *Order { // return pointer to Order struct
	myOrder := Order{
		id:        id,
		amount:    amount,
		status:    status,
		createdAt: time.Now(),
	}
	return &myOrder // return address of myOrder
}

func (o *Order) changeStatus(status string) { // (o type) is a receiver - use asterisk to change original struct
	o.status = status // automatically modifies the original struct - dereferencing
}

func main() {
	order := Order{ // creating an instance of struct
		id:     1,
		amount: 230,
		status: "Pending",
	}

	// if you don't assign createdAt field, it will have zero value

	order.createdAt = time.Now() // assign current time to createdAt field

	fmt.Println(order.amount) // get specific field

	fmt.Println(order)

	order2 := Order{
		id:        2,
		amount:    240,
		status:    "Received",
		createdAt: time.Now(),
	}

	order.changeStatus("Confirmed")

	fmt.Println(order2)

	fmt.Println("Order:", order)

	makeOrder := newOrder(3, 250, "Shipped") // using newOrder function to create an instance

	fmt.Println(makeOrder) // don't need to dereference pointer to print struct fields

	language := struct { // anonymous struct and initialization at the same time - inline struct
		name   string
		isGood bool
	}{
		name:   "Golang",
		isGood: true,
	}

	fmt.Println(language)
}
