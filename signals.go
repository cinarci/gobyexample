package main

import (
	"fmt"       // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"        // os paketi, işletim sistemi işlevlerine erişim sağlar
	"os/signal" // os/signal paketi, gelen sinyalleri işlemeye yarar
	"syscall"   // syscall paketi, düşük seviyeli sistem çağrılarını içerir
)

func signals() {

	sigs := make(chan os.Signal, 1) // Bir sinyal kanalı oluşturulur, kapasitesi 1

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM) // Bu kanal, SIGINT ve SIGTERM sinyalleri ile bilgilendirilir

	done := make(chan bool, 1) // İşlemin tamamlandığını bildirmek için bir boolean kanalı oluşturulur, kapasitesi 1

	go func() { // Bir goroutine başlatılır

		sig := <-sigs    // Kanalda bir sinyal alana kadar bekler
		fmt.Println()    // Yeni bir satır ekler
		fmt.Println(sig) // Alınan sinyali yazdırır
		done <- true     // İşlemin tamamlandığını bildirmek için done kanalına true değeri gönderir
	}()

	fmt.Println("awaiting signal") // "awaiting signal" mesajını yazdırır
	<-done                         // done kanalında bir değer alana kadar bekler
	fmt.Println("exiting")         // "exiting" mesajını yazdırır
}
