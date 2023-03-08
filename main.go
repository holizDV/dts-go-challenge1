package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const age int = 21
	const cr string = "%"
	const j bool = 10 > 5
	const unicode string = "\u042F (ya)"
	const base8 int8 = 20
	const base10 int8 = 21
	const base16 int16 = 15
	const bases16 int16 = 13
	unicodeChar, _ := utf8.DecodeRuneInString("Я")
	const k float64 = 123.456

	fmt.Printf("nilai : %d\n", age)
	fmt.Printf("tipe data : %T\n", age)
	fmt.Printf("character : %s\n", cr)
	fmt.Printf("boolean : %v\n", j)
	fmt.Printf("unicode russian : %s\n", unicode)
	fmt.Printf("base8 : %o\n", base8)
	fmt.Printf("base10 : %d\n", base10)
	fmt.Printf("base16 : %x\n", base16)
	fmt.Printf("bases16 : %X %d\n", base16, bases16)
	fmt.Printf("unicode character Я : %#U\n", unicodeChar)
	fmt.Printf("float64 : %g\n", k)
	fmt.Printf("float : %f\n", k)
	fmt.Printf("float scientific : %e\n", k)
}
