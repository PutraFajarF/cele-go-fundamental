package main

import (
	"encoding/json"
	"fmt"
)

/*
Basic program yang akan kita buat pada malam ini adalah aplikasi pemesanan buku, dimana nantinya kita memesan buku dengan input nama buku dan jumlah bukunya. Jika nama buku tersedia dan stock buku tersedia maka pemesanan akan berhasil dan sebaliknya.
*/

// 2. Definisikan struct Books yang akan menyimpan data array dari BookList yg sudah didefinisikan sebelumnya
type Book struct {
	Item []*BookList
}

// 1. Definisikan struct Booklist yang akan menyimpan Field seperti nama buku, genre buku, dan Amount (jumlah buku)
type BookList struct {
	Name   string
	Genre  string
	Amount int32
}

func main() {
	book, bookList := BookStore("Java Programming for Kids", 20)

	fmt.Println(book, string(bookList))
}

// 3. Buat fungsi BookStore untuk melakukan pemesanan buku
func BookStore(bookName string, amount int32) (result string, bookList []byte) {
	// 4. Definisikan isi data dari struct Book dan BookList yang sudah dibuat sebelumnya
	book := Book{
		Item: []*BookList{
			{
				Name:   "Harry Potter",
				Genre:  "Fantasy",
				Amount: 20,
			},
			{
				Name:   "Bobo",
				Genre:  "Kids",
				Amount: 30,
			},
			{
				Name:   "Tsubasa",
				Genre:  "Soccer",
				Amount: 14,
			},
			{
				Name:   "Java Programming for Kids",
				Genre:  "Technology",
				Amount: 45,
			},
		},
	}

	// 5. Selanjutnya kita akan looping data book diatas dengan metode for range, semua value yg ada di book.Item akan di assign ke variable v
	for _, v := range book.Item {
		// 6. lakukan pengecekan apabila value name yg ada di data book tersebut sama seperti dengan nama buku yg dipesan, jika tidak maka akan menuju kondisi else
		if v.Name == bookName {
			// 7. Jika buku yg dipesan sesuai dengan data book yg tersedia lakukan pengecekan apakah stock buku cukup dengan jumlah buku yg dipesan
			if v.Amount < amount {
				result = fmt.Sprintf("Book named %s is out of stock", v.Name)
				// json.MarshalIndent() adalah fungsi yg dibuat untuk mengubah data menjadi format data JSON dengan indentasi dan format yg terstruktur
				bookList, _ = json.MarshalIndent(book.Item, "", "\t")

				return result, bookList
			}

			if v.Amount >= amount {
				v.Amount -= amount
				result = fmt.Sprintf("Book name %s successfully purchased", v.Name)
				bookList, _ = json.MarshalIndent(book.Item, "", "\t")

				return result, bookList
			}
		} else {
			result = fmt.Sprintf("Book not found")
			bookList, _ = json.MarshalIndent(book.Item, "", "\t")
			continue
		}
	}

	return
}
