package main

import (
	"os"      // os paketi, işletim sistemi işlevlerine erişim sağlar
	"os/exec" // os/exec paketi, harici komutları çalıştırmak için kullanılır
	"syscall" // syscall paketi, düşük seviyeli sistem çağrılarını içerir
)

func execingprocess() {

	binary, lookErr := exec.LookPath("ls") // 'ls' komutunun binary dosyasının yolunu bulur
	if lookErr != nil {
		panic(lookErr) // Eğer hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}

	args := []string{"ls", "-a", "-l", "-h"} // 'ls' komutuna verilecek argümanları tanımlar

	env := os.Environ() // Mevcut ortam değişkenlerini alır

	execErr := syscall.Exec(binary, args, env) // Mevcut süreci 'ls' komutunu çalıştıracak şekilde değiştirir
	if execErr != nil {
		panic(execErr) // Eğer hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}
}
