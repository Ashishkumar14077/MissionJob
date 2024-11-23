package basics

import "fmt"

func Count_digit(num int) {
	count := 0
	for num>0{
		Dummy(&num)
		count++
	}
	fmt.Println("Digits : ",count)
}

func Dummy(n *int){
	*n = *n/10;
}