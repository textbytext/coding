package main

import (
	"fmt"
	"log"
	"os"
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
	fmt.Println("doFor")

	aSlice := []int{-1, 2, 1, -1, 2, -2}
	for i, v := range aSlice {
		fmt.Println("index:", i, "value: ", v)
	}

	i := 0
	for ok := true; ok; ok = (i != 10) {
		fmt.Print(i*i, " ")
		i++
	}
}

func printArgs() {
	fmt.Println("printArgs")

	fmt.Print("app arguments: ")
	for _, v := range os.Args {
		fmt.Print("'", v, "'")
	}
}

func fatalPanic() {
	if len(os.Args) == 1 {
		log.Fatal("Fatal: the fatal message")
	}

	log.Panic("Panic: a panic message")
}

func main() {
	// doPrint()
	//doSwitch()

	//fmt.Println("")
	//doFor()

	//fmt.Println("")
	//printArgs()

	fmt.Println("")
	fatalPanic()
}
