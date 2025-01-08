package main

import "fmt"

func getLargest(arr []int) int {
	largest := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}
	return largest
}

func getSecondLargest(arr []int) int {
	largest := arr[0]
	secondLargest := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > largest {
			largest = arr[i]
		}
	}

	for i := 0; i < len(arr); i++ {
		if arr[i] != largest && arr[i] > secondLargest {
			secondLargest = arr[i]
		}
	}
	return secondLargest
}

func isSorted(arr []int) bool {
	isSort := true
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			isSort = false
		}
	}
	return isSort
}

func removeDuplicates(arr []int) []int {
	obj := map[int]int{}
	arr2 := []int{}
	for i := 0; i < len(arr)-1; i++ {
		if obj[arr[i]] == 0 {
			obj[arr[i]] = arr[i]
		}
	}

	for key := range obj {
		arr2 = append(arr2, key)
	}

	return arr2
}

func leftRotateByOne(arr []int) []int {
	first := arr[0]
	for i := 0; i < len(arr)-1; i++ {
		arr[i] = arr[i+1]
	}
	arr[len(arr)-1] = first

	return arr
}

func leftRotateByDPlaces(arr []int, d int) []int {
	dArr := []int{}

	for i := 0; i < d; i++ {
		dArr = append(dArr, arr[i])
	}

	for i := d; i < len(arr); i++ {
		arr[i-d] = arr[i]
	}

	for i := 0; i < d; i++ {
		arr[(len(arr)-d)+i] = dArr[i]
	}

	return arr
}

func rightRotateDPlaces(arr []int, d int) []int {
	dArr := []int{}

	for i := len(arr) - d; i < len(arr); i++ {
		dArr = append(dArr, arr[i])
	}

	for i := len(arr) - 1; i >= d; i-- {
		arr[i] = arr[i-d]
	}

	for i := 0; i < d; i++ {
		arr[i] = dArr[i]
	}

	return arr
}

func moveZeroToEnd(arr []int) []int {
	tempArr := []int{}

	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			tempArr = append(tempArr, arr[i])
		}
	}

	for i := 0; i < len(tempArr); i++ {
		arr[i] = tempArr[i]
	}

	for i := len(tempArr); i < len(arr); i++ {
		arr[i] = 0
	}

	return arr
}

func unionArr(arr []int, arr2 []int) []int {
	obj := map[int]int{}
	arr3 := []int{}

	for i := 0; i < len(arr); i++ {
		if obj[arr[i]] == 0 {
			obj[arr[i]] = 1
		}
	}

	for i := 0; i < len(arr2); i++ {
		if obj[arr2[i]] == 0 {
			obj[arr2[i]] = 1
		}
	}

	for key := range obj {
		arr3 = append(arr3, key)
	}

	return arr3
}

func intersectionArr(arr []int, arr2 []int) []int {
	obj := map[int]int{}
	arr3 := []int{}

	for i := 0; i < len(arr); i++ {
		if obj[arr[i]] == 0 {
			obj[arr[i]] = 1
		}
	}

	for i := 0; i < len(arr2); i++ {
		if obj[arr2[i]] != 0 {
			arr3 = append(arr3, arr2[i])
		}
	}

	return arr3
}

func main() {
	arr := []int{1, 3, 0, 2, 4, 5, 0, 1, 6, 7, 0, 2, 12, 0, 4, 10, 5, 3, 9, 54}
	arr2 := []int{12, 0, 4, 10, 5, 3, 9, 54, 800, 542}
	fmt.Println(arr)
	// fmt.Println(getLargest(arr))
	// fmt.Println(getSecondLargest(arr))
	// fmt.Println(isSorted(arr))
	// fmt.Println(removeDuplicates(arr))
	// fmt.Println(leftRotateByOne(arr))
	// fmt.Println(leftRotateByDPlaces(arr, 4))
	// fmt.Println(rightRotateDPlaces(arr, 4))
	// fmt.Println(moveZeroToEnd(arr))
	fmt.Println(unionArr(arr, arr2))
	fmt.Println(intersectionArr(arr, arr2))
}
