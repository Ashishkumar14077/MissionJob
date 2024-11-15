package main

import "fmt"

func func1(x int){
	//var x int 
	//x = 5
	fmt.Printf("Address of X is %v after Call By Value where x is %d\n",&x,x)
}

func func2(y *int){
	fmt.Printf("Address of Y is %v after Call By Refernce where value of y is %d\n",y,*y)
}

func main(){
	var x,y int
	x = 5
	y = 10

	fmt.Printf("Address Of X is %v and Y is %v\n",&x,&y)
	// fmt.Println("Value of X=",x)
	// fmt.Println("Refernce of X=",&x)
	fmt.Println("Value of Y=",y)
	fmt.Println("Refernce of Y=",&y)
	func1(x)  // call by Value
	func2(&y)	// call by Refernce(Address)
}