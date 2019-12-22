package sort

func Merge(arr []int, left int, right int) []int {

	if right > left {
		// find middle
		middle := (right + left) / 2

		// left Merge
		arr = Merge(arr, left, middle)

		// right merge
		arr = Merge(arr, middle+1, right)

		// do merge

		arr = doMerge(arr, left, middle, right)

	}

	return arr
}

func doMerge(arr []int, left int, middle int, right int) []int {
	sizeLeft := middle - left + 1
	arrLeft := make([]int, sizeLeft)

	sizeRight := right - middle
	arrRight := make([]int, sizeRight)

	for i := 0; i < sizeLeft; i++ {
		arrLeft[i] = arr[i+left]
	}

	for j := 0; j < sizeRight; j++ {
		arrRight[j] = arr[j+middle+1]
	}

	i := 0
	j := 0
	k := left

	for i < sizeLeft && j < sizeRight {
		if arrLeft[i] > arrRight[j] {
			arr[k] = arrRight[j]
			j++
			k++
		} else {
			arr[k] = arrLeft[i]
			i++
			k++
		}
	}

	for i < sizeLeft {
		arr[k] = arrLeft[i]
		k++
		i++
	}

	for j < sizeRight {
		arr[k] = arrRight[j]
		k++
		j++
	}

	return arr
}
