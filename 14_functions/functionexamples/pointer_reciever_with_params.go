package functionexamples

import "fmt"

type shapes interface {
	area() int
	area1(int, int) int
}

type rectangle struct {
	width  int
	height int
}

func (r rectangle) area() int {
	return r.width * r.height
}

func (r *rectangle) area1(width int, height int) int {
	return r.width * r.height
}

func measurement(s shapes) { // helper function
	fmt.Println("area1", s.area())
}
func ReceiverImplementaionParams() {
	re := &rectangle{1, 2}
	measurement(re)
	fmt.Println("area2", re.area1(45, 67))
}
