package main

import "fmt"

func main() {
	//name := "bugrahan"
	//fmt.Println(name)
	//fmt.Println(&name) //  & -> address operator
	//
	//num := 17
	//fmt.Println(num)
	//fmt.Println(&num)

	//test := &num
	//fmt.Printf("%T, %v", test, test)

	x := 5
	fmt.Println(x)     // => 5
	fmt.Println(&x)    // => 5'in adresi
	fmt.Println(*(&x)) // => deferencing (5)
}
