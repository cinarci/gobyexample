package main

import (
	"fmt" // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"  // os paketi, işletim sistemi işlevlerine erişim sağlar
)

func exit() {
	defer fmt.Println("!") // defer ifadesi, bu satırı fonksiyon sona erdikten sonra çalıştırılmak üzere erteler

	os.Exit(3) // os.Exit fonksiyonu programı belirtilen çıkış kodu ile sonlandırır, bu durumda çıkış kodu 3
	// Dikkat: os.Exit(3) çağrısı defer edilmiş fonksiyonları çalıştırmaz.
}
