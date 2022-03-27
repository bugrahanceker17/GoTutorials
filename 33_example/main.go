// 1-) 5 string elemandan oluşan bir array ve 5 int elemandan oluşan bir slices oluşturup index numaralarıyla birlikte gösterin
//package main
//
//import "fmt"
//
//func main() {
//
//	strArray := [5]string{"Bu", "bir", "string", "array", "denemesidir"}
//
//	for i, v := range strArray {
//		fmt.Printf("%v, %v", i, v)
//		fmt.Println("")
//	}
//	fmt.Println("")
//
//	intSlices := []int{100, 200, 300}
//	for ind, val := range intSlices {
//		fmt.Printf("%v, %v", ind, val)
//		fmt.Println("")
//	}
//}

// 2-) [0,1,2,3,4,5,6,7,8] slice dan [0,1,2,3], [4,5,6], [6,7,8], [2,3,6,7] slicelarını oluşturunuz.
//package main
//
//import "fmt"
//
//func main() {
//	slices := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
//	fmt.Println(slices[:4])  // => [0 1 2 3]
//	fmt.Println(slices[4:7]) // => [4 5 6]
//	fmt.Println(slices[6:])  // => [6 7 8]
//
//	editedSlices := append(slices[2:4], slices[6:8]...)
//	fmt.Println(editedSlices) // => [2,3,6,7]
//}

// 3-) slicelar için copy metodunu ve assign ( = ) ile farkını açıklayınız.
package main

import "fmt"

func main() {
	slices1 := []int{5, 8, 16, 48}
	slices2 := make([]int, 2)
	fmt.Println(slices1) // => [5 8 16 48]
	fmt.Println(slices2) // => [0 0]
	copy(slices2, slices1)
	fmt.Println(slices1) // => [5 8 16 48]
	fmt.Println(slices2) // => [5 8] -> Sadece 2 elemanı kopyaladı çünkü uzunluğu 2

	slices1[0] = 100
	fmt.Println(slices1) // => [100 8 16 48]
	fmt.Println(slices2) // => [5 8] Değişmedi çünkü aynı refeansta değiller. Değerleri kopyaladı, referans almadı

	//ancak assing (=) işlemi yaparsak referans olarak alır ve 1.slices'da değişen eleman, 2.slices'da da değişir
}

// 4-) map gösterimi ile yazar ve yazarlara ait kitapların isimlerini for range ile gösterin.
//package main
//
//import "fmt"
//
//func main() {
//
//	values := map[string]string{
//		"Sabahattin Ali":       "Kürk Mantolu Madonna",
//		"Peyami Safa":          "Dokuzuncu Hariciye Koğuşu",
//		"Namık Kemal":          "İntibah",
//		"Halid Ziya Uşaklıgil": "Mai ve Siyah",
//	}
//
//	for author, book := range values {
//		fmt.Printf("Yazar: %v, Kitap: %v", author, book)
//		fmt.Println()
//	}
//}
