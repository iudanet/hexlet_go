package solution

import "sort"

// BEGIN (write your solution here)
func UniqueSortedUserIDs(userIDs []int64) []int64 {
	elementMap := make(map[int64]int64)
	for _, s := range userIDs {
        elementMap[s] = s
        // or just keys, without values: elementMap[s] = ""
    }
	v := make([]int64, 0, len(elementMap))

	for  _, value := range elementMap {
	v = append(v, value)
	}
	sort.Slice(v, func(i, j int) bool {
		return v[i] < v[j]
	})
	return v
}

// END
