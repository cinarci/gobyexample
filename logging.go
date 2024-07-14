package main

import (
	"bytes" // bytes paketi, byte slice'ları manipüle etmek için kullanılır
	"fmt"   // fmt paketi, formatlı I/O işlemleri için kullanılır
	"log"   // log paketi, logging işlevleri sağlar
	"os"    // os paketi, işletim sistemi işlevlerine erişim sağlar

	"log/slog" // log/slog paketi, yapılandırılabilir ve yapılandırılmış logging işlevleri sağlar
)

func logging() {

	log.Println("standard logger") // Varsayılan logger ile standart bir log mesajı yazdırır

	log.SetFlags(log.LstdFlags | log.Lmicroseconds) // Logger'ın bayraklarını, standart bayraklar ve mikro saniyelerle ayarlar
	log.Println("with micro")                       // Mikro saniyelerle log mesajı yazdırır

	log.SetFlags(log.LstdFlags | log.Lshortfile) // Logger'ın bayraklarını, standart bayraklar ve kısa dosya adı ve satır numarası ile ayarlar
	log.Println("with file/line")                // Dosya adı ve satır numarası ile log mesajı yazdırır

	mylog := log.New(os.Stdout, "my:", log.LstdFlags) // Yeni bir logger oluşturur ve standart bayraklarla stdout'a yazar
	mylog.Println("from mylog")                       // Yeni logger ile log mesajı yazdırır

	mylog.SetPrefix("ohmy:")    // Logger'ın prefix'ini değiştirir
	mylog.Println("from mylog") // Yeni prefix ile log mesajı yazdırır

	var buf bytes.Buffer                           // Bir byte buffer oluşturur
	buflog := log.New(&buf, "buf:", log.LstdFlags) // Buffer'a yazacak bir logger oluşturur

	buflog.Println("hello") // Buffer'a log mesajı yazar

	fmt.Print("from buflog:", buf.String()) // Buffer'daki içeriği yazdırır

	jsonHandler := slog.NewJSONHandler(os.Stderr, nil) // stderr'e JSON formatında log yazacak bir handler oluşturur
	myslog := slog.New(jsonHandler)                    // Bu handler ile yeni bir structured logger oluşturur
	myslog.Info("hi there")                            // Structured logger ile bir info mesajı yazdırır

	myslog.Info("hello again", "key", "val", "age", 25) // Structured logger ile anahtar-değer çiftleri içeren bir info mesajı yazdırır
}

/* Bu kod, Go'da logging işlemlerini nasıl yapılandırabileceğinizi ve kullanabileceğinizi gösterir. İşte adım adım açıklaması:

log.Println("standard logger") - Varsayılan logger ile standart bir log mesajı yazdırır.
log.SetFlags(log.LstdFlags | log.Lmicroseconds) - Logger'ın bayraklarını standart bayraklar ve mikro saniyelerle ayarlar.
log.Println("with micro") - Mikro saniyelerle log mesajı yazdırır.
log.SetFlags(log.LstdFlags | log.Lshortfile) - Logger'ın bayraklarını standart bayraklar ve kısa dosya adı ve satır numarası ile ayarlar.
log.Println("with file/line") - Dosya adı ve satır numarası ile log mesajı yazdırır.
mylog := log.New(os.Stdout, "my:", log.LstdFlags) - Yeni bir logger oluşturur ve standart bayraklarla stdout'a yazar.
mylog.Println("from mylog") - Yeni logger ile log mesajı yazdırır.
mylog.SetPrefix("ohmy:") - Logger'ın prefix'ini değiştirir.
mylog.Println("from mylog") - Yeni prefix ile log mesajı yazdırır.
var buf bytes.Buffer - Bir byte buffer oluşturur.
buflog := log.New(&buf, "buf:", log.LstdFlags) - Buffer'a yazacak bir logger oluşturur.
buflog.Println("hello") - Buffer'a log mesajı yazar.
fmt.Print("from buflog:", buf.String()) - Buffer'daki içeriği yazdırır.
jsonHandler := slog.NewJSONHandler(os.Stderr, nil) - stderr'e JSON formatında log yazacak bir handler oluşturur.
myslog := slog.New(jsonHandler) - Bu handler ile yeni bir structured logger oluşturur.
myslog.Info("hi there") - Structured logger ile bir info mesajı yazdırır.
myslog.Info("hello again", "key", "val", "age", 25) - Structured logger ile anahtar-değer çiftleri içeren bir info mesajı yazdırır.
Bu kod, farklı log formatları, prefix'ler ve yapılandırılmış log mesajları ile nasıl çalışılacağını göstermektedir.
*/
