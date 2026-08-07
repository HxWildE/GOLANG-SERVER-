package main

import (
	"fmt"
)

func runMaths() {
	fmt.Println("starting math")

	// var k int  = 5
	// var p int32 = 6
	// fmt.Println(k + p) :unmathced types cant have a common operand

	var nm string = "Ss"
	kk := 4
	fmt.Print(nm)
	fmt.Println(kk)
	//nm + kk dont work
	str := "check"
	fmt.Println(str + nm)

}

func runMod(a, b int) int {
	return a % b
}
