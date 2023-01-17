package main

import (
	"fmt"
	
	"github.com/AnggaArdhinata/basic_go/src/task1"
	"github.com/AnggaArdhinata/basic_go/src/task2"
	"github.com/AnggaArdhinata/basic_go/src/task3"
)

func main() {

	//! Tugas 1 - Pembulatan
	task1.Pembulatan(4.37)
	task1.Pembulatan(4.32)
	task1.Pembulatan(4.324)

	//! Tugas 2 - Deret Bilangan
	task2.Nilai = 20

	fmt.Println("Angka: ", task2.Nilai)

	task2.Deret.Ganjil()
	task2.Deret.Genap()
	task2.Deret.Prima()
	task2.Deret.Fibonacci()

	//! Tugas 3 - Interface
	task3.Panjang = 9.7
	task3.Lebar = 6.0
	task3.Tinggi = 5.0

	task3.Bangun.Hitung(&task3.Bangun)
}
