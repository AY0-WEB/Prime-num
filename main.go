package main

import (
	"fmt"
)

func collectDetails() {
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
func main() {
	if ok {
		fmt.Printf("\n%d is a prime number\n", checkPrime)
	}
	if !ok {
		fmt.Printf("\n%d is not a prime number\n", checkPrime)
	}

}
