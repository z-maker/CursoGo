package main

import (
	"fmt"
)

func main() {

	x :=10
	y:=5
	z:=10

	if x == z{
		fmt.Printf("%d ok es igual que %d \n",x,z)
	}
	if x != y{
		fmt.Printf("%d ok es diferente que %d \n",x,y)
	}
	if y < x{
		fmt.Printf("%d ok es menor que %d \n",y,x)
	}
	if z < y{
		fmt.Printf("%d ok es mayor que %d \n",z,y)
	}else{
		fmt.Printf("else \n")
	}
}