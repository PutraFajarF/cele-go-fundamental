package main

import "fmt"

func main() {
	// For with clause => sebuah perulangan berdasarkan kondisi yang telah ditetapkan
	// print angka dari 0 sampai 9
	for i := 0; i < 10; i++ {
		fmt.Println("Angka i:", i)
	}

	// cara for loop yang memiliki efek yang sama dengan while loop di bahasa pemrograman lain
	// di golang tidak ada while loop
	j := 10
	for j >= 0 {
		fmt.Println("Nilai j:", j)
		j--
	}

	// handling multiple variable di dalam for loop
	for i, j := 0, 100; i < 10; i, j = i+1, j+1 {
		fmt.Println("nilai i dan j :", i, j)
	}

	// For with Range => sebuah perulangan berdasarkan data array atau slice yang telah dideklarasikan sebelumnya
	var cities = []string{"Jakarta", "Bogor", "Depok", "Tangerang", "Bekasi", "Bandung"}

	for i, kota := range cities {
		fmt.Println("Nomor:", i, "Kota:", kota)
		fmt.Printf("Nomor: %d, Kota: %s\n", i, kota)
	}
}
