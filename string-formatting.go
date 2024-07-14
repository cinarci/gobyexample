package main

import (
	"fmt"
	"os"
)

type point struct {
	x, y int
}

func strf() {

	// Bir struct tipinde örnek oluşturuluyor
	p := point{1, 2}

	// `%v` formatı: struct'ın değerlerini yazdırır
	fmt.Printf("struct1: %v\n", p)

	// `%+v` formatı: struct'ın alan adları ile birlikte değerlerini yazdırır
	fmt.Printf("struct2: %+v\n", p)

	// `%#v` formatı: struct'ın tipini ve alan adları ile birlikte değerlerini yazdırır
	fmt.Printf("struct3: %#v\n", p)

	// `%T` formatı: değişkenin tipini yazdırır
	fmt.Printf("type: %T\n", p)

	// `%t` formatı: boolean değeri yazdırır
	fmt.Printf("bool: %t\n", true)

	// `%d` formatı: ondalık sayıyı (int) yazdırır
	fmt.Printf("int: %d\n", 123)

	// `%b` formatı: binary formatında sayıyı yazdırır
	fmt.Printf("bin: %b\n", 14)

	// `%c` formatı: unicode karakterini yazdırır
	fmt.Printf("char: %c\n", 33)

	// `%x` formatı: hexadecimal formatında sayıyı yazdırır
	fmt.Printf("hex: %x\n", 456)

	// `%f` formatı: float sayıyı yazdırır
	fmt.Printf("float1: %f\n", 78.9)

	// `%e` formatı: scientific notation (exponential formatında) float sayıyı yazdırır
	fmt.Printf("float2: %e\n", 123400000.0)

	// `%E` formatı: büyük harfli scientific notation float sayıyı yazdırır
	fmt.Printf("float3: %E\n", 123400000.0)

	// `%s` formatı: string'i yazdırır
	fmt.Printf("str1: %s\n", "\"string\"")

	// `%q` formatı: double-quoted string'i (string içinde tırnak işaretleri ile) yazdırır
	fmt.Printf("str2: %q\n", "\"string\"")

	// `%x` formatı: string'i hexadecimal formatında yazdırır
	fmt.Printf("str3: %x\n", "hex this")

	// `%p` formatı: pointer'ın hexadecimal adresini yazdırır
	fmt.Printf("pointer: %p\n", &p)

	// `%6d` formatı: 6 karakter genişliğinde sağa hizalanmış integer sayıyı yazdırır
	fmt.Printf("width1: |%6d|%6d|\n", 12, 345)

	// `%6.2f` formatı: 6 karakter genişliğinde ve 2 ondalık hane ile float sayıyı yazdırır
	fmt.Printf("width2: |%6.2f|%6.2f|\n", 1.2, 3.45)

	// `%-6.2f` formatı: 6 karakter genişliğinde ve 2 ondalık hane ile sola hizalanmış float sayıyı yazdırır
	fmt.Printf("width3: |%-6.2f|%-6.2f|\n", 1.2, 3.45)

	// `%6s` formatı: 6 karakter genişliğinde sağa hizalanmış string'i yazdırır
	fmt.Printf("width4: |%6s|%6s|\n", "foo", "b")

	// `%-6s` formatı: 6 karakter genişliğinde sola hizalanmış string'i yazdırır
	fmt.Printf("width5: |%-6s|%-6s|\n", "foo", "b")

	// `Sprintf` fonksiyonu: formatlı string oluşturur ve bu string'i değişkene atar
	s := fmt.Sprintf("sprintf: a %s", "string")
	fmt.Println(s)

	// `Fprintf` fonksiyonu: formatlı string'i belirtilen `io.Writer` nesnesine yazar
	fmt.Fprintf(os.Stderr, "io: an %s\n", "error")
}
