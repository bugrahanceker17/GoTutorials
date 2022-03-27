package main

import "fmt"

func main() {
	// ARRAY VE SLİCES ARASINDAKİ EN TEMEL FARK : ARRAYLER OLUŞTURULURKEN ELEMAN SAYISINI BİLİR
	// FAKAT SLİCES LAR BİLMEZ

	//myArr := [3]int{1, 2, 3}
	//myArr2 := [...]int{1, 2, 3, 4}
	//fmt.Println(myArr)
	//fmt.Println(myArr2)
	//fmt.Println(len(myArr))
	//fmt.Println(len(myArr2))

	//mySlc := []int{1, 2, 3}
	//fmt.Println(mySlc)
	//fmt.Println(len(mySlc))

	var slicess []int // boş slices oluşur
	fmt.Println(slicess)
	slicess = make([]int, 4) //burada kaç elemanlı olduğunu atarız--make metodu ile oluşturulan slices'larda zero value olur
	fmt.Println(slicess)
	slicess[0] = 10
	fmt.Println(slicess)

	//array ve slices farkları
	// 1 - Arraylerin index sınırları belli olduğundan compile time'da değişmeyeceği için daha güvenilirdir
	// 2 - Biz bir array'in kopyasını aldığımızda, elemanların kopyasını alırız yani array'in kendisiniz kopyalarız
	// fakat slices'larda bu durum geçerli değildir. Slices'larda referansların kopyasını alırız

	//myArr := [3]int{1, 2, 3}
	//fmt.Println(myArr) // -> [1 2 3]
	//var myArr2 = myArr
	//fmt.Println(myArr2) // -> [1 2 3]
	//
	//myArr2[0] = 100
	//fmt.Println(myArr2) // -> [100 2 3]
	//fmt.Println(myArr)  // -> [1 2 3]
	//herhangi bir değişiklik yok çünkü değerlerini kopyaladık
	//---------- PASS BY VALUE -----------

	mySlc := []int{1, 2, 3}
	fmt.Println(mySlc) // -> [1 2 3]
	mySlc2 := mySlc
	fmt.Println(mySlc2) // -> [1 2 3]

	mySlc2[0] = 33
	fmt.Println(mySlc)  // -> [33 2 3]
	fmt.Println(mySlc2) // -> [33 2 3]
	// Slices'lar ise referans tutar. Yani bir slices'da yapılan değişiklik, referansını tuttuğu diğer slices'da da değişir
	//---------- PASS BY REFERENCE -----------

}
