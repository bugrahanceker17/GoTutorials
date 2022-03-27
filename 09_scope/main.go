package main

import "fmt"

var packageVar = "Package Scope"

func main() {

	if true {
		blockVar := "Block Scope"
		fmt.Println(blockVar)
	}

	var funcVar = "Function Scope"
	fmt.Println(funcVar)

	test()
}

func test() {
	packageVar = "Package Scope"
	fmt.Println(packageVar)
}
