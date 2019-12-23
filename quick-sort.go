package sort

import "github.com/gitflash/go-helpers"

func Quick(arr []int, low int, high int) []int {

	if high > low {
		// create partition
		pi := partition(arr, low, high)

		// left side
		arr = Quick(arr, low, pi-1)

		// right side
		arr = Quick(arr, pi+1, high)
	}

	return arr
}

func partition(arr []int, low int, high int) int {

	pivot := arr[high]
	i := low - 1
	for j := low; j <= high-1; j++ {
		if arr[j] < pivot {
			i++

			helpers.Swap(&arr[i], &arr[j])
		}
	}

	helpers.Swap(&arr[i+1], &arr[high])
	return i + 1
}
