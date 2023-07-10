package main

import "fmt"

func main() {
	harga, isDiscount := 5000, true

	// if condition_that_evaluates_to_boolean{
	//      perform action1
	// }else if condition_that_evaluates_to_boolean{
	//      perform action2
	// }else{
	//      perform action3
	// }

	// If-else seringkali dikombinasikan dengan hasil operator-operator yang sudah dijelaskan sebelumnya yang menjadi acuan eksekusi kode berikutnya.
	if harga >= 20000 {
		fmt.Println("Terlalu mahal")
	} else if harga <= 15000 && harga > 10000 && isDiscount == false {
		fmt.Println("Tidak ada discount")
	} else if harga <= 10000 && isDiscount == true {
		fmt.Println("Beli sekarang!")
	} else {
		fmt.Println("Boleh dibeli")
	}
}
