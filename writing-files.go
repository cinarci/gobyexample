package main

import (
	"bufio" // bufio paketi, tamponlu yazma işlemleri için kullanılır
	"fmt"   // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"    // os paketi, işletim sistemi işlevlerine erişim sağlar
)

func checkw(e error) {
	if e != nil {
		panic(e) // Hata oluştuğunda programı durdurur ve hata mesajını gösterir
	}
}

func writingfiles() {
	// İlk dosyaya yazılacak veri
	d1 := []byte("hello\ngo\n")
	err := os.WriteFile("/tmp/dat1", d1, 0644) // Dosyaya veriyi yazma
	check(err)

	// İkinci dosya oluşturuluyor
	f, err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close() // main fonksiyonun sonunda dosyayı kapat

	// İkinci dosyaya yazılacak veri
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2) // Dosyaya veriyi yazma
	check(err)
	fmt.Printf("wrote %d bytes\n", n2)

	// Dosyaya string olarak yazma
	n3, err := f.WriteString("writes\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", n3)

	// Veriyi tamponlayarak dosyaya yazma
	f.Sync() // Tamponlanan veriyi diske yazar

	w := bufio.NewWriter(f)                // Yeni bir bufio.Writer oluştur, f dosyasına yazma işlemi için
	n4, err := w.WriteString("buffered\n") // Tampon üzerinden veriyi dosyaya yaz
	check(err)
	fmt.Printf("wrote %d bytes\n", n4)

	w.Flush() // Tampondaki veriyi dosyaya yaz

}
