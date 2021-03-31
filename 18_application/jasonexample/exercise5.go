package jasonexample

import (
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type byAge []User
type byName []User

func (a byAge) Len() int           { return len(a) }
func (a byAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a byName) Len() int           { return len(a) }
func (a byName) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byName) Less(i, j int) bool { return a[i].Name < a[j].Name }

func Exercise5() {
	user1 := User{
		Name: "ajay",
		Age:  43,
	}
	user2 := User{
		Name: "vijay",
		Age:  13,
	}
	user3 := User{
		Name: "sanjay",
		Age:  78,
	}

	users := []User{user1, user2, user3}
	fmt.Println(users)
	sort.Sort(byAge(users))
	fmt.Println("-----sort by age-----")
	fmt.Println(users)
	sort.Sort(byName(users))
	fmt.Println("----sort by Name----")
	fmt.Println(users)

}
