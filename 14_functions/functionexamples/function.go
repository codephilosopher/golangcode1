package functionexamples

import "fmt"

//FunctionExample return function examples
func FunctionExample() {
	// x := foo(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	x := foo("pradeek", arr...)
	fmt.Println("the total sum is, ", x)

	a1 := actors{
		person: person{
			firstName: "ashik",
			lastName:  "ajith",
			active:    true,
		},
		yearsOfActive: 3,
	}

	a2 := actors{
		person: person{
			firstName: "sunu",
			lastName:  "babu",
			active:    true,
		},
		yearsOfActive: 3,
	}

	// a1.speak()
	bar(a1)
	bar(a2)
	sfsdfsdf(a1)
	sfsdfsdf(a2)
}

func foo(s string, x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
		fmt.Println(v, sum)
	}
	return sum
}

type person struct {
	firstName string
	lastName  string
	active    bool
}

type actors struct {
	person
	yearsOfActive int
}

//speak() method becomes the part actors type
func (a actors) speak() {
	fmt.Println("I'm", a.firstName, a.lastName, "active for ", a.yearsOfActive, "years", "status: ", a.active)
}

func (a actors) leaker() {
	fmt.Println(a.firstName, a.lastName, a.yearsOfActive, a.active)
}

type german interface {
	leaker()
}

func sfsdfsdf(g german) {
	g.leaker()
}

//interface

type humans interface {
	speak()
}

func bar(h humans) {
	h.speak()
}
