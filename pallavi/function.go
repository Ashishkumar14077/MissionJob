package main

import "fmt"

func Sum(x int, y *int){
	fmt.Printf("address of x %v\n", &x)
	fmt.Printf("address of y %v\n", y)
	fmt.Printf("value of x %d\n", x)
	fmt.Printf("value of y %d\n", *y)

	x = x+5
	*y = *y+20
	fmt.Printf("after adding value %d\n",x)
	fmt.Printf("after adding value %d\n",*y)
}

func main(){
	x := 10
	y := 20
	fmt.Printf("address of x: %v\n", &x)
	fmt.Printf("address of y: %v\n", &y)
	Sum(x, &y)
	fmt.Printf("value of x: %v\n", x)
	fmt.Printf("value of y: %v\n",y)
}
