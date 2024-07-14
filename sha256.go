package main

import (
	"crypto/sha256" // crypto/sha256 paketi, SHA-256 hash fonksiyonunu sağlar
	"fmt"           // fmt paketi, formatlı I/O işlemleri için kullanılır
)

func sha() {
	s := "sha256 this string" // Hash'ini hesaplayacağımız string

	h := sha256.New() // Yeni bir SHA-256 hash nesnesi oluştur

	h.Write([]byte(s)) // String'i byte dizisine dönüştürüp hash nesnesine yaz

	bs := h.Sum(nil) // Hash'lenmiş veriyi al (nil parametresi, ekstra veri eklenmemesini sağlar)

	fmt.Println(s)         // İlk olarak, işlenen string'i ekrana yazdır
	fmt.Printf("%x\n", bs) // Hash'in hexadecimal formatında çıktısını ekrana yazdır
}
