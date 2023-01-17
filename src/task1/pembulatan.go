package task1

import (
	"fmt"
	"math"
)

func Pembulatan(num float64) float64{

	round := float64(math.Round(num * 10) / 10)
	fmt.Print("Nilai Pembulatan : ")
	fmt.Printf("%.2f", round)
	
	fmt.Println()
	return round

}