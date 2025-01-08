package main

import (
	"fmt"
)

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

func merge(arr []int, low int, mid int, high int) {
	tempArr := []int{}
	left := low
	right := mid + 1
	for left <= mid && right <= high {
		if arr[left] <= arr[right] {
			tempArr = append(tempArr, arr[left])
			left++
		} else {
			tempArr = append(tempArr, arr[right])
			right++
		}
	}

	for left <= mid {
		tempArr = append(tempArr, arr[left])
		left++
	}

	for right <= high {
		tempArr = append(tempArr, arr[right])
		right++
	}

	for i := low; i <= high; i++ {
		arr[i] = tempArr[i-low]
	}
}

func mergeSort(arr []int, low int, high int) {
	if low >= high {
		return
	}
	mid := (low + high) / 2
	mergeSort(arr, low, mid)
	mergeSort(arr, mid+1, high)
	merge(arr, low, mid, high)
}

func partition(arr []int, low int, high int) int {
	pivot := arr[low]
	i := low
	j := high

	for i < j {
		for arr[i] <= pivot && i <= high-1 {
			i++
		}
		for arr[j] > pivot && j >= low+1 {
			j--
		}
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	arr[low], arr[j] = arr[j], arr[low]
	return j
}

func quickSort(arr []int, low int, high int) {
	if low < high {
		pIndex := partition(arr, low, high)
		quickSort(arr, low, pIndex-1)
		quickSort(arr, pIndex+1, high)
	}
}

// Select and Swap
func main() {
	arr := []int{25, 12, 65, 34, 87, 1, 3, 23}
	// selectionSort(arr)
	// bubbleSort(arr)
	// insertionSort(arr)
	// mergeSort(arr, 0, len(arr)-1)
	quickSort(arr, 0, len(arr)-1)
	fmt.Print(arr)
}
