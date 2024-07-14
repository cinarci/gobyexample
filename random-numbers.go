package main

import (
	"fmt"
	"math/rand/v2" // rand paketinin 2. sürümü import ediliyor
)

func randomnumbers() {

	// 0 ile 99 arasında rastgele iki tamsayı üretir ve yazdırır
	fmt.Print(rand.IntN(100), ",")
	fmt.Print(rand.IntN(100))
	fmt.Println()

	// 0.0 ile 1.0 arasında rastgele bir float64 sayı üretir ve yazdırır
	fmt.Println(rand.Float64())

	// 5.0 ile 10.0 arasında rastgele iki float sayı üretir ve yazdırır
	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()

	// Belirli tohumlarla yeni bir PCG (Permuted Congruential Generator) üretici örneği oluşturur
	s2 := rand.NewPCG(42, 1024)
	r2 := rand.New(s2)

	// Yeni üretici örneği kullanarak 0 ile 99 arasında rastgele iki tamsayı üretir ve yazdırır
	fmt.Print(r2.IntN(100), ",")
	fmt.Print(r2.IntN(100))
	fmt.Println()

	// Aynı tohumları yeniden kullanarak başka bir üretici örneği oluşturur
	s3 := rand.NewPCG(42, 1024)
	r3 := rand.New(s3)

	// Yeniden kullanılan üretici örneği ile 0 ile 99 arasında rastgele iki tamsayı üretir ve yazdırır
	fmt.Print(r3.IntN(100), ",")
	fmt.Print(r3.IntN(100))
	fmt.Println()
}
