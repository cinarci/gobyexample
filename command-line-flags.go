package main

import (
	"flag" // flag paketi, komut satırı bayraklarını (flags) işlemek için kullanılır
	"fmt"  // fmt paketi, formatlı I/O işlemleri için kullanılır
)

func clf() {

	// Komut satırı bayrağı olarak bir string değişkeni tanımlar
	wordPtr := flag.String("word", "foo", "a string")
	// Komut satırı bayrağı olarak bir integer değişkeni tanımlar
	numbPtr := flag.Int("numb", 42, "an int")
	// Komut satırı bayrağı olarak bir boolean değişkeni tanımlar
	forkPtr := flag.Bool("fork", false, "a bool")

	// Komut satırı bayrağı olarak bir string değişkeni tanımlar ve bu değişkeni önceden tanımlanmış bir değişkene bağlar
	var svar string
	flag.StringVar(&svar, "svar", "bar", "a string var")

	// Komut satırı bayraklarını ayrıştırır
	flag.Parse()

	// Ayrıştırılan bayrakların değerlerini ekrana yazdırır
	fmt.Println("word:", *wordPtr)    // word bayrağının değeri
	fmt.Println("numb:", *numbPtr)    // numb bayrağının değeri
	fmt.Println("fork:", *forkPtr)    // fork bayrağının değeri
	fmt.Println("svar:", svar)        // svar bayrağının değeri
	fmt.Println("tail:", flag.Args()) // Ayrıştırılmamış kalan argümanları ekrana yazdırır
}
