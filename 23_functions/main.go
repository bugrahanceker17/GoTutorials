package main

import "fmt"

// func functionName (params) return value { code }

func main() {

	fmt.Println(sum(4, 8))

	sum := sum(5, 99)
	fmt.Println(sum)
	sayHi()
}

func sum(x, y int) int {
	return x + y
}

func sayHi() {
	fmt.Println("hello there!")
}
