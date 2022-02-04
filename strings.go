package main

import (
	"fmt"
	"html"
	"io/ioutil"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.Contains("Uğur AKÇIL", "ÇIL"))
	fmt.Println(strings.ContainsAny("Uğur AKÇIL", "Jğ"))
	fmt.Println(strings.Count("Uğur AKÇIL", "Ç"))
	fmt.Println(strings.Index("Uğur AKÇIL", "Ç")) // + IndexAny + IndexByte + LastIndex
	fmt.Println(readFile("demo1.txt"))

	greetings := []string{"hello", "world", "wide", "web"}
	fmt.Println(strings.Join(greetings, "-"))

	fmt.Printf("Fields are: %q \n", strings.Fields("  Uğur AKÇIL  is A  ++ Developer    "))

	fieldsFuncDemoFunc := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}

	fmt.Printf("FieldsFunc : %q \n", strings.FieldsFunc(" Uğ*ur /:.. 2AKÇIL ", fieldsFuncDemoFunc))

	fmt.Println(strings.HasPrefix("Gopher", "Go")) // + HasSuffix

	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber)) // + IndexFunc

	fmt.Println("ba" + strings.Repeat("na", 2))

	hackStyle := func(r rune) rune {
		switch r {
		case 'A':
			return '4'
		case 'I':
			return '1'
		case 'İ':
			return '1'
		case 'E':
			return '3'
		case 'O':
			return '0'
		case 'Z':
			return '7'
		case 'Q':
			return '9'
		case 'B':
			return '8'
		case 'G':
			return '6'
		case 'Ğ':
			return '6'
		case 'R':
			return '2'
		default:
			return r
		}
	}

	fmt.Println(strings.Map(hackStyle, "UĞUR AKÇIL"))

	fmt.Println(strings.Replace("Uğur AKÇIL is a full stack developer, full stack alien", "full stack", "golang", -1)) // + ReplaceAll

	fmt.Printf("%q\n", strings.Split("Uğur AKÇIL, Gökçe AKÇIL, Defne AKÇIL", ",")) // + SplitAfter

	fmt.Println(strings.Title("uğur akçıl is a developer ~since 2005"))

	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "UĞUR AKÇIL BUGÜN ÇELİŞKİLİ DEĞİL")) // + ToTitle & ToUpper

	fmt.Println(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})) // + TrimLeft + TrimLeftFunc

	fmt.Println(strings.TrimSpace(" \t\n Hello, Gophers \n\t\r\n"))

	fmt.Println(html.EscapeString(`"Fran & Freddie's Diner" <tasty@example.com>`)) // + UnescapeString
}

func readFile(fileFullPath string) string {
	bs, err := ioutil.ReadFile(fileFullPath)

	if err != nil {
		return "Dosya bulunamadı"
	}

	str := string(bs)

	return str
}
