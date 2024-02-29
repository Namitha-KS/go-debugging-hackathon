package main

// find average given function parameters
// no return statements allowed
func average(a, b, c int) int {
	average := (a + b + c) / 3
	return average
}

// find min and maximum from the given function parameters
//
func min_max(a, b, c int) (int, int) {
	values := []int{a, b, c}
	min := values[0]
	max := values[0]
	for _, value := range values {
		if value < min {
			min = value
		}
		if value > max {
			max = value
		}
	}
	return min, max
}
