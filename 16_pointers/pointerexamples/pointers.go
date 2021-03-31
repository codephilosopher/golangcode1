package pointerexamples

import "fmt"

//PointerExamples returns pointer examples
func PointerExamples() {
	a := 4
	b := &a
	fmt.Println(*b)
	var c *int
	c = &a
	fmt.Println(*c)
	fmt.Println(*&a)
}
