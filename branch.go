package main

import "errors"

func PrimeNumber(x int) (bool, error) {
	if x < 2 {
		return false, errors.New("Pls enter a number greater than 1")
	} else {
		for i := 2; i*i <= x; i++ {
			if x%i == 0 {
				return false, nil
			}
		}
	}
	return true, nil
}
