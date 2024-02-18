package solution

import "fmt"

type IVoiceable interface {
    Voice() string
}

type Cat struct {
    // … 
}

type Cow struct {
    // … 
}

type Dog struct {
	// … 
}

// BEGIN (write your solution here)
func (c Cat) Voice() string {
    return fmt.Sprint("Мяу")
}
func (c Cow) Voice() string {
    return fmt.Sprint("Мууу")
}
func (c Dog) Voice() string {
    return fmt.Sprint("Гав")
}
// END
