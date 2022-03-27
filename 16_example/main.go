// 1)    x =x-10 vsx-=10 ---------------------------------------------------------------------------
//package main
//
//import "fmt"
//
//func main() {
//	x := 50
//	// x = x - 10 // assignment statement
//	x -= 10 // assignment operation
//	fmt.Printf("%T, %v", x, x)
//}

// 2)   K = F -32 / 1.8 + 273  => -40 F kaç K derecedir? ---------------------------------------------
//package main
//
//import "fmt"
//
//func main() {
//	F := -40
//	K := float64(F-32)/1.8 + 273
//	fmt.Printf("%T, %v", K, K)
//}

// 3) Sabitlerde Shadowing çalışır mı ?  ------------------------------ ------------------------------
//package main
//
//import "fmt"
//
//const x = 17
//
//func main() {
//	const x = 110
//
//	fmt.Printf("%v, %T\n", x, x)
//}

// 4) const x = 4 , const y = 5.2,  x + y ??  ------------------------------ ------------------------------
//package main
//
//import "fmt"
//
//func main() {
//	const x = 4 //typless şuanda
//	const y = 5.2
//	fmt.Printf("%T, %v", x+y, x+y) // x in veritipi float64 olur burada
//}

// 5) const x float64 = 6.4,  y:=4 + x,	 y = ?????   ------------------------------ ------------------------------
package main

import "fmt"

func main() {
	const x float64 = 7.1
	y := 4 + x
	fmt.Printf("%v, %T", y, y)
}
