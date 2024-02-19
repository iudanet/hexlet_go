package solution

// BEGIN (write your solution here)
func Remove(nums []int, i int) []int {
	// fmt.Println(len(nums))
	// fmt.Println( len(nums) -1 > i)
	// fmt.Println(i < 0)

 if len(nums) -1 < i || i < 0 {
	return nums
 }
 if i == 0 {
	return nums[i+1:]
 }
 if i == len(nums) -1 {
	return nums[:i]
 }
 nums = append(nums[:i], nums[i+1:]...)
 return nums
}
// END
