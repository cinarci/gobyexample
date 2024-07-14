package main

import (
	"bufio"   // bufio paketi, veriyi satır satır okumak için kullanılır
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"      // os paketi, işletim sistemi işlevlerine erişim sağlar
	"strings" // strings paketi, string işlemleri için kullanılır
)

func linefilters() {
	// bufio.NewScanner(os.Stdin) ile standart girdiyi (os.Stdin) okuyacak bir tarayıcı (scanner) oluşturulur
	scanner := bufio.NewScanner(os.Stdin)

	// Tarayıcı (scanner) her bir satır için Scan() fonksiyonuyla çalıştırılır
	for scanner.Scan() {
		// Satırın metnini alır ve büyük harfe dönüştürür
		ucl := strings.ToUpper(scanner.Text())

		// Büyük harfe dönüştürülmüş satırı ekrana yazdırır
		fmt.Println(ucl)
	}

	// scanner.Err() ile tarama sırasında bir hata olup olmadığı kontrol edilir
	if err := scanner.Err(); err != nil {
		// Hata varsa hata mesajını standart hata çıkışına (stderr) yazdırır ve programı hata kodu ile sonlandırır
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}
