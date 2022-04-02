//package main
//
//import "fmt"
//
//func main() {
//	mySlc := []int{1, 10, 100}
//	fmt.Println(mySlc) // => [1 10 100]
//
//	double(mySlc) // => [2 20 200]
//	// Çünkü slices'lar pass by reference olduğu için, double fonksiyonuna mySlc slices'ının bir kopyası gönderilir
//	// mySlc referansı değiştiği için, kendisi de değişir
//	fmt.Println(mySlc) // => [2 20 200]
//
//}
//
//func double(slc []int) {
//	for i := 0; i < len(slc); i++ {
//		slc[i] *= 2
//	}
//
//	fmt.Println(slc)
//}

package main

import "fmt"

func main() {
	x := 5
	fmt.Println(x)
	double(&x)
	fmt.Println(x)
}

// normal şartlarda referansının kopyasını alıp değiştirirdik ancak x yine aynı şekilde kalırdı
// fakat pointer yardımıyla x değerinin adresini alıp onun üzerinde işlem yaptığımız için
// artık x'in de değerini değiştirmiş olduk
func double(num *int) { // -> pointer method
	fmt.Println(num)
	*num *= 2
	fmt.Println(*num)
}
