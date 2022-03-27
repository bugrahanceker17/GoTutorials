package main

import "fmt"

func main() {

	if x := 5; x < 0 {
		fmt.Println(x, "negatif sayıdır")
	} else {
		fmt.Println(x, "pozitif sayıdır")
	}

	//fmt.Println(x) // ÇALIŞMAZ ÇÜNKÜ SADECE İF BLOĞU İÇİN TANIMLANMIŞTIR

}
