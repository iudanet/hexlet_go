package solution

// BEGIN (write your solution here)

func IntsCopy(src []int, maxLen int) []int {
	
	if maxLen == 0 || maxLen < 0 {
		return []int{}
	}
	if maxLen > len(src) {
		copySlice := make([]int, len(src))
		copy(copySlice, src)
		return copySlice
	}
	c := make([]int, maxLen)
	copy( c, src)
	return c
}
// END
