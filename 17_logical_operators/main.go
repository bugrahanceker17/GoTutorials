package main

import "fmt"

func main() {
	//x, y := 5, 8
	//fmt.Printf("%v, %T\n", x == y, x == y)
	//fmt.Printf("%v, %T\n", x != y, x != y)
	//fmt.Printf("%v, %T\n", x <= y, x <= y)
	//fmt.Printf("%v, %T\n", x >= y, x >= y)
	//fmt.Printf("%v, %T\n", x > y, x > y)
	//fmt.Printf("%v, %T\n", x < y, x < y)

	x, y := 14, 20
	set1 := (x == y) // false
	set2 := (x < y)  // true

	fmt.Printf("%v\n", set1 && set2) // ikisi de true ise true döner
	fmt.Printf("%v\n", set1 || set2) // herhangi birisi true olsa true döner
	fmt.Printf("%v\n", !set1)        // herhangi birisi true olsa true döner

}
