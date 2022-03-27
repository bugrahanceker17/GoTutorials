package main

import "fmt"

func main() {
	//myMap := map[string]int{ // LITERAL
	//	"Ahmet": 11,
	//	"Vali":  17,
	//	"Sude":  22,
	//	"Can":   19,
	//}
	//
	//fmt.Println(myMap)
	//fmt.Println(myMap["Sude"])

	//---------------------------------------------------------------------------------------------
	//demo := make(map[string]int)
	//fmt.Println(demo)

	//---------------------------------------------------------------------------------------------
	studentGrades := map[string]int{ // => maps : pass by reference
		"Çiğdem":   96,
		"Buğrahan": 80,
		"Furkan":   58,
		"Zeynep":   77,
	}
	//fmt.Println(studentGrades["Çiğdem"])

	//BİR DEĞERİN MAP İÇİNDE OLUP OLMADIĞINI KONTROL ETME ------------------------------------------
	//value, ok := studentGrades["Alptürk"]
	//_, ok2 := studentGrades["Çiğdem"]
	//fmt.Println(value, ok)
	//
	//if ok2 {
	//	fmt.Println("bu değer mevcuttur")
	//}

	//MAP E YENİ DEĞER EKLEME ------------------------------------------------------------------
	//fmt.Println(studentGrades) // => map[Buğrahan:80 Furkan:58 Zeynep:77 Çiğdem:96]
	//studentGrades["Ali"] = 99
	//fmt.Println(studentGrades) // => map[Ali:99 Buğrahan:80 Furkan:58 Zeynep:77 Çiğdem:96]

	//MAP DEN DEĞER SİLME ------------------------------------------------------------------
	//delete(studentGrades, "Furkan")
	//fmt.Println(studentGrades)

	// TEK TEK ELEMANLARA ERİŞME
	for k, v := range studentGrades {
		fmt.Println(k, v)
	}
}
