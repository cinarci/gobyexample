package main

import (
	"bufio"    // bufio paketi, buffered I/O işlemleri için kullanılır
	"fmt"      // fmt paketi, formatlı I/O işlemleri için kullanılır
	"net/http" // net/http paketi, HTTP istemcisi ve sunucusu işlevleri sağlar
)

func httpclient() {

	// Belirtilen URL'ye bir GET isteği gönderir
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err) // Hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}
	defer resp.Body.Close() // Fonksiyon sona erdiğinde HTTP yanıt gövdesini kapatır

	// HTTP yanıt durum kodunu yazdırır
	fmt.Println("Response status:", resp.Status)

	// HTTP yanıt gövdesini satır satır okumak için bir scanner oluşturur
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ { // İlk 5 satırı okur
		fmt.Println(scanner.Text()) // Okunan her satırı yazdırır
	}

	// Scanner kullanırken oluşan hataları kontrol eder
	if err := scanner.Err(); err != nil {
		panic(err) // Hata varsa, program panik durumuna geçer ve hata mesajını gösterir
	}
}

/*
Bu kod, basit bir HTTP sunucusu oluşturur ve iki farklı yolu ("/hello" ve "/headers") işler:

hello fonksiyonu:

/hello yoluna gelen istekleri işler.
Yanıt olarak "hello" metnini gönderir.
headers fonksiyonu:

/headers yoluna gelen istekleri işler.
İstek başlıklarını (headers) alır ve her başlığı ad-değer çiftleri olarak yanıt olarak gönderir.
main fonksiyonu:

http.HandleFunc("/hello", hello) ile /hello yolunu hello fonksiyonuna yönlendirir.
http.HandleFunc("/headers", headers) ile /headers yolunu headers fonksiyonuna yönlendirir.
http.ListenAndServe(":8090", nil) ile 8090 portunda HTTP sunucusunu başlatır. Bu sunucu, yukarıdaki iki yolu dinler ve gelen istekleri ilgili handler fonksiyonlarına yönlendirir.
Bu kodu çalıştırdığınızda, tarayıcı veya curl gibi bir araçla http://localhost:8090/hello ve http://localhost:8090/headers URL'lerine istek göndererek sonuçları görebilirsiniz.
*/
