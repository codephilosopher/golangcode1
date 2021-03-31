package forloopexample

import "fmt"

//ForloopExample will display for loops
func ForloopExample() {
	x := 0
	for x < 10 {
		fmt.Println(x, "loop from comparison forloop")
		x++
	}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		fmt.Println(i, "loop from combination forloop")
	}
	fmt.Println("")
	for i := 0; i < 10; i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(j, "loop from inside")
		}
		fmt.Println("")
	}

	fmt.Println("")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if j == 5 {
				break
			}
			fmt.Println(j, "loop from inside")
		}
		fmt.Println("")
	}
	// continue skip the process, in this example skip the process while even number
	fmt.Println("")
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println(i)
	}

	fmt.Println("")
	for i := 33; i < 122; i++ {
		fmt.Printf("%v\t\t%#U\n", i, i)
	}

	fmt.Println("")
	switch {
	case (4 == 4):
		fmt.Println("both four are equal")
		fallthrough
	case (3 == 1):
		fmt.Println("both 3 and 1 are not equal")
	default:
		fmt.Println("this is default")
	}
}
