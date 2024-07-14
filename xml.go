package main

import (
	"encoding/xml"
	"fmt"
)

// Plant struct, XMLName alanı ile XML element adını belirtir
type Plantx struct {
	XMLName xml.Name `xml:"plant"`   // XML element adı
	Id      int      `xml:"id,attr"` // id alanı, XML attributu olarak tanımlanır
	Name    string   `xml:"name"`    // name alanı
	Origin  []string `xml:"origin"`  // origin alanı, birden fazla değer içerebilir
}

// String metodu, Plant yapısının string temsilini döndürür
func (p Plantx) String() string {
	return fmt.Sprintf("Plant id=%v, name=%v, origin=%v",
		p.Id, p.Name, p.Origin)
}

func xmlparsing() {
	// Coffee bitkisini oluşturur
	coffee := &Plantx{Id: 27, Name: "Coffee"}
	coffee.Origin = []string{"Ethiopia", "Brazil"} // Origin alanını doldurur

	// XML'e kodlar ve düzenli bir şekilde çıktıyı alır
	out, _ := xml.MarshalIndent(coffee, " ", "  ")
	fmt.Println(string(out))

	// XML verisinin başına standart XML başlığını ekler
	fmt.Println(xml.Header + string(out))

	// XML verisini geri çözümleyerek Plant yapısını oluşturur
	var p Plantx
	if err := xml.Unmarshal(out, &p); err != nil {
		panic(err)
	}
	fmt.Println(p) // Geri çözümlenen bitki yapısını yazdırır

	// Domates bitkisini oluşturur
	tomato := &Plantx{Id: 81, Name: "Tomato"}
	tomato.Origin = []string{"Mexico", "California"} // Origin alanını doldurur

	// Birden fazla bitkiyi içeren bir yapı oluşturur
	type Nesting struct {
		XMLName xml.Name  `xml:"nesting"`            // XML element adı
		Plants  []*Plantx `xml:"parent>child>plant"` // Plants alanı, içinde birden fazla bitki yapısı içerir
	}

	nesting := &Nesting{}
	nesting.Plants = []*Plantx{coffee, tomato} // Nesting yapısını doldurur

	// XML'e kodlar ve düzenli bir şekilde çıktıyı alır
	out, _ = xml.MarshalIndent(nesting, " ", "  ")
	fmt.Println(string(out))
}
