package main

import "fmt"

// inputs := [][]int{
// 	{10, 20, 30},
// 	{-10, -20, -30},
// 	{0, -1, 2, -3},
// }

// for i, v := range inputs {
// 	fmt.Println("inputs", i+1)
// 	fmt.Println("value =", v)
// 	fmt.Println("min", v, "=", min(v))
// 	fmt.Println("min", v, "=", mid(v))
// 	fmt.Println("max", v, "=", max(v))

// }

//==============================================

//input := []int{100, 20, 30, 40, 50, 60, 200, 80}
//input := []int{}

func mid(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	var value int
	mid := len(input) / 2
	if len(input)%2 == 0 {
		mid = mid - 1
	}

	for i, v := range input {
		if i == mid {
			value = v
		}
	}
	return value
}

func min(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	min := input[0]
	for _, v := range input {
		if v <= min {
			min = v
		}
	}
	return min
}

func max(input []int) int {
	if len(input) == 0 {
		panic("array no Value")
	}

	max := input[0]
	for _, v := range input {
		if v >= max {
			max = v
		}
	}
	return max
}

func interest(amt float32, r float32) float32 {
	// x := amt * (r)
	// return amt + x                                // posture1 (r is float)
	// return amt * (1 + r)   // = (amt(1) + amt(r)) // posture2 (r is float)
	return amt + (amt * r / 100) // posture3 (r is %)
}

func interest2(amt float32, r float32) float32 {
	x := amt + (amt * r / 100)
	fmt.Println(x)
	return x
}

// 100, 10, 3
func interestY(amt float32, r float32, year int) float32 {
	for i := 0; i < year; i++ {
		amt = interest(amt, r)
		amt = amt + (amt * r / 100)
	}
	return amt
}

func loop() {
	var input string
	for {
		fmt.Scanf("%s", &input)
		fmt.Println("input = ", input)
		if input == "exit" {
			break
		}
	}
}

func indexInt() []int {
	//var box []int
	box := []int{}
	for i := 1; i <= 20; i++ {
		box = append(box, i)
	}
	return box
}

func loopcars() {
	cars := []string{"mazda", "toyota", "honda"}
	for i := 0; i <= len(cars); i++ {
		fmt.Println(cars[i])
	}
	for i, c := range cars {
		fmt.Println(c, cars[i])
	}

}

func sudkun5(limit int, sud int) {
	fmt.Println("sudkun", sud)
	fmt.Println("=====")
	for i := 1; i <= limit; i++ {
		fmt.Println(i, "*", sud, "=", i*sud)
	}
}

func average(input []int) int { //ค่าเฉลี่ย
	sum := 0
	for i := range input {
		sum = sum + input[i]
	}
	return sum / len(input)
}

func xbar(list []int) int {
	sumxbar := 0
	for i := range list {
		sumxbar = sumxbar + list[i]
	}
	// for i := range list {
	// 	i = average(list)
	// 	sumxbar = (sumxbar + i)
	// }
	return sumxbar / len(list)
}

func zixma(list []int) int {
	xi := 0
	sumzixma := 0
	for i := range list {
		xi = (list[i] - xbar(list)) * (list[i] - xbar(list))
		sumzixma = sumzixma + xi
	}
	return sumzixma
}

// cc := math.Sqrt(float64(bb))
// fmt.Printf("Result 1: %.1f", cc)

func Product(input []int) int { // รับ[]; * value in array[]int; return ผลรวมการคูณเป็น int
	kun := 1
	for i := range input {
		kun = kun * input[i]
	}
	return kun
}

func percentOf(part int, total int) float64 { //หา%
	// fmt.Println("===ผลสอบ===")
	// fmt.Println("คะแนนที่ได้ =", part, ": คะแนนเต็ม =", total)
	//fmt.Println("คุณสอบได้ =", (float64(part)*float64(100))/float64(total), "%")
	return (float64(part) * float64(100)) / float64(total)
}

func listPercentOf(points []int, total int) []float64 { //หา % in array[]
	result := []float64{}
	for i := range points {
		v := percentOf(points[i], total)
		result = append(result, float64(v))
		// result = append(result, percentOf(points[i], total))
	}
	return result

	// result := make([]float64, len(points))
	// for i, v := range points {
	// 	// // v == points[i]
	// 	result[i] = percentOf(v, total)
	// }
	// return result
}

//============ listPercentOf() Print in main =========================
// box1 := []int{5, 10, 2}
// n := 10
// aa := listPercentOf(box1, n)

// for i := 0; i < len(aa); i++ {
// 	fmt.Println("===ผลสอบ===")
// 	fmt.Println("คะแนนที่ได้ =", box1[i], ": คะแนนเต็ม =", n)
// 	//fmt.Println("คุณสอบได้ =", percentOf(box1[i], n), "%")
// 	fmt.Println("คุณสอบได้", aa[i], "%")
// 	if box1[i] == n {
// 		fmt.Println("เก่งมาก")
// 	} else if box1[i] <= 3 {
// 		fmt.Println("คุณสอบตก")
// 	}

// }

//==================================

//input1 := loop20
//[3 6 9 12 15 18 21 24 27 30]

func loop0() []int { // ไม่รับค่า,loop next ไปจนถึง i ที่กำหนด โดยกำหนดเงื่อนไข
	box := []int{}
	for i := 1; i <= 30; i++ {
		if i%3 == 0 {
			box = append(box, i)
		}
	}
	return box
}

func loop1(limit int) []int { // รับค่า limit int ,loop next to limit
	var box []int
	for i := 1; i <= limit; i++ {
		box = append(box, i)
	}
	return box
}

func loop2(input1 []int) []int { // รับ input1 []int ,loop next ,loop in input, ใช้ range(1)index,2)value)
	box := []int{}
	// for i := range input1 {
	// 	v := input1[i]
	// 	box = append(box, v)
	// }
	for _, v := range input1 {
		box = append(box, v)
	}
	return box
}

func loop3(input1 []int) []int { // รับ input1 []int, loop back
	box := []int{}
	for i := len(input1) - 1; i >= 0; i-- {
		v := input1[i]
		box = append(box, v)
	}
	return box
}

func loop4(input1 []int) []int { // รับ input1 []int ,loop next ,loop in input ถึง mid(input1)
	mid := len(input1) / 2
	if len(input1)%2 == 0 {
		mid = mid - 1
	}

	box := []int{}
	for i := range input1 {
		v := input1[i]
		box = append(box, v)
		if i == mid {
			break
		}
	}
	return box
}

func loop5(input1 []int) []int { // รับ input1 []int ,loop back ,loop in input ถึง mid(input1)
	mid := len(input1) / 2
	if len(input1)%2 == 0 {
		mid = mid - 1
	}

	box := []int{}
	for i := len(input1) - 1; i >= 0; i-- {
		v := input1[i]
		box = append(box, v)
		if i == mid+1 {
			break
		}
	}
	return box
}

func loop6(input1 []int, index int) []int { // รับ input1 []int, index int , loop input1 * index, return []
	box := []int{}
	for _, v := range input1 {
		vv := v * index
		box = append(box, vv)
	}
	return box
}

func loop7(input1 []int, index int) []int { // รับ input1 []int, index int , loop input1 * index, add if,return []
	box := []int{}
	for _, v := range input1 {
		v = v * index
		if v%4 == 0 {
			box = append(box, v)
		}
	}
	return box
}

/*
1. average find mean from array of int
	- input = [1,2,3] output = 2
2. find duplicate numbers from two arrays
	-  if input1 = [1,2,3] input2 = [2,3,7] this function will return [2,3]
3. find prime number that less than limit
	- func will return array of prime number that less than limit
	- input = 7 output =[2,3,5]

	[1,2,3,4]
	[5,2,9,1,8,32,246462,4]
	[]

*/
