package task3

import (
	"fmt"

	"github.com/AnggaArdhinata/basic_go/src/task3/interfaces"
)


type BangunStruct struct {
	Panjang, Lebar, Tinggi *float64
}

var Panjang float64
var Lebar float64
var Tinggi float64

var Bangun = BangunStruct{Panjang: &Panjang, Lebar: &Lebar, Tinggi: &Tinggi}

func (b *BangunStruct) Luas()  float64{
	p := *b.Panjang
	l := *b.Lebar

	result := p * l

	fmt.Println("Luas: ", result)
	return result
}

func (b *BangunStruct) Keliling()  float64{
	p := *b.Panjang
	l := *b.Lebar

	result := 2*p + 2*l

	fmt.Println("Keliling: ", result)
	return result
}

func (b *BangunStruct) Volume()  float64{
	p := *b.Panjang
	l := *b.Lebar
	t := *b.Tinggi

	result := p * l * t

	fmt.Println("Volume: ", result)
	return result
	
}

func (b *BangunStruct) Hitung(hasil interfaces.Hitung) {
	hasil.Luas()
	hasil.Keliling()
	hasil.Volume()
}
