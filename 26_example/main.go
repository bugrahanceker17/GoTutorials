//------------------------------------------------------------------------------------------------
// iki rakam arasında toplama, çıkarma ve çarpma işlemlerinin yapıldığı fonk yaz

//package main
//
//import "fmt"
//
//func main() {
//	sum, dif, prod := calculation(5, 8)
//	fmt.Println("Toplama :", sum)
//	fmt.Println("Çıkarma :", dif)
//	fmt.Println("Çarpımı :", prod)
//}
//
//func calculation(num1, num2 int) (sum, dif, prod int) {
//	sum = num2 + num1
//	dif = num1 - num2
//	prod = num1 * num2
//	return sum, dif, prod
//}

//------------------------------------------------------------------------------------------------
//klavyeden girilen nota göre geçtiniz veya kaldınız dönüşü yapan fonk yaz

//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	fmt.Print("Lütfen notunuzu giriniz : ")
//	response, _ := calc()
//
//	var result string
//
//	if response >= 50 {
//		result = "Geçtin"
//	} else {
//		result = "Kaldın"
//	}
//
//	fmt.Println(result)
//}
//
//func calc() (int, error) {
//	reader := bufio.NewReader(os.Stdin)
//	input, err := reader.ReadString('\n')
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	input = strings.TrimSpace(input)
//	num, err := strconv.Atoi(input)
//	if err != nil {
//		fmt.Println(err)
//	}
//
//	return num, nil
//}

//------------------------------------------------------------------------------------------------
// 1-100 arasında bir sayıyı tahmin etme uygulaması yazınız. Toplam hakkınız 10 olsun

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println("SAYI TAHMİN OYUNU")

	target := getRandom(1, 100)

	reader := bufio.NewReader(os.Stdin)

	for attemps := 0; attemps < 10; attemps++ {
		fmt.Println(10-attemps, "hakkınız kaldı")
		fmt.Print("Lütfen tahmininizi giriniz : ")

		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		input = strings.TrimSpace(input)
		number, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println(err)
		}
		if attemps != 9 {
			if number > target {
				fmt.Print("Daha küçük bir sayı giriniz... ")
			} else if number < target {
				fmt.Print("Daha büyük bir sayı giriniz... ")
			} else {
				fmt.Printf("DOĞRU TAHMİN! %d. denemede buldunuz", attemps)
				break
			}
		} else {
			fmt.Printf("10 denemede bulamadınız... Aranan sayı : %d", target)
			break
		}
	}

}

func getRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
