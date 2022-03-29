package main

import "fmt"

type mile float64
type kilometer float64

func main() {

	var m1 mile
	m1 = 3.4
	fmt.Println(m1)

	f1 := float64(4.5)
	fmt.Println(m1 + mile(f1))

	k1 := kilometer(8.9)
	fmt.Println(k1)
}
