package main

import (
	"fmt"      // fmt paketi, formatlı I/O işlemleri için kullanılır
	"net/http" // net/http paketi, HTTP sunucusu ve istemcisi işlevleri sağlar
)

func hello(w http.ResponseWriter, req *http.Request) {
	// Bu handler, /hello yoluna gelen HTTP isteklerini işler
	fmt.Fprintf(w, "hello\n") // HTTP yanıtına "hello" metnini yazar
}

func headers(w http.ResponseWriter, req *http.Request) {
	// Bu handler, /headers yoluna gelen HTTP isteklerini işler
	for name, headers := range req.Header { // HTTP isteğinin başlıklarını döngüyle gezer
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h) // Her başlık adını ve değerini HTTP yanıtına yazar
		}
	}
}

func httpserver() {
	http.HandleFunc("/hello", hello)     // /hello yoluna gelen istekleri hello fonksiyonu ile işler
	http.HandleFunc("/headers", headers) // /headers yoluna gelen istekleri headers fonksiyonu ile işler

	http.ListenAndServe(":8090", nil) // 8090 portunda HTTP sunucusunu başlatır
	// nil parametresi, varsayılan HTTP multiplexer'ın kullanılacağını belirtir
}
