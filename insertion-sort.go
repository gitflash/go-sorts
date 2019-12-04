package sort

import (
	"github.com/gitflash/go-helpers"
)

func Insertion(arr []int) []int {

	if len(arr) == 0 {
		return arr
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j > 0; j-- {
			if arr[j] < arr[j-1] {
				helpers.Swap(&arr[j], &arr[j-1])
			}
		}
	}

	return arr
}
