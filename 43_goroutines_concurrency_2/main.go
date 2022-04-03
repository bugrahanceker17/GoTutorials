package main

import (
	"fmt"
	"math"
	"sync"
)

type circle struct {
	r float64
}

func (c circle) display() {
	fmt.Println("A Circle")
	await.Done()
}

func area(c circle, channel chan float64) {
	result := math.Pi * c.r * c.r
	channel <- result
}

var await sync.WaitGroup

func main() {
	await.Add(1)
	c1 := circle{6}
	channel := make(chan float64)
	go c1.display()
	await.Wait()
	go area(c1, channel)
	fmt.Printf("Yarıçapı %v olan dairenin alanı : %.2f\n", c1.r, <-channel)

}

// GoLang'da return işlemi yapılan bir fonksiyon için goroutines kullanılamaz
// bunun için channel lar kullanılır
// ---------------------------------------- CHANNEL OLUŞTURMA İŞLEMi ----------------------------------------

//package main
//
//import "fmt"
//
//func merhaba(chan1 chan string) {
//	chan1 <- "Merhaba" //Kanala bu değeri gönderiyoruz
//}
//
//func main() {
//	myChannel := make(chan string) // kanalı oluşturuyoruz
//
//	go merhaba(myChannel) // Değer almak için return kullanamayız. Kanal üzerinden alabiliriz
//
//	fmt.Println(myChannel)
//	fmt.Println(<-myChannel)
//
//}

// ---------------------------------------- ikinci örnek ----------------------------------------

//package main
//
//import "fmt"
//
//func main() {
//	myCn := make(chan string)
//
//	myCn <- "GOLANG IS AWESOME"
//	// Channel -> bulunduğu goroutine den, değer dönmeden bir sonraki işleme geçmez, bloklar
//	fmt.Println(<-myCn)
//}
