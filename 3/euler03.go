package main

import "fmt"

func main() {

	var n = 600851475143
	var i = 2
	for i*i < n {
		fmt.Println(i * i)
		for n%i == 0 {

			n = n / i
			fmt.Println(n)
		}

		i = i + 1
	}
}
