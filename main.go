package main

import (
	"fmt"
	basics "goLearn/mission_job/Basics"
)

func main(){
	var x,y int
	x = 5
	y = 10

	fmt.Printf("Address Of X is %v and Y is %v\n",&x,&y)
	// fmt.Println("Value of X=",x)
	// fmt.Println("Refernce of X=",&x)
	fmt.Println("Value of Y=",y)
	fmt.Println("Refernce of Y=",&y)
	basics.Func1(x)  // call by Value
	basics.Func2(&y)	// call by Refernce(Address)
}