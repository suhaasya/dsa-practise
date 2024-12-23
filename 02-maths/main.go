package main

import (
	"fmt"
	"math"
	"slices"
)

func count(num int) int {
	count := 0
	n := num

	for n > 0 {
		n = n / 10
		count++
	}

	return count
}

func reverse(count int) int {
	revNum := 0
	for count > 0 {
		rem := count % 10
		revNum = revNum*10 + rem
		count = count / 10
	}
	return revNum
}

func palindrome(num int) bool {
	revNum := reverse(num)

	return revNum == num
}

func armstrongNum(num int) bool {
	armNum := 0
	n := num
	for n > 0 {
		rem := n % 10
		armNum = armNum + rem*rem*rem
		n = n / 10
	}

	return armNum == num
}

func printAllDivisions(num int) {
	var values = []int{}
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			values = append(values, i)
			if i != num/i {
				values = append(values, num/i)
			}
		}
	}

	slices.Sort(values)
	for _, val := range values {
		fmt.Print(val)
		fmt.Print(" ")
	}

}

func isPrime(num int) bool {
	counter := 0
	for i := 1; i <= int(math.Sqrt(float64(num))); i++ {
		if num%i == 0 {
			counter++
			if (num / i) != i {
				counter++
			}
		}
	}
	fmt.Println(counter)
	return counter == 2
}

func findGCD(num1 int, num2 int) int {
	gcd := 1
	for i := 2; i <= int(math.Min(float64(num1), float64(num2))); i++ {
		if num1%i == 0 && num2%i == 0 {
			gcd = i
		}
	}
	return gcd
}

func main() {
	// fmt.Println(count(123))
	// fmt.Println(reverse(123))
	// fmt.Println(palindrome(123))
	// fmt.Println(palindrome(121))
	// fmt.Println(armstrongNum(371))
	// fmt.Println(armstrongNum(123))
	// printAllDivisions(345)
	// printAllDivisions(36)
	// fmt.Println(isPrime(11))
	// fmt.Println(isPrime(36))
	// fmt.Println(isPrime(1))
	fmt.Println(findGCD(12, 16))
	fmt.Println(findGCD(12, 100))
}
