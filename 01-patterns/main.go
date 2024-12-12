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

func printPattern8(n int){
    for i := 0; i < n; i++ {
		for j:=0; j<i;j++{
			fmt.Print(" ")
		}
		for j := 0; j<((n*2)-1-(i*2)); j++ {
            fmt.Print("*")
		}
        for j:=0; j<i;j++{
			fmt.Print(" ")
		}

		fmt.Println("")
	}
}

func printPattern9(n int){
	for i:=0; i< (2*n -1); i++{
		if (i<=n){
			for j:=0;j<i+1;j++{
				fmt.Print("*")
			}
		}else{
			for j:=0;j<(2*n-i);j++{
				fmt.Print("*")
			}
		}
		fmt.Println("")
	}
}

func printPattern10(n int){
	for i:=1;i<=n;i++{
		for j:=1; j<=i;j++{
			if(j%2==0){
				fmt.Print(0)
			}else{
				fmt.Print(1)
			}
		}
		fmt.Println("")
	}
}

func printPattern11(n int){
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Print(i+1)
		}

		for j:=0;j<(2*n-2*(i+1));j++{
			fmt.Print(" ")
		}

		for j:=0;j<=i;j++{
			fmt.Print(i+1)
		}
		fmt.Println("")
	}
}

func printPattern13(n int){
	p:=0;

	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			p= p+1
			fmt.Print(p," ")
		}
		fmt.Println("")
	}
}

func printPattern14(n int){
	for i:=1;i<=n;i++{
		for j:=0;j<i;j++{
			fmt.Printf("%c",rune(65+j))
		}
		fmt.Println("")
	}
}

func printPattern15(n int){
	for i:=0;i<=n;i++{
		for j:=0;j<(n-i);j++{
			fmt.Printf("%c",rune(65+j))
		}
		fmt.Println("")
	}
}

func printPattern16(n int){
	for i:=0;i<n;i++{
		for j:=0;j<=i;j++{
			fmt.Printf("%c",rune(65+i))
		}
		fmt.Println("")
	}
}

func printPattern17(n int){
	for i:=1;i<=n;i++{

		for j:=0;j<n-i;j++{
			fmt.Print(" ")
		}

		a:= 64
		breakP:= (2*i+1)/2
		for j:=1;j<=(2*i)-1;j++{
			if(j<=breakP){
				a=a+1
			}else{
				a=a-1
			}
			fmt.Printf("%c",rune(a))
		}

		for j:=0;j<n-i;j++{
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
	// printPattern8(5)
	// printPattern9(5)
	// printPattern10(5)
	// printPattern11(5)
	// printPattern13(5)
	// printPattern14(5)
	// printPattern15(5)
	// printPattern16(5)
	printPattern17(5)
}