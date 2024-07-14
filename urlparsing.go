package main

import (
	"fmt"     // fmt paketi, formatlı I/O işlemleri için kullanılır
	"net"     // net paketi, network işlemleri için kullanılır (SplitHostPort)
	"net/url" // net/url paketi, URL işlemleri için kullanılır
)

func urlparsing() {
	s := "postgres://user:pass@host.com:5432/path?k=v#f" // Parçalanacak URL

	u, err := url.Parse(s) // URL'i parçala
	if err != nil {
		panic(err)
	}

	fmt.Println(u.Scheme) // URL'nin şemasını (scheme) ekrana yazdırır ("postgres")

	fmt.Println(u.User)            // Kullanıcı bilgilerini ekrana yazdırır (username ve password)
	fmt.Println(u.User.Username()) // Kullanıcı adını ekrana yazdırır ("user")
	p, _ := u.User.Password()      // Şifreyi alır (hata kontrolü yok, URL'nin doğru olması varsayılmış)
	fmt.Println(p)                 // Şifreyi ekrana yazdırır ("pass")

	fmt.Println(u.Host)                        // Host bilgisini ekrana yazdırır ("host.com:5432")
	host, port, _ := net.SplitHostPort(u.Host) // Host ve port'u ayırır
	fmt.Println(host)                          // Host adını ekrana yazdırır ("host.com")
	fmt.Println(port)                          // Port numarasını ekrana yazdırır ("5432")

	fmt.Println(u.Path)     // Yol bilgisini ekrana yazdırır ("/path")
	fmt.Println(u.Fragment) // Fragment bilgisini ekrana yazdırır ("f")

	fmt.Println(u.RawQuery)            // Raw (ham) sorgu bilgisini ekrana yazdırır ("k=v")
	m, _ := url.ParseQuery(u.RawQuery) // Sorgu parametrelerini ayrıştırır
	fmt.Println(m)                     // Sorgu parametrelerini ekrana yazdırır (map[k:[v]])
	fmt.Println(m["k"][0])             // "k" parametresinin değerini ekrana yazdırır ("v")
}
