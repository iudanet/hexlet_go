package main

import (
	"fmt"

	"github.com/iudanet/hexlet_go/solution"
)
func main() {

a := solution.Remove([]int{1, 2, 3}, 1)
fmt.Println(a)
b := solution.Remove([]int{1, 2, 3}, 0)
fmt.Println(b)
}
