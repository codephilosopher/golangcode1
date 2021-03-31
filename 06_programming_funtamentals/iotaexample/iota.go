package iotaexample

import "fmt"

//IotaExample gives some return of iota increment
func IotaExample() {
	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB
		PB
		EB
		ZB
		YB
	)

	fmt.Printf("%b\n", KB)
	fmt.Printf("%b\n", MB)
	fmt.Printf("%b\n", GB)
}
