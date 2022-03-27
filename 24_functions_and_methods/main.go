//package main
//
//func main() {
//
//	//hello("Bugrahan") // Argument
//
//   var x int = 10
//   var today time.Time = time.Now()  // Now() -> Method
//   fmt.Println(today)
//   fmt.Println(x)
//   fmt.Println(today)
//}

//func hello(name string) { // Parameter
//	fmt.Printf("Benim adım %s", name)
//}

//----------------------------------------------------------------------------------------------------
//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	fmt.Print("Lütfen sınav sonucunu giriniz : ")
//	reader := bufio.NewReader(os.Stdin)
//	value, _ := reader.ReadString('\n') // _ blank identifier
//	fmt.Println(value)
//}

//----------------------------------------------------------------------------------------------------

package main

import "fmt"

func main() {
	bolum, kalan := test(104, 5)
	fmt.Printf("Bölüm : %d, Kalan: %d", bolum, kalan)
}

func test(v1, v2 int) (bolum, kalan int) {
	bolum = v1 / v2
	kalan = v1 % v2
	return bolum, kalan
}
