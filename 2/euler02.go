package main

import "fmt"

func main() {

	var sum = 0
	var fib = [3]int{1, 1, 2}
	for fib[2] < 4000000 {
		fmt.Print(fib, "\n")
		if fib[2]%2 == 0 {
			fmt.Printf("Adding %v to the mix \n", fib[2])
			sum += fib[2]
		}
		fib[0] = fib[1]
		fib[1] = fib[2]
		fib[2] = fib[0] + fib[1]

	}

	fmt.Print(sum, "\n")
}
