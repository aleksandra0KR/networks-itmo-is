package main

import (
	"fmt"
	"log"

	"github.com/aleksandra0KR/networks-itmo-is/internal/calculate"
)

type NetworkConfiguration struct {
	ip       string
	subnet   string
	networks []int
}

func main() {
	variants := []NetworkConfiguration{
		{
			"194.85.32.19",
			"255.255.255.0",
			[]int{10, 6, 1, 18, 100},
		},
		{
			"10.12.12.15",
			"255.255.254.0",
			[]int{25, 16, 240, 117, 1},
		},
		{
			"212.24.15.199",
			"255.255.255.192",
			[]int{7, 0, 0, 11, 10},
		},
		{
			"120.13.120.120",
			"255.255.255.224",
			[]int{5, 2, 2, 1, 1},
		},
	}

	for i, variant := range variants {
		data, err := calculate.NewData(variant.ip, variant.subnet, variant.networks)
		if err != nil {
			log.Fatalf("Error processing variant %d: %v", i+1, err)
		}

		result := calculate.Calculate(data)
		fmt.Printf("Configuration %d:\n%v\n", i+1, result)
	}
}
