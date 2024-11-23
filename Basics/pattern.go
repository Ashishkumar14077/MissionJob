package basics

import (
	"fmt"
)


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

func Pattern4(){
	// 1
	// 22
	// 333
	// 4444
	// 55555

	for i:=0;i<5;i++{
		for j:=0;j<=i;j++{
			fmt.Printf("%d",i+1)
		}
		fmt.Printf("\n")
	}
}

func Pattern5(){
	// *****
	// ****
	// ***
	// **
	// *

	for i:=5;i>=1;i--{
		for j:=0;j<i;j++{
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}

func Pattern6(){
	// 12345
	// 1234
	// 123
	// 12
	// 1

	for i:=5;i>=1;i--{
		for j:=0;j<i;j++{
			fmt.Printf("%d",j+1)
		}
		fmt.Printf("\n")
	}
}

func Pattern7(){
/*
	ssss*	
	sss***
	ss*****
	s*******
	*********

	1. space <-
	2. right tringle
	3. left tringle

	SSSS* 
	SSS@**
	SS@@***
	S@@@****
	@@@@*****

*/
 for i:=0;i<5;i++{
	for j:=0;j<(4-i);j++{
		fmt.Printf(" ")
	}
	for k:=0;k<i;k++{
		fmt.Printf("*")
	}
	for m:=0;m<i+1;m++{
		fmt.Printf("*")
	}
	fmt.Printf("\n")
 }
}

func Pattern8(){
	// *****@@@@
	// S****@@@
	// SS***@@
	// SSS**@
	// SSSS*
	for i:=0;i<5;i++{
		for j:=0;j<i;j++{
			fmt.Printf(" ")
		}
		for k:=0;k<5-i;k++{
			fmt.Printf("*")
		}
		for m:=0;m<4-i;m++{
			fmt.Printf("*")
		}	
		fmt.Printf("\n")
	}
}

func Pattern9(){
	//   1
	//   01
	//   101
	//   0101
	//   10101

	for i:=0;i<5;i++{
		for j:=0;j<=i;j++{
			if (i+j)%2==0{
				//even
				fmt.Printf("1")
			} else if (i+j)%2!=0{
				//odd
				fmt.Printf("0")
			}
		}
		fmt.Printf("\n")
	}
}