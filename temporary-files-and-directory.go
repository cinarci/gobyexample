package main

import (
	"fmt"           // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"            // os paketi, işletim sistemi işlevlerine erişim sağlar
	"path/filepath" // filepath paketi, dosya ve dizin yollarını işlemek için kullanılır
)

// check fonksiyonu, hata kontrolü yapar ve hata durumunda panik oluşturur
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func tfd() {

	// Geçici bir dosya oluşturur
	f, err := os.CreateTemp("", "sample")
	check(err) // Hata kontrolü yapar

	// Geçici dosyanın ismini ekrana yazdırır
	fmt.Println("Temp file name:", f.Name())

	// Program bitiminde geçici dosyayı siler
	defer os.Remove(f.Name())

	// Geçici dosyaya 1, 2, 3, 4 baytlarını yazar
	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err) // Hata kontrolü yapar

	// Geçici bir dizin oluşturur
	dname, err := os.MkdirTemp("", "sampledir")
	check(err) // Hata kontrolü yapar
	// Geçici dizinin ismini ekrana yazdırır
	fmt.Println("Temp dir name:", dname)

	// Program bitiminde geçici dizini ve içeriğini siler
	defer os.RemoveAll(dname)

	// Geçici dizin içinde bir dosya yolu oluşturur
	fname := filepath.Join(dname, "file1")
	// Bu dosyaya 1, 2 baytlarını yazar ve dosya izinlerini 0666 olarak ayarlar
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err) // Hata kontrolü yapar
}
