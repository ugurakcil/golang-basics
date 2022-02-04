package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"runtime"
	"runtime/debug"
	"time"
)

func saveLog(message interface{}) {
	fmt.Println(message)
}

func openFile(filename interface{}) {
	defer saveLog("Dosya açma işlemi loglandı")

	if _, ok := filename.(string); !ok {
		fmt.Println("Dosya adı sayısal değer olamaz : ", filename, ok)
		return
	}

	file, err := os.Open(filename.(string)) // fmt.Sprintf("%v", filename)

	if err != nil {
		if pathError, ok := err.(*os.PathError); ok {
			fmt.Println("Hatalı path : ", pathError.Path)
			return
		}

		fmt.Println("Dosya açılamadı")
		return
	}

	fmt.Println("Dosya bulundu : ", file.Name())
}

type InputError struct {
	err error
}

func (e *InputError) Error() string {
	if reflect.TypeOf(e.err).String() == "*errors.errorString" {
		if e.err.Error() == "expected integer" {
			return "Sadece sayı girmeniz gerekmektedir"
		}
	}

	return e.err.Error()
}

func guessingGame() (string, error) {
	defer saveLog("Tahmin oyunu çalıştırıldı")

	var number int
	randNumber := rand.Intn(9) + 1 // 1-10 (max-min)+min

	fmt.Print("1 ile 10 arasında bir sayı giriniz : ")
	_, err := fmt.Scanf("%d", &number)

	if err != nil {
		return "", &InputError{err}
	}

	if number < 1 || number > 10 {
		return "", errors.New("Girdiğiniz sayı 1-10 arasında değil")
	}

	fmt.Println("Girdiğiniz sayı :", number)
	fmt.Println("Gelen rakam : ", randNumber)

	if number == randNumber {
		return "Doğru tahmin ettiniz. Tebrikler!", nil
	} else {
		return "Yanlış tahmin. Tekrar deneyin", nil
	}

	return "", nil
}

func main() {
	rand.Seed(time.Now().UnixNano() + 666) // global configuration for rand.Intn

	openFile("test.txt")

	guessingMessage, err := guessingGame()

	if err != nil {
		fmt.Println("HATA : ", err)
	} else {
		fmt.Println(guessingMessage)
	}
}

/*
*
*
*
*
* */

func foo(fn func()) {
	defer func() {
		if e := recover(); e != nil {
			switch e := e.(type) {
			case string:
				fmt.Println("recovered (string) panic:", e)
			case runtime.Error:
				fmt.Println("recovered (runtime.Error) panic:", e.Error())
			case error:
				fmt.Println("recovered (error) panic:", e.Error())
			default:
				fmt.Println("recovered (default) panic:", e)
			}
			fmt.Println(string(debug.Stack()))
			return
		}
		fmt.Println("no panic recovered")
	}()
	fn()
}
