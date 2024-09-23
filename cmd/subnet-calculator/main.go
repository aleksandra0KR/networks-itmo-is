package main

import "github.com/aleksandra0KR/networks-itmo-is/internal/calculate"

func main() {
	ip := "194.85.32.19"
	subnet := "255.255.255.0"
	networks := []int{10, 6, 1, 18, 10}

	calculate.Calculate()
}
