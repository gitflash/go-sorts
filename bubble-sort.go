package sort

import "github.com/gitflash/go-helpers"

func Bubble(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr); i++ {
		isSwapped := false
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				helpers.Swap(&arr[j], &arr[j+1])
				isSwapped = true
			}
		}

		if !isSwapped {
			return arr
		}
	}

	return arr
}
