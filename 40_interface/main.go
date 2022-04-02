package main

import (
	"fmt"
	"math"
)

type rectangle struct {
	a float64
	b float64
}

type circle struct {
	r float64
}

type shape interface {
	area() float64
	cimcumference() float64
}

func interfaceFunc(i shape) {
	fmt.Println(i)
	fmt.Println(i.area())
	fmt.Println(i.cimcumference())
	fmt.Printf("%T", i)
	fmt.Println()
}

func main() {
	r1 := rectangle{3, 8}
	fmt.Println("Alanı :", r1.area())
	fmt.Println("Çevresi :", r1.circumference())

	//c1 := circle{5}
	//interfaceFunc(c1)
}

func (r rectangle) display() {
	fmt.Println("Kenar uzunlukları", r.a, "ve", r.b, "olan dikdörtgen")
}

func (r rectangle) area() float64 {
	return r.a * r.b
}

func (r rectangle) circumference() float64 {
	return 2 * (r.a + r.b)
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) circumference() float64 {
	return 2 * math.Pi * c.r
}
