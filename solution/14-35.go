package solution

// BEGIN (write your solution here)
func  SafeWrite(nums [5]int, i, val int) [5]int {
	if i > len(nums) - 1 || i < 0 {
		return nums
	}
	nums[i] = val
	return nums
}
// END
