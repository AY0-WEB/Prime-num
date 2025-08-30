package main

import "fmt"

func CollectDetails() (bool, error) {
	var checkPrime int
	fmt.Print("Enter the number: ")
	if _, err := fmt.Scan(&checkPrime); err != nil {
		fmt.Println("pls enter an integer")
	}
	ok, err := PrimeNumber(checkPrime)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
