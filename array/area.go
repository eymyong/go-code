package array

import "fmt"

func TriA(h float64, b float64) float64 {
	return (h * b) * 0.5
}

func TriA2(h float64, b float64) {
	x := (h * b) * 0.5
	fmt.Println(x)
}

func CircleA(r float64) float64 {
	return (r * r) * 3.14
}

func CircleA2(r float64) {
	x := (r * r) * 3.14
	fmt.Println(x)
}

func PrmA(b float32, h float32) float32 {
	return (b * b) * h * 1 / 3
}

func PrmA2(b float32, h float32) {
	x := (b * b) * h * 1 / 3
	fmt.Println(x)
}
