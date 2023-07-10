package main

import "fmt"

func main() {

	// Switch Case Statement memiliki kegunaan yg sama dengan if else, perbedaan terletak dicara penulisannya yg berbeda
	bahasaPemrograman := "Golang"

	switch bahasaPemrograman {
	case "JavaScript":
		fmt.Println("Kamu sedang belajar bahasa pemrograman JavaScript")
	case "Python":
		fmt.Println("Kamu sedang belajar bahasa pemrograman Python")
	case "Java":
		fmt.Println("Kamu sedang belajar bahasa pemrograman Java")
	case "Golang", "Go":
		fmt.Println("Kamu sedang belajar bahasa pemrograman Golang")
	default:
		fmt.Println("Kamu sedang belajar bahasa pemrograman lain")
	}

	// switch case dengan deklarasi variable langsung, "true" disini untuk membandingkan boolean expressions yang menghasilkan nilai true. Kita bisa menghapus "true" tersebut karena akan direturn secara default
	switch n := 10; true {
	case n > 0:
		fmt.Println("Positive")
	case n < 0:
		fmt.Println("Negative")
	default:
		fmt.Println("Zero")
	}
}
