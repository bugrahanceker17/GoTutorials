package main

import "fmt"

func main() {

	//var name string = "Bugrahan" //var keyword -- name of variable -- static type
	//fmt.Println("Hello", name)
	//
	//var age uint8 // variable declaration
	//age = 24      // variable assign
	//
	//var isMarried bool
	//isMarried = false
	//
	//var firstName, lastName string = "Cigdem", "Selvi"
	//fmt.Println(firstName, lastName)
	//fmt.Println(age)
	//fmt.Println(isMarried)
	//
	//lang := "Go"
	//fmt.Println(lang)
	//
	//fmt.Printf("%T\n", name)
	//fmt.Printf("%T\n", age)
	//fmt.Printf("%T\n", isMarried)
	//fmt.Printf("%T\n", lang)

	//---------------------------------------------------------

	//var (
	//	name     string = "Go"
	//	isUseful bool   = true
	//)
	//
	//fmt.Println(name)
	//fmt.Println(isUseful)

	//---------------------------------------------------------

	//var age, name, sort = 24, "Batuhan", 2
	//
	//fmt.Println(age)
	//fmt.Println(name)
	//fmt.Println(sort)

	//---------------------------------------------------------

	age, name, sort := 24, "Batuhan", 2
	fmt.Println(age)
	fmt.Println(name)
	fmt.Println(sort)

	//---------------------------------------------------------

	var isExists bool
	fmt.Println(isExists) // Zero value => false

	var lastName string
	fmt.Println(lastName) // Zero value => ""

	var number int
	fmt.Println(number) // Zero value => 0
}
