//package main
//
//import (
//	"errors"
//	"fmt"
//)
//
//func main() {
//	result, err := evenNum(10)
//	if err != nil {
//		fmt.Println(err)
//	} else {
//		fmt.Println("Sonuç :", result)
//	}
//}
//
//func evenNum(num int) (int, error) {
//	if num%2 != 0 {
//		return 0, errors.New("HATA: Tek sayı girdiniz")
//	}
//	return num, nil
//}

//-----------------------------------------------------------------------------------------------------

// KAREKÖK ALMA İŞLEMİ YAPILACAK
//POZİTİF SAYILARIN KAREKÖKÜ ALINABİLİR, NEGATİF SAYILARIN KAREKÖKÜ ALINAMAZ
package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	result, err := squareRoot(17)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Karekökü :", result)
	}

}

func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("HATA! Negatif sayıların karekökü alınamaz")
	}

	return math.Sqrt(num), nil
}
