package search

// LinearSearch 按序查找
func LinearSearch(xs []int, target int) bool {
	for _, v := range xs {
		if v == target {
			return true
		}
	}
	return false
}

// BinarySearch 二分查找
func BinarySearch(xs []int, target int) bool {
	lo, hi := 0, len(xs)
	for lo < hi {
		mid := (lo + hi) / 2
		if xs[mid] == target {
			return true
		} else if xs[mid] > target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return false
}
