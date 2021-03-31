package functionexamples

import (
	"fmt"
)

type Foo interface {
	Fizz()
}

type Bar struct{}

func (b *Bar) Fizz() { // either you pass the pointer or use b.method()
	fmt.Println("fizz")
}

func Fizzy(foo Foo) { // pass the pointer here
	foo.Fizz()
}

func ReceiverImplementaion() {
	b := Bar{}
	//Fizzy(b)
	b.Fizz()
}
