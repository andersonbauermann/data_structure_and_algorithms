package sorting

func sort(nums []int) []int {
	n := len(nums)
	for i := 0; i < n-1; i++ {
		isSorted := true
		for j := 0; j < n-i-1; j++ {
			if nums[j] > nums[j+1] {
				isSorted = false
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}

		if isSorted {
			break
		}
	}
	return nums
}

func PrintBubbleSort() {
	nums := []int{64, 34, 25, 12, 22, 11, 90}
	sorted := sort(nums)
	for _, num := range sorted {
		print(num, " ")
	}
	println()
}
