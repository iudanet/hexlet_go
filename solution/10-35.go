package solution

import "strings"

// BEGIN (write your solution here)
func  ModifySpaces(s, mode string) string {
	switch {
	default:
		return strings.ReplaceAll(s, " ", "*")
	case mode  == "dash" :
		return strings.ReplaceAll(s, " ", "-")
	case mode == "underscore" :
		return strings.ReplaceAll(s, " ", "_")

	}
}
// END
