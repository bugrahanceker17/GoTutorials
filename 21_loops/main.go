package main

import "fmt"

func main() {

	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}
	// ---------------------------------  ---------------------------------  ---------------------------------
	//j := 0
	//for ; j <= 10; j += 5 { // -------------------FARKLI YAZIMI
	//	fmt.Println(j)
	//}
	//fmt.Println(j)
	// ---------------------------------  ---------------------------------  ---------------------------------
	//for { //----------Infinite loop - Sonsuz döngü
	//	fmt.Println("test")
	//}
	// ---------------------------------  ---------------------------------  ---------------------------------
	//for i := 0; true; i++ { // ------- true olduğu her koşulda yani sonsuz döngü
	//	fmt.Println("sonsuz döngü", i)
	//}
	// ---------------------------------  ---------------------------------  ---------------------------------
	//for i := 0; false; i++ { // ------- true koşulunu sağlamadığı için döngü hiç çalışmaz
	//	fmt.Println("sonsuz döngü", i)
	//}
	// ---------------------------------  ---------------------------------  ---------------------------------
	//i := 10
	//for i > 0 { // -------- i nin conditinal'ı 0 dan büyük olması olduğu için
	//	fmt.Println(i)
	//	i--
	//}
	// ---------------------------------  ---------------------------------  ---------------------------------
	//for i := 0; i <= 10; i++ {
	//	if i%3 == 0 {
	//		continue // bu şartı sağlarsa döngüdeki bir sonraki adıma geç
	//	}
	//	fmt.Println(i)
	//}// ---------------------------------  ---------------------------------  ---------------------------------
	for i := 0; i <= 10; i++ {
		if i == 3 {
			break // bu şartı sağlarsa döngüden çık
		}
		fmt.Println(i)
	}
}
