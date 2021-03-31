package structexample

import "fmt"

//StructExample returns StructExample
func StructExample() {

	type Address struct {
		address string
	}
	type Person struct {
		Address
		name string
		age  int
	}

	person := Person{
		//Embeded
		Address: Address{
			address: "some address",
		},
		name: "pradeek",
		age:  60,
	}

	//Anonymous
	car := struct {
		manufacture string
		model       string
	}{
		manufacture: "Porche",
		model:       "twin turbo carbulet",
	}

	fmt.Println(person)
	fmt.Println(person.name, person.age, person.address)
	fmt.Println(car)

	type Vehicle struct {
		doors string
		color string
	}

	type Truck struct {
		Vehicle
		fourWheel bool
	}

	type Sedan struct {
		Vehicle
		luxury bool
	}

	truck := Truck{
		Vehicle: Vehicle{
			doors: "2",
			color: "#95699",
		},
		fourWheel: true,
	}

	sedan := Sedan{
		Vehicle: Vehicle{
			doors: "4",
			color: "#45645",
		},
		luxury: false,
	}

	fmt.Println(truck)
	fmt.Println(sedan)

}
