package main

import "fmt"

func main() {

	//underArray := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(underArray)
	//
	//mySlc := underArray[2:5] // underArray'in başlangıc ve bitişini belirtiriz. (başlangıç dahil, bitiş dahil değil)
	//fmt.Println(mySlc)       // => [2 3 4]
	//mySlc2 := underArray[4:]
	//fmt.Println(mySlc2) // => [4 5 6 7 8 9]
	//mySlc3 := underArray[:7]
	//fmt.Println(mySlc3) // => [0 1 2 3 4 5 6]
	//mySlc4 := underArray[:]
	//fmt.Println(mySlc4) // => [0 1 2 3 4 5 6 7 8 9]

	//--------------------------SLİCE'A YENİ ELEMAN EKLEME---------------------------------------
	//mySlc := []int{1, 2, 3}
	//fmt.Println(mySlc)
	//
	//mySlc = append(mySlc, 4, 5)
	//fmt.Println(mySlc)
	//
	//myslc2 := append(mySlc, 4, 5)
	//fmt.Println(myslc2)

	//--------------------İKİ FARKLI SLİCE'IN ELEMANLARINI BİRBİRİNE EKLEME------------------------------

	//mySlc := []int{2, 4, 6}
	//fmt.Println(mySlc)
	//
	//mySlc2 := []int{8, 10, 12}
	//fmt.Println(mySlc2)
	//
	//mySlc = append(mySlc, mySlc2...) // (üç nokta) slice tipini elemanlarına ayırır
	//fmt.Println(mySlc)

	//----------------------------------------- ELEMAN SİLME ---------------------------------------------------
	mySlc := []int{1, 2, 3, 4, 5, 6}
	mySlc2 := mySlc[3:] // => ilk 3 elemanı silmek için
	fmt.Println(mySlc2)
	mySlc3 := mySlc[:len(mySlc)-2] //  => son 2 elemeanı silmek için
	fmt.Println(mySlc3)

}
