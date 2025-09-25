package validator

import (
	"strconv"
)

func IsValidLuhn(numberOrder string) bool {
	num, err := strconv.Atoi(numberOrder)
	if err != nil {
		return false
	}

	if num < 1 {
		return false
	}
	sum := 0
	arr := intToArr(num)

	parity := len(arr) % 2

	for i, val := range arr {
		if i%2 == parity {
			val *= 2
			if val > 9 {
				val -= 9
			}
		}
		sum += val
	}
	if sum%10 == 0 {
		return true
	}
	return false
}

func intToArr(num int) []int {
	if num < 1 {
		return []int{}
	}
	var arr []int
	for num > 0 {
		arr = append(arr, num%10)
		num = num / 10
	}

	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}

	return arr
}
