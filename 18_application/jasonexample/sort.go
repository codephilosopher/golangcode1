package jasonexample

import (
	"fmt"
	"sort"
)

type person struct {
	Name string
	Age  int
}

func (p person) String() string {
	return fmt.Sprintf("%s: %d", p.Name, p.Age)
}

type ByAge []person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age > a[j].Age }

//SortExample will return sort
func SortExample() {
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	fmt.Println("before sorting: ", s)
	sort.Strings(s)
	fmt.Println("after sorting: ", s)
	s1 := []int{5, 2, 6, 3, 1, 4} // unsorted
	fmt.Println("before sorting: ", s1)
	sort.Ints(s1)
	fmt.Println("after sorting: ", s1)

	//custom sorting
	people := []person{
		{"Bob", 31},
		{"John", 42},
		{"Michael", 17},
		{"Jenny", 26},
	}

	fmt.Println(people)
	// There are two ways to sort a slice. First, one can define
	// a set of methods for the slice type, as with ByAge, and
	// call sort.Sort. In this first example we use that technique.
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
