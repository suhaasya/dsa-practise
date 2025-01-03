package main

import "fmt"

func selectionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr)-1; j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		temp := arr[min]
		arr[min] = arr[i]
		arr[i] = temp
	}
	return arr
}

func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		starter := 0
		for j := starter + 1; j < len(arr)-1-i; j++ {
			if arr[starter] > arr[j] {
				temp := arr[starter]
				arr[starter] = arr[j]
				arr[j] = temp
				starter++
			}
		}

	}
	return arr
}

func insertionSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		j := i
		for j > 0 && arr[j-1] > arr[j] {
			temp := arr[j-1]
			arr[j-1] = arr[j]
			arr[j] = temp
			j--
		}

	}
	return arr
}

// Select and Swap
func main() {
	arr := []int{25, 12, 65, 34, 87, 1, 3, 23}
	// selectionSort(arr)
	// bubbleSort(arr)
	insertionSort(arr)
	fmt.Print(arr)
}
