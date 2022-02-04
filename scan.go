package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func FullScan() (output string, err error) {
	reader := bufio.NewReader(os.Stdin)
	output, err = reader.ReadString('\n')
	return output, err
}

func ScanJustInt() (output int, err error) {
	var input string
	var joinedInput interface{}

	input, _ = FullScan()

	joinedInput = strings.Join(strings.FieldsFunc(input, func(c rune) bool {
		return !unicode.IsNumber(c)
	}), "")

	output, _ = strconv.Atoi(joinedInput.(string))

	return output, err
}

func main() {
	fmt.Println("Write hello;")
	var sayHi string
	fmt.Scanln(&sayHi)

	fmt.Printf("T: %T - d: %d - s: %s - v: %v\n", sayHi, sayHi, sayHi, sayHi)

	fmt.Println("Write how are u;")
	howAreU, _ := FullScan()

	fmt.Printf("T: %T - d: %d - s: %s - v: %v\n", howAreU, howAreU, howAreU, howAreU)

	fmt.Println("Write text with a number;")
	foundNumber, _ := ScanJustInt()

	fmt.Printf("T: %T - d: %d - s: %s - v: %v\n", foundNumber, foundNumber, foundNumber, foundNumber)
}
