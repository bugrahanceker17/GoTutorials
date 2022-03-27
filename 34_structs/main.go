//package main
//
//import "fmt"
//
//func main() {
//	person1 := Employee
//	person1.name = "Bugrahan"
//	person1.isMarried = false
//	person1.age = 24
//
//	fmt.Println(person1)
//
//	person2 := Employee
//	person2.name = "Alptürk"
//	person2.isMarried = true
//	person2.age = 17
//
//	fmt.Println(person2)
//}
//
//var Employee struct {
//	name      string
//	age       int
//	isMarried bool
//}

package main

import "fmt"

type employee struct { // underlying type
	name      string
	age       int
	isMarried bool
	languages []string
}

func main() {

	var person1 employee
	person1.name = "Çınar"
	person1.age = 24
	person1.isMarried = false
	person1.languages = []string{"Typescript", "Php", "Ruby"}

	var person2 employee
	person2.name = "Cem"
	person2.age = 17
	person2.isMarried = true
	person2.languages = []string{"C#", "Golang", "Javascript"}

	person3 := employee{
		age:       14,
		name:      "Eylül Lilya",
		isMarried: false,
		languages: []string{
			"Phyton",
			"C++",
		},
	}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)
	fmt.Printf("%#v\n", person1)
	fmt.Printf("%#v\n", person2)
	fmt.Printf("%#v\n", person3)
}
