package main

import (
	"fmt" // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"  // os paketi, işletim sistemi işlevlerine erişim sağlar
)

func cla() {

	// Tüm komut satırı argümanlarını içeren bir dilim oluşturur (program ismi dahil)
	argsWithProg := os.Args

	// Program ismi hariç, sadece komut satırı argümanlarını içeren bir dilim oluşturur
	argsWithoutProg := os.Args[1:]

	// Üçüncü komut satırı argümanını alır (sıfırdan başlayarak sayılır)
	arg := os.Args[3]

	// Tüm komut satırı argümanlarını ekrana yazdırır (program ismi dahil)
	fmt.Println(argsWithProg)

	// Program ismi hariç, sadece komut satırı argümanlarını ekrana yazdırır
	fmt.Println(argsWithoutProg)

	// Üçüncü komut satırı argümanını ekrana yazdırır
	fmt.Println(arg)
}
