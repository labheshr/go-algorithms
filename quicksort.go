package main

import "fmt"

func partition(arr[] int) int {

	start, end := 0, len(arr) - 1
	pIndex, pivot := start, arr[end]

	for i := start ; i < end; i++ {
		if arr[i] <= pivot {
			arr[i], arr[pIndex] = arr[pIndex], arr[i]
			pIndex++
		}
	}

	arr[pIndex], arr[end] = arr[end], arr[pIndex]

	return pIndex
}

func quickSort(arr[] int) []int {

	start, end := 0, len(arr) - 1
	if start < end {
		idx := partition(arr[start:])
		quickSort(arr[:idx])
		quickSort(arr[idx+1:])
	}
	return arr
}

func main() {

	test := [...]int{7,2,1,6,8,5,3,4}
	out := quickSort(test[:])
	fmt.Println(out)
}