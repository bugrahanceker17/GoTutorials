package main

import (
	"fmt"
	"sync"
)

//Goroutines -> Birden fazla işin aynı anda yapılması (Async)
//Eş zamanlı (concurrency) olarak yapılan her bir göreve 'Goroutines' denir
var wg sync.WaitGroup

func main() { // main çalışmaya başladığında main goroutine çalışmaya başlıyor
	wg.Add(2)

	go printX()
	wg.Wait()
	fmt.Println()
	printY()
}

func printX() {
	for i := 0; i < 20; i++ {
		fmt.Print("X")
	}
	wg.Done()
}

func printY() {
	for i := 0; i < 20; i++ {
		fmt.Print("Y")
	}
}
