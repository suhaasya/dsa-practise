package main

import "fmt"

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

func main() {
	// print(5, "suhas")
	// printNNumbers(12)
	// printNNumbersRev(12)
	printNNumbersRev2(12)
}
