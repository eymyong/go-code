package forloop

import "fmt"

// data   = [[1,2,3], [4,5,6]]
// output = [2,4,6]
func filterEven(data1 [][]int) []int {
	result := []int{}

	for _, v := range data1 {
		for _, vv := range v {
			if vv%2 == 0 {
				result = append(result, vv)
			}
		}
	}
	return result

	// for i := range data1 {
	// 	v := data1[i]
	// 	for ii := range v {
	// 		vv := v[ii]
	// 		if vv%2 == 0 {
	// 			box = append(result, vv)
	// 		}
	// 	}
	// }
	// return result
}

// data   = [[1,2,3], [4,5,6], [1,3,5], [10,20]]
// output = [[2], [4,6], [], [10,20]]
func filterEven2(data2 [][]int) [][]int {
	result := [][]int{}
	for _, v := range data2 {
		box := []int{}
		for _, vv := range v {
			if vv%2 == 0 {
				box = append(box, vv)
			}
		}
		result = append(result, box)
	}
	return result
}

func pattern1() {

	for i := 1; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Println(i)
		}
		fmt.Println(i)
	}
}

func pattern2() {

	for i := 0; i <= 3; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Println(i, j)
		}
		fmt.Println("finish loop j")
	}
}

func pattern3() {

	for i := 0; i <= 2; i++ {
		for j := 0; j <= 4; j++ {
			fmt.Printf("%d%d", i, j)
		}
		fmt.Println("finish loop j")
	}
}
func pattern4() {
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%d ", j)
		}
		fmt.Println("aa")
	}
}
