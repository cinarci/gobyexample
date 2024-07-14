package main

import (
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"      // os paketi, işletim sistemi işlevlerine erişim sağlar
	"strings" // strings paketi, string işlemleri için kullanılır
)

func envvariables() {

	os.Setenv("FOO", "1")                 // FOO isimli bir ortam değişkeni oluşturur ve değerini 1 olarak ayarlar
	fmt.Println("FOO:", os.Getenv("FOO")) // FOO ortam değişkeninin değerini ekrana yazdırır
	fmt.Println("BAR:", os.Getenv("BAR")) // BAR ortam değişkeninin değerini ekrana yazdırır (eğer yoksa boş string döner)

	fmt.Println()
	for _, e := range os.Environ() { // Tüm ortam değişkenlerini alır ve döngü ile gezer
		pair := strings.SplitN(e, "=", 2) // Her bir ortam değişkenini 'key=value' formatında ayırır, en fazla iki parçaya böler
		fmt.Println(pair[0])              // Ortam değişkeninin anahtarını ekrana yazdırır
	}
}
