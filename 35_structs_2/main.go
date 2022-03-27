package main

import "fmt"

type employee struct {
	name      string
	age       int
	isMarried bool
}

type manager struct {
	employee
	hasDegree bool
}

func main() {
	p1 := employee{
		name:      "Bugrahan",
		age:       24,
		isMarried: false,
	}

	fmt.Println(p1)

	mngr := manager{
		employee: employee{
			name:      "Batuhan",
			age:       20,
			isMarried: false,
		},
		hasDegree: true,
	}

	fmt.Println(mngr)

	// anonim struct
	theBoss := struct {
		name   string
		isRich bool
	}{name: "THE BOSS", isRich: true}

	fmt.Println(theBoss)
}
