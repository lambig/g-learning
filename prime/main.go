package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var primes []int
	Individable := func(divident int, divisors []int) bool {
		for i := 0; i < len(primes) && primes[i] <= divident/2; i++ {
			if divident%divisors[i] == 0 {
				return false
			}
		}
		return true
	}
	ToStringSlice := func(intSlice []int) []string {
		stringSlice := make([]string, len(intSlice))
		for i := 0; i < len(intSlice); i++ {
			stringSlice[i] = strconv.Itoa(intSlice[i])
		}
		return stringSlice
	}

	var max int
	fmt.Scan(&max)
	for i := 2; i <= max; i++ {
		if Individable(i, primes) {
			primes = append(primes, i)
		}
	}
	stringPrimes := ToStringSlice(primes)

	fmt.Println(strconv.Itoa(max) + "までの素数は " + strings.Join(stringPrimes, ", "))
	fmt.Println("個数は " + strconv.Itoa(len(primes)))
}
