package task2

import (
	"fmt"
	"strconv"
	
)


type DeretBilang struct {
	Limit *int
}

var Nilai int

var Deret = DeretBilang{Limit: &Nilai}


func (n *DeretBilang) Ganjil()  string{
	result := ""
	for i := 0; i < *n.Limit; i++ {
		if i%2 !=0 {
			result += strconv.Itoa(i)+" "
		}

	}
	fmt.Println("Bilangan Ganjil: ", result)

	return result
}

func (n *DeretBilang) Genap()  string{
	result := ""
	for i := 1; i <= *n.Limit; i++ {
		if i%2 == 0 {
			result += strconv.Itoa(i)+" "
		}
	}
	fmt.Println("Bilangan Genap: ", result)

	return result
}

func (n *DeretBilang) Prima() string{
	result := ""
	var i, j, flag int

	for i = 2; i < *n.Limit; i++ {
		flag = 0
		for j = 2; j < i; j++ {
			if i % j == 0 {
				flag = 1
				break
			}
		}
		if i > 1 && flag == 0 {
			result += strconv.Itoa(i)+" "
		}
	
	}
	
	fmt.Println("Bilangan Prima: ", result)

	return result

	
}

func (n *DeretBilang) Fibonacci()  string{
	result := ""
	n1 := 0
	n2 := 1
	n3 := n2

	for i := 1; i <= *n.Limit; i++ {
		result += strconv.Itoa(n2)+" "
		n3 = n1 + n2
		n1 = n2
		n2 = n3
		if n2 > *n.Limit {
			break
		}
	}
	
	fmt.Println("Bilangan Fibonacci: ", result)
	
	return result
}


//? Cara 2 Fibonacci 
// for {
// 	n3 = n2
// 	n2 = n1 + n2
// 	if n2 >= *n.Limit {
// 		break
// 	}
// 	n1 = n3
// 	result += strconv.Itoa(n2)+" "
// }

// fmt.Println("Bilangan Fibonacci: ", strings.Split(result, " "))
