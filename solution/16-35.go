package solution

// BEGIN (write your solution here)
func Map(strs []string, mapFunc func(s string) string) []string {
	slise := strs
	var r []string
	for _, value := range slise {
		r = append(r, mapFunc(value))
	}
	return r
}

// END
