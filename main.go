package main

import "fmt"

// BinarySearch 二分查找
func BinarySearch(arr []int, value int) bool {
	if arr == nil {
		return false
	}
	if len(arr) == 0 {
		return false
	}
	left := 0
	right := len(arr)
	found := false

	// 查找 [left, right) 范围内是否存在 value
	// 画个图可以看出来，如果 [x, x, x ..., left, right, x, x, x ...] 这种情况的话，left 和 right 相连，就可以直接判断为不存在了
	for (left + 1) < right {
		mid := (left + right) / 2
		fmt.Println(left, right, mid)
		if arr[mid] > value {
			right = mid
		} else if arr[mid] < value {
			left = mid
		} else {
			found = true
			break
		}
	}
	return found
}

func main() {
	arr := BinarySearch([]int{1, 4, 7, 11, 19, 27, 33, 34, 45, 61, 87, 125, 174, 258, 369}, 35)
	fmt.Println(arr)
}
