package main

import (
	"bufio" // bufio paketi, tamponlu okuma işlemleri için kullanılır
	"fmt"   // fmt paketi, formatlı I/O işlemleri için kullanılır
	"io"    // io paketi, temel I/O işlemleri için kullanılır
	"os"    // os paketi, işletim sistemi işlevlerine erişim sağlar
)

func checkrf(e error) {
	if e != nil {
		panic(e) // Hata oluştuğunda programı durdurur ve hata mesajını gösterir
	}
}

func readingfiles() {
	// Dosya içeriğini tamamını oku ve ekrana yazdır
	dat, err := os.ReadFile("/tmp/dat")
	check(err)
	fmt.Print(string(dat))

	// Dosyayı aç ve ilk 5 byte'ı oku
	f, err := os.Open("/tmp/dat")
	check(err)
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes: %s\n", n1, string(b1[:n1]))

	// Dosyanın 6. byte'ından itibaren 2 byte oku
	o2, err := f.Seek(6, io.SeekStart)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n2, o2, string(b2[:n2]))

	// Dosyanın mevcut konumundan 4 byte ileri git
	_, err = f.Seek(4, io.SeekCurrent)
	check(err)

	// Dosyanın sonundan 10 byte geri git
	_, err = f.Seek(-10, io.SeekEnd)
	check(err)

	// Dosyanın 6. byte'ından itibaren en az 2 byte oku
	o3, err := f.Seek(6, io.SeekStart)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// Dosyanın başlangıcına git
	_, err = f.Seek(0, io.SeekStart)
	check(err)

	// bufio.NewReader ile tamponlu okuma yap
	r4 := bufio.NewReader(f)
	b4, err := r4.Peek(5)
	check(err)
	fmt.Printf("5 bytes: %s\n", string(b4))

	f.Close() // Dosyayı kapat
}
