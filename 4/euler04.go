package main

import (
	"fmt"
	"strconv"
)

func main() {

	var largestPalindrome = 0
	for i := 0; i < 1000; i++ {
		for y := 0; y < 1000; y++ {
			var num = i * y
			if isPalindrome(num) && num > largestPalindrome {
				largestPalindrome = num
			}
		}
	}
	fmt.Println(largestPalindrome)
}

func isPalindrome(number int) bool {
	var stringNum = strconv.Itoa(number)
	if reverseString(stringNum) == stringNum {
		return true
	}
	return false
}

func reverseString(str string) string {
	byte_str := []rune(str)
	for i, j := 0, len(byte_str)-1; i < j; i, j = i+1, j-1 {
		byte_str[i], byte_str[j] = byte_str[j], byte_str[i]
	}
	return string(byte_str)
}
