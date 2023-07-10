package main

import "fmt"

/*
Di Go, pointer (*) adalah tipe data yang digunakan untuk menyimpan alamat memori dari suatu variabel. Pointer memungkinkan kita untuk mengakses dan memanipulasi nilai yang tersimpan di lokasi memori yang sama, oleh beberapa variabel atau fungsi, sehingga memungkinkan penggunaan efisien dan perubahan nilai secara langsung.
*/

func main() {
	var num int = 42
	var ptr *int

	ptr = &num // Mengambil alamat memori dari variabel num

	fmt.Println("Nilai dari num:", num)
	fmt.Println("Alamat memori dari num:", &num)
	fmt.Println("Nilai yang ditunjuk oleh pointer:", *ptr)
	fmt.Println("Alamat memori yang ditunjuk oleh pointer:", ptr)

	*ptr = 100 // Mengubah nilai yang ditunjuk oleh pointer
	fmt.Println("Nilai yang ditunjuk oleh pointer setelah diubah:", *ptr)
	fmt.Println("Nilai dari num setelah diubah:", num)
}
