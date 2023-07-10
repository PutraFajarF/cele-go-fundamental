package main

import (
	"fmt"
	"math"
)

// Function adalah sebuah code pemrograman yang dibuat untuk mengerjakan suatu tugas tertentu

// cara 1, mendefinisikan sebuah function tanpa parameter
func tugas1() {
	fmt.Println("Ini adalah fungsi tugas1()")
}

// cara 2, mendefinisikan sebuah function dengan 2 parameter yaitu a dan b
func tugas2(a int, b int) {
	fmt.Println("Hasil Jumlah :", a+b)
}

// cara 3, mendefinisikan sebuah function dengan menggunakan shorthand parameter notation
func tugas3(a, b, c int, d, e float64, s string) {
	fmt.Println(a, b, c, d, e, s)
}

// cara 4, mendefisikan sebuah function yang memiliki 1 parameter dan menghasilkan 1 return value
func tugas4(a float64) float64 {
	return math.Pow(a, a)
	// semua code setelah return tidak akan dieksekusi
}

// cara 5, mendefinisikan sebuah function yang memiliki 2 parameter dan menghasilkan 2 value yg direturn
func tugas5(a, b int) (int, int) {
	return a + b, a * b
}

// cara 6, mendefinisikan sebuah function yang memiliki 2 parameter dan mengembalikan sebuah parameter yang memiliki "nama"
func tugas6(a, b int) (s int) {
	fmt.Println("s:", s) // variable s disini akan diinisialisasi dengan nilai 0 dalam fungsi ini
	s = a + b

	// secara otomatis akan mengembalikan nilai s
	return // ini dikenal dengan "naked" return
}

func main() {

	// memanggil fungsi2 yang sudah dibuat sebelumnya
	tugas1() // => "Ini adalah fungsi tugas1()"

	tugas2(1, 2) // => "Hasil Jumlah : 3"

	tugas3(1, 2, 3, 1.1, 2.2, "ok") // 1, 2, 3, 1.1, 2.2, "ok"

	fmt.Println(tugas4(2)) // 4

	hasilTambah, hasilKali := tugas5(2, 3) // 5, 6
	fmt.Println("Hasil Tambah & Hasil Bagi :", hasilTambah, hasilKali)

	hasilJumlah := tugas6(4, 5) // 9
	fmt.Println("Hasil Jumlah :", hasilJumlah)
}
