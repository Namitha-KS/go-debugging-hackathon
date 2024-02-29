package main

import "fmt"

func main() {
	// define variable a
	// define variable b
	// define variable c

	// read three variable inputs from console
	// use fmt.Scan to read
	// use fmt.Println to write
	var a, b, c int
	fmt.Println("enter 3 numbers :")
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Println(a, b, c)

	// call add function
	// print result here
	sum := add(a, b, c)
	fmt.Println("sum : ", sum)

	// call average function from second.go
	// without returing a variable get the result here
	// modifying average functions deinfiniton is allowed
	// print the average here
	fmt.Println("average : ", average(a, b, c))

	// call min_max function from second.go
	// print min and max here
	min, max := min_max(a, b, c)
	fmt.Println("minimum : ", min)
	fmt.Println("maximum : ", max)

	// create a variable radius and assign value of `b` to it
	// call calculate_circle_area from third.go
	radius := b
	area := calculate_circle_area(float64(radius))
	fmt.Println("Area of Circle = ", area)
}
