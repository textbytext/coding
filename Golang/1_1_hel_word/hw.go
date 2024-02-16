package main

import (
	"fmt"
)

const PrintGlobal = 12

func doPrint() {
	var str1 = "str1"
	fmt.Println(str1)

	str2 := "str2"
	fmt.Println(str2)

	fmt.Println("global: ", PrintGlobal, PrintGlobal)
}

func doSwitch() {
	v1 := 12
	v2 := "123"

	switch {
	case v1 == 12:
		fmt.Println(12)
		// fallthrough

	case v2 == "123":
		fmt.Println(v2)
	}
}

func doFor() {
	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}
}

func main() {
	// doPrint()
	//doSwitch()
	doFor()
}
