package main

import (
	"embed" // embed paketi, Go 1.16 ve sonrasında dosyaları derleme zamanında ikili dosyaya gömmek için kullanılır
)

/*
// single_file.txt dosyasını string değişkenine gömer
//
//go:embed folder/single_file.txt
*/
var fileString string

/*
// single_file.txt dosyasını byte dilimi (slice) değişkenine gömer
//
//go:embed folder/single_file.txt
*/
var fileByte []byte

/*
	folder icindeki single_file.txt ve *.hash uzantılı dosyaları embed.FS tipindeki folder değişkenine gömer

//
//go:embed folder/single_file.txt
//go:embed folder/*.hash
*/
var folder embed.FS

func embeddirective() {

	// fileString içeriğini ekrana yazdırır
	print(fileString)
	// fileByte içeriğini string olarak ekrana yazdırır
	print(string(fileByte))

	// folder içindeki file1.hash dosyasını okur ve içeriğini ekrana yazdırır
	content1, _ := folder.ReadFile("folder/file1.hash")
	print(string(content1))

	// folder içindeki file2.hash dosyasını okur ve içeriğini ekrana yazdırır
	content2, _ := folder.ReadFile("folder/file2.hash")
	print(string(content2))
}
