package main

import "fmt"

func main() {

	grade := 3

	switch grade {
	case 5:
		fmt.Println("AA")
	case 4:
		fmt.Println("BB")
	case 3:
		fmt.Println("CC")
	case 2:
		fmt.Println("DD")
	case 1:
		fmt.Println("FF")
	}
}
