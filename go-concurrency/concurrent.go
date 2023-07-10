package main

import (
	"fmt"
	"time"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	// Membuat channel untuk mengirim hasil penjumlahan
	sumChannel := make(chan int)

	// Membuat goroutine untuk menjumlahkan angka-angka secara konkuren
	go sum(numbers, sumChannel)

	// Menerima hasil penjumlahan dari channel
	result := <-sumChannel

	fmt.Println("Hasil penjumlahan:", result)
}

func sum(numbers []int, resultChannel chan<- int) {
	sum := 0
	for _, num := range numbers {
		sum += num
		time.Sleep(time.Millisecond * 100) // Simulasi operasi yang membutuhkan waktu
	}

	// Mengirim hasil penjumlahan ke channel
	resultChannel <- sum
}
