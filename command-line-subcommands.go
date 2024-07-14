package main

import (
	"flag" // flag paketi, komut satırı bayraklarını (flags) işlemek için kullanılır
	"fmt"  // fmt paketi, formatlı I/O işlemleri için kullanılır
	"os"   // os paketi, işletim sistemi işlevlerine erişim sağlar
)

func clsc() {

	// 'foo' isimli alt komut için bir FlagSet oluşturur
	fooCmd := flag.NewFlagSet("foo", flag.ExitOnError)
	// 'foo' alt komutu için bir boolean bayrak (flag) tanımlar
	fooEnable := fooCmd.Bool("enable", false, "enable")
	// 'foo' alt komutu için bir string bayrak (flag) tanımlar
	fooName := fooCmd.String("name", "", "name")

	// 'bar' isimli alt komut için bir FlagSet oluşturur
	barCmd := flag.NewFlagSet("bar", flag.ExitOnError)
	// 'bar' alt komutu için bir integer bayrak (flag) tanımlar
	barLevel := barCmd.Int("level", 0, "level")

	// Eğer komut satırı argümanları yeterince değilse bir hata mesajı verir
	if len(os.Args) < 2 {
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}

	// Komut satırı argümanlarının ikinci elemanına (alt komut) göre işlem yapar
	switch os.Args[1] {

	case "foo": // 'foo' alt komutu için
		fooCmd.Parse(os.Args[2:]) // 'foo' komutunun bayraklarını ayrıştırır
		fmt.Println("subcommand 'foo'")
		fmt.Println("  enable:", *fooEnable)  // 'enable' bayrağının değerini yazdırır
		fmt.Println("  name:", *fooName)      // 'name' bayrağının değerini yazdırır
		fmt.Println("  tail:", fooCmd.Args()) // Kalan argümanları yazdırır

	case "bar": // 'bar' alt komutu için
		barCmd.Parse(os.Args[2:]) // 'bar' komutunun bayraklarını ayrıştırır
		fmt.Println("subcommand 'bar'")
		fmt.Println("  level:", *barLevel)    // 'level' bayrağının değerini yazdırır
		fmt.Println("  tail:", barCmd.Args()) // Kalan argümanları yazdırır

	default: // Beklenmeyen bir alt komut girildiğinde
		fmt.Println("expected 'foo' or 'bar' subcommands")
		os.Exit(1)
	}
}
