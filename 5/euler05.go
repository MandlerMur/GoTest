package main

import (
	"fmt"
)

func main() {
	var keepGoing = true
	for i := 1; keepGoing; i++ {
		for y := 2; i%y == 0; y++ {
			if y >= 20 {
				fmt.Println(i)
				keepGoing = false
			}
		}
	}
}
