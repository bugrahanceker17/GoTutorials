package main

import (
	"fmt"
)

func main() {
	const x = 5 // => compile time
	var y = 1   // => run time
	fmt.Println(x + y)

}
