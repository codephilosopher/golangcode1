package stringsexample

import "fmt"

//StringsExample is a function to inspect value as byte, binary,ocatal
func StringsExample() {
	helloMessages := "Hello How are you"
	byteStrings := []byte(helloMessages)
	for _, v := range byteStrings {
		fmt.Printf("%d ", v)
	}
	fmt.Println("")
	for _, v := range byteStrings {
		fmt.Printf("%#U ", v)
	}
	fmt.Println("")
	for _, v := range byteStrings {
		fmt.Printf("%#b ", v)
	}
}
