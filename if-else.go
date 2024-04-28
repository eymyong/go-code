package main

import "fmt"

func buy(a int) {
	if a > 50 {
		fmt.Println("ซื้อได้")
	} else {
		fmt.Println("ซื้อไม่ได้")
	}
}

func grade(a int) {
	if a >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else if a >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if a >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if a >= 50 {
		fmt.Println("คุณได้เกรด D")
	} else {
		fmt.Println("คุณตก")
	}

}

func grade1() {
	var score int
	fmt.Println("คะแนน")
	fmt.Scanf("%d", &score)
	fmt.Println("score = ", score)
	if score >= 80 {
		fmt.Println("คุณได้เกรด A")
	} else if score >= 70 {
		fmt.Println("คุณได้เกรด B")
	} else if score >= 60 {
		fmt.Println("คุณได้เกรด C")
	} else if score >= 50 {
		fmt.Println("คุณได้เกรด D")
	} else {
		fmt.Println("คุณตก")
	}

}
