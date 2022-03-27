package main

import "fmt"

func main() {
	//x := 12
	//if x%2 == 0 {
	//	fmt.Printf("%v sayısının 2'ye bölümünden kalan 0 dır", x)
	//} else {
	//	fmt.Printf("%v sayısı 2'ye tam bölünmez", x)
	//}

	a := 11

	if a < 0 {
		fmt.Println(a, "negatif sayıdır")
	} else if a%2 == 0 {
		fmt.Println(a, "pozitif ve çift sayıdır")
	} else {
		fmt.Println(a, "pozitif ve tek sayıdır")
	}

}
