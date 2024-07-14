package main

import (
	"fmt"
	s "strings" // Paketi `s` olarak kısaltılmış olarak import ediyoruz
)

var p = fmt.Println // fmt.Println fonksiyonunu `p` olarak kısaltıyoruz

func strff() {

	// `Contains` fonksiyonu: bir stringin başka bir stringi içerip içermediğini kontrol eder
	p("Contains:  ", s.Contains("test", "es"))

	// `Count` fonksiyonu: bir stringin içinde belirtilen alt stringin kaç kez geçtiğini sayar
	p("Count:     ", s.Count("test", "t"))

	// `HasPrefix` fonksiyonu: bir stringin belirtilen prefix ile başlayıp başlamadığını kontrol eder
	p("HasPrefix: ", s.HasPrefix("test", "te"))

	// `HasSuffix` fonksiyonu: bir stringin belirtilen suffix ile bitip bitmediğini kontrol eder
	p("HasSuffix: ", s.HasSuffix("test", "st"))

	// `Index` fonksiyonu: bir string içinde belirtilen alt stringin ilk indeksini döner
	p("Index:     ", s.Index("test", "e"))

	// `Join` fonksiyonu: bir string slice'ını belirtilen ayraç ile birleştirir
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))

	// `Repeat` fonksiyonu: belirtilen stringi belirtilen sayıda tekrar eder
	p("Repeat:    ", s.Repeat("a", 5))

	// `Replace` fonksiyonu: bir string içinde belirtilen alt stringleri belirtilen yeni string ile değiştirir
	// `-1` değeri tüm alt stringleri değiştirir, `1` değeri ise sadece ilk alt stringi değiştirir
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))

	// `Split` fonksiyonu: bir stringi belirtilen ayraçla böler ve bir string slice olarak döner
	p("Split:     ", s.Split("a-b-c-d-e", "-"))

	// `ToLower` fonksiyonu: bir stringin tüm karakterlerini küçük harfe çevirir
	p("ToLower:   ", s.ToLower("TEST"))

	// `ToUpper` fonksiyonu: bir stringin tüm karakterlerini büyük harfe çevirir
	p("ToUpper:   ", s.ToUpper("test"))
}
