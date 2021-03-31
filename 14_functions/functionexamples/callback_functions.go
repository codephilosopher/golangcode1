package functionexamples

import "fmt"

//CallbackFunctions will returns callback functions
func CallbackFunctions() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	even := even(sum, arr...)
	fmt.Println(even)
}

func sum(i ...int) int {
	sum := 0
	for _, v := range i {
		sum += v
	}
	return sum
}

func even(f func(i ...int) int, i ...int) int {
	var eveArray []int
	for _, v := range i {
		if v%2 == 0 {
			eveArray = append(eveArray, v)
		}
	}
	return f(eveArray...)
}
