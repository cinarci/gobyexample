package main

import (
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"strconv" // strconv paketi, string dönüşümleri için kullanılır
)

func numberparsing() {

	// ParseFloat: Verilen string'i float64 tipine dönüştürür
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f) // Çıktı: 1.234

	// ParseInt: Verilen string'i int64 tipine dönüştürür
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i) // Çıktı: 123

	// 0x ile başlayan string'i hexadeciaml olarak ParseInt kullanarak dönüştürür
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d) // Çıktı: 456

	// ParseUint: Verilen string'i uint64 tipine dönüştürür
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u) // Çıktı: 789

	// Atoi: Verilen string'i integer (int) tipine dönüştürür
	k, _ := strconv.Atoi("135")
	fmt.Println(k) // Çıktı: 135

	// Atoi: Dönüşüm hatası durumunda dönen hata
	_, e := strconv.Atoi("wat")
	fmt.Println(e) // Çıktı: strconv.Atoi: parsing "wat": invalid syntax
}
