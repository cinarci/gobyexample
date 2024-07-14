package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func regularexpressions() {
	// Verilen metinde "p([a-z]+)ch" düzenli ifadesine uyan bir eşleşme arar
	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")
	fmt.Println(match) // true

	// "p([a-z]+)ch" düzenli ifadesini derler
	r, _ := regexp.Compile("p([a-z]+)ch")

	// Verilen metinde eşleşme arar ve true/false döner
	fmt.Println(r.MatchString("peach")) // true

	// Verilen metinde ilk eşleşmeyi döner
	fmt.Println(r.FindString("peach punch")) // peach

	// Verilen metinde ilk eşleşmenin indekslerini döner
	fmt.Println("idx:", r.FindStringIndex("peach punch")) // [0 5]

	// Verilen metinde tüm eşleşmeyi ve alt eşleşmeleri döner
	fmt.Println(r.FindStringSubmatch("peach punch")) // [peach ea]

	// Verilen metinde tüm eşleşmenin ve alt eşleşmelerin indekslerini döner
	fmt.Println(r.FindStringSubmatchIndex("peach punch")) // [0 5 1 3]

	// Verilen metinde tüm eşleşmeleri döner
	fmt.Println(r.FindAllString("peach punch pinch", -1)) // [peach punch pinch]

	// Verilen metinde tüm eşleşmeleri ve alt eşleşmelerin indekslerini döner
	fmt.Println("all:", r.FindAllStringSubmatchIndex(
		"peach punch pinch", -1)) // [[0 5 1 3] [6 11 7 9] [12 17 13 15]]

	// Verilen metinde belirli sayıda eşleşme döner (2 eşleşme)
	fmt.Println(r.FindAllString("peach punch pinch", 2)) // [peach punch]

	// Verilen byte dizisinde eşleşme arar ve true/false döner
	fmt.Println(r.Match([]byte("peach"))) // true

	// Yeni bir düzenli ifade derler ve r'ye atar
	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println("regexp:", r) // regexp:p([a-z]+)ch

	// Verilen metinde tüm eşleşmeleri değiştirir
	fmt.Println(r.ReplaceAllString("a peach", "<fruit>")) // a <fruit>

	// Verilen byte dizisinde ReplaceAllFunc ile tüm eşleşmeleri büyük harfe çevirir
	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(string(out)) // a PEACH
}
