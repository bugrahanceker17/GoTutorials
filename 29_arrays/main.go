package main

import "fmt"

func main() {
	//ARRAY -> Composite value type

	//cities := [4]string{"Bursa", "Ankara", "Osmaniye", "İstanbul"}
	//fmt.Println(cities)
	//fmt.Println(len(cities))
	//------------------------------------------------------------------------------------------------------------------
	//var myArr [5]int
	//fmt.Println(myArr) // -> [0 0 0 0 0]
	//------------------------------------------------------------------------------------------------------------------
	//cities := [4]string{"Bursa", "Ankara", "Osmaniye", "İstanbul"}
	//for i := 0; i < len(cities); i++ {
	//	fmt.Println(cities[i])
	//}

	//1'DEN 10'A KADAR OLAN SAYILARIN KARESİNİ ALMA----------------------------------------------------------
	//numbers := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//numbers = getSquare(numbers) // First class function
	//fmt.Println(numbers)

	//------------------------------------------------------------------------------------------------------------
	//FOREACH
	cities := [4]string{"Bursa", "Ankara", "Osmaniye", "İstanbul"}
	for index, city := range cities {
		fmt.Println(index, city)
	}

	fmt.Println()

	for _, city := range cities { // => blank identifier
		fmt.Println(city)
	}
}

func getSquare(numbers [10]int) [10]int {

	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] * numbers[i]
	}

	return numbers
}
