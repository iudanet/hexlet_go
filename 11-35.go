package solution

import "strings"

// UserCreateRequest is a request to create a new user.
type UserCreateRequest struct {
	FirstName string
	Age       int
}

// BEGIN (write your solution here)
func Validate(req UserCreateRequest) string {
	if strings.Contains(req.FirstName, " ") || req.FirstName == "" {
		return "invalid request"
	}
	if req.Age <= 0 || req.Age > 150 {
		return "invalid request"
	}
	return ""
}

// END
