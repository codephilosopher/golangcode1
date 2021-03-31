package jasonexample

import (
	"encoding/json"
	"fmt"
	"os"
)

type Profile struct {
	Name string
	Age  int
}

func Exercise3() {
	p1 := Profile{
		Name: "Ajay",
		Age:  34,
	}

	p2 := Profile{
		Name: "Vijay",
		Age:  24,
	}

	p3 := Profile{
		Name: "Sanjay",
		Age:  54,
	}

	profiles := []Profile{p1, p2, p3}
	fmt.Println(profiles)
	err := json.NewEncoder(os.Stdout).Encode(profiles)
	if err != nil {
		fmt.Println(err)
	}

}
