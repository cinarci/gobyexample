package main

import (
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"testing" // testing paketi, test yazma ve çalıştırma işlevlerini sağlar
)

// IntMin fonksiyonu, verilen iki tamsayıdan daha küçük olanı döndürür
func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// TestIntMinBasic fonksiyonu, IntMin fonksiyonunu temel bir test ile sınar
func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2) // IntMin(2, -2) çağrısı yapılır
	if ans != -2 {       // Sonuç beklenen değer olan -2 değilse hata verir
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

// TestIntMinTableDriven fonksiyonu, IntMin fonksiyonunu çeşitli test vakaları ile sınar
func TestIntMinTableDriven(t *testing.T) {
	// Test vakalarını içeren bir dilim oluşturur
	var tests = []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
		{0, -1, -1},
		{-1, 0, -1},
	}

	// Her bir test vakasını işler
	for _, tt := range tests {
		// Test ismini oluşturur
		testname := fmt.Sprintf("%d,%d", tt.a, tt.b)
		// Her bir test vakası için ayrı bir alt test çalıştırır
		t.Run(testname, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b) // IntMin fonksiyonunu çağırır
			if ans != tt.want {       // Sonuç beklenen değere eşit değilse hata verir
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// BenchmarkIntMin fonksiyonu, IntMin fonksiyonunun performansını sınar
func BenchmarkIntMin(b *testing.B) {
	// B.N defa IntMin fonksiyonunu çağırır
	for i := 0; i < b.N; i++ {
		IntMin(1, 2)
	}
}

/*
go test -y

go test -bench=.
*/
