package arrayexample

import "fmt"

// ArrayExample will return array programs
func ArrayExample() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[1:]
	c := a[3:6]
	fmt.Println(b)
	for i := 0; i < len(c); i++ {
		fmt.Println(i, c[i])
	}
	byteExample := []byte{}
	stringExample := "hello"
	newSlice := append(byteExample, stringExample...)
	for _, v := range newSlice {
		fmt.Printf("%c", v)
	}
	fmt.Println("")
	// delete elements
	newDeletedSlice := append(a[:2], a[4:]...)
	fmt.Println(newDeletedSlice)

	//make
	makeArray := make([]int, 10, 12)
	fmt.Println(makeArray)
	fmt.Println(len(makeArray))
	makeArray = append(makeArray, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(makeArray)
	fmt.Println(len(makeArray))
	fmt.Println(cap(makeArray))
	makeArray = append(makeArray, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	fmt.Println(makeArray)
	fmt.Println(len(makeArray))
	fmt.Println(cap(makeArray))
	// multidemnsion array

	array1 := []int{1, 2, 3, 4, 5, 6}
	array2 := []int{7, 8, 9}
	newArraySlice := [][]int{array1, array2}
	fmt.Println(newArraySlice)

	//Map

	mapSlice := map[string]int{
		"pradeek": 67,
		"afsal":   23,
		"raneesh": 19,
		"vipin":   45,
	}
	fmt.Println(mapSlice)
	fmt.Println(mapSlice["afsal"])

	if v, ok := mapSlice["raneesh"]; ok {
		fmt.Println(v)
	}
	// loop over map

	for k, v := range mapSlice {
		fmt.Println(k, v)
	}

	// delete
	delete(mapSlice, "vipin")
	fmt.Println(mapSlice)

}
