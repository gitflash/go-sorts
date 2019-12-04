package sort

import "github.com/gitflash/go-helpers"

func Selection(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}

		helpers.Swap(&arr[i], &arr[min])
	}

	return arr
}
