package basics

import "fmt"


func Func1(x int){
	//var x int 
	//x = 5
	fmt.Printf("Address of X is %v after Call By Value where x is %d\n",&x,x)
}

func Func2(y *int){
	fmt.Printf("Address of Y is %v after Call By Refernce where value of y is %d\n",y,*y)
}
