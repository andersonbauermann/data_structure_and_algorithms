package sorting

func quick_sort(nums []int, left int, right int) []int {
	if left < right {
		pi := partition(nums, left, right)
		quick_sort(nums, left, pi-1)
		quick_sort(nums, pi+1, right)
	}

	return nums
}

func partition(nums []int, left int, right int) int {
	pivot := nums[right]
	i := left - 1

	for j := left; j < right; j++ {
		if nums[j] <= pivot {
			i++
			nums[i], nums[j] = nums[j], nums[i]
		}
	}

	nums[i+1], nums[right] = nums[right], nums[i+1]

	return i + 1
}

func PrintQuickSort() {
	nums := []int{77, 5, 70, 23, 27, 13, 90}
	sorted := quick_sort(nums, 0, len(nums)-1)
	for _, num := range sorted {
		print(num, " ")
	}
	println()
}
