// 1 -) Underlying Type 'int' olacak şekilde kendi veri tipinizi oluşturunuz
// veri tipi ve değerini '10' yazdırınız.

//package main
//
//import "fmt"
//
//type myValue int
//
//func main() {
//	var x myValue = 17
//	fmt.Printf("%T, %v", x, x)
//}

// 2 -) Başlangıç değeri 10 olan bir X değişkeni oluşturun sonrasında
// bulunduğu adres üzerinden y değişkenini tanımlayıp x değerini 20 yapınız.

//package main
//func main() {
//	x := 10
//	fmt.Println(x)
//	y := &x
//	*y = 20
//	fmt.Println(x)
//	fmt.Println(*y)
//
//}

// 3 -) Underlying Type struct olan Rectangle type oluşturunuz.
// display, area, circumference metodlarını yazınız.

//package main
//
//import "fmt"
//
//type rectangle struct {
//	a int
//	b int
//}
//
//func (r rectangle) display() {
//	fmt.Println("Kenar uzunlukları", r.a, "ve", r.b, "olan dikdörtgen")
//}
//
//func (r rectangle) area() int {
//	return r.a * r.b
//}
//
//func (r rectangle) circumference() int {
//	return 2 * (r.a + r.b)
//}
//
//func main() {
//	r1 := rectangle{3, 4}
//	r1.display()
//	fmt.Println("Alanı :", r1.area())
//	fmt.Println("Çevresi :", r1.circumference())
//}

// 4-) family aile bireyleri şeklinde veri tipi oluşturalım, underlying struct. Aile bireylerinin hepsini farklı
// şekilde tanımlayınız. Sonrasında for döngüsünde yazdırınız. field age, name, isMarried field.

package main

import "fmt"

type family struct {
	age       int
	name      string
	isMarried bool
}

func show() []family {
	f1 := family{
		age:       24,
		name:      "Bugrahan Alpturk",
		isMarried: false,
	}

	f2 := family{
		age:       20,
		name:      "Batuhan Gokturk",
		isMarried: false,
	}

	f3 := family{
		age:       14,
		name:      "Bilgehan Serturk",
		isMarried: false,
	}

	f4 := family{
		age:       50,
		name:      "Hatice",
		isMarried: true,
	}

	f5 := family{
		age:       49,
		name:      "Mehmet",
		isMarried: true,
	}

	return []family{f1, f2, f3, f4, f5}
}

func main() {
	familyMembers := show()

	for i := 0; i < len(familyMembers); i++ {
		fmt.Println(familyMembers[i])
	}
}
