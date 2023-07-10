package main

import "fmt"

// variable digunakan untuk menyimpan sebuah data dalam pemrograman dan data tersebut bisa diganti

func main() {
	// 1. Deklarasi variable dengan menggunakan keyword var
	var x string
	x = "Belajar Golang!"
	fmt.Println(x)

	// 2. Deklarasi variable tanpa tipe data (type inference)
	var nama = "Budi" // tipe data string
	umur := 25        // tipe data int
	fmt.Println(nama, umur)

	//3.  Deklarasi variable dengan multiple assignments
	var angka1, angka2 int
	angka1, angka2 = 50, 100

	// swapping 2 variable dengan menggunakan multiple assignments
	angka1, angka2 = angka2, angka1
	fmt.Println(angka1, angka2)

	// 4. Multipe short declarations
	hewan, jumlah := "kucing", 2
	fmt.Println(hewan, jumlah)

	// 5. Deklarasi dengan menggunakan const
	var handphone = "samsung"
	handphone = "xiaomi"
	fmt.Println(handphone)

	const laptop = "asus"
	fmt.Println(laptop)

}
