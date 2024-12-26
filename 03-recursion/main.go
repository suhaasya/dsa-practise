package main

import (
	"fmt"
)

func print(times int, name string) {
	if times == 0 {
		return
	}
	fmt.Println(name)
	print(times-1, name)
}

var counter = 0

func printNNumbers(n int) {
	if n == 0 {
		return
	}
	counter++
	fmt.Println(counter)
	printNNumbers((n - 1))
}

func printNNumbersRev(n int) {
	if n == 0 {
		return
	}
	fmt.Println(n)
	printNNumbersRev((n - 1))
}

func printNNumbersRev2(n int) {
	if n == 0 {
		return
	}
	printNNumbersRev2((n - 1))
	fmt.Println(n)
}

func printSumNNumberss(n int) int {
	if n == 1 {
		return 1
	}

	return n + printSumNNumberss(n-1)
}

func factorial(n int) int {
	if n == 1 {
		return 1
	}
	return n * factorial(n-1)
}

func isPalindrome(i int, str string) bool {
	if i >= len(str) {
		return true
	}
	if str[i] != str[len(str)-1-i] {
		return false
	}
	return isPalindrome(i+1, str)
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	// print(5, "suhas")
	// printNNumbers(12)
	// printNNumbersRev(12)
	// printNNumbersRev2(12)
	// n := printSumNNumberss(3)
	// strings.len
	// fmt.Println(factorial(4))
	// fmt.Println(isPalindrome(0, "madam"))
	// fmt.Println(isPalindrome(0, "madman"))
	fmt.Println(fibonacci(10))
}
