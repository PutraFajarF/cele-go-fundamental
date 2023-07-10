package main

import "fmt"

func main() {
	a, b := 10, 5.5

	//** ARITHMETIC OPERATORS => sebuah operator yg digunakan untuk menghitung sebuah angka pada variable **//
	//  +       Penjumlahan
	// -        Pengurangan
	// *        Perkalian
	// /        Pembagian
	// %        Sisa Bagi
	// tidak ada operator pangkat di Golang, biasanya menggunakan math.Pow(a, b) dari library bawaan Golang

	fmt.Println(a + 5)   // => 15
	fmt.Println(3.1 - b) // => -2.4
	fmt.Println(a * a)   // => 100
	fmt.Println(a / a)   // => 1
	fmt.Println(11 % 5)  // => 2

	// Go merupakan Strong Typed Language
	// fmt.Println(a * b)          // =>  invalid operation: a * b (mismatched types int and float64)
	fmt.Println(a * int(b))     // => 50
	fmt.Println(float64(a) * b) // => 55

	// IncDec Statements
	// Pernyataan "++" dan "--" menambah atau mengurangi valuenya dengan nilai 1
	x := 10
	x++ // x is 11. Same as: x += 1
	x-- // x is 10. Same as: x -= 1

	//** ASSIGNMENT OPERATORS => operator yang digunakan untuk memberikan nilai ke suatu variabel dengan menggunakan operasi matematika atau operasi lainnya. Operator ini mempersingkat penulisan ekspresi dan memberikan hasil operasi langsung kepada variabel yang ditentukan**//
	//  =   (simple assignment)
	// +=   (increment assignment)
	// -=   (decrement assignment)
	// *=   (multiplication assignment)
	// /=   (division assignment)
	// %=   (modulus assignment)

	a = 10
	a += 2 // => a is 12
	a -= 3 // => a is 9
	a *= 2 // => a is 18
	a /= 3 // => a is 6
	a %= 5 // => a is 1

	//** COMPARISON OPERATORS => operator yang digunakan untuk membandingkan variable satu dengan yg lainnya sehingga menghasilkan output TRUE or FALSE **//
	//  ==      Sama Dengan
	// !=       Tidak Sama Dengan
	// >        Lebih Besar Dari
	// <        Lebih Kecil Dari
	// >=       Lebih Besar Sama Dengan Dari
	// <=       Lebih Kecil Sama Dengan Dari

	fmt.Println(5 == 6)   // => false
	fmt.Println(5 != 6)   // => true
	fmt.Println(10 > 10)  // => false
	fmt.Println(10 >= 10) // => true
	fmt.Println(5 < 5)    // => false
	fmt.Println(5 <= 5)   // => true

	//** LOGICAL OPERATORS => Operator yg digunakan untuk membandingkan antara comparison satu dengan comparison lainnya dan menghasilkan output TRUE atau FALSE **//
	// &&       logical AND
	// ||       logical OR
	// !        logical Negation / NOT

	fmt.Println(0 < 2 && 4 > 1) // => true
	fmt.Println(1 > 5 || 4 > 5) // => false
	fmt.Println(!(1 > 2))       // => true

}
