package main

func main() {
	nums := []int{1, 2, 3}
	nums = append(nums, 4)
	nums = append(nums, 5)

	nums2 := []int{1, 2, 3}
	nums2 = append(nums2, 4, 5)

	nums3 := []int{1, 2, 3}
	nums3 = append(nums3, nums2...)

}
