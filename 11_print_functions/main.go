package main

import (
	"fmt"
	"strconv"
)

func main() {
	name := "Bugrahan"
	fmt.Printf("Benim adım %v", name) // "Benim adım Bugrahan
	fmt.Println("")
	fmt.Printf("%T", name)
	fmt.Println("")
	x := 80
	y := strconv.Itoa(x)
	fmt.Println(y)
}
