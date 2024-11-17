package basics

import "fmt"


func Patter2(){
	/*
	*
	**
	***
	****
	*****
	*/

	star := "*"
	for i:=0;i<5;i++{
		for j:=0;j<=i;j++{
			fmt.Printf(star)
		}
		fmt.Printf("\n")
	}
}


func Pattern3(){
	// 1
	// 12
	// 123
	// 1234
	// 12345

   for i:=0;i<5;i++{
	for j:=0;j<=i;j++{
		fmt.Printf("%d",j+1)
	}
	fmt.Printf("\n")
   }
}