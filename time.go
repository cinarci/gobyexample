package main

import (
	"fmt"
	"time"
)

func timef() {
	p := fmt.Println

	// Şu anki zamanı alır
	now := time.Now()
	p(now)

	// Belirli bir tarihi oluşturur (2009-11-17 20:34:58.651387237 UTC)
	then := time.Date(
		2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	p(then)

	// Tarihin yıl, ay, gün, saat, dakika, saniye ve nanosaniye cinsinden değerlerini yazdırır
	p(then.Year())
	p(then.Month())
	p(then.Day())
	p(then.Hour())
	p(then.Minute())
	p(then.Second())
	p(then.Nanosecond())

	// Tarihin bulunduğu zaman dilimini yazdırır
	p(then.Location())

	// Tarihin haftanın hangi gününe denk geldiğini yazdırır (Pazartesi, Salı, vs.)
	p(then.Weekday())

	// Bir tarihin başka bir tarihten önce, sonra veya eşit mi olduğunu kontrol eder
	p(then.Before(now))
	p(then.After(now))
	p(then.Equal(now))

	// İki zaman arasındaki farkı hesaplar
	diff := now.Sub(then)
	p(diff)

	// Farkı saat, dakika, saniye ve nanosaniye cinsinden yazdırır
	p(diff.Hours())
	p(diff.Minutes())
	p(diff.Seconds())
	p(diff.Nanoseconds())

	// Bir tarihe belirli bir zaman aralığı ekler veya çıkarır
	p(then.Add(diff))
	p(then.Add(-diff))
}
