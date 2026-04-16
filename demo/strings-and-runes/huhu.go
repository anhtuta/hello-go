package stringsandrunes

import (
	"fmt"
	"unicode/utf8"
)

// https://gobyexample.com/strings-and-runes
// A Go string is a read-only slice of bytes. The language and the standard library treat strings specially -
// as containers of text encoded in UTF-8. In other languages, strings are made of “characters”.
// In Go, the concept of a character is called a rune - it’s an integer that represents a Unicode code point
func RuneDemo() {
	fmt.Print("\n======= Strings and Runes =======\n\n")

	// const s = "สวัสดี"
	const s = "单身情歌 - 林志炫"

	// strings are equivalent to []byte, this will produce the length of the raw bytes stored within
	// Mặc dù s có 10 ký tự (10 rune), nhưng có len(s) = 24 byte, vì các ký tự Chinese được mã hóa bằng nhiều byte trong UTF-8
	fmt.Println("Len:", len(s), "bytes")

	// Indexing into a string produces the raw byte values at each index.
	// This loop generates the hex values of all the bytes that constitute the code points in s.
	// Dùng for i theo len của s sẽ trả về byte, không phải rune
	fmt.Print("Bytes (Hex): ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%3x ", s[i]) // byte được in ra ở dạng hex
	}
	fmt.Print("\nBytes (Dec): ")
	for i := 0; i < len(s); i++ {
		// fmt.Print(s[i], " ") // byte được in ra ở dạng decimal
		fmt.Printf("%3d ", s[i]) // byte được in ra ở dạng decimal với 3 chữ số
	}
	fmt.Println()

	// Mỗi 1 từ/ký tự tiếng Chinese = 1 rune
	fmt.Println("Rune count:", utf8.RuneCountInString(s))

	// A range loop handles strings specially and decodes each rune along with its offset in the string.
	// For range sẽ trả về index của byte đầu tiên của rune và giá trị rune đó, không phải byte
	for idx, runeValue := range s {
		fmt.Printf("%#U starts at %d\n", runeValue, idx)
	}

	// We can achieve the same iteration by using the utf8.DecodeRuneInString function explicitly
	fmt.Println("\nUsing DecodeRuneInString")
	for i, w := 0, 0; i < len(s); i += w {
		runeValue, width := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%#U starts at %d\n", runeValue, i)
		w = width

		examineRune(runeValue)
	}
}

func examineRune(r rune) {
	// Values enclosed in single quotes are rune literals. We can compare a rune value to a rune literal directly
	if r == 't' {
		fmt.Println("found tee")
	} else if r == '林' { // cannot compare a rune to a string literal "林"
		fmt.Println("found forest")
	}
}
