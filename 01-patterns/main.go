package main

import "fmt"

func printPattern1(n int){
    for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("*")
		}
        fmt.Println()
	}
}

func printPattern2(n int){
    for i := 0; i < n; i++ {
		for j := 0; j<=i; j++ {
            fmt.Print("*")
		}
        fmt.Println("")
	}
}

func printPattern3(n int){
    for i := 0; i < n; i++ {
		for j := 0; j<=i; j++ {
            fmt.Print(j+1)
		}
        fmt.Println("")
	}
}

func printPattern4(n int){
    for i := 0; i < n; i++ {
		for j := 0; j<=i; j++ {
            fmt.Print(i+1)
		}
        fmt.Println("")
	}
}

func printPattern5(n int){
    for i := 0; i <= n; i++ {
		for j := 0; j<n-i+1; j++ {
            fmt.Print("*")
		}
        fmt.Println("")
	}
}

func printPattern6(n int){
    for i := 0; i <= n; i++ {
		for j := 0; j<n-i+1; j++ {
            fmt.Print(j+1)
		}
        fmt.Println("")
	}
}

func printPattern7(n int){
    for i := 0; i < n; i++ {
		for j:=0; j<n-i-1;j++{
			fmt.Print(" ")
		}
		for j := 0; j<((i*2)+1); j++ {
            fmt.Print("*")
		}
        for j:=0; j<n-i-1;j++{
			fmt.Print(" ")
		}

		fmt.Println("")
	}
}

func main() {
	// printPattern1(5)
	// printPattern2(5)
	// printPattern3(5)
	// printPattern4(5)
	// printPattern5(5)
	// printPattern6(5)
	printPattern7(5)
}
