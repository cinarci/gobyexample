package main

import (
	"fmt"
	"time"
)

func timeformatting() {
	p := fmt.Println

	t := time.Now()
	p(t.Format(time.RFC3339)) // Şu anki zamanı RFC3339 formatında biçimlendirir

	t1, e := time.Parse(
		time.RFC3339,
		"2012-11-01T22:08:41+00:00")
	p(t1) // Verilen zaman dizisini RFC3339 formatında ayrıştırır

	p(t.Format("3:04PM"))                           // Şu anki zamanı özel bir saat formatında biçimlendirir
	p(t.Format("Mon Jan _2 15:04:05 2006"))         // Şu anki zamanı başka bir özel formatla biçimlendirir
	p(t.Format("2006-01-02T15:04:05.999999-07:00")) // Şu anki zamanı belirli bir düzende biçimlendirir

	form := "3 04 PM"
	t2, e := time.Parse(form, "8 41 PM") // Özel bir formata ("3 04 PM") göre zaman dizisini ayrıştırır
	p(t2)

	// Şu anki zamanı belirli bir düzende elle biçimlendirerek yazdırır
	fmt.Printf("%d-%02d-%02dT%02d:%02d:%02d-00:00\n",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	ansic := "Mon Jan _2 15:04:05 2006"
	_, e = time.Parse(ansic, "8:41PM") // Zaman dizesini "Mon Jan _2 15:04:05 2006" formatında ayrıştırmaya çalışır
	p(e)                               // Ayrıştırma sırasında oluşabilecek hataları yazdırır
}
