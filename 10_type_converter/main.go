package main

import "fmt"

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v  %T\n", x, x)
	fmt.Printf("%v  %T\n", y, y)

	// Type conversion =>  dönüştürmek istenilen tip||int(y)||parantex içine değeri yaz

	fmt.Println(x + int(y))
}
