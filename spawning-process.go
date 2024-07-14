package main

import (
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"io"      // io paketi, temel I/O işlevleri için kullanılır
	"os/exec" // os/exec paketi, harici komutları çalıştırmak için kullanılır
)

func spawningprocess() {

	dateCmd := exec.Command("date") // "date" komutunu çalıştıracak bir Command nesnesi oluşturur

	dateOut, err := dateCmd.Output() // "date" komutunu çalıştırır ve çıktısını alır
	if err != nil {
		panic(err) // Eğer bir hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}
	fmt.Println("> date")        // "date" komutunun başlığını yazdırır
	fmt.Println(string(dateOut)) // "date" komutunun çıktısını yazdırır

	_, err = exec.Command("date", "-x").Output() // Yanlış bir argümanla "date" komutunu çalıştırır
	if err != nil {
		switch e := err.(type) { // Hata türüne göre farklı işlemler yapar
		case *exec.Error:
			fmt.Println("failed executing:", err) // Komut çalıştırılamadıysa hata mesajını yazdırır
		case *exec.ExitError:
			fmt.Println("command exit rc =", e.ExitCode()) // Komut çalıştı fakat başarısız olduysa çıkış kodunu yazdırır
		default:
			panic(err) // Diğer hatalar için panik durumuna geçer ve hata mesajını gösterir
		}
	}

	grepCmd := exec.Command("grep", "hello") // "grep hello" komutunu çalıştıracak bir Command nesnesi oluşturur

	grepIn, _ := grepCmd.StdinPipe()                 // "grep" komutuna veri yazmak için bir stdin pipe oluşturur
	grepOut, _ := grepCmd.StdoutPipe()               // "grep" komutunun çıktısını almak için bir stdout pipe oluşturur
	grepCmd.Start()                                  // "grep" komutunu başlatır
	grepIn.Write([]byte("hello grep\ngoodbye grep")) // "grep" komutunun stdin'ine veri yazar
	grepIn.Close()                                   // stdin pipe'ını kapatır
	grepBytes, _ := io.ReadAll(grepOut)              // "grep" komutunun çıktısını okur
	grepCmd.Wait()                                   // "grep" komutunun tamamlanmasını bekler

	fmt.Println("> grep hello")    // "grep hello" komutunun başlığını yazdırır
	fmt.Println(string(grepBytes)) // "grep" komutunun çıktısını yazdırır

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h") // "bash" shell üzerinden "ls -a -l -h" komutunu çalıştıracak bir Command nesnesi oluşturur
	lsOut, err := lsCmd.Output()                       // "ls -a -l -h" komutunu çalıştırır ve çıktısını alır
	if err != nil {
		panic(err) // Eğer bir hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}
	fmt.Println("> ls -a -l -h") // "ls -a -l -h" komutunun başlığını yazdırır
	fmt.Println(string(lsOut))   // "ls -a -l -h" komutunun çıktısını yazdırır
}
