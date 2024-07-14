package main

import (
	"fmt"           // fmt paketi, formatlı I/O işlemleri için kullanılır
	"path/filepath" // filepath paketi, dosya ve dizin yollarını işlemek için kullanılır
	"strings"       // strings paketi, string işlemleri için kullanılır
)

func filepathf() {

	// filepath.Join ile dizin ve dosya isimlerini birleştirir
	p := filepath.Join("dir1", "dir2", "filename")
	fmt.Println("p:", p) // "dir1/dir2/filename"

	// Birleştirirken gereksiz '/' karakterlerini düzeltir
	fmt.Println(filepath.Join("dir1//", "filename")) // "dir1/filename"
	// "../" işlemini değerlendirir ve doğru yolu bulur
	fmt.Println(filepath.Join("dir1/../dir1", "filename")) // "dir1/filename"

	// Verilen yolun sadece dizin kısmını alır
	fmt.Println("Dir(p):", filepath.Dir(p)) // "dir1/dir2"
	// Verilen yolun sadece dosya ismini alır
	fmt.Println("Base(p):", filepath.Base(p)) // "filename"

	// Yolların mutlak olup olmadığını kontrol eder
	fmt.Println(filepath.IsAbs("dir/file"))  // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true

	// Dosya ismini belirler
	filename := "config.json"

	// Dosya uzantısını alır
	ext := filepath.Ext(filename)
	fmt.Println(ext) // ".json"

	// Dosya isminin uzantısız halini alır
	fmt.Println(strings.TrimSuffix(filename, ext)) // "config"

	// İlk yoldan ikinci yola olan göreceli yolu hesaplar
	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // "t/file"

	// Farklı dizinler arasındaki göreceli yolu hesaplar
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel) // "../c/t/file"
}
