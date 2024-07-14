package main

import (
	"fmt"           // fmt paketi, formatlı I/O işlemleri için kullanılır
	"io/fs"         // io/fs paketi, dosya sistemi arayüzlerini sağlar
	"os"            // os paketi, işletim sistemi işlevlerine erişim sağlar
	"path/filepath" // filepath paketi, dosya ve dizin yollarını işlemek için kullanılır
)

// check fonksiyonu, hata kontrolü yapar ve hata durumunda panik oluşturur
func checkd(e error) {
	if e != nil {
		panic(e)
	}
}

// directories fonksiyonu, dizin ve dosya işlemlerini gerçekleştirir
func directories() {

	// "subdir" adında bir dizin oluşturur
	err := os.Mkdir("subdir", 0755)
	check(err)

	// Program bitiminde "subdir" dizinini ve içeriğini siler
	defer os.RemoveAll("subdir")

	// Boş bir dosya oluşturmak için yardımcı bir fonksiyon
	createEmptyFile := func(name string) {
		d := []byte("")
		check(os.WriteFile(name, d, 0644))
	}

	// "subdir/file1" adında boş bir dosya oluşturur
	createEmptyFile("subdir/file1")

	// "subdir/parent/child" adında bir dizin oluşturur
	err = os.MkdirAll("subdir/parent/child", 0755)
	check(err)

	// "subdir/parent" dizininde boş dosyalar oluşturur
	createEmptyFile("subdir/parent/file2")
	createEmptyFile("subdir/parent/file3")
	createEmptyFile("subdir/parent/child/file4")

	// "subdir/parent" dizinini okur
	c, err := os.ReadDir("subdir/parent")
	check(err)

	// "subdir/parent" içeriğini listeleyerek ekrana yazdırır
	fmt.Println("Listing subdir/parent")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Çalışma dizinini "subdir/parent/child" olarak değiştirir
	err = os.Chdir("subdir/parent/child")
	check(err)

	// Geçerli dizini (şu anda "subdir/parent/child") okur
	c, err = os.ReadDir(".")
	check(err)

	// "subdir/parent/child" içeriğini listeleyerek ekrana yazdırır
	fmt.Println("Listing subdir/parent/child")
	for _, entry := range c {
		fmt.Println(" ", entry.Name(), entry.IsDir())
	}

	// Çalışma dizinini kök dizine ("subdir" dizinine geri dönmek için) değiştirir
	err = os.Chdir("../../..")
	check(err)

	// "subdir" dizinini ziyaret eder ve içeriğini listeleyerek ekrana yazdırır
	fmt.Println("Visiting subdir")
	err = filepath.WalkDir("subdir", visit)
}

// visit fonksiyonu, WalkDir tarafından her dosya veya dizin ziyaret edildiğinde çağrılır
func visit(path string, d fs.DirEntry, err error) error {
	if err != nil {
		return err
	}
	fmt.Println(" ", path, d.IsDir())
	return nil
}
