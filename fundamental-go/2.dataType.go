package main

import "fmt"

func main() {
	// Tipe data di Golang = String, Integer, Boolean, Float, Array, Slice, Map
	// 1. Tipe Data Integer
	var nilai int
	fmt.Println(nilai)

	// 2. Tipe Data float
	var harga float64
	fmt.Println(harga)

	// 3. Tipe Data String
	var name string
	fmt.Println(name)

	// 4. Tipe Data boolean
	var done bool
	fmt.Println(done)

	// 5. Tipe Data Array => tipe data yang dapat menampung multiple values dengan cara perlu mendefinisikan jumlah data yang dapat ditampung (ukuran tetap)
	var numbers = [4]int{7, 8, 99, 100}
	fmt.Println(numbers)

	// 6. Tipe Data Slice => tipe data yang dapat menyimpan multiple values tanpa perlu mendefinisikan jumlah data yang dapat disimpan (ukuran dinamis + lebih fleksible)
	var cities = []string{"Jakarta", "Bogor", "Depok", "Tangerang", "Bekasi"}
	fmt.Println("slice1 :", cities)
	fmt.Println("slice2 :", cities[1])
	fmt.Println("slice3 :", cities[2:])
	fmt.Println("slice4 :", cities[:3])
	fmt.Println("Jumlah data cities :", len(cities))

	// 7. Tipe Data Map => tipe data yang digunakan untuk merepresentasikan data pasangan nilai (Key-value pairs) yang tidak diurutkan. Dalam tipe data map, setiap key unit terkait dengan satu value
	balances := map[string]int{
		"USD": 200,
		"EUR": 150,
	}
	fmt.Println(balances)

	// 8. Tipe Data Struct => tipe data yang digunakan untuk mengelompokan beberapa nilai yang diterkait menjadi satu entitas tunggal. Struct dapat menggabungan berbagai tipe data termasuk struct itu sendiri.
	type Person struct {
		name string
		age  int
	}
	var data Person
	data.name = "joko"
	data.age = 28
	fmt.Println(data)
}
