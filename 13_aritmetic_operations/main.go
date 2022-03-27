package main

import "fmt"

func main() {
	//x, y := 15, 5 // 15, 5 => operand & + => operator

	//fmt.Printf("%T, %v\n", x, x)
	//fmt.Printf("%T, %v\n", y, y)

	//fmt.Printf("%T, %v\n", x+y, x+y)
	//fmt.Printf("%T, %v\n", x-y, x-y)
	//fmt.Printf("%T, %v\n", x/y, x/y) // => result : 3
	//fmt.Printf("%T, %v\n", x/y, x/y) // => result :

	//z := 5.0 / 2 // => Float sonuç için operand'ı float şeklinde yaz
	//fmt.Printf("%T, %v\n", z, z)

	//INCREMENT AND DECREMENT------------------------------------------------------

	x := 10
	fmt.Println(x + 1)
	x += 2
	fmt.Println(x)
	x += 1
	fmt.Println(x)
}
